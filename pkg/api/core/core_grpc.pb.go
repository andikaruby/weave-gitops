// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	// GetObject gets data about a single primary object from a cluster.
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	// ListObjects gets data about primary objects.
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
	// ListFluxRuntimeObjects lists the flux runtime deployments from a cluster.
	ListFluxRuntimeObjects(ctx context.Context, in *ListFluxRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListFluxRuntimeObjectsResponse, error)
	ListFluxCrds(ctx context.Context, in *ListFluxCrdsRequest, opts ...grpc.CallOption) (*ListFluxCrdsResponse, error)
	// ListRuntimeObjects lists Weave GitOps runtime components from a cluster. Weave GitOps runtime is composed of Flux runtime
	// but also the other components in the ecosystem like TF-controller or Policy Agent.
	ListRuntimeObjects(ctx context.Context, in *ListRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListRuntimeObjectsResponse, error)
	// ListRuntimeCrds lists Weave GitOps runtime components CRDs from a cluster. Weave GitOps runtime is composed of Flux runtime
	// but also the other components in the ecosystem like TF-controller or Policy Agent.
	ListRuntimeCrds(ctx context.Context, in *ListRuntimeCrdsRequest, opts ...grpc.CallOption) (*ListRuntimeCrdsResponse, error)
	// GetReconciledObjects returns a list of objects that were created
	// as a result of reconciling a Flux automation.
	// This list is derived by looking at the Kustomization or HelmRelease
	// specified in the request body.
	GetReconciledObjects(ctx context.Context, in *GetReconciledObjectsRequest, opts ...grpc.CallOption) (*GetReconciledObjectsResponse, error)
	// GetChildObjects returns the children of a given object,
	// specified by a GroupVersionKind.
	// Not all Kubernets objects have children. For example, a Deployment
	// has a child ReplicaSet, but a Service has no child objects.
	GetChildObjects(ctx context.Context, in *GetChildObjectsRequest, opts ...grpc.CallOption) (*GetChildObjectsResponse, error)
	// GetFluxNamespace returns with a namespace with a specific label.
	GetFluxNamespace(ctx context.Context, in *GetFluxNamespaceRequest, opts ...grpc.CallOption) (*GetFluxNamespaceResponse, error)
	// ListNamespaces returns with the list of available namespaces.
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	// ListEvents returns with a list of events
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	// SyncResource forces a reconciliation of a Flux resource
	SyncFluxObject(ctx context.Context, in *SyncFluxObjectRequest, opts ...grpc.CallOption) (*SyncFluxObjectResponse, error)
	// GetVersion returns version information about the server
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
	// GetFeatureFlags returns configuration information about the server
	GetFeatureFlags(ctx context.Context, in *GetFeatureFlagsRequest, opts ...grpc.CallOption) (*GetFeatureFlagsResponse, error)
	// ToggleSuspendResource suspends or resumes a flux object.
	ToggleSuspendResource(ctx context.Context, in *ToggleSuspendResourceRequest, opts ...grpc.CallOption) (*ToggleSuspendResourceResponse, error)
	// GetSessionLogs returns the logs for a given session
	GetSessionLogs(ctx context.Context, in *GetSessionLogsRequest, opts ...grpc.CallOption) (*GetSessionLogsResponse, error)
	// IsCRDAvailable returns with a hashmap where the keys are the names of
	// the clusters, and the value is a boolean indicating whether given CRD is
	// installed or not on that cluster.
	IsCRDAvailable(ctx context.Context, in *IsCRDAvailableRequest, opts ...grpc.CallOption) (*IsCRDAvailableResponse, error)
	GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryResponse, error)
	// ListPolicies list policies available on the cluster
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
	// GetPolicy gets a policy by name
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error)
	// ListPolicyValidations lists policy validations
	ListPolicyValidations(ctx context.Context, in *ListPolicyValidationsRequest, opts ...grpc.CallOption) (*ListPolicyValidationsResponse, error)
	// GetPolicyValidation gets a policy validation by id
	GetPolicyValidation(ctx context.Context, in *GetPolicyValidationRequest, opts ...grpc.CallOption) (*GetPolicyValidationResponse, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListFluxRuntimeObjects(ctx context.Context, in *ListFluxRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListFluxRuntimeObjectsResponse, error) {
	out := new(ListFluxRuntimeObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListFluxRuntimeObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListFluxCrds(ctx context.Context, in *ListFluxCrdsRequest, opts ...grpc.CallOption) (*ListFluxCrdsResponse, error) {
	out := new(ListFluxCrdsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListFluxCrds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListRuntimeObjects(ctx context.Context, in *ListRuntimeObjectsRequest, opts ...grpc.CallOption) (*ListRuntimeObjectsResponse, error) {
	out := new(ListRuntimeObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListRuntimeObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListRuntimeCrds(ctx context.Context, in *ListRuntimeCrdsRequest, opts ...grpc.CallOption) (*ListRuntimeCrdsResponse, error) {
	out := new(ListRuntimeCrdsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListRuntimeCrds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetReconciledObjects(ctx context.Context, in *GetReconciledObjectsRequest, opts ...grpc.CallOption) (*GetReconciledObjectsResponse, error) {
	out := new(GetReconciledObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetReconciledObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetChildObjects(ctx context.Context, in *GetChildObjectsRequest, opts ...grpc.CallOption) (*GetChildObjectsResponse, error) {
	out := new(GetChildObjectsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetChildObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetFluxNamespace(ctx context.Context, in *GetFluxNamespaceRequest, opts ...grpc.CallOption) (*GetFluxNamespaceResponse, error) {
	out := new(GetFluxNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetFluxNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) SyncFluxObject(ctx context.Context, in *SyncFluxObjectRequest, opts ...grpc.CallOption) (*SyncFluxObjectResponse, error) {
	out := new(SyncFluxObjectResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/SyncFluxObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetFeatureFlags(ctx context.Context, in *GetFeatureFlagsRequest, opts ...grpc.CallOption) (*GetFeatureFlagsResponse, error) {
	out := new(GetFeatureFlagsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetFeatureFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ToggleSuspendResource(ctx context.Context, in *ToggleSuspendResourceRequest, opts ...grpc.CallOption) (*ToggleSuspendResourceResponse, error) {
	out := new(ToggleSuspendResourceResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ToggleSuspendResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetSessionLogs(ctx context.Context, in *GetSessionLogsRequest, opts ...grpc.CallOption) (*GetSessionLogsResponse, error) {
	out := new(GetSessionLogsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetSessionLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) IsCRDAvailable(ctx context.Context, in *IsCRDAvailableRequest, opts ...grpc.CallOption) (*IsCRDAvailableResponse, error) {
	out := new(IsCRDAvailableResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/IsCRDAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetInventory(ctx context.Context, in *GetInventoryRequest, opts ...grpc.CallOption) (*GetInventoryResponse, error) {
	out := new(GetInventoryResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error) {
	out := new(GetPolicyResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListPolicyValidations(ctx context.Context, in *ListPolicyValidationsRequest, opts ...grpc.CallOption) (*ListPolicyValidationsResponse, error) {
	out := new(ListPolicyValidationsResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/ListPolicyValidations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) GetPolicyValidation(ctx context.Context, in *GetPolicyValidationRequest, opts ...grpc.CallOption) (*GetPolicyValidationResponse, error) {
	out := new(GetPolicyValidationResponse)
	err := c.cc.Invoke(ctx, "/gitops_core.v1.Core/GetPolicyValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
// All implementations must embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	// GetObject gets data about a single primary object from a cluster.
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	// ListObjects gets data about primary objects.
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
	// ListFluxRuntimeObjects lists the flux runtime deployments from a cluster.
	ListFluxRuntimeObjects(context.Context, *ListFluxRuntimeObjectsRequest) (*ListFluxRuntimeObjectsResponse, error)
	ListFluxCrds(context.Context, *ListFluxCrdsRequest) (*ListFluxCrdsResponse, error)
	// ListRuntimeObjects lists Weave GitOps runtime components from a cluster. Weave GitOps runtime is composed of Flux runtime
	// but also the other components in the ecosystem like TF-controller or Policy Agent.
	ListRuntimeObjects(context.Context, *ListRuntimeObjectsRequest) (*ListRuntimeObjectsResponse, error)
	// ListRuntimeCrds lists Weave GitOps runtime components CRDs from a cluster. Weave GitOps runtime is composed of Flux runtime
	// but also the other components in the ecosystem like TF-controller or Policy Agent.
	ListRuntimeCrds(context.Context, *ListRuntimeCrdsRequest) (*ListRuntimeCrdsResponse, error)
	// GetReconciledObjects returns a list of objects that were created
	// as a result of reconciling a Flux automation.
	// This list is derived by looking at the Kustomization or HelmRelease
	// specified in the request body.
	GetReconciledObjects(context.Context, *GetReconciledObjectsRequest) (*GetReconciledObjectsResponse, error)
	// GetChildObjects returns the children of a given object,
	// specified by a GroupVersionKind.
	// Not all Kubernets objects have children. For example, a Deployment
	// has a child ReplicaSet, but a Service has no child objects.
	GetChildObjects(context.Context, *GetChildObjectsRequest) (*GetChildObjectsResponse, error)
	// GetFluxNamespace returns with a namespace with a specific label.
	GetFluxNamespace(context.Context, *GetFluxNamespaceRequest) (*GetFluxNamespaceResponse, error)
	// ListNamespaces returns with the list of available namespaces.
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	// ListEvents returns with a list of events
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	// SyncResource forces a reconciliation of a Flux resource
	SyncFluxObject(context.Context, *SyncFluxObjectRequest) (*SyncFluxObjectResponse, error)
	// GetVersion returns version information about the server
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	// GetFeatureFlags returns configuration information about the server
	GetFeatureFlags(context.Context, *GetFeatureFlagsRequest) (*GetFeatureFlagsResponse, error)
	// ToggleSuspendResource suspends or resumes a flux object.
	ToggleSuspendResource(context.Context, *ToggleSuspendResourceRequest) (*ToggleSuspendResourceResponse, error)
	// GetSessionLogs returns the logs for a given session
	GetSessionLogs(context.Context, *GetSessionLogsRequest) (*GetSessionLogsResponse, error)
	// IsCRDAvailable returns with a hashmap where the keys are the names of
	// the clusters, and the value is a boolean indicating whether given CRD is
	// installed or not on that cluster.
	IsCRDAvailable(context.Context, *IsCRDAvailableRequest) (*IsCRDAvailableResponse, error)
	GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryResponse, error)
	// ListPolicies list policies available on the cluster
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	// GetPolicy gets a policy by name
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	// ListPolicyValidations lists policy validations
	ListPolicyValidations(context.Context, *ListPolicyValidationsRequest) (*ListPolicyValidationsResponse, error)
	// GetPolicyValidation gets a policy validation by id
	GetPolicyValidation(context.Context, *GetPolicyValidationRequest) (*GetPolicyValidationResponse, error)
	mustEmbedUnimplementedCoreServer()
}

// UnimplementedCoreServer must be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedCoreServer) ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedCoreServer) ListFluxRuntimeObjects(context.Context, *ListFluxRuntimeObjectsRequest) (*ListFluxRuntimeObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFluxRuntimeObjects not implemented")
}
func (UnimplementedCoreServer) ListFluxCrds(context.Context, *ListFluxCrdsRequest) (*ListFluxCrdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFluxCrds not implemented")
}
func (UnimplementedCoreServer) ListRuntimeObjects(context.Context, *ListRuntimeObjectsRequest) (*ListRuntimeObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeObjects not implemented")
}
func (UnimplementedCoreServer) ListRuntimeCrds(context.Context, *ListRuntimeCrdsRequest) (*ListRuntimeCrdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuntimeCrds not implemented")
}
func (UnimplementedCoreServer) GetReconciledObjects(context.Context, *GetReconciledObjectsRequest) (*GetReconciledObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReconciledObjects not implemented")
}
func (UnimplementedCoreServer) GetChildObjects(context.Context, *GetChildObjectsRequest) (*GetChildObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildObjects not implemented")
}
func (UnimplementedCoreServer) GetFluxNamespace(context.Context, *GetFluxNamespaceRequest) (*GetFluxNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFluxNamespace not implemented")
}
func (UnimplementedCoreServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedCoreServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedCoreServer) SyncFluxObject(context.Context, *SyncFluxObjectRequest) (*SyncFluxObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncFluxObject not implemented")
}
func (UnimplementedCoreServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedCoreServer) GetFeatureFlags(context.Context, *GetFeatureFlagsRequest) (*GetFeatureFlagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureFlags not implemented")
}
func (UnimplementedCoreServer) ToggleSuspendResource(context.Context, *ToggleSuspendResourceRequest) (*ToggleSuspendResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleSuspendResource not implemented")
}
func (UnimplementedCoreServer) GetSessionLogs(context.Context, *GetSessionLogsRequest) (*GetSessionLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionLogs not implemented")
}
func (UnimplementedCoreServer) IsCRDAvailable(context.Context, *IsCRDAvailableRequest) (*IsCRDAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCRDAvailable not implemented")
}
func (UnimplementedCoreServer) GetInventory(context.Context, *GetInventoryRequest) (*GetInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventory not implemented")
}
func (UnimplementedCoreServer) ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedCoreServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedCoreServer) ListPolicyValidations(context.Context, *ListPolicyValidationsRequest) (*ListPolicyValidationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyValidations not implemented")
}
func (UnimplementedCoreServer) GetPolicyValidation(context.Context, *GetPolicyValidationRequest) (*GetPolicyValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyValidation not implemented")
}
func (UnimplementedCoreServer) mustEmbedUnimplementedCoreServer() {}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListFluxRuntimeObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFluxRuntimeObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListFluxRuntimeObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListFluxRuntimeObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListFluxRuntimeObjects(ctx, req.(*ListFluxRuntimeObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListFluxCrds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFluxCrdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListFluxCrds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListFluxCrds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListFluxCrds(ctx, req.(*ListFluxCrdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListRuntimeObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuntimeObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListRuntimeObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListRuntimeObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListRuntimeObjects(ctx, req.(*ListRuntimeObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListRuntimeCrds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuntimeCrdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListRuntimeCrds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListRuntimeCrds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListRuntimeCrds(ctx, req.(*ListRuntimeCrdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetReconciledObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReconciledObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetReconciledObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetReconciledObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetReconciledObjects(ctx, req.(*GetReconciledObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetChildObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChildObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetChildObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetChildObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetChildObjects(ctx, req.(*GetChildObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetFluxNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFluxNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetFluxNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetFluxNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetFluxNamespace(ctx, req.(*GetFluxNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_SyncFluxObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncFluxObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).SyncFluxObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/SyncFluxObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).SyncFluxObject(ctx, req.(*SyncFluxObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetFeatureFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeatureFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetFeatureFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetFeatureFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetFeatureFlags(ctx, req.(*GetFeatureFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ToggleSuspendResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleSuspendResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ToggleSuspendResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ToggleSuspendResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ToggleSuspendResource(ctx, req.(*ToggleSuspendResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetSessionLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetSessionLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetSessionLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetSessionLogs(ctx, req.(*GetSessionLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_IsCRDAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsCRDAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).IsCRDAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/IsCRDAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).IsCRDAvailable(ctx, req.(*IsCRDAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetInventory(ctx, req.(*GetInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListPolicyValidations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyValidationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ListPolicyValidations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/ListPolicyValidations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ListPolicyValidations(ctx, req.(*ListPolicyValidationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_GetPolicyValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).GetPolicyValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitops_core.v1.Core/GetPolicyValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).GetPolicyValidation(ctx, req.(*GetPolicyValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitops_core.v1.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObject",
			Handler:    _Core_GetObject_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _Core_ListObjects_Handler,
		},
		{
			MethodName: "ListFluxRuntimeObjects",
			Handler:    _Core_ListFluxRuntimeObjects_Handler,
		},
		{
			MethodName: "ListFluxCrds",
			Handler:    _Core_ListFluxCrds_Handler,
		},
		{
			MethodName: "ListRuntimeObjects",
			Handler:    _Core_ListRuntimeObjects_Handler,
		},
		{
			MethodName: "ListRuntimeCrds",
			Handler:    _Core_ListRuntimeCrds_Handler,
		},
		{
			MethodName: "GetReconciledObjects",
			Handler:    _Core_GetReconciledObjects_Handler,
		},
		{
			MethodName: "GetChildObjects",
			Handler:    _Core_GetChildObjects_Handler,
		},
		{
			MethodName: "GetFluxNamespace",
			Handler:    _Core_GetFluxNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _Core_ListNamespaces_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _Core_ListEvents_Handler,
		},
		{
			MethodName: "SyncFluxObject",
			Handler:    _Core_SyncFluxObject_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Core_GetVersion_Handler,
		},
		{
			MethodName: "GetFeatureFlags",
			Handler:    _Core_GetFeatureFlags_Handler,
		},
		{
			MethodName: "ToggleSuspendResource",
			Handler:    _Core_ToggleSuspendResource_Handler,
		},
		{
			MethodName: "GetSessionLogs",
			Handler:    _Core_GetSessionLogs_Handler,
		},
		{
			MethodName: "IsCRDAvailable",
			Handler:    _Core_IsCRDAvailable_Handler,
		},
		{
			MethodName: "GetInventory",
			Handler:    _Core_GetInventory_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _Core_ListPolicies_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _Core_GetPolicy_Handler,
		},
		{
			MethodName: "ListPolicyValidations",
			Handler:    _Core_ListPolicyValidations_Handler,
		},
		{
			MethodName: "GetPolicyValidation",
			Handler:    _Core_GetPolicyValidation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/core/core.proto",
}
