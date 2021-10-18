// Code generated by counterfeiter. DO NOT EDIT.
package fakegitprovider

import (
	"context"
	"sync"

	"github.com/fluxcd/go-git-providers/gitprovider"
)

type UserRepository struct {
	APIObjectStub        func() interface{}
	aPIObjectMutex       sync.RWMutex
	aPIObjectArgsForCall []struct {
	}
	aPIObjectReturns struct {
		result1 interface{}
	}
	aPIObjectReturnsOnCall map[int]struct {
		result1 interface{}
	}
	BranchesStub        func() gitprovider.BranchClient
	branchesMutex       sync.RWMutex
	branchesArgsForCall []struct {
	}
	branchesReturns struct {
		result1 gitprovider.BranchClient
	}
	branchesReturnsOnCall map[int]struct {
		result1 gitprovider.BranchClient
	}
	CommitsStub        func() gitprovider.CommitClient
	commitsMutex       sync.RWMutex
	commitsArgsForCall []struct {
	}
	commitsReturns struct {
		result1 gitprovider.CommitClient
	}
	commitsReturnsOnCall map[int]struct {
		result1 gitprovider.CommitClient
	}
	DeleteStub        func(context.Context) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeployKeysStub        func() gitprovider.DeployKeyClient
	deployKeysMutex       sync.RWMutex
	deployKeysArgsForCall []struct {
	}
	deployKeysReturns struct {
		result1 gitprovider.DeployKeyClient
	}
	deployKeysReturnsOnCall map[int]struct {
		result1 gitprovider.DeployKeyClient
	}
	GetStub        func() gitprovider.RepositoryInfo
	getMutex       sync.RWMutex
	getArgsForCall []struct {
	}
	getReturns struct {
		result1 gitprovider.RepositoryInfo
	}
	getReturnsOnCall map[int]struct {
		result1 gitprovider.RepositoryInfo
	}
	PullRequestsStub        func() gitprovider.PullRequestClient
	pullRequestsMutex       sync.RWMutex
	pullRequestsArgsForCall []struct {
	}
	pullRequestsReturns struct {
		result1 gitprovider.PullRequestClient
	}
	pullRequestsReturnsOnCall map[int]struct {
		result1 gitprovider.PullRequestClient
	}
	ReconcileStub        func(context.Context) (bool, error)
	reconcileMutex       sync.RWMutex
	reconcileArgsForCall []struct {
		arg1 context.Context
	}
	reconcileReturns struct {
		result1 bool
		result2 error
	}
	reconcileReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RepositoryStub        func() gitprovider.RepositoryRef
	repositoryMutex       sync.RWMutex
	repositoryArgsForCall []struct {
	}
	repositoryReturns struct {
		result1 gitprovider.RepositoryRef
	}
	repositoryReturnsOnCall map[int]struct {
		result1 gitprovider.RepositoryRef
	}
	SetStub        func(gitprovider.RepositoryInfo) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 gitprovider.RepositoryInfo
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(context.Context) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UserRepository) APIObject() interface{} {
	fake.aPIObjectMutex.Lock()
	ret, specificReturn := fake.aPIObjectReturnsOnCall[len(fake.aPIObjectArgsForCall)]
	fake.aPIObjectArgsForCall = append(fake.aPIObjectArgsForCall, struct {
	}{})
	stub := fake.APIObjectStub
	fakeReturns := fake.aPIObjectReturns
	fake.recordInvocation("APIObject", []interface{}{})
	fake.aPIObjectMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) APIObjectCallCount() int {
	fake.aPIObjectMutex.RLock()
	defer fake.aPIObjectMutex.RUnlock()
	return len(fake.aPIObjectArgsForCall)
}

func (fake *UserRepository) APIObjectCalls(stub func() interface{}) {
	fake.aPIObjectMutex.Lock()
	defer fake.aPIObjectMutex.Unlock()
	fake.APIObjectStub = stub
}

