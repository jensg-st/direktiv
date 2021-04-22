// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ingress

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DirektivIngressClient is the client API for DirektivIngress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirektivIngressClient interface {
	AddNamespace(ctx context.Context, in *AddNamespaceRequest, opts ...grpc.CallOption) (*AddNamespaceResponse, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	GetNamespaces(ctx context.Context, in *GetNamespacesRequest, opts ...grpc.CallOption) (*GetNamespacesResponse, error)
	AddWorkflow(ctx context.Context, in *AddWorkflowRequest, opts ...grpc.CallOption) (*AddWorkflowResponse, error)
	DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*DeleteWorkflowResponse, error)
	GetWorkflowByName(ctx context.Context, in *GetWorkflowByNameRequest, opts ...grpc.CallOption) (*GetWorkflowByNameResponse, error)
	GetWorkflowByUid(ctx context.Context, in *GetWorkflowByUidRequest, opts ...grpc.CallOption) (*GetWorkflowByUidResponse, error)
	GetWorkflowInstance(ctx context.Context, in *GetWorkflowInstanceRequest, opts ...grpc.CallOption) (*GetWorkflowInstanceResponse, error)
	GetWorkflowInstances(ctx context.Context, in *GetWorkflowInstancesRequest, opts ...grpc.CallOption) (*GetWorkflowInstancesResponse, error)
	GetInstancesByWorkflow(ctx context.Context, in *GetInstancesByWorkflowRequest, opts ...grpc.CallOption) (*GetInstancesByWorkflowResponse, error)
	GetWorkflowInstanceLogs(ctx context.Context, in *GetWorkflowInstanceLogsRequest, opts ...grpc.CallOption) (*GetWorkflowInstanceLogsResponse, error)
	CancelWorkflowInstance(ctx context.Context, in *CancelWorkflowInstanceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetWorkflows(ctx context.Context, in *GetWorkflowsRequest, opts ...grpc.CallOption) (*GetWorkflowsResponse, error)
	InvokeWorkflow(ctx context.Context, in *InvokeWorkflowRequest, opts ...grpc.CallOption) (*InvokeWorkflowResponse, error)
	UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*UpdateWorkflowResponse, error)
	BroadcastEvent(ctx context.Context, in *BroadcastEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	StoreSecret(ctx context.Context, in *StoreSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetRegistries(ctx context.Context, in *GetRegistriesRequest, opts ...grpc.CallOption) (*GetRegistriesResponse, error)
	DeleteRegistry(ctx context.Context, in *DeleteRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	StoreRegistry(ctx context.Context, in *StoreRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	WorkflowMetrics(ctx context.Context, in *WorkflowMetricsRequest, opts ...grpc.CallOption) (*WorkflowMetricsResponse, error)
}

type direktivIngressClient struct {
	cc grpc.ClientConnInterface
}

func NewDirektivIngressClient(cc grpc.ClientConnInterface) DirektivIngressClient {
	return &direktivIngressClient{cc}
}

func (c *direktivIngressClient) AddNamespace(ctx context.Context, in *AddNamespaceRequest, opts ...grpc.CallOption) (*AddNamespaceResponse, error) {
	out := new(AddNamespaceResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/AddNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetNamespaces(ctx context.Context, in *GetNamespacesRequest, opts ...grpc.CallOption) (*GetNamespacesResponse, error) {
	out := new(GetNamespacesResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) AddWorkflow(ctx context.Context, in *AddWorkflowRequest, opts ...grpc.CallOption) (*AddWorkflowResponse, error) {
	out := new(AddWorkflowResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/AddWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*DeleteWorkflowResponse, error) {
	out := new(DeleteWorkflowResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/DeleteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflowByName(ctx context.Context, in *GetWorkflowByNameRequest, opts ...grpc.CallOption) (*GetWorkflowByNameResponse, error) {
	out := new(GetWorkflowByNameResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflowByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflowByUid(ctx context.Context, in *GetWorkflowByUidRequest, opts ...grpc.CallOption) (*GetWorkflowByUidResponse, error) {
	out := new(GetWorkflowByUidResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflowByUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflowInstance(ctx context.Context, in *GetWorkflowInstanceRequest, opts ...grpc.CallOption) (*GetWorkflowInstanceResponse, error) {
	out := new(GetWorkflowInstanceResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflowInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflowInstances(ctx context.Context, in *GetWorkflowInstancesRequest, opts ...grpc.CallOption) (*GetWorkflowInstancesResponse, error) {
	out := new(GetWorkflowInstancesResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflowInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetInstancesByWorkflow(ctx context.Context, in *GetInstancesByWorkflowRequest, opts ...grpc.CallOption) (*GetInstancesByWorkflowResponse, error) {
	out := new(GetInstancesByWorkflowResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetInstancesByWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflowInstanceLogs(ctx context.Context, in *GetWorkflowInstanceLogsRequest, opts ...grpc.CallOption) (*GetWorkflowInstanceLogsResponse, error) {
	out := new(GetWorkflowInstanceLogsResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflowInstanceLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) CancelWorkflowInstance(ctx context.Context, in *CancelWorkflowInstanceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/CancelWorkflowInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetWorkflows(ctx context.Context, in *GetWorkflowsRequest, opts ...grpc.CallOption) (*GetWorkflowsResponse, error) {
	out := new(GetWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) InvokeWorkflow(ctx context.Context, in *InvokeWorkflowRequest, opts ...grpc.CallOption) (*InvokeWorkflowResponse, error) {
	out := new(InvokeWorkflowResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/InvokeWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*UpdateWorkflowResponse, error) {
	out := new(UpdateWorkflowResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/UpdateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) BroadcastEvent(ctx context.Context, in *BroadcastEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/BroadcastEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error) {
	out := new(GetSecretsResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) StoreSecret(ctx context.Context, in *StoreSecretRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/StoreSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) GetRegistries(ctx context.Context, in *GetRegistriesRequest, opts ...grpc.CallOption) (*GetRegistriesResponse, error) {
	out := new(GetRegistriesResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/GetRegistries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) DeleteRegistry(ctx context.Context, in *DeleteRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/DeleteRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) StoreRegistry(ctx context.Context, in *StoreRegistryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/StoreRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivIngressClient) WorkflowMetrics(ctx context.Context, in *WorkflowMetricsRequest, opts ...grpc.CallOption) (*WorkflowMetricsResponse, error) {
	out := new(WorkflowMetricsResponse)
	err := c.cc.Invoke(ctx, "/ingress.DirektivIngress/WorkflowMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirektivIngressServer is the server API for DirektivIngress service.
// All implementations must embed UnimplementedDirektivIngressServer
// for forward compatibility
type DirektivIngressServer interface {
	AddNamespace(context.Context, *AddNamespaceRequest) (*AddNamespaceResponse, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	GetNamespaces(context.Context, *GetNamespacesRequest) (*GetNamespacesResponse, error)
	AddWorkflow(context.Context, *AddWorkflowRequest) (*AddWorkflowResponse, error)
	DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*DeleteWorkflowResponse, error)
	GetWorkflowByName(context.Context, *GetWorkflowByNameRequest) (*GetWorkflowByNameResponse, error)
	GetWorkflowByUid(context.Context, *GetWorkflowByUidRequest) (*GetWorkflowByUidResponse, error)
	GetWorkflowInstance(context.Context, *GetWorkflowInstanceRequest) (*GetWorkflowInstanceResponse, error)
	GetWorkflowInstances(context.Context, *GetWorkflowInstancesRequest) (*GetWorkflowInstancesResponse, error)
	GetInstancesByWorkflow(context.Context, *GetInstancesByWorkflowRequest) (*GetInstancesByWorkflowResponse, error)
	GetWorkflowInstanceLogs(context.Context, *GetWorkflowInstanceLogsRequest) (*GetWorkflowInstanceLogsResponse, error)
	CancelWorkflowInstance(context.Context, *CancelWorkflowInstanceRequest) (*empty.Empty, error)
	GetWorkflows(context.Context, *GetWorkflowsRequest) (*GetWorkflowsResponse, error)
	InvokeWorkflow(context.Context, *InvokeWorkflowRequest) (*InvokeWorkflowResponse, error)
	UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*UpdateWorkflowResponse, error)
	BroadcastEvent(context.Context, *BroadcastEventRequest) (*empty.Empty, error)
	GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*empty.Empty, error)
	StoreSecret(context.Context, *StoreSecretRequest) (*empty.Empty, error)
	GetRegistries(context.Context, *GetRegistriesRequest) (*GetRegistriesResponse, error)
	DeleteRegistry(context.Context, *DeleteRegistryRequest) (*empty.Empty, error)
	StoreRegistry(context.Context, *StoreRegistryRequest) (*empty.Empty, error)
	WorkflowMetrics(context.Context, *WorkflowMetricsRequest) (*WorkflowMetricsResponse, error)
	mustEmbedUnimplementedDirektivIngressServer()
}

// UnimplementedDirektivIngressServer must be embedded to have forward compatible implementations.
type UnimplementedDirektivIngressServer struct {
}

func (UnimplementedDirektivIngressServer) AddNamespace(context.Context, *AddNamespaceRequest) (*AddNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNamespace not implemented")
}
func (UnimplementedDirektivIngressServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedDirektivIngressServer) GetNamespaces(context.Context, *GetNamespacesRequest) (*GetNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaces not implemented")
}
func (UnimplementedDirektivIngressServer) AddWorkflow(context.Context, *AddWorkflowRequest) (*AddWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorkflow not implemented")
}
func (UnimplementedDirektivIngressServer) DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*DeleteWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflowByName(context.Context, *GetWorkflowByNameRequest) (*GetWorkflowByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowByName not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflowByUid(context.Context, *GetWorkflowByUidRequest) (*GetWorkflowByUidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowByUid not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflowInstance(context.Context, *GetWorkflowInstanceRequest) (*GetWorkflowInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowInstance not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflowInstances(context.Context, *GetWorkflowInstancesRequest) (*GetWorkflowInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowInstances not implemented")
}
func (UnimplementedDirektivIngressServer) GetInstancesByWorkflow(context.Context, *GetInstancesByWorkflowRequest) (*GetInstancesByWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstancesByWorkflow not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflowInstanceLogs(context.Context, *GetWorkflowInstanceLogsRequest) (*GetWorkflowInstanceLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowInstanceLogs not implemented")
}
func (UnimplementedDirektivIngressServer) CancelWorkflowInstance(context.Context, *CancelWorkflowInstanceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelWorkflowInstance not implemented")
}
func (UnimplementedDirektivIngressServer) GetWorkflows(context.Context, *GetWorkflowsRequest) (*GetWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflows not implemented")
}
func (UnimplementedDirektivIngressServer) InvokeWorkflow(context.Context, *InvokeWorkflowRequest) (*InvokeWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeWorkflow not implemented")
}
func (UnimplementedDirektivIngressServer) UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*UpdateWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkflow not implemented")
}
func (UnimplementedDirektivIngressServer) BroadcastEvent(context.Context, *BroadcastEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastEvent not implemented")
}
func (UnimplementedDirektivIngressServer) GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedDirektivIngressServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedDirektivIngressServer) StoreSecret(context.Context, *StoreSecretRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreSecret not implemented")
}
func (UnimplementedDirektivIngressServer) GetRegistries(context.Context, *GetRegistriesRequest) (*GetRegistriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistries not implemented")
}
func (UnimplementedDirektivIngressServer) DeleteRegistry(context.Context, *DeleteRegistryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegistry not implemented")
}
func (UnimplementedDirektivIngressServer) StoreRegistry(context.Context, *StoreRegistryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRegistry not implemented")
}
func (UnimplementedDirektivIngressServer) WorkflowMetrics(context.Context, *WorkflowMetricsRequest) (*WorkflowMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkflowMetrics not implemented")
}
func (UnimplementedDirektivIngressServer) mustEmbedUnimplementedDirektivIngressServer() {}

// UnsafeDirektivIngressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirektivIngressServer will
// result in compilation errors.
type UnsafeDirektivIngressServer interface {
	mustEmbedUnimplementedDirektivIngressServer()
}

func RegisterDirektivIngressServer(s grpc.ServiceRegistrar, srv DirektivIngressServer) {
	s.RegisterService(&DirektivIngress_ServiceDesc, srv)
}

func _DirektivIngress_AddNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).AddNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/AddNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).AddNamespace(ctx, req.(*AddNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetNamespaces(ctx, req.(*GetNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_AddWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).AddWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/AddWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).AddWorkflow(ctx, req.(*AddWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/DeleteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).DeleteWorkflow(ctx, req.(*DeleteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflowByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflowByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflowByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflowByName(ctx, req.(*GetWorkflowByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflowByUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowByUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflowByUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflowByUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflowByUid(ctx, req.(*GetWorkflowByUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflowInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflowInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflowInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflowInstance(ctx, req.(*GetWorkflowInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflowInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflowInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflowInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflowInstances(ctx, req.(*GetWorkflowInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetInstancesByWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstancesByWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetInstancesByWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetInstancesByWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetInstancesByWorkflow(ctx, req.(*GetInstancesByWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflowInstanceLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowInstanceLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflowInstanceLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflowInstanceLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflowInstanceLogs(ctx, req.(*GetWorkflowInstanceLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_CancelWorkflowInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelWorkflowInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).CancelWorkflowInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/CancelWorkflowInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).CancelWorkflowInstance(ctx, req.(*CancelWorkflowInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetWorkflows(ctx, req.(*GetWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_InvokeWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).InvokeWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/InvokeWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).InvokeWorkflow(ctx, req.(*InvokeWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_UpdateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).UpdateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/UpdateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).UpdateWorkflow(ctx, req.(*UpdateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_BroadcastEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).BroadcastEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/BroadcastEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).BroadcastEvent(ctx, req.(*BroadcastEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetSecrets(ctx, req.(*GetSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_StoreSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).StoreSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/StoreSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).StoreSecret(ctx, req.(*StoreSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_GetRegistries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).GetRegistries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/GetRegistries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).GetRegistries(ctx, req.(*GetRegistriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_DeleteRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).DeleteRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/DeleteRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).DeleteRegistry(ctx, req.(*DeleteRegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_StoreRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).StoreRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/StoreRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).StoreRegistry(ctx, req.(*StoreRegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivIngress_WorkflowMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivIngressServer).WorkflowMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ingress.DirektivIngress/WorkflowMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivIngressServer).WorkflowMetrics(ctx, req.(*WorkflowMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirektivIngress_ServiceDesc is the grpc.ServiceDesc for DirektivIngress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirektivIngress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ingress.DirektivIngress",
	HandlerType: (*DirektivIngressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNamespace",
			Handler:    _DirektivIngress_AddNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _DirektivIngress_DeleteNamespace_Handler,
		},
		{
			MethodName: "GetNamespaces",
			Handler:    _DirektivIngress_GetNamespaces_Handler,
		},
		{
			MethodName: "AddWorkflow",
			Handler:    _DirektivIngress_AddWorkflow_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _DirektivIngress_DeleteWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflowByName",
			Handler:    _DirektivIngress_GetWorkflowByName_Handler,
		},
		{
			MethodName: "GetWorkflowByUid",
			Handler:    _DirektivIngress_GetWorkflowByUid_Handler,
		},
		{
			MethodName: "GetWorkflowInstance",
			Handler:    _DirektivIngress_GetWorkflowInstance_Handler,
		},
		{
			MethodName: "GetWorkflowInstances",
			Handler:    _DirektivIngress_GetWorkflowInstances_Handler,
		},
		{
			MethodName: "GetInstancesByWorkflow",
			Handler:    _DirektivIngress_GetInstancesByWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflowInstanceLogs",
			Handler:    _DirektivIngress_GetWorkflowInstanceLogs_Handler,
		},
		{
			MethodName: "CancelWorkflowInstance",
			Handler:    _DirektivIngress_CancelWorkflowInstance_Handler,
		},
		{
			MethodName: "GetWorkflows",
			Handler:    _DirektivIngress_GetWorkflows_Handler,
		},
		{
			MethodName: "InvokeWorkflow",
			Handler:    _DirektivIngress_InvokeWorkflow_Handler,
		},
		{
			MethodName: "UpdateWorkflow",
			Handler:    _DirektivIngress_UpdateWorkflow_Handler,
		},
		{
			MethodName: "BroadcastEvent",
			Handler:    _DirektivIngress_BroadcastEvent_Handler,
		},
		{
			MethodName: "GetSecrets",
			Handler:    _DirektivIngress_GetSecrets_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _DirektivIngress_DeleteSecret_Handler,
		},
		{
			MethodName: "StoreSecret",
			Handler:    _DirektivIngress_StoreSecret_Handler,
		},
		{
			MethodName: "GetRegistries",
			Handler:    _DirektivIngress_GetRegistries_Handler,
		},
		{
			MethodName: "DeleteRegistry",
			Handler:    _DirektivIngress_DeleteRegistry_Handler,
		},
		{
			MethodName: "StoreRegistry",
			Handler:    _DirektivIngress_StoreRegistry_Handler,
		},
		{
			MethodName: "WorkflowMetrics",
			Handler:    _DirektivIngress_WorkflowMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/ingress/protocol.proto",
}
