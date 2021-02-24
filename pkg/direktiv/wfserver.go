package direktiv

import (
	"context"

	"github.com/vorteil/direktiv/pkg/flow"
	"github.com/vorteil/direktiv/pkg/ingress"
	"github.com/vorteil/direktiv/pkg/isolate"
	"github.com/vorteil/direktiv/pkg/secrets"
	"google.golang.org/grpc"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres for ent
	log "github.com/sirupsen/logrus"
	"github.com/vorteil/direktiv/pkg/dlog"
)

const (
	workflowOnly   = "wfo"
	workflowRunner = "wfr"
	workflowAll    = "wf"
)

const (
	lockID   = 2610
	lockWait = 10
)

type component interface {
	start() error
	stop()
	name() string
	setClient(*WorkflowServer) error
}

type componentAPIs struct {
	flowClient    flow.DirektivFlowClient
	secretsClient secrets.SecretsServiceClient
	isolateClient isolate.DirektivIsolateClient
	ingressClient ingress.DirektivIngressClient

	conns []*grpc.ClientConn
}

// WorkflowServer is a direktiv server
type WorkflowServer struct {
	id         uuid.UUID
	ctx        context.Context
	config     *Config
	serverType string

	dbManager     *dbManager
	tmManager     *timerManager
	engine        *workflowEngine
	actionManager *actionManager

	LifeLine       chan bool
	instanceLogger dlog.Log
	secrets        secrets.SecretsServiceClient

	components    map[string]component
	componentAPIs componentAPIs
}

func (s *WorkflowServer) initWorkflowServer() error {

	var err error

	// prep timers
	s.tmManager, err = newTimerManager(s)
	if err != nil {
		return err
	}

	s.engine, err = newWorkflowEngine(s)
	if err != nil {
		return err
	}

	// register the timer functions
	var timerFunctions = map[string]func([]byte) error{
		timerCleanOneShot:         s.tmManager.cleanOneShot,
		timerCleanInstanceRecords: s.tmManager.cleanInstanceRecords,
	}

	for n, f := range timerFunctions {
		err := s.tmManager.registerFunction(n, f)
		if err != nil {
			return err
		}
	}

	addCron := func(name, cron string) {
		// add clean up timers
		_, err := s.dbManager.getTimer(name)

		// on error we assuming it is not in the database
		if err != nil {
			s.tmManager.addCron(name, name, cron, []byte(""))
		}

	}

	addCron(timerCleanOneShot, "*/10 * * * *")
	addCron(timerCleanInstanceRecords, "0 * * * *")

	return nil

}

// NewWorkflowServer creates a new workflow server
func NewWorkflowServer(config *Config, serverType string) (*WorkflowServer, error) {

	log.Debugf("server type: %s", serverType)
	ctx := context.Background()

	var (
		err error
	)

	s := &WorkflowServer{
		id:         uuid.New(),
		ctx:        ctx,
		LifeLine:   make(chan bool),
		serverType: serverType,
		config:     config,
		components: make(map[string]component),
	}

	s.dbManager, err = newDBManager(ctx, s.config.Database.DB)
	if err != nil {
		return nil, err
	}

	if s.isWorkflowServer() {
		err = s.initWorkflowServer()
		if err != nil {
			return nil, err
		}
	}

	// if s.isRunnerServer() {

	am, err := newActionManager(s.config, s.dbManager, &s.instanceLogger)
	if err != nil {
		return nil, err
	}
	s.actionManager = am
	s.components[isolateComponent] = am
	// }

	ingressServer := newIngressServer(s)
	s.components[ingressComponent] = ingressServer

	healthServer := newHealthServer(config)
	s.components[healthComponent] = healthServer

	flowServer := newFlowServer(config, s.engine)
	s.components[flowComponent] = flowServer

	secretsServer, err := newSecretsServer(config)
	if err != nil {
		log.Errorf("can not create secret server: %v", err)
		return nil, err
	}
	s.components[secretsComponent] = secretsServer

	for _, comp := range s.components {
		log.Debugf("starting %s component", comp.name())
		err := comp.start()
		if err != nil {
			log.Errorf("can not start: %v", err)
			return nil, err
		}
	}

	for _, comp := range s.components {
		log.Debugf("creating client for %s component", comp.name())
		err = comp.setClient(s)
		if err != nil {
			log.Errorf("can not create client: %v", err)
			return nil, err
		}
	}

	// TODO: that move in isolate manager
	am.grpcFlow = s.componentAPIs.flowClient

	return s, nil

}

func (s *WorkflowServer) isWorkflowServer() bool {
	if s.serverType == workflowOnly || s.serverType == workflowAll {
		return true
	}
	return false
}

func (s *WorkflowServer) isRunnerServer() bool {
	if s.serverType == workflowRunner || s.serverType == workflowAll {
		return true
	}
	return false
}

// SetInstanceLogger set logger for direktiv for firecracker instances
func (s *WorkflowServer) SetInstanceLogger(l dlog.Log) {
	s.instanceLogger = l
}

// Lifeline interface impl
func (s *WorkflowServer) Lifeline() chan bool {
	return s.LifeLine
}

func (s *WorkflowServer) cleanup() {

	// closing db at the end
	if s.dbManager != nil {
		defer s.dbManager.dbEnt.Close()
	}

	if s.tmManager != nil {
		s.tmManager.stopTimers()
	}

	// stop components
	for _, comp := range s.components {
		comp.stop()
	}

	// close grpc if client has been used
	for _, conn := range s.componentAPIs.conns {
		conn.Close()
	}

}

// Stop stops the server gracefully
func (s *WorkflowServer) Stop() {

	go func() {

		log.Printf("stopping workflow server")
		s.cleanup()
		s.LifeLine <- true

	}()
}

// Kill kills the server
func (s *WorkflowServer) Kill() {

	go func() {

		defer func() {
			_ = recover()
		}()

		s.cleanup()
		s.LifeLine <- true

	}()

}

// Run starts all components of direktiv
func (s *WorkflowServer) Run() error {

	// subscribe to cmds
	if s.isWorkflowServer() {

		log.Debugf("subscribing to sync queue")
		err := s.startDatabaseListener()
		if err != nil {
			s.Kill()
			return err
		}

		// start timers
		err = s.tmManager.startTimers()
		if err != nil {
			s.Kill()
			return err
		}

	}

	if s.actionManager != nil {
		err := s.actionManager.start()
		if err != nil {
			s.Kill()
			return err
		}
	}

	return nil

}
