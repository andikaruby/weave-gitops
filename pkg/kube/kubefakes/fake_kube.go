// Code generated by counterfeiter. DO NOT EDIT.
package kubefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/kube"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeKube struct {
	ApplyStub        func(context.Context, []byte, string) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
		arg3 string
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, []byte) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteByNameStub        func(context.Context, string, schema.GroupVersionResource, string) error
	deleteByNameMutex       sync.RWMutex
	deleteByNameArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 schema.GroupVersionResource
		arg4 string
	}
	deleteByNameReturns struct {
		result1 error
	}
	deleteByNameReturnsOnCall map[int]struct {
		result1 error
	}
	FetchNamespaceWithLabelStub        func(context.Context, string, string) (*v1.Namespace, error)
	fetchNamespaceWithLabelMutex       sync.RWMutex
	fetchNamespaceWithLabelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	fetchNamespaceWithLabelReturns struct {
		result1 *v1.Namespace
		result2 error
	}
	fetchNamespaceWithLabelReturnsOnCall map[int]struct {
		result1 *v1.Namespace
		result2 error
	}
	FluxPresentStub        func(context.Context) (bool, error)
	fluxPresentMutex       sync.RWMutex
	fluxPresentArgsForCall []struct {
		arg1 context.Context
	}
	fluxPresentReturns struct {
		result1 bool
		result2 error
	}
	fluxPresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetClusterNameStub        func(context.Context) (string, error)
	getClusterNameMutex       sync.RWMutex
	getClusterNameArgsForCall []struct {
		arg1 context.Context
	}
	getClusterNameReturns struct {
		result1 string
		result2 error
	}
	getClusterNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetClusterStatusStub        func(context.Context) kube.ClusterStatus
	getClusterStatusMutex       sync.RWMutex
	getClusterStatusArgsForCall []struct {
		arg1 context.Context
	}
	getClusterStatusReturns struct {
		result1 kube.ClusterStatus
	}
	getClusterStatusReturnsOnCall map[int]struct {
		result1 kube.ClusterStatus
	}
	GetResourceStub        func(context.Context, types.NamespacedName, kube.Resource) error
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 kube.Resource
	}
	getResourceReturns struct {
		result1 error
	}
	getResourceReturnsOnCall map[int]struct {
		result1 error
	}
	GetSecretStub        func(context.Context, types.NamespacedName) (*v1.Secret, error)
	getSecretMutex       sync.RWMutex
	getSecretArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}
	getSecretReturns struct {
		result1 *v1.Secret
		result2 error
	}
	getSecretReturnsOnCall map[int]struct {
		result1 *v1.Secret
		result2 error
	}
	GetWegoConfigStub        func(context.Context, string) (*kube.WegoConfig, error)
	getWegoConfigMutex       sync.RWMutex
	getWegoConfigArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getWegoConfigReturns struct {
		result1 *kube.WegoConfig
		result2 error
	}
	getWegoConfigReturnsOnCall map[int]struct {
		result1 *kube.WegoConfig
		result2 error
	}
	NamespacePresentStub        func(context.Context, string) (bool, error)
	namespacePresentMutex       sync.RWMutex
	namespacePresentArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	namespacePresentReturns struct {
		result1 bool
		result2 error
	}
	namespacePresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RawStub        func() client.Client
	rawMutex       sync.RWMutex
	rawArgsForCall []struct {
	}
	rawReturns struct {
		result1 client.Client
	}
	rawReturnsOnCall map[int]struct {
		result1 client.Client
	}
	SecretPresentStub        func(context.Context, string, string) (bool, error)
	secretPresentMutex       sync.RWMutex
	secretPresentArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	secretPresentReturns struct {
		result1 bool
		result2 error
	}
	secretPresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetResourceStub        func(context.Context, kube.Resource) error
	setResourceMutex       sync.RWMutex
	setResourceArgsForCall []struct {
		arg1 context.Context
		arg2 kube.Resource
	}
	setResourceReturns struct {
		result1 error
	}
	setResourceReturnsOnCall map[int]struct {
		result1 error
	}
	SetWegoConfigStub        func(context.Context, kube.WegoConfig, string) (*v1.ConfigMap, error)
	setWegoConfigMutex       sync.RWMutex
	setWegoConfigArgsForCall []struct {
		arg1 context.Context
		arg2 kube.WegoConfig
		arg3 string
	}
	setWegoConfigReturns struct {
		result1 *v1.ConfigMap
		result2 error
	}
	setWegoConfigReturnsOnCall map[int]struct {
		result1 *v1.ConfigMap
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKube) Apply(arg1 context.Context, arg2 []byte, arg3 string) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
		arg3 string
	}{arg1, arg2Copy, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2Copy, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeKube) ApplyCalls(stub func(context.Context, []byte, string) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeKube) ApplyArgsForCall(i int) (context.Context, []byte, string) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) Delete(arg1 context.Context, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2Copy})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeKube) DeleteCalls(stub func(context.Context, []byte) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeKube) DeleteArgsForCall(i int) (context.Context, []byte) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) DeleteByName(arg1 context.Context, arg2 string, arg3 schema.GroupVersionResource, arg4 string) error {
	fake.deleteByNameMutex.Lock()
	ret, specificReturn := fake.deleteByNameReturnsOnCall[len(fake.deleteByNameArgsForCall)]
	fake.deleteByNameArgsForCall = append(fake.deleteByNameArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 schema.GroupVersionResource
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.DeleteByNameStub
	fakeReturns := fake.deleteByNameReturns
	fake.recordInvocation("DeleteByName", []interface{}{arg1, arg2, arg3, arg4})
	fake.deleteByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) DeleteByNameCallCount() int {
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	return len(fake.deleteByNameArgsForCall)
}

func (fake *FakeKube) DeleteByNameCalls(stub func(context.Context, string, schema.GroupVersionResource, string) error) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = stub
}

