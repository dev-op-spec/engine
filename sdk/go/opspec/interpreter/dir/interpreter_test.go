package dir

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/data"
	"github.com/opctl/sdk-golang/model"
	"github.com/opctl/sdk-golang/opspec/interpreter/interpolater"
	"github.com/opctl/sdk-golang/opspec/interpreter/reference"
)

var _ = Context("Interpret", func() {
	Context("expression is scope ref", func() {
		It("should call referenceInterpreter.Interpret w/ expected args", func() {
			/* arrange */
			providedScope := map[string]*model.Value{"": new(model.Value)}

			providedExpression := "$(providedReference)"
			providedOpRef := new(data.FakeHandle)

			expectedReference := strings.TrimSuffix(strings.TrimPrefix(providedExpression, interpolater.RefStart), interpolater.RefEnd)

			fakeReferenceInterpreter := new(reference.FakeInterpreter)
			// err to trigger immediate return
			fakeReferenceInterpreter.InterpretReturns(nil, errors.New("dummyError"))

			objectUnderTest := _interpreter{
				referenceInterpreter: fakeReferenceInterpreter,
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedExpression,
				providedOpRef,
			)

			/* assert */
			expectedReference,
				actualScope,
				actualOpRef := fakeReferenceInterpreter.InterpretArgsForCall(0)

			Expect(expectedReference).To(Equal(expectedReference))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualOpRef).To(Equal(providedOpRef))
		})
		Context("referenceInterpreter.Interpret errs", func() {
			It("should return expected result", func() {

				/* arrange */
				providedScope := map[string]*model.Value{"": new(model.Value)}

				providedExpression := "$(providedReference)"

				fakeReferenceInterpreter := new(reference.FakeInterpreter)
				interpretErr := errors.New("dummyError")
				fakeReferenceInterpreter.InterpretReturns(nil, interpretErr)

				expectedErr := fmt.Errorf("unable to interpret %+v to file; error was %v", providedExpression, interpretErr)

				objectUnderTest := _interpreter{
					referenceInterpreter: fakeReferenceInterpreter,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.Interpret(
					providedScope,
					providedExpression,
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(expectedErr))

			})
		})
		Context("referenceInterpreter.Interpret doesn't error", func() {
			It("should return expected result", func() {
				/* arrange */
				expectedValue := new(model.Value)
				fakeReferenceInterpreter := new(reference.FakeInterpreter)
				fakeReferenceInterpreter.InterpretReturns(expectedValue, nil)

				objectUnderTest := _interpreter{
					referenceInterpreter: fakeReferenceInterpreter,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					"$(providedReference)",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualValue).To(Equal(expectedValue))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