func (fake *UserRepository) APIObjectReturns(result1 interface{}) {
	fake.aPIObjectMutex.Lock()
	defer fake.aPIObjectMutex.Unlock()
	fake.APIObjectStub = nil
	fake.aPIObjectReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *UserRepository) APIObjectReturnsOnCall(i int, result1 interface{}) {
	fake.aPIObjectMutex.Lock()
	defer fake.aPIObjectMutex.Unlock()
	fake.APIObjectStub = nil
	if fake.aPIObjectReturnsOnCall == nil {
		fake.aPIObjectReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.aPIObjectReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *UserRepository) Branches() gitprovider.BranchClient {
	fake.branchesMutex.Lock()
	ret, specificReturn := fake.branchesReturnsOnCall[len(fake.branchesArgsForCall)]
	fake.branchesArgsForCall = append(fake.branchesArgsForCall, struct {
	}{})
	stub := fake.BranchesStub
	fakeReturns := fake.branchesReturns
	fake.recordInvocation("Branches", []interface{}{})
	fake.branchesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) BranchesCallCount() int {
	fake.branchesMutex.RLock()
	defer fake.branchesMutex.RUnlock()
	return len(fake.branchesArgsForCall)
}

func (fake *UserRepository) BranchesCalls(stub func() gitprovider.BranchClient) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = stub
}

func (fake *UserRepository) BranchesReturns(result1 gitprovider.BranchClient) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = nil
	fake.branchesReturns = struct {
		result1 gitprovider.BranchClient
	}{result1}
}

func (fake *UserRepository) BranchesReturnsOnCall(i int, result1 gitprovider.BranchClient) {
	fake.branchesMutex.Lock()
	defer fake.branchesMutex.Unlock()
	fake.BranchesStub = nil
	if fake.branchesReturnsOnCall == nil {
		fake.branchesReturnsOnCall = make(map[int]struct {
			result1 gitprovider.BranchClient
		})
	}
	fake.branchesReturnsOnCall[i] = struct {
		result1 gitprovider.BranchClient
	}{result1}
}

func (fake *UserRepository) Commits() gitprovider.CommitClient {
	fake.commitsMutex.Lock()
	ret, specificReturn := fake.commitsReturnsOnCall[len(fake.commitsArgsForCall)]
	fake.commitsArgsForCall = append(fake.commitsArgsForCall, struct {
	}{})
	stub := fake.CommitsStub
	fakeReturns := fake.commitsReturns
	fake.recordInvocation("Commits", []interface{}{})
	fake.commitsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) CommitsCallCount() int {
	fake.commitsMutex.RLock()
	defer fake.commitsMutex.RUnlock()
	return len(fake.commitsArgsForCall)
}

func (fake *UserRepository) CommitsCalls(stub func() gitprovider.CommitClient) {
	fake.commitsMutex.Lock()
	defer fake.commitsMutex.Unlock()
	fake.CommitsStub = stub
}

func (fake *UserRepository) CommitsReturns(result1 gitprovider.CommitClient) {
	fake.commitsMutex.Lock()
	defer fake.commitsMutex.Unlock()
	fake.CommitsStub = nil
	fake.commitsReturns = struct {
		result1 gitprovider.CommitClient
	}{result1}
}

func (fake *UserRepository) CommitsReturnsOnCall(i int, result1 gitprovider.CommitClient) {
	fake.commitsMutex.Lock()
	defer fake.commitsMutex.Unlock()
	fake.CommitsStub = nil
	if fake.commitsReturnsOnCall == nil {
		fake.commitsReturnsOnCall = make(map[int]struct {
			result1 gitprovider.CommitClient
		})
	}
	fake.commitsReturnsOnCall[i] = struct {
		result1 gitprovider.CommitClient
	}{result1}
}

func (fake *UserRepository) Delete(arg1 context.Context) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *UserRepository) DeleteCalls(stub func(context.Context) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *UserRepository) DeleteArgsForCall(i int) context.Context {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UserRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *UserRepository) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *UserRepository) DeployKeys() gitprovider.DeployKeyClient {
	fake.deployKeysMutex.Lock()
	ret, specificReturn := fake.deployKeysReturnsOnCall[len(fake.deployKeysArgsForCall)]
	fake.deployKeysArgsForCall = append(fake.deployKeysArgsForCall, struct {
	}{})
	stub := fake.DeployKeysStub
	fakeReturns := fake.deployKeysReturns
	fake.recordInvocation("DeployKeys", []interface{}{})
	fake.deployKeysMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) DeployKeysCallCount() int {
	fake.deployKeysMutex.RLock()
	defer fake.deployKeysMutex.RUnlock()
	return len(fake.deployKeysArgsForCall)
}

