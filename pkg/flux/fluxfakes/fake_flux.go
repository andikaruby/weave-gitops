// Code generated by counterfeiter. DO NOT EDIT.
package fluxfakes

import (
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/flux"
)

type FakeFlux struct {
	GetAllResourcesStatusStub        func(string, string) ([]byte, error)
	getAllResourcesStatusMutex       sync.RWMutex
	getAllResourcesStatusArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getAllResourcesStatusReturns struct {
		result1 []byte
		result2 error
	}
	getAllResourcesStatusReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetLatestStatusAllNamespacesStub        func() ([]string, error)
	getLatestStatusAllNamespacesMutex       sync.RWMutex
	getLatestStatusAllNamespacesArgsForCall []struct {
	}
	getLatestStatusAllNamespacesReturns struct {
		result1 []string
		result2 error
	}
	getLatestStatusAllNamespacesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetVersionStub        func() (string, error)
	getVersionMutex       sync.RWMutex
	getVersionArgsForCall []struct {
	}
	getVersionReturns struct {
		result1 string
		result2 error
	}
	getVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	PreCheckStub        func() (string, error)
	preCheckMutex       sync.RWMutex
	preCheckArgsForCall []struct {
	}
	preCheckReturns struct {
		result1 string
		result2 error
	}
	preCheckReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFlux) GetAllResourcesStatus(arg1 string, arg2 string) ([]byte, error) {
	fake.getAllResourcesStatusMutex.Lock()
	ret, specificReturn := fake.getAllResourcesStatusReturnsOnCall[len(fake.getAllResourcesStatusArgsForCall)]
	fake.getAllResourcesStatusArgsForCall = append(fake.getAllResourcesStatusArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetAllResourcesStatusStub
	fakeReturns := fake.getAllResourcesStatusReturns
	fake.recordInvocation("GetAllResourcesStatus", []interface{}{arg1, arg2})
	fake.getAllResourcesStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) GetAllResourcesStatusCallCount() int {
	fake.getAllResourcesStatusMutex.RLock()
	defer fake.getAllResourcesStatusMutex.RUnlock()
	return len(fake.getAllResourcesStatusArgsForCall)
}

func (fake *FakeFlux) GetAllResourcesStatusCalls(stub func(string, string) ([]byte, error)) {
	fake.getAllResourcesStatusMutex.Lock()
	defer fake.getAllResourcesStatusMutex.Unlock()
	fake.GetAllResourcesStatusStub = stub
}

func (fake *FakeFlux) GetAllResourcesStatusArgsForCall(i int) (string, string) {
	fake.getAllResourcesStatusMutex.RLock()
	defer fake.getAllResourcesStatusMutex.RUnlock()
	argsForCall := fake.getAllResourcesStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFlux) GetAllResourcesStatusReturns(result1 []byte, result2 error) {
	fake.getAllResourcesStatusMutex.Lock()
	defer fake.getAllResourcesStatusMutex.Unlock()
	fake.GetAllResourcesStatusStub = nil
	fake.getAllResourcesStatusReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) GetAllResourcesStatusReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getAllResourcesStatusMutex.Lock()
	defer fake.getAllResourcesStatusMutex.Unlock()
	fake.GetAllResourcesStatusStub = nil
	if fake.getAllResourcesStatusReturnsOnCall == nil {
		fake.getAllResourcesStatusReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getAllResourcesStatusReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) GetLatestStatusAllNamespaces() ([]string, error) {
	fake.getLatestStatusAllNamespacesMutex.Lock()
	ret, specificReturn := fake.getLatestStatusAllNamespacesReturnsOnCall[len(fake.getLatestStatusAllNamespacesArgsForCall)]
	fake.getLatestStatusAllNamespacesArgsForCall = append(fake.getLatestStatusAllNamespacesArgsForCall, struct {
	}{})
	stub := fake.GetLatestStatusAllNamespacesStub
	fakeReturns := fake.getLatestStatusAllNamespacesReturns
	fake.recordInvocation("GetLatestStatusAllNamespaces", []interface{}{})
	fake.getLatestStatusAllNamespacesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) GetLatestStatusAllNamespacesCallCount() int {
	fake.getLatestStatusAllNamespacesMutex.RLock()
	defer fake.getLatestStatusAllNamespacesMutex.RUnlock()
	return len(fake.getLatestStatusAllNamespacesArgsForCall)
}

func (fake *FakeFlux) GetLatestStatusAllNamespacesCalls(stub func() ([]string, error)) {
	fake.getLatestStatusAllNamespacesMutex.Lock()
	defer fake.getLatestStatusAllNamespacesMutex.Unlock()
	fake.GetLatestStatusAllNamespacesStub = stub
}

func (fake *FakeFlux) GetLatestStatusAllNamespacesReturns(result1 []string, result2 error) {
	fake.getLatestStatusAllNamespacesMutex.Lock()
	defer fake.getLatestStatusAllNamespacesMutex.Unlock()
	fake.GetLatestStatusAllNamespacesStub = nil
	fake.getLatestStatusAllNamespacesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) GetLatestStatusAllNamespacesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getLatestStatusAllNamespacesMutex.Lock()
	defer fake.getLatestStatusAllNamespacesMutex.Unlock()
	fake.GetLatestStatusAllNamespacesStub = nil
	if fake.getLatestStatusAllNamespacesReturnsOnCall == nil {
		fake.getLatestStatusAllNamespacesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getLatestStatusAllNamespacesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) GetVersion() (string, error) {
	fake.getVersionMutex.Lock()
	ret, specificReturn := fake.getVersionReturnsOnCall[len(fake.getVersionArgsForCall)]
	fake.getVersionArgsForCall = append(fake.getVersionArgsForCall, struct {
	}{})
	stub := fake.GetVersionStub
	fakeReturns := fake.getVersionReturns
	fake.recordInvocation("GetVersion", []interface{}{})
	fake.getVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) GetVersionCallCount() int {
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	return len(fake.getVersionArgsForCall)
}

func (fake *FakeFlux) GetVersionCalls(stub func() (string, error)) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = stub
}

