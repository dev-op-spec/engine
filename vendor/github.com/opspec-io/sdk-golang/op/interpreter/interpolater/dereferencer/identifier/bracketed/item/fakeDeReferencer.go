// Code generated by counterfeiter. DO NOT EDIT.
package item

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type FakeDeReferencer struct {
	DeReferenceStub        func(identifier string, value model.Value) (*model.Value, error)
	deReferenceMutex       sync.RWMutex
	deReferenceArgsForCall []struct {
		identifier string
		value      model.Value
	}
	deReferenceReturns struct {
		result1 *model.Value
		result2 error
	}
	deReferenceReturnsOnCall map[int]struct {
		result1 *model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeReferencer) DeReference(identifier string, value model.Value) (*model.Value, error) {
	fake.deReferenceMutex.Lock()
	ret, specificReturn := fake.deReferenceReturnsOnCall[len(fake.deReferenceArgsForCall)]
	fake.deReferenceArgsForCall = append(fake.deReferenceArgsForCall, struct {
		identifier string
		value      model.Value
	}{identifier, value})
	fake.recordInvocation("DeReference", []interface{}{identifier, value})
	fake.deReferenceMutex.Unlock()
	if fake.DeReferenceStub != nil {
		return fake.DeReferenceStub(identifier, value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deReferenceReturns.result1, fake.deReferenceReturns.result2
}

func (fake *FakeDeReferencer) DeReferenceCallCount() int {
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
	return len(fake.deReferenceArgsForCall)
}

func (fake *FakeDeReferencer) DeReferenceArgsForCall(i int) (string, model.Value) {
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
	return fake.deReferenceArgsForCall[i].identifier, fake.deReferenceArgsForCall[i].value
}

func (fake *FakeDeReferencer) DeReferenceReturns(result1 *model.Value, result2 error) {
	fake.DeReferenceStub = nil
	fake.deReferenceReturns = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeDeReferencer) DeReferenceReturnsOnCall(i int, result1 *model.Value, result2 error) {
	fake.DeReferenceStub = nil
	if fake.deReferenceReturnsOnCall == nil {
		fake.deReferenceReturnsOnCall = make(map[int]struct {
			result1 *model.Value
			result2 error
		})
	}
	fake.deReferenceReturnsOnCall[i] = struct {
		result1 *model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeDeReferencer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deReferenceMutex.RLock()
	defer fake.deReferenceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeReferencer) recordInvocation(key string, args []interface{}) {
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

var _ DeReferencer = new(FakeDeReferencer)
