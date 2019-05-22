// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	sync "sync"

	api "code.cloudfoundry.org/cli/cf/api"
)

type FakeCurlRepository struct {
	RequestStub        func(string, string, string, string, bool) (string, string, error)
	requestMutex       sync.RWMutex
	requestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}
	requestReturns struct {
		result1 string
		result2 string
		result3 error
	}
	requestReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCurlRepository) Request(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool) (string, string, error) {
	fake.requestMutex.Lock()
	ret, specificReturn := fake.requestReturnsOnCall[len(fake.requestArgsForCall)]
	fake.requestArgsForCall = append(fake.requestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Request", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.requestMutex.Unlock()
	if fake.RequestStub != nil {
		return fake.RequestStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.requestReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCurlRepository) RequestCallCount() int {
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	return len(fake.requestArgsForCall)
}

func (fake *FakeCurlRepository) RequestCalls(stub func(string, string, string, string, bool) (string, string, error)) {
	fake.requestMutex.Lock()
	defer fake.requestMutex.Unlock()
	fake.RequestStub = stub
}

func (fake *FakeCurlRepository) RequestArgsForCall(i int) (string, string, string, string, bool) {
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	argsForCall := fake.requestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeCurlRepository) RequestReturns(result1 string, result2 string, result3 error) {
	fake.requestMutex.Lock()
	defer fake.requestMutex.Unlock()
	fake.RequestStub = nil
	fake.requestReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCurlRepository) RequestReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.requestMutex.Lock()
	defer fake.requestMutex.Unlock()
	fake.RequestStub = nil
	if fake.requestReturnsOnCall == nil {
		fake.requestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.requestReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCurlRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestMutex.RLock()
	defer fake.requestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCurlRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.CurlRepository = new(FakeCurlRepository)
