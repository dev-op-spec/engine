package predicate

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/predicates/predicate/eq"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/predicates/predicate/exists"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/predicates/predicate/ne"
	"github.com/opctl/opctl/sdks/go/types"
)

var _ = Context("Interpreter", func() {
	Context("NewInterpreter", func() {
		It("shouldn't return nil", func() {
			/* arrange/act/assert */
			Expect(NewInterpreter()).To(Not(BeNil()))
		})
	})
	Context("Interpret", func() {
		Context("Eq Predicate", func() {
			It("should call eqInterpreter.Interpret w/ expected args & return result", func() {
				/* arrange */
				providedOpHandle := new(data.FakeHandle)

				providedScgPredicate := &types.SCGPredicate{
					Eq: new([]interface{}),
				}

				providedScope := map[string]*types.Value{}

				fakeEqInterpreter := new(eq.FakeInterpreter)
				expectedResult := true
				expectedError := errors.New("expectedErr")
				fakeEqInterpreter.InterpretReturns(true, expectedError)

				objectUnderTest := _interpreter{
					eqInterpreter: fakeEqInterpreter,
				}

				/* act */
				actualResult, actualError := objectUnderTest.Interpret(
					providedOpHandle,
					providedScgPredicate,
					providedScope,
				)

				/* assert */
				actualExpression,
					actualOpHandle,
					actualScope := fakeEqInterpreter.InterpretArgsForCall(0)

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualExpression).To(Equal(*providedScgPredicate.Eq))
				Expect(actualOpHandle).To(Equal(providedOpHandle))

				Expect(actualResult).To(Equal(expectedResult))
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("Exists Predicate", func() {
			It("should call existsInterpreter.Interpret w/ expected args & return result", func() {
				/* arrange */
				providedOpHandle := new(data.FakeHandle)

				providedScgPredicate := &types.SCGPredicate{
					Exists: new(string),
				}

				providedScope := map[string]*types.Value{}

				fakeExistsInterpreter := new(exists.FakeInterpreter)
				expectedResult := true
				expectedError := errors.New("expectedErr")
				fakeExistsInterpreter.InterpretReturns(true, expectedError)

				objectUnderTest := _interpreter{
					existsInterpreter: fakeExistsInterpreter,
				}

				/* act */
				actualResult, actualError := objectUnderTest.Interpret(
					providedOpHandle,
					providedScgPredicate,
					providedScope,
				)

				/* assert */
				actualExpression,
					actualOpHandle,
					actualScope := fakeExistsInterpreter.InterpretArgsForCall(0)

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualExpression).To(Equal(*providedScgPredicate.Exists))
				Expect(actualOpHandle).To(Equal(providedOpHandle))

				Expect(actualResult).To(Equal(expectedResult))
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("Ne predicate", func() {
			It("should call neInterpreter.Interpret w/ expected args & return result", func() {
				/* arrange */
				providedOpHandle := new(data.FakeHandle)

				providedScgPredicate := &types.SCGPredicate{
					Ne: new([]interface{}),
				}

				providedScope := map[string]*types.Value{}

				fakeNeInterpreter := new(ne.FakeInterpreter)
				expectedResult := true
				expectedError := errors.New("expectedErr")
				fakeNeInterpreter.InterpretReturns(true, expectedError)

				objectUnderTest := _interpreter{
					neInterpreter: fakeNeInterpreter,
				}

				/* act */
				actualResult, actualError := objectUnderTest.Interpret(
					providedOpHandle,
					providedScgPredicate,
					providedScope,
				)

				/* assert */
				actualExpression,
					actualOpHandle,
					actualScope := fakeNeInterpreter.InterpretArgsForCall(0)

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualExpression).To(Equal(*providedScgPredicate.Ne))
				Expect(actualOpHandle).To(Equal(providedOpHandle))

				Expect(actualResult).To(Equal(expectedResult))
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("NotExists Predicate", func() {
			It("should call existsInterpreter.Interpret w/ expected args & return result", func() {
				/* arrange */
				providedOpHandle := new(data.FakeHandle)

				providedScgPredicate := &types.SCGPredicate{
					NotExists: new(string),
				}

				providedScope := map[string]*types.Value{}

				fakeNotExistsInterpreter := new(exists.FakeInterpreter)
				expectedResult := true
				expectedError := errors.New("expectedErr")
				fakeNotExistsInterpreter.InterpretReturns(true, expectedError)

				objectUnderTest := _interpreter{
					notExistsInterpreter: fakeNotExistsInterpreter,
				}

				/* act */
				actualResult, actualError := objectUnderTest.Interpret(
					providedOpHandle,
					providedScgPredicate,
					providedScope,
				)

				/* assert */
				actualExpression,
					actualOpHandle,
					actualScope := fakeNotExistsInterpreter.InterpretArgsForCall(0)

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualExpression).To(Equal(*providedScgPredicate.NotExists))
				Expect(actualOpHandle).To(Equal(providedOpHandle))

				Expect(actualResult).To(Equal(expectedResult))
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("Unexpected predicate", func() {
			It("should return expected result", func() {
				/* arrange */
				providedScgPredicate := &types.SCGPredicate{}
				objectUnderTest := _interpreter{}

				expectedError := fmt.Errorf("unable to interpret predicate; predicate was unexpected type %+v", providedScgPredicate)

				/* act */
				_, actualError := objectUnderTest.Interpret(
					new(data.FakeHandle),
					providedScgPredicate,
					map[string]*types.Value{},
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
	})
})