func (fake *FakeKube) DeleteByNameArgsForCall(i int) (context.Context, string, schema.GroupVersionResource, string) {
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	argsForCall := fake.deleteByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeKube) DeleteByNameReturns(result1 error) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = nil
	fake.deleteByNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) DeleteByNameReturnsOnCall(i int, result1 error) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = nil
	if fake.deleteByNameReturnsOnCall == nil {
		fake.deleteByNameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteByNameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) FetchNamespaceWithLabel(arg1 context.Context, arg2 string, arg3 string) (*v1.Namespace, error) {
	fake.fetchNamespaceWithLabelMutex.Lock()
	ret, specificReturn := fake.fetchNamespaceWithLabelReturnsOnCall[len(fake.fetchNamespaceWithLabelArgsForCall)]
	fake.fetchNamespaceWithLabelArgsForCall = append(fake.fetchNamespaceWithLabelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FetchNamespaceWithLabelStub
	fakeReturns := fake.fetchNamespaceWithLabelReturns
	fake.recordInvocation("FetchNamespaceWithLabel", []interface{}{arg1, arg2, arg3})
	fake.fetchNamespaceWithLabelMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) FetchNamespaceWithLabelCallCount() int {
	fake.fetchNamespaceWithLabelMutex.RLock()
	defer fake.fetchNamespaceWithLabelMutex.RUnlock()
	return len(fake.fetchNamespaceWithLabelArgsForCall)
}

func (fake *FakeKube) FetchNamespaceWithLabelCalls(stub func(context.Context, string, string) (*v1.Namespace, error)) {
	fake.fetchNamespaceWithLabelMutex.Lock()
	defer fake.fetchNamespaceWithLabelMutex.Unlock()
	fake.FetchNamespaceWithLabelStub = stub
}

func (fake *FakeKube) FetchNamespaceWithLabelArgsForCall(i int) (context.Context, string, string) {
	fake.fetchNamespaceWithLabelMutex.RLock()
	defer fake.fetchNamespaceWithLabelMutex.RUnlock()
	argsForCall := fake.fetchNamespaceWithLabelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) FetchNamespaceWithLabelReturns(result1 *v1.Namespace, result2 error) {
	fake.fetchNamespaceWithLabelMutex.Lock()
	defer fake.fetchNamespaceWithLabelMutex.Unlock()
	fake.FetchNamespaceWithLabelStub = nil
	fake.fetchNamespaceWithLabelReturns = struct {
		result1 *v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FetchNamespaceWithLabelReturnsOnCall(i int, result1 *v1.Namespace, result2 error) {
	fake.fetchNamespaceWithLabelMutex.Lock()
	defer fake.fetchNamespaceWithLabelMutex.Unlock()
	fake.FetchNamespaceWithLabelStub = nil
	if fake.fetchNamespaceWithLabelReturnsOnCall == nil {
		fake.fetchNamespaceWithLabelReturnsOnCall = make(map[int]struct {
			result1 *v1.Namespace
			result2 error
		})
	}
	fake.fetchNamespaceWithLabelReturnsOnCall[i] = struct {
		result1 *v1.Namespace
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresent(arg1 context.Context) (bool, error) {
	fake.fluxPresentMutex.Lock()
	ret, specificReturn := fake.fluxPresentReturnsOnCall[len(fake.fluxPresentArgsForCall)]
	fake.fluxPresentArgsForCall = append(fake.fluxPresentArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FluxPresentStub
	fakeReturns := fake.fluxPresentReturns
	fake.recordInvocation("FluxPresent", []interface{}{arg1})
	fake.fluxPresentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) FluxPresentCallCount() int {
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	return len(fake.fluxPresentArgsForCall)
}

func (fake *FakeKube) FluxPresentCalls(stub func(context.Context) (bool, error)) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = stub
}

func (fake *FakeKube) FluxPresentArgsForCall(i int) context.Context {
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	argsForCall := fake.fluxPresentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) FluxPresentReturns(result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	fake.fluxPresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	if fake.fluxPresentReturnsOnCall == nil {
		fake.fluxPresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fluxPresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterName(arg1 context.Context) (string, error) {
	fake.getClusterNameMutex.Lock()
	ret, specificReturn := fake.getClusterNameReturnsOnCall[len(fake.getClusterNameArgsForCall)]
	fake.getClusterNameArgsForCall = append(fake.getClusterNameArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetClusterNameStub
	fakeReturns := fake.getClusterNameReturns
	fake.recordInvocation("GetClusterName", []interface{}{arg1})
	fake.getClusterNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetClusterNameCallCount() int {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	return len(fake.getClusterNameArgsForCall)
}

func (fake *FakeKube) GetClusterNameCalls(stub func(context.Context) (string, error)) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = stub
}

func (fake *FakeKube) GetClusterNameArgsForCall(i int) context.Context {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	argsForCall := fake.getClusterNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) GetClusterNameReturns(result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	fake.getClusterNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	if fake.getClusterNameReturnsOnCall == nil {
		fake.getClusterNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getClusterNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterStatus(arg1 context.Context) kube.ClusterStatus {
	fake.getClusterStatusMutex.Lock()
	ret, specificReturn := fake.getClusterStatusReturnsOnCall[len(fake.getClusterStatusArgsForCall)]
	fake.getClusterStatusArgsForCall = append(fake.getClusterStatusArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetClusterStatusStub
	fakeReturns := fake.getClusterStatusReturns
	fake.recordInvocation("GetClusterStatus", []interface{}{arg1})
	fake.getClusterStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) GetClusterStatusCallCount() int {
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	return len(fake.getClusterStatusArgsForCall)
}

func (fake *FakeKube) GetClusterStatusCalls(stub func(context.Context) kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = stub
}

func (fake *FakeKube) GetClusterStatusArgsForCall(i int) context.Context {
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	argsForCall := fake.getClusterStatusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) GetClusterStatusReturns(result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	fake.getClusterStatusReturns = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) GetClusterStatusReturnsOnCall(i int, result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	if fake.getClusterStatusReturnsOnCall == nil {
		fake.getClusterStatusReturnsOnCall = make(map[int]struct {
			result1 kube.ClusterStatus
		})
	}
	fake.getClusterStatusReturnsOnCall[i] = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) GetResource(arg1 context.Context, arg2 types.NamespacedName, arg3 kube.Resource) error {
	fake.getResourceMutex.Lock()
	ret, specificReturn := fake.getResourceReturnsOnCall[len(fake.getResourceArgsForCall)]
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 kube.Resource
	}{arg1, arg2, arg3})
	stub := fake.GetResourceStub
	fakeReturns := fake.getResourceReturns
	fake.recordInvocation("GetResource", []interface{}{arg1, arg2, arg3})
	fake.getResourceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeKube) GetResourceCalls(stub func(context.Context, types.NamespacedName, kube.Resource) error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = stub
}

func (fake *FakeKube) GetResourceArgsForCall(i int) (context.Context, types.NamespacedName, kube.Resource) {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	argsForCall := fake.getResourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) GetResourceReturns(result1 error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) GetResourceReturnsOnCall(i int, result1 error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = nil
	if fake.getResourceReturnsOnCall == nil {
		fake.getResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) GetSecret(arg1 context.Context, arg2 types.NamespacedName) (*v1.Secret, error) {
	fake.getSecretMutex.Lock()
	ret, specificReturn := fake.getSecretReturnsOnCall[len(fake.getSecretArgsForCall)]
	fake.getSecretArgsForCall = append(fake.getSecretArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}{arg1, arg2})
	stub := fake.GetSecretStub
	fakeReturns := fake.getSecretReturns
	fake.recordInvocation("GetSecret", []interface{}{arg1, arg2})
	fake.getSecretMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetSecretCallCount() int {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	return len(fake.getSecretArgsForCall)
}

func (fake *FakeKube) GetSecretCalls(stub func(context.Context, types.NamespacedName) (*v1.Secret, error)) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = stub
}

func (fake *FakeKube) GetSecretArgsForCall(i int) (context.Context, types.NamespacedName) {
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	argsForCall := fake.getSecretArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) GetSecretReturns(result1 *v1.Secret, result2 error) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	fake.getSecretReturns = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetSecretReturnsOnCall(i int, result1 *v1.Secret, result2 error) {
	fake.getSecretMutex.Lock()
	defer fake.getSecretMutex.Unlock()
	fake.GetSecretStub = nil
	if fake.getSecretReturnsOnCall == nil {
		fake.getSecretReturnsOnCall = make(map[int]struct {
			result1 *v1.Secret
			result2 error
		})
	}
	fake.getSecretReturnsOnCall[i] = struct {
		result1 *v1.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetWegoConfig(arg1 context.Context, arg2 string) (*kube.WegoConfig, error) {
	fake.getWegoConfigMutex.Lock()
	ret, specificReturn := fake.getWegoConfigReturnsOnCall[len(fake.getWegoConfigArgsForCall)]
	fake.getWegoConfigArgsForCall = append(fake.getWegoConfigArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetWegoConfigStub
	fakeReturns := fake.getWegoConfigReturns
	fake.recordInvocation("GetWegoConfig", []interface{}{arg1, arg2})
	fake.getWegoConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetWegoConfigCallCount() int {
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	return len(fake.getWegoConfigArgsForCall)
}

func (fake *FakeKube) GetWegoConfigCalls(stub func(context.Context, string) (*kube.WegoConfig, error)) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = stub
}

func (fake *FakeKube) GetWegoConfigArgsForCall(i int) (context.Context, string) {
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	argsForCall := fake.getWegoConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) GetWegoConfigReturns(result1 *kube.WegoConfig, result2 error) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = nil
	fake.getWegoConfigReturns = struct {
		result1 *kube.WegoConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetWegoConfigReturnsOnCall(i int, result1 *kube.WegoConfig, result2 error) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = nil
	if fake.getWegoConfigReturnsOnCall == nil {
		fake.getWegoConfigReturnsOnCall = make(map[int]struct {
			result1 *kube.WegoConfig
			result2 error
		})
	}
	fake.getWegoConfigReturnsOnCall[i] = struct {
		result1 *kube.WegoConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) NamespacePresent(arg1 context.Context, arg2 string) (bool, error) {
	fake.namespacePresentMutex.Lock()
	ret, specificReturn := fake.namespacePresentReturnsOnCall[len(fake.namespacePresentArgsForCall)]
	fake.namespacePresentArgsForCall = append(fake.namespacePresentArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.NamespacePresentStub
	fakeReturns := fake.namespacePresentReturns
	fake.recordInvocation("NamespacePresent", []interface{}{arg1, arg2})
	fake.namespacePresentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) NamespacePresentCallCount() int {
	fake.namespacePresentMutex.RLock()
	defer fake.namespacePresentMutex.RUnlock()
	return len(fake.namespacePresentArgsForCall)
}

func (fake *FakeKube) NamespacePresentCalls(stub func(context.Context, string) (bool, error)) {
	fake.namespacePresentMutex.Lock()
	defer fake.namespacePresentMutex.Unlock()
	fake.NamespacePresentStub = stub
}

func (fake *FakeKube) NamespacePresentArgsForCall(i int) (context.Context, string) {
	fake.namespacePresentMutex.RLock()
	defer fake.namespacePresentMutex.RUnlock()
	argsForCall := fake.namespacePresentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) NamespacePresentReturns(result1 bool, result2 error) {
	fake.namespacePresentMutex.Lock()
	defer fake.namespacePresentMutex.Unlock()
	fake.NamespacePresentStub = nil
	fake.namespacePresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) NamespacePresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.namespacePresentMutex.Lock()
	defer fake.namespacePresentMutex.Unlock()
	fake.NamespacePresentStub = nil
	if fake.namespacePresentReturnsOnCall == nil {
		fake.namespacePresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.namespacePresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) Raw() client.Client {
	fake.rawMutex.Lock()
	ret, specificReturn := fake.rawReturnsOnCall[len(fake.rawArgsForCall)]
	fake.rawArgsForCall = append(fake.rawArgsForCall, struct {
	}{})
	stub := fake.RawStub
	fakeReturns := fake.rawReturns
	fake.recordInvocation("Raw", []interface{}{})
	fake.rawMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) RawCallCount() int {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return len(fake.rawArgsForCall)
}

func (fake *FakeKube) RawCalls(stub func() client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = stub
}

func (fake *FakeKube) RawReturns(result1 client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	fake.rawReturns = struct {
		result1 client.Client
	}{result1}
}

func (fake *FakeKube) RawReturnsOnCall(i int, result1 client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	if fake.rawReturnsOnCall == nil {
		fake.rawReturnsOnCall = make(map[int]struct {
			result1 client.Client
		})
	}
	fake.rawReturnsOnCall[i] = struct {
		result1 client.Client
	}{result1}
}

func (fake *FakeKube) SecretPresent(arg1 context.Context, arg2 string, arg3 string) (bool, error) {
	fake.secretPresentMutex.Lock()
	ret, specificReturn := fake.secretPresentReturnsOnCall[len(fake.secretPresentArgsForCall)]
	fake.secretPresentArgsForCall = append(fake.secretPresentArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SecretPresentStub
	fakeReturns := fake.secretPresentReturns
	fake.recordInvocation("SecretPresent", []interface{}{arg1, arg2, arg3})
	fake.secretPresentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) SecretPresentCallCount() int {
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	return len(fake.secretPresentArgsForCall)
}

func (fake *FakeKube) SecretPresentCalls(stub func(context.Context, string, string) (bool, error)) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = stub
}

func (fake *FakeKube) SecretPresentArgsForCall(i int) (context.Context, string, string) {
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	argsForCall := fake.secretPresentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) SecretPresentReturns(result1 bool, result2 error) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = nil
	fake.secretPresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) SecretPresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = nil
	if fake.secretPresentReturnsOnCall == nil {
		fake.secretPresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.secretPresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) SetResource(arg1 context.Context, arg2 kube.Resource) error {
	fake.setResourceMutex.Lock()
	ret, specificReturn := fake.setResourceReturnsOnCall[len(fake.setResourceArgsForCall)]
	fake.setResourceArgsForCall = append(fake.setResourceArgsForCall, struct {
		arg1 context.Context
		arg2 kube.Resource
	}{arg1, arg2})
	stub := fake.SetResourceStub
	fakeReturns := fake.setResourceReturns
	fake.recordInvocation("SetResource", []interface{}{arg1, arg2})
	fake.setResourceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) SetResourceCallCount() int {
	fake.setResourceMutex.RLock()
	defer fake.setResourceMutex.RUnlock()
	return len(fake.setResourceArgsForCall)
}

func (fake *FakeKube) SetResourceCalls(stub func(context.Context, kube.Resource) error) {
	fake.setResourceMutex.Lock()
	defer fake.setResourceMutex.Unlock()
	fake.SetResourceStub = stub
}

func (fake *FakeKube) SetResourceArgsForCall(i int) (context.Context, kube.Resource) {
	fake.setResourceMutex.RLock()
	defer fake.setResourceMutex.RUnlock()
	argsForCall := fake.setResourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) SetResourceReturns(result1 error) {
	fake.setResourceMutex.Lock()
	defer fake.setResourceMutex.Unlock()
	fake.SetResourceStub = nil
	fake.setResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) SetResourceReturnsOnCall(i int, result1 error) {
	fake.setResourceMutex.Lock()
	defer fake.setResourceMutex.Unlock()
	fake.SetResourceStub = nil
	if fake.setResourceReturnsOnCall == nil {
		fake.setResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) SetWegoConfig(arg1 context.Context, arg2 kube.WegoConfig, arg3 string) (*v1.ConfigMap, error) {
	fake.setWegoConfigMutex.Lock()
	ret, specificReturn := fake.setWegoConfigReturnsOnCall[len(fake.setWegoConfigArgsForCall)]
	fake.setWegoConfigArgsForCall = append(fake.setWegoConfigArgsForCall, struct {
		arg1 context.Context
		arg2 kube.WegoConfig
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SetWegoConfigStub
	fakeReturns := fake.setWegoConfigReturns
	fake.recordInvocation("SetWegoConfig", []interface{}{arg1, arg2, arg3})
	fake.setWegoConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) SetWegoConfigCallCount() int {
	fake.setWegoConfigMutex.RLock()
	defer fake.setWegoConfigMutex.RUnlock()
	return len(fake.setWegoConfigArgsForCall)
}

func (fake *FakeKube) SetWegoConfigCalls(stub func(context.Context, kube.WegoConfig, string) (*v1.ConfigMap, error)) {
	fake.setWegoConfigMutex.Lock()
	defer fake.setWegoConfigMutex.Unlock()
	fake.SetWegoConfigStub = stub
}

func (fake *FakeKube) SetWegoConfigArgsForCall(i int) (context.Context, kube.WegoConfig, string) {
	fake.setWegoConfigMutex.RLock()
	defer fake.setWegoConfigMutex.RUnlock()
	argsForCall := fake.setWegoConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) SetWegoConfigReturns(result1 *v1.ConfigMap, result2 error) {
	fake.setWegoConfigMutex.Lock()
	defer fake.setWegoConfigMutex.Unlock()
	fake.SetWegoConfigStub = nil
	fake.setWegoConfigReturns = struct {
		result1 *v1.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) SetWegoConfigReturnsOnCall(i int, result1 *v1.ConfigMap, result2 error) {
	fake.setWegoConfigMutex.Lock()
	defer fake.setWegoConfigMutex.Unlock()
	fake.SetWegoConfigStub = nil
	if fake.setWegoConfigReturnsOnCall == nil {
		fake.setWegoConfigReturnsOnCall = make(map[int]struct {
			result1 *v1.ConfigMap
			result2 error
		})
	}
	fake.setWegoConfigReturnsOnCall[i] = struct {
		result1 *v1.ConfigMap
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	fake.fetchNamespaceWithLabelMutex.RLock()
	defer fake.fetchNamespaceWithLabelMutex.RUnlock()
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	fake.getSecretMutex.RLock()
	defer fake.getSecretMutex.RUnlock()
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	fake.namespacePresentMutex.RLock()
	defer fake.namespacePresentMutex.RUnlock()
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	fake.setResourceMutex.RLock()
	defer fake.setResourceMutex.RUnlock()
	fake.setWegoConfigMutex.RLock()
	defer fake.setWegoConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKube) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ kube.Kube = new(FakeKube)
