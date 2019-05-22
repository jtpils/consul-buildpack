// Code generated by counterfeiter. DO NOT EDIT.
package ccv2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
)

type FakeConnectionWrapper struct {
	MakeStub        func(*cloudcontroller.Request, *cloudcontroller.Response) error
	makeMutex       sync.RWMutex
	makeArgsForCall []struct {
		arg1 *cloudcontroller.Request
		arg2 *cloudcontroller.Response
	}
	makeReturns struct {
		result1 error
	}
	makeReturnsOnCall map[int]struct {
		result1 error
	}
	WrapStub        func(cloudcontroller.Connection) cloudcontroller.Connection
	wrapMutex       sync.RWMutex
	wrapArgsForCall []struct {
		arg1 cloudcontroller.Connection
	}
	wrapReturns struct {
		result1 cloudcontroller.Connection
	}
	wrapReturnsOnCall map[int]struct {
		result1 cloudcontroller.Connection
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnectionWrapper) Make(arg1 *cloudcontroller.Request, arg2 *cloudcontroller.Response) error {
	fake.makeMutex.Lock()
	ret, specificReturn := fake.makeReturnsOnCall[len(fake.makeArgsForCall)]
	fake.makeArgsForCall = append(fake.makeArgsForCall, struct {
		arg1 *cloudcontroller.Request
		arg2 *cloudcontroller.Response
	}{arg1, arg2})
	fake.recordInvocation("Make", []interface{}{arg1, arg2})
	fake.makeMutex.Unlock()
	if fake.MakeStub != nil {
		return fake.MakeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.makeReturns
	return fakeReturns.result1
}

func (fake *FakeConnectionWrapper) MakeCallCount() int {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	return len(fake.makeArgsForCall)
}

func (fake *FakeConnectionWrapper) MakeCalls(stub func(*cloudcontroller.Request, *cloudcontroller.Response) error) {
	fake.makeMutex.Lock()
	defer fake.makeMutex.Unlock()
	fake.MakeStub = stub
}

func (fake *FakeConnectionWrapper) MakeArgsForCall(i int) (*cloudcontroller.Request, *cloudcontroller.Response) {
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	argsForCall := fake.makeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConnectionWrapper) MakeReturns(result1 error) {
	fake.makeMutex.Lock()
	defer fake.makeMutex.Unlock()
	fake.MakeStub = nil
	fake.makeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) MakeReturnsOnCall(i int, result1 error) {
	fake.makeMutex.Lock()
	defer fake.makeMutex.Unlock()
	fake.MakeStub = nil
	if fake.makeReturnsOnCall == nil {
		fake.makeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnectionWrapper) Wrap(arg1 cloudcontroller.Connection) cloudcontroller.Connection {
	fake.wrapMutex.Lock()
	ret, specificReturn := fake.wrapReturnsOnCall[len(fake.wrapArgsForCall)]
	fake.wrapArgsForCall = append(fake.wrapArgsForCall, struct {
		arg1 cloudcontroller.Connection
	}{arg1})
	fake.recordInvocation("Wrap", []interface{}{arg1})
	fake.wrapMutex.Unlock()
	if fake.WrapStub != nil {
		return fake.WrapStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.wrapReturns
	return fakeReturns.result1
}

func (fake *FakeConnectionWrapper) WrapCallCount() int {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return len(fake.wrapArgsForCall)
}

func (fake *FakeConnectionWrapper) WrapCalls(stub func(cloudcontroller.Connection) cloudcontroller.Connection) {
	fake.wrapMutex.Lock()
	defer fake.wrapMutex.Unlock()
	fake.WrapStub = stub
}

func (fake *FakeConnectionWrapper) WrapArgsForCall(i int) cloudcontroller.Connection {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	argsForCall := fake.wrapArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConnectionWrapper) WrapReturns(result1 cloudcontroller.Connection) {
	fake.wrapMutex.Lock()
	defer fake.wrapMutex.Unlock()
	fake.WrapStub = nil
	fake.wrapReturns = struct {
		result1 cloudcontroller.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) WrapReturnsOnCall(i int, result1 cloudcontroller.Connection) {
	fake.wrapMutex.Lock()
	defer fake.wrapMutex.Unlock()
	fake.WrapStub = nil
	if fake.wrapReturnsOnCall == nil {
		fake.wrapReturnsOnCall = make(map[int]struct {
			result1 cloudcontroller.Connection
		})
	}
	fake.wrapReturnsOnCall[i] = struct {
		result1 cloudcontroller.Connection
	}{result1}
}

func (fake *FakeConnectionWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeMutex.RLock()
	defer fake.makeMutex.RUnlock()
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConnectionWrapper) recordInvocation(key string, args []interface{}) {
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

var _ ccv2.ConnectionWrapper = new(FakeConnectionWrapper)
