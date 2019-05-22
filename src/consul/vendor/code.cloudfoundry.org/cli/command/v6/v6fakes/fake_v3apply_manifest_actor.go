// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3ApplyManifestActor struct {
	ApplyApplicationManifestStub        func(v3action.ManifestParser, string) (v3action.Warnings, error)
	applyApplicationManifestMutex       sync.RWMutex
	applyApplicationManifestArgsForCall []struct {
		arg1 v3action.ManifestParser
		arg2 string
	}
	applyApplicationManifestReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	applyApplicationManifestReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifest(arg1 v3action.ManifestParser, arg2 string) (v3action.Warnings, error) {
	fake.applyApplicationManifestMutex.Lock()
	ret, specificReturn := fake.applyApplicationManifestReturnsOnCall[len(fake.applyApplicationManifestArgsForCall)]
	fake.applyApplicationManifestArgsForCall = append(fake.applyApplicationManifestArgsForCall, struct {
		arg1 v3action.ManifestParser
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ApplyApplicationManifest", []interface{}{arg1, arg2})
	fake.applyApplicationManifestMutex.Unlock()
	if fake.ApplyApplicationManifestStub != nil {
		return fake.ApplyApplicationManifestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.applyApplicationManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifestCallCount() int {
	fake.applyApplicationManifestMutex.RLock()
	defer fake.applyApplicationManifestMutex.RUnlock()
	return len(fake.applyApplicationManifestArgsForCall)
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifestCalls(stub func(v3action.ManifestParser, string) (v3action.Warnings, error)) {
	fake.applyApplicationManifestMutex.Lock()
	defer fake.applyApplicationManifestMutex.Unlock()
	fake.ApplyApplicationManifestStub = stub
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifestArgsForCall(i int) (v3action.ManifestParser, string) {
	fake.applyApplicationManifestMutex.RLock()
	defer fake.applyApplicationManifestMutex.RUnlock()
	argsForCall := fake.applyApplicationManifestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifestReturns(result1 v3action.Warnings, result2 error) {
	fake.applyApplicationManifestMutex.Lock()
	defer fake.applyApplicationManifestMutex.Unlock()
	fake.ApplyApplicationManifestStub = nil
	fake.applyApplicationManifestReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3ApplyManifestActor) ApplyApplicationManifestReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.applyApplicationManifestMutex.Lock()
	defer fake.applyApplicationManifestMutex.Unlock()
	fake.ApplyApplicationManifestStub = nil
	if fake.applyApplicationManifestReturnsOnCall == nil {
		fake.applyApplicationManifestReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.applyApplicationManifestReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3ApplyManifestActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyApplicationManifestMutex.RLock()
	defer fake.applyApplicationManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3ApplyManifestActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3ApplyManifestActor = new(FakeV3ApplyManifestActor)
