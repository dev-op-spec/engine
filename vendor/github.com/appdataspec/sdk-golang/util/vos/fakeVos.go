// This file was generated by counterfeiter
package vos

import (
	"sync"
)

type FakeVOS struct {
	GetenvStub        func(key string) string
	getenvMutex       sync.RWMutex
	getenvArgsForCall []struct {
		key string
	}
	getenvReturns struct {
		result1 string
	}
	SetenvStub        func(key, value string) error
	setenvMutex       sync.RWMutex
	setenvArgsForCall []struct {
		key   string
		value string
	}
	setenvReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVOS) Getenv(key string) string {
	fake.getenvMutex.Lock()
	fake.getenvArgsForCall = append(fake.getenvArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Getenv", []interface{}{key})
	fake.getenvMutex.Unlock()
	if fake.GetenvStub != nil {
		return fake.GetenvStub(key)
	} else {
		return fake.getenvReturns.result1
	}
}

func (fake *FakeVOS) GetenvCallCount() int {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return len(fake.getenvArgsForCall)
}

func (fake *FakeVOS) GetenvArgsForCall(i int) string {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return fake.getenvArgsForCall[i].key
}

func (fake *FakeVOS) GetenvReturns(result1 string) {
	fake.GetenvStub = nil
	fake.getenvReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVOS) Setenv(key string, value string) error {
	fake.setenvMutex.Lock()
	fake.setenvArgsForCall = append(fake.setenvArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("Setenv", []interface{}{key, value})
	fake.setenvMutex.Unlock()
	if fake.SetenvStub != nil {
		return fake.SetenvStub(key, value)
	} else {
		return fake.setenvReturns.result1
	}
}

func (fake *FakeVOS) SetenvCallCount() int {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return len(fake.setenvArgsForCall)
}

func (fake *FakeVOS) SetenvArgsForCall(i int) (string, string) {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return fake.setenvArgsForCall[i].key, fake.setenvArgsForCall[i].value
}

func (fake *FakeVOS) SetenvReturns(result1 error) {
	fake.SetenvStub = nil
	fake.setenvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVOS) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVOS) recordInvocation(key string, args []interface{}) {
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

var _ VOS = new(FakeVOS)