func (fake *UserRepository) DeployKeysCalls(stub func() gitprovider.DeployKeyClient) {
	fake.deployKeysMutex.Lock()
	defer fake.deployKeysMutex.Unlock()
	fake.DeployKeysStub = stub
}

func (fake *UserRepository) DeployKeysReturns(result1 gitprovider.DeployKeyClient) {
	fake.deployKeysMutex.Lock()
	defer fake.deployKeysMutex.Unlock()
	fake.DeployKeysStub = nil
	fake.deployKeysReturns = struct {
		result1 gitprovider.DeployKeyClient
	}{result1}
}

func (fake *UserRepository) DeployKeysReturnsOnCall(i int, result1 gitprovider.DeployKeyClient) {
	fake.deployKeysMutex.Lock()
	defer fake.deployKeysMutex.Unlock()
	fake.DeployKeysStub = nil
	if fake.deployKeysReturnsOnCall == nil {
		fake.deployKeysReturnsOnCall = make(map[int]struct {
			result1 gitprovider.DeployKeyClient
		})
	}
	fake.deployKeysReturnsOnCall[i] = struct {
		result1 gitprovider.DeployKeyClient
	}{result1}
}

func (fake *UserRepository) Get() gitprovider.RepositoryInfo {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
	}{})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *UserRepository) GetCalls(stub func() gitprovider.RepositoryInfo) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *UserRepository) GetReturns(result1 gitprovider.RepositoryInfo) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 gitprovider.RepositoryInfo
	}{result1}
}

func (fake *UserRepository) GetReturnsOnCall(i int, result1 gitprovider.RepositoryInfo) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 gitprovider.RepositoryInfo
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 gitprovider.RepositoryInfo
	}{result1}
}

func (fake *UserRepository) PullRequests() gitprovider.PullRequestClient {
	fake.pullRequestsMutex.Lock()
	ret, specificReturn := fake.pullRequestsReturnsOnCall[len(fake.pullRequestsArgsForCall)]
	fake.pullRequestsArgsForCall = append(fake.pullRequestsArgsForCall, struct {
	}{})
	stub := fake.PullRequestsStub
	fakeReturns := fake.pullRequestsReturns
	fake.recordInvocation("PullRequests", []interface{}{})
	fake.pullRequestsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) PullRequestsCallCount() int {
	fake.pullRequestsMutex.RLock()
	defer fake.pullRequestsMutex.RUnlock()
	return len(fake.pullRequestsArgsForCall)
}

func (fake *UserRepository) PullRequestsCalls(stub func() gitprovider.PullRequestClient) {
	fake.pullRequestsMutex.Lock()
	defer fake.pullRequestsMutex.Unlock()
	fake.PullRequestsStub = stub
}

func (fake *UserRepository) PullRequestsReturns(result1 gitprovider.PullRequestClient) {
	fake.pullRequestsMutex.Lock()
	defer fake.pullRequestsMutex.Unlock()
	fake.PullRequestsStub = nil
	fake.pullRequestsReturns = struct {
		result1 gitprovider.PullRequestClient
	}{result1}
}

func (fake *UserRepository) PullRequestsReturnsOnCall(i int, result1 gitprovider.PullRequestClient) {
	fake.pullRequestsMutex.Lock()
	defer fake.pullRequestsMutex.Unlock()
	fake.PullRequestsStub = nil
	if fake.pullRequestsReturnsOnCall == nil {
		fake.pullRequestsReturnsOnCall = make(map[int]struct {
			result1 gitprovider.PullRequestClient
		})
	}
	fake.pullRequestsReturnsOnCall[i] = struct {
		result1 gitprovider.PullRequestClient
	}{result1}
}

func (fake *UserRepository) Reconcile(arg1 context.Context) (bool, error) {
	fake.reconcileMutex.Lock()
	ret, specificReturn := fake.reconcileReturnsOnCall[len(fake.reconcileArgsForCall)]
	fake.reconcileArgsForCall = append(fake.reconcileArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ReconcileStub
	fakeReturns := fake.reconcileReturns
	fake.recordInvocation("Reconcile", []interface{}{arg1})
	fake.reconcileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UserRepository) ReconcileCallCount() int {
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	return len(fake.reconcileArgsForCall)
}

func (fake *UserRepository) ReconcileCalls(stub func(context.Context) (bool, error)) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = stub
}