func (fake *FakeFlux) GetVersionReturns(result1 string, result2 error) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	fake.getVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) GetVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	if fake.getVersionReturnsOnCall == nil {
		fake.getVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) PreCheck() (string, error) {
	fake.preCheckMutex.Lock()
	ret, specificReturn := fake.preCheckReturnsOnCall[len(fake.preCheckArgsForCall)]
	fake.preCheckArgsForCall = append(fake.preCheckArgsForCall, struct {
	}{})
	stub := fake.PreCheckStub
	fakeReturns := fake.preCheckReturns
	fake.recordInvocation("PreCheck", []interface{}{})
	fake.preCheckMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlux) PreCheckCallCount() int {
	fake.preCheckMutex.RLock()
	defer fake.preCheckMutex.RUnlock()
	return len(fake.preCheckArgsForCall)
}

func (fake *FakeFlux) PreCheckCalls(stub func() (string, error)) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = stub
}

func (fake *FakeFlux) PreCheckReturns(result1 string, result2 error) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = nil
	fake.preCheckReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) PreCheckReturnsOnCall(i int, result1 string, result2 error) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = nil
	if fake.preCheckReturnsOnCall == nil {
		fake.preCheckReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.preCheckReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFlux) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAllResourcesStatusMutex.RLock()
	defer fake.getAllResourcesStatusMutex.RUnlock()
	fake.getLatestStatusAllNamespacesMutex.RLock()
	defer fake.getLatestStatusAllNamespacesMutex.RUnlock()
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	fake.preCheckMutex.RLock()
	defer fake.preCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFlux) recordInvocation(key string, args []interface{}) {
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

var _ flux.Flux = new(FakeFlux)
