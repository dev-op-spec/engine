// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type FakeCallStore struct {
	AddStub        func(*model.DCG)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 *model.DCG
	}
	ListWithParentIDStub        func(string) []*model.DCG
	listWithParentIDMutex       sync.RWMutex
	listWithParentIDArgsForCall []struct {
		arg1 string
	}
	listWithParentIDReturns struct {
		result1 []*model.DCG
	}
	listWithParentIDReturnsOnCall map[int]struct {
		result1 []*model.DCG
	}
	SetIsKilledStub        func(string)
	setIsKilledMutex       sync.RWMutex
	setIsKilledArgsForCall []struct {
		arg1 string
	}
	TryGetStub        func(string) *model.DCG
	tryGetMutex       sync.RWMutex
	tryGetArgsForCall []struct {
		arg1 string
	}
	tryGetReturns struct {
		result1 *model.DCG
	}
	tryGetReturnsOnCall map[int]struct {
		result1 *model.DCG
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCallStore) Add(arg1 *model.DCG) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 *model.DCG
	}{arg1})
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(arg1)
	}
}

func (fake *FakeCallStore) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeCallStore) AddCalls(stub func(*model.DCG)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeCallStore) AddArgsForCall(i int) *model.DCG {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCallStore) ListWithParentID(arg1 string) []*model.DCG {
	fake.listWithParentIDMutex.Lock()
	ret, specificReturn := fake.listWithParentIDReturnsOnCall[len(fake.listWithParentIDArgsForCall)]
	fake.listWithParentIDArgsForCall = append(fake.listWithParentIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListWithParentID", []interface{}{arg1})
	fake.listWithParentIDMutex.Unlock()
	if fake.ListWithParentIDStub != nil {
		return fake.ListWithParentIDStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listWithParentIDReturns
	return fakeReturns.result1
}

func (fake *FakeCallStore) ListWithParentIDCallCount() int {
	fake.listWithParentIDMutex.RLock()
	defer fake.listWithParentIDMutex.RUnlock()
	return len(fake.listWithParentIDArgsForCall)
}

func (fake *FakeCallStore) ListWithParentIDCalls(stub func(string) []*model.DCG) {
	fake.listWithParentIDMutex.Lock()
	defer fake.listWithParentIDMutex.Unlock()
	fake.ListWithParentIDStub = stub
}

func (fake *FakeCallStore) ListWithParentIDArgsForCall(i int) string {
	fake.listWithParentIDMutex.RLock()
	defer fake.listWithParentIDMutex.RUnlock()
	argsForCall := fake.listWithParentIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCallStore) ListWithParentIDReturns(result1 []*model.DCG) {
	fake.listWithParentIDMutex.Lock()
	defer fake.listWithParentIDMutex.Unlock()
	fake.ListWithParentIDStub = nil
	fake.listWithParentIDReturns = struct {
		result1 []*model.DCG
	}{result1}
}

func (fake *FakeCallStore) ListWithParentIDReturnsOnCall(i int, result1 []*model.DCG) {
	fake.listWithParentIDMutex.Lock()
	defer fake.listWithParentIDMutex.Unlock()
	fake.ListWithParentIDStub = nil
	if fake.listWithParentIDReturnsOnCall == nil {
		fake.listWithParentIDReturnsOnCall = make(map[int]struct {
			result1 []*model.DCG
		})
	}
	fake.listWithParentIDReturnsOnCall[i] = struct {
		result1 []*model.DCG
	}{result1}
}

func (fake *FakeCallStore) SetIsKilled(arg1 string) {
	fake.setIsKilledMutex.Lock()
	fake.setIsKilledArgsForCall = append(fake.setIsKilledArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetIsKilled", []interface{}{arg1})
	fake.setIsKilledMutex.Unlock()
	if fake.SetIsKilledStub != nil {
		fake.SetIsKilledStub(arg1)
	}
}

func (fake *FakeCallStore) SetIsKilledCallCount() int {
	fake.setIsKilledMutex.RLock()
	defer fake.setIsKilledMutex.RUnlock()
	return len(fake.setIsKilledArgsForCall)
}

func (fake *FakeCallStore) SetIsKilledCalls(stub func(string)) {
	fake.setIsKilledMutex.Lock()
	defer fake.setIsKilledMutex.Unlock()
	fake.SetIsKilledStub = stub
}

func (fake *FakeCallStore) SetIsKilledArgsForCall(i int) string {
	fake.setIsKilledMutex.RLock()
	defer fake.setIsKilledMutex.RUnlock()
	argsForCall := fake.setIsKilledArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCallStore) TryGet(arg1 string) *model.DCG {
	fake.tryGetMutex.Lock()
	ret, specificReturn := fake.tryGetReturnsOnCall[len(fake.tryGetArgsForCall)]
	fake.tryGetArgsForCall = append(fake.tryGetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("TryGet", []interface{}{arg1})
	fake.tryGetMutex.Unlock()
	if fake.TryGetStub != nil {
		return fake.TryGetStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tryGetReturns
	return fakeReturns.result1
}

func (fake *FakeCallStore) TryGetCallCount() int {
	fake.tryGetMutex.RLock()
	defer fake.tryGetMutex.RUnlock()
	return len(fake.tryGetArgsForCall)
}

func (fake *FakeCallStore) TryGetCalls(stub func(string) *model.DCG) {
	fake.tryGetMutex.Lock()
	defer fake.tryGetMutex.Unlock()
	fake.TryGetStub = stub
}

func (fake *FakeCallStore) TryGetArgsForCall(i int) string {
	fake.tryGetMutex.RLock()
	defer fake.tryGetMutex.RUnlock()
	argsForCall := fake.tryGetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCallStore) TryGetReturns(result1 *model.DCG) {
	fake.tryGetMutex.Lock()
	defer fake.tryGetMutex.Unlock()
	fake.TryGetStub = nil
	fake.tryGetReturns = struct {
		result1 *model.DCG
	}{result1}
}

func (fake *FakeCallStore) TryGetReturnsOnCall(i int, result1 *model.DCG) {
	fake.tryGetMutex.Lock()
	defer fake.tryGetMutex.Unlock()
	fake.TryGetStub = nil
	if fake.tryGetReturnsOnCall == nil {
		fake.tryGetReturnsOnCall = make(map[int]struct {
			result1 *model.DCG
		})
	}
	fake.tryGetReturnsOnCall[i] = struct {
		result1 *model.DCG
	}{result1}
}

func (fake *FakeCallStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.listWithParentIDMutex.RLock()
	defer fake.listWithParentIDMutex.RUnlock()
	fake.setIsKilledMutex.RLock()
	defer fake.setIsKilledMutex.RUnlock()
	fake.tryGetMutex.RLock()
	defer fake.tryGetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCallStore) recordInvocation(key string, args []interface{}) {
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