func (fake *UserRepository) ReconcileArgsForCall(i int) context.Context {
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	argsForCall := fake.reconcileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UserRepository) ReconcileReturns(result1 bool, result2 error) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = nil
	fake.reconcileReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *UserRepository) ReconcileReturnsOnCall(i int, result1 bool, result2 error) {
	fake.reconcileMutex.Lock()
	defer fake.reconcileMutex.Unlock()
	fake.ReconcileStub = nil
	if fake.reconcileReturnsOnCall == nil {
		fake.reconcileReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reconcileReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *UserRepository) Repository() gitprovider.RepositoryRef {
	fake.repositoryMutex.Lock()
	ret, specificReturn := fake.repositoryReturnsOnCall[len(fake.repositoryArgsForCall)]
	fake.repositoryArgsForCall = append(fake.repositoryArgsForCall, struct {
	}{})
	stub := fake.RepositoryStub
	fakeReturns := fake.repositoryReturns
	fake.recordInvocation("Repository", []interface{}{})
	fake.repositoryMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) RepositoryCallCount() int {
	fake.repositoryMutex.RLock()
	defer fake.repositoryMutex.RUnlock()
	return len(fake.repositoryArgsForCall)
}

func (fake *UserRepository) RepositoryCalls(stub func() gitprovider.RepositoryRef) {
	fake.repositoryMutex.Lock()
	defer fake.repositoryMutex.Unlock()
	fake.RepositoryStub = stub
}

func (fake *UserRepository) RepositoryReturns(result1 gitprovider.RepositoryRef) {
	fake.repositoryMutex.Lock()
	defer fake.repositoryMutex.Unlock()
	fake.RepositoryStub = nil
	fake.repositoryReturns = struct {
		result1 gitprovider.RepositoryRef
	}{result1}
}

func (fake *UserRepository) RepositoryReturnsOnCall(i int, result1 gitprovider.RepositoryRef) {
	fake.repositoryMutex.Lock()
	defer fake.repositoryMutex.Unlock()
	fake.RepositoryStub = nil
	if fake.repositoryReturnsOnCall == nil {
		fake.repositoryReturnsOnCall = make(map[int]struct {
			result1 gitprovider.RepositoryRef
		})
	}
	fake.repositoryReturnsOnCall[i] = struct {
		result1 gitprovider.RepositoryRef
	}{result1}
}

func (fake *UserRepository) Set(arg1 gitprovider.RepositoryInfo) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 gitprovider.RepositoryInfo
	}{arg1})
	stub := fake.SetStub
	fakeReturns := fake.setReturns
	fake.recordInvocation("Set", []interface{}{arg1})
	fake.setMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *UserRepository) SetCalls(stub func(gitprovider.RepositoryInfo) error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *UserRepository) SetArgsForCall(i int) gitprovider.RepositoryInfo {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UserRepository) SetReturns(result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *UserRepository) SetReturnsOnCall(i int, result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *UserRepository) Update(arg1 context.Context) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *UserRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *UserRepository) UpdateCalls(stub func(context.Context) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *UserRepository) UpdateArgsForCall(i int) context.Context {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UserRepository) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *UserRepository) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *UserRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIObjectMutex.RLock()
	defer fake.aPIObjectMutex.RUnlock()
	fake.branchesMutex.RLock()
	defer fake.branchesMutex.RUnlock()
	fake.commitsMutex.RLock()
	defer fake.commitsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deployKeysMutex.RLock()
	defer fake.deployKeysMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.pullRequestsMutex.RLock()
	defer fake.pullRequestsMutex.RUnlock()
	fake.reconcileMutex.RLock()
	defer fake.reconcileMutex.RUnlock()
	fake.repositoryMutex.RLock()
	defer fake.repositoryMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *UserRepository) recordInvocation(key string, args []interface{}) {
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

var _ gitprovider.UserRepository = new(UserRepository)
