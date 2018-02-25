package expression

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg"
)

var _ = Context("EvalToObject", func() {
	Context("expression is map[string]interface{}", func() {
		It("should call evalObjectInitializer.Eval w/ expected args", func() {

			/* arrange */
			providedScope := map[string]*model.Value{"dummyName": {}}
			providedExpression := map[string]interface{}{
				"prop1Name": "prop1Value",
			}
			providedPkgRef := new(pkg.FakeHandle)

			fakeEvalObjectInitializerer := new(fakeEvalObjectInitializerer)
			// err to trigger immediate return
			evalErr := errors.New("evalErr")
			fakeEvalObjectInitializerer.EvalReturns(map[string]interface{}{}, evalErr)

			objectUnderTest := _evalObjecter{
				evalObjectInitializerer: fakeEvalObjectInitializerer,
			}

			/* act */
			objectUnderTest.EvalToObject(
				providedScope,
				providedExpression,
				providedPkgRef,
			)

			/* assert */
			actualExpression,
				actualScope,
				actualPkgRef := fakeEvalObjectInitializerer.EvalArgsForCall(0)

			Expect(actualExpression).To(Equal(providedExpression))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualPkgRef).To(Equal(providedPkgRef))

		})
		Context("evalObjectInitializer.Eval errs", func() {
			It("should return expected result", func() {

				/* arrange */
				providedExpression := map[string]interface{}{
					"prop1Name": "prop1Value",
				}

				fakeEvalObjectInitializerer := new(fakeEvalObjectInitializerer)
				// err to trigger immediate return
				evalErr := errors.New("evalErr")
				fakeEvalObjectInitializerer.EvalReturns(map[string]interface{}{}, evalErr)

				expectedErr := fmt.Errorf("unable to evaluate %+v to object; error was %v", providedExpression, evalErr)

				objectUnderTest := _evalObjecter{
					evalObjectInitializerer: fakeEvalObjectInitializerer,
				}

				/* act */
				_, actualErr := objectUnderTest.EvalToObject(
					map[string]*model.Value{},
					providedExpression,
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))

			})

		})
		Context("evalObjectInitializer.Eval doesn't err", func() {
			It("should return expected result", func() {

				/* arrange */
				fakeEvalObjectInitializerer := new(fakeEvalObjectInitializerer)
				expectedValue := map[string]interface{}{
					"prop1Name": "prop1Value",
				}
				fakeEvalObjectInitializerer.EvalReturns(expectedValue, nil)

				objectUnderTest := _evalObjecter{
					evalObjectInitializerer: fakeEvalObjectInitializerer,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToObject(
					map[string]*model.Value{},
					map[string]interface{}{},
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(*actualValue).To(Equal(model.Value{Object: expectedValue}))
				Expect(actualErr).To(BeNil())
			})
		})
	})
	Context("expression is string", func() {
		It("should call interpolater.Interpolate w/ expected args", func() {
			/* arrange */
			providedScope := map[string]*model.Value{"dummyName": {}}
			providedExpression := "dummyExpression"
			providedPkgRef := new(pkg.FakeHandle)

			fakeInterpolater := new(interpolater.Fake)
			// err to trigger immediate return
			fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

			objectUnderTest := _evalObjecter{
				interpolater: fakeInterpolater,
			}

			/* act */
			objectUnderTest.EvalToObject(
				providedScope,
				providedExpression,
				providedPkgRef,
			)

			/* assert */
			actualExpression,
				actualScope,
				actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

			Expect(actualExpression).To(Equal(providedExpression))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualPkgRef).To(Equal(providedPkgRef))

		})
		Context("interpolater.Interpolate errs", func() {
			It("should return expected err", func() {
				/* arrange */
				fakeInterpolater := new(interpolater.Fake)
				interpolateErr := errors.New("dummyError")
				fakeInterpolater.InterpolateReturns("", interpolateErr)

				objectUnderTest := _evalObjecter{
					interpolater: fakeInterpolater,
				}

				/* act */
				_, actualErr := objectUnderTest.EvalToObject(
					map[string]*model.Value{},
					"dummyExpression",
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(interpolateErr))

			})
		})
		Context("interpolater.Interpolate doesn't err", func() {
			It("should call data.CoerceToObject w/ expected args & return result", func() {
				/* arrange */
				fakeInterpolater := new(interpolater.Fake)

				interpolatedValue := "dummyString"
				fakeInterpolater.InterpolateReturns(interpolatedValue, nil)

				fakeData := new(data.Fake)

				coercedValue := model.Value{Object: map[string]interface{}{"dummyName": "dummyValue"}}
				fakeData.CoerceToObjectReturns(&coercedValue, nil)

				objectUnderTest := _evalObjecter{
					data:         fakeData,
					interpolater: fakeInterpolater,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToObject(
					map[string]*model.Value{},
					"dummyExpression",
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(*fakeData.CoerceToObjectArgsForCall(0)).To(Equal(model.Value{String: &interpolatedValue}))

				Expect(*actualValue).To(Equal(coercedValue))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
