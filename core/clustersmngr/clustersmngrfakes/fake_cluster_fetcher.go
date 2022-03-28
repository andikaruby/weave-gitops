// Code generated by counterfeiter. DO NOT EDIT.
package clustersmngrfakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/core/clustersmngr"
)

type FakeClusterFetcher struct {
	FetchStub        func(context.Context) ([]clustersmngr.Cluster, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 context.Context
	}
	fetchReturns struct {
		result1 []clustersmngr.Cluster
		result2 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 []clustersmngr.Cluster
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClusterFetcher) Fetch(arg1 context.Context) ([]clustersmngr.Cluster, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FetchStub
	fakeReturns := fake.fetchReturns
	fake.recordInvocation("Fetch", []interface{}{arg1})
	fake.fetchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClusterFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeClusterFetcher) FetchCalls(stub func(context.Context) ([]clustersmngr.Cluster, error)) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = stub
}

func (fake *FakeClusterFetcher) FetchArgsForCall(i int) context.Context {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	argsForCall := fake.fetchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClusterFetcher) FetchReturns(result1 []clustersmngr.Cluster, result2 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 []clustersmngr.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterFetcher) FetchReturnsOnCall(i int, result1 []clustersmngr.Cluster, result2 error) {
	fake.fetchMutex.Lock()
	defer fake.fetchMutex.Unlock()
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 []clustersmngr.Cluster
			result2 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 []clustersmngr.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeClusterFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClusterFetcher) recordInvocation(key string, args []interface{}) {
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

var _ clustersmngr.ClusterFetcher = new(FakeClusterFetcher)
