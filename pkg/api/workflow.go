package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/vorteil/direktiv/pkg/ingress"
)

func (h *Handler) getUIDforName(ns, name string) (string, error) {

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByName(ctx, &ingress.GetWorkflowByNameRequest{
		Namespace: &ns,
		Name:      &name,
	})

	if err != nil {
		return "", err
	}

	return resp.GetUid(), nil

}

func (h *Handler) workflows(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflows(ctx, &ingress.GetWorkflowsRequest{
		Namespace: &n,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	// wip uids
	for _, t := range resp.Workflows {
		t.Uid = nil
	}

	writeData(resp, w)

}

func (h *Handler) getWorkflow(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByName(ctx, &ingress.GetWorkflowByNameRequest{
		Namespace: &n,
		Name:      &name,
	})

	if err != nil {
		ErrResponse(w, err)
		return
	}

	// wipe uid
	resp.Uid = nil

	writeData(resp, w)

}

func (h *Handler) updateWorkflow(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	uid, err := h.getUIDforName(ns, name)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	var useRevision bool
	rev, err := strconv.Atoi(r.URL.Query().Get("revision"))
	if err == nil {
		useRevision = true
	}

	var active bool
	var useActive bool
	if val, ok := r.URL.Query()["active"]; ok {
		if val[0] == "true" {
			active = true
			useActive = true
		} else if val[0] == "false" {
			useActive = true
		}
	}

	var logEvent string
	var useLogEvent bool
	if val, ok := r.URL.Query()["logEvent"]; ok {
		logEvent = val[0]
		useLogEvent = true
	}

	revision := int32(rev)

	// Read Body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	var contentType string
	if typeMap, ok := r.Header["Content-Type"]; ok {
		contentType = typeMap[0]
	}

	switch contentType {
	case "text/yaml":
	default:
		ErrResponse(w, fmt.Errorf("content type '%s' is not supported. supported media types: 'text/yaml'", contentType))
		return
	}

	// Construct direktiv GRPC Request
	request := ingress.UpdateWorkflowRequest{
		Uid:      &uid,
		Workflow: b,
	}

	if useActive {
		request.Active = &active
	}

	if useLogEvent {
		request.LogToEvents = &logEvent
	}

	if useRevision {
		request.Revision = &revision
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.UpdateWorkflow(ctx, &request)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	writeData(resp, w)
}

func (h *Handler) toggleWorkflow(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	uid, err := h.getUIDforName(ns, name)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByUid(ctx, &ingress.GetWorkflowByUidRequest{
		Uid: &uid,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	active := !*resp.Active

	resp2, err := h.s.direktiv.UpdateWorkflow(ctx, &ingress.UpdateWorkflowRequest{
		Uid:      &uid,
		Active:   &active,
		Workflow: resp.Workflow,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	writeData(resp2, w)

}

func (h *Handler) createWorkflow(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	active := true

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	var contentType string
	if typeMap, ok := r.Header["Content-Type"]; ok {
		contentType = typeMap[0]
	}

	if contentType != "text/yaml" {
		ErrResponse(w, fmt.Errorf("content type '%s' is not supported. supported media types: 'text/yaml'", contentType))
		return
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.AddWorkflow(ctx, &ingress.AddWorkflowRequest{
		Namespace: &n,
		Active:    &active,
		Workflow:  b,
	})

	if err != nil {
		ErrResponse(w, err)
		return
	}

	// empty UID which omits it
	resp.Uid = nil

	retData, err := json.Marshal(resp)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(retData)

}

func (h *Handler) deleteWorkflow(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	uid, err := h.getUIDforName(ns, name)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.DeleteWorkflow(ctx, &ingress.DeleteWorkflowRequest{
		Uid: &uid,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	writeData(resp, w)

}

func (h *Handler) downloadWorkflow(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	uid, err := h.getUIDforName(ns, name)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByUid(ctx, &ingress.GetWorkflowByUidRequest{
		Uid: &uid,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+*resp.Id+".yaml")
	w.Header().Set("Content-Type", "application/x-yaml")

	writeData(resp, w)

}

func (h *Handler) executeWorkflow(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	name := mux.Vars(r)["workflowTarget"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrResponse(w, err)
		return
	}

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.InvokeWorkflow(ctx, &ingress.InvokeWorkflowRequest{
		Namespace: &ns,
		Name:      &name,
		Input:     b,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		ErrResponse(w, err)
		return
	}

}

func (h *Handler) workflowInstances(w http.ResponseWriter, r *http.Request) {

	ns := mux.Vars(r)["namespace"]
	wf := mux.Vars(r)["workflowTarget"]

	o, l := paginationParams(r)
	offset := int32(o)
	limit := int32(l)

	ctx, cancel := CtxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetInstancesByWorkflow(ctx, &ingress.GetInstancesByWorkflowRequest{
		Offset:    &offset,
		Limit:     &limit,
		Namespace: &ns,
		Workflow:  &wf,
	})
	if err != nil {
		ErrResponse(w, err)
		return
	}

	writeData(resp, w)

}
