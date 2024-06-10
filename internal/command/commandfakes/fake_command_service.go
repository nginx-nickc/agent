// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"context"
	"sync"

	v1 "github.com/nginx/agent/v3/api/grpc/mpi/v1"
)

type FakeCommandService struct {
	CancelSubscriptionStub        func(context.Context)
	cancelSubscriptionMutex       sync.RWMutex
	cancelSubscriptionArgsForCall []struct {
		arg1 context.Context
	}
	UpdateDataPlaneHealthStub        func(context.Context, []*v1.InstanceHealth) error
	updateDataPlaneHealthMutex       sync.RWMutex
	updateDataPlaneHealthArgsForCall []struct {
		arg1 context.Context
		arg2 []*v1.InstanceHealth
	}
	updateDataPlaneHealthReturns struct {
		result1 error
	}
	updateDataPlaneHealthReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateDataPlaneStatusStub        func(context.Context, *v1.Resource) error
	updateDataPlaneStatusMutex       sync.RWMutex
	updateDataPlaneStatusArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Resource
	}
	updateDataPlaneStatusReturns struct {
		result1 error
	}
	updateDataPlaneStatusReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandService) CancelSubscription(arg1 context.Context) {
	fake.cancelSubscriptionMutex.Lock()
	fake.cancelSubscriptionArgsForCall = append(fake.cancelSubscriptionArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.CancelSubscriptionStub
	fake.recordInvocation("CancelSubscription", []interface{}{arg1})
	fake.cancelSubscriptionMutex.Unlock()
	if stub != nil {
		fake.CancelSubscriptionStub(arg1)
	}
}

func (fake *FakeCommandService) CancelSubscriptionCallCount() int {
	fake.cancelSubscriptionMutex.RLock()
	defer fake.cancelSubscriptionMutex.RUnlock()
	return len(fake.cancelSubscriptionArgsForCall)
}

func (fake *FakeCommandService) CancelSubscriptionCalls(stub func(context.Context)) {
	fake.cancelSubscriptionMutex.Lock()
	defer fake.cancelSubscriptionMutex.Unlock()
	fake.CancelSubscriptionStub = stub
}

func (fake *FakeCommandService) CancelSubscriptionArgsForCall(i int) context.Context {
	fake.cancelSubscriptionMutex.RLock()
	defer fake.cancelSubscriptionMutex.RUnlock()
	argsForCall := fake.cancelSubscriptionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCommandService) UpdateDataPlaneHealth(arg1 context.Context, arg2 []*v1.InstanceHealth) error {
	var arg2Copy []*v1.InstanceHealth
	if arg2 != nil {
		arg2Copy = make([]*v1.InstanceHealth, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateDataPlaneHealthMutex.Lock()
	ret, specificReturn := fake.updateDataPlaneHealthReturnsOnCall[len(fake.updateDataPlaneHealthArgsForCall)]
	fake.updateDataPlaneHealthArgsForCall = append(fake.updateDataPlaneHealthArgsForCall, struct {
		arg1 context.Context
		arg2 []*v1.InstanceHealth
	}{arg1, arg2Copy})
	stub := fake.UpdateDataPlaneHealthStub
	fakeReturns := fake.updateDataPlaneHealthReturns
	fake.recordInvocation("UpdateDataPlaneHealth", []interface{}{arg1, arg2Copy})
	fake.updateDataPlaneHealthMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCommandService) UpdateDataPlaneHealthCallCount() int {
	fake.updateDataPlaneHealthMutex.RLock()
	defer fake.updateDataPlaneHealthMutex.RUnlock()
	return len(fake.updateDataPlaneHealthArgsForCall)
}

func (fake *FakeCommandService) UpdateDataPlaneHealthCalls(stub func(context.Context, []*v1.InstanceHealth) error) {
	fake.updateDataPlaneHealthMutex.Lock()
	defer fake.updateDataPlaneHealthMutex.Unlock()
	fake.UpdateDataPlaneHealthStub = stub
}

func (fake *FakeCommandService) UpdateDataPlaneHealthArgsForCall(i int) (context.Context, []*v1.InstanceHealth) {
	fake.updateDataPlaneHealthMutex.RLock()
	defer fake.updateDataPlaneHealthMutex.RUnlock()
	argsForCall := fake.updateDataPlaneHealthArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommandService) UpdateDataPlaneHealthReturns(result1 error) {
	fake.updateDataPlaneHealthMutex.Lock()
	defer fake.updateDataPlaneHealthMutex.Unlock()
	fake.UpdateDataPlaneHealthStub = nil
	fake.updateDataPlaneHealthReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandService) UpdateDataPlaneHealthReturnsOnCall(i int, result1 error) {
	fake.updateDataPlaneHealthMutex.Lock()
	defer fake.updateDataPlaneHealthMutex.Unlock()
	fake.UpdateDataPlaneHealthStub = nil
	if fake.updateDataPlaneHealthReturnsOnCall == nil {
		fake.updateDataPlaneHealthReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateDataPlaneHealthReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandService) UpdateDataPlaneStatus(arg1 context.Context, arg2 *v1.Resource) error {
	fake.updateDataPlaneStatusMutex.Lock()
	ret, specificReturn := fake.updateDataPlaneStatusReturnsOnCall[len(fake.updateDataPlaneStatusArgsForCall)]
	fake.updateDataPlaneStatusArgsForCall = append(fake.updateDataPlaneStatusArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Resource
	}{arg1, arg2})
	stub := fake.UpdateDataPlaneStatusStub
	fakeReturns := fake.updateDataPlaneStatusReturns
	fake.recordInvocation("UpdateDataPlaneStatus", []interface{}{arg1, arg2})
	fake.updateDataPlaneStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCommandService) UpdateDataPlaneStatusCallCount() int {
	fake.updateDataPlaneStatusMutex.RLock()
	defer fake.updateDataPlaneStatusMutex.RUnlock()
	return len(fake.updateDataPlaneStatusArgsForCall)
}

func (fake *FakeCommandService) UpdateDataPlaneStatusCalls(stub func(context.Context, *v1.Resource) error) {
	fake.updateDataPlaneStatusMutex.Lock()
	defer fake.updateDataPlaneStatusMutex.Unlock()
	fake.UpdateDataPlaneStatusStub = stub
}

func (fake *FakeCommandService) UpdateDataPlaneStatusArgsForCall(i int) (context.Context, *v1.Resource) {
	fake.updateDataPlaneStatusMutex.RLock()
	defer fake.updateDataPlaneStatusMutex.RUnlock()
	argsForCall := fake.updateDataPlaneStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommandService) UpdateDataPlaneStatusReturns(result1 error) {
	fake.updateDataPlaneStatusMutex.Lock()
	defer fake.updateDataPlaneStatusMutex.Unlock()
	fake.UpdateDataPlaneStatusStub = nil
	fake.updateDataPlaneStatusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandService) UpdateDataPlaneStatusReturnsOnCall(i int, result1 error) {
	fake.updateDataPlaneStatusMutex.Lock()
	defer fake.updateDataPlaneStatusMutex.Unlock()
	fake.UpdateDataPlaneStatusStub = nil
	if fake.updateDataPlaneStatusReturnsOnCall == nil {
		fake.updateDataPlaneStatusReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateDataPlaneStatusReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelSubscriptionMutex.RLock()
	defer fake.cancelSubscriptionMutex.RUnlock()
	fake.updateDataPlaneHealthMutex.RLock()
	defer fake.updateDataPlaneHealthMutex.RUnlock()
	fake.updateDataPlaneStatusMutex.RLock()
	defer fake.updateDataPlaneStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommandService) recordInvocation(key string, args []interface{}) {
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
