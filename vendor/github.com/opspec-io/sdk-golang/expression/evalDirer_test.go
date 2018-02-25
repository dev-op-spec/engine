package expression

import (
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg"
	"path/filepath"
)

var _ = Context("EvalToDir", func() {
	Context("expression is scope ref", func() {
		It("should return expected result", func() {
			/* arrange */
			scopeName := "dummyScopeName"
			providedExpression := fmt.Sprintf("$(%v)", scopeName)
			scopeValue := model.Value{Dir: new(string)}
			providedScope := map[string]*model.Value{
				scopeName: &scopeValue,
			}

			objectUnderTest := _evalDirer{}

			/* act */
			actualValue, actualErr := objectUnderTest.EvalToDir(
				providedScope,
				providedExpression,
				new(pkg.FakeHandle),
			)

			/* assert */
			Expect(*actualValue).To(Equal(scopeValue))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("expression is deprecated pkg fs ref", func() {
		It("should call interpolater.Interpolate w/ expected args", func() {
			/* arrange */
			providedScope := map[string]*model.Value{}
			providedExpression := "/dummy/deprecated/pkg-fs-ref"
			providedPkgRef := new(pkg.FakeHandle)

			fakeInterpolater := new(interpolater.Fake)
			// err to trigger immediate return
			fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

			objectUnderTest := _evalDirer{
				interpolater: fakeInterpolater,
			}

			/* act */
			objectUnderTest.EvalToDir(
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
		Context("interpolater.Interpolate errors", func() {
			It("should return expected result", func() {
				/* arrange */
				providedExpression := "/deprecatedPkgFsRef"

				fakeInterpolater := new(interpolater.Fake)
				interpolateError := errors.New("dummyError")
				fakeInterpolater.InterpolateReturns("", interpolateError)

				expectedErr := fmt.Errorf(
					"unable to evaluate %v to dir; error was %v",
					providedExpression,
					interpolateError.Error(),
				)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					map[string]*model.Value{},
					providedExpression,
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("interpolater.Interpolate doesn't error", func() {
			It("should return expected result", func() {
				/* arrange */
				interpolatedExpression := "dummyExpression"

				fakeInterpolater := new(interpolater.Fake)
				fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				fakeHandle := new(pkg.FakeHandle)
				pkgRef := "dummyPkgRef"
				fakeHandle.RefReturns(pkgRef)

				expectedFileValue := filepath.Join(pkgRef, interpolatedExpression)

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					map[string]*model.Value{},
					"/deprecatedPkgFsRef",
					fakeHandle,
				)

				/* assert */
				Expect(*actualValue).To(Equal(model.Value{Dir: &expectedFileValue}))
				Expect(actualErr).To(BeNil())
			})
		})
	})
	Context("expression is scope ref w/ path", func() {
		It("should call interpolater.Interpolate w/ expected args", func() {
			/* arrange */
			scopeName := "dummyScopeName"
			providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

			providedPath := "dummyPath"
			providedExpression := fmt.Sprintf("$(%v/%v)", scopeName, providedPath)
			providedPkgRef := new(pkg.FakeHandle)

			fakeInterpolater := new(interpolater.Fake)
			// err to trigger immediate return
			fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

			objectUnderTest := _evalDirer{
				interpolater: fakeInterpolater,
			}

			/* act */
			objectUnderTest.EvalToDir(
				providedScope,
				providedExpression,
				providedPkgRef,
			)

			/* assert */
			actualExpression,
				actualScope,
				actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

			Expect(actualExpression).To(Equal(providedPath))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualPkgRef).To(Equal(providedPkgRef))
		})
		Context("interpolater.Interpolate errs", func() {
			It("should return expected result", func() {

				/* arrange */
				scopeName := "dummyScopeName"
				providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

				providedPath := "dummyPath"
				providedExpression := fmt.Sprintf("$(%v/%v)", scopeName, providedPath)

				fakeInterpolater := new(interpolater.Fake)
				interpolateErr := errors.New("dummyError")
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				expectedErr := fmt.Errorf(
					"unable to evaluate path %v; error was %v",
					providedPath,
					interpolateErr.Error(),
				)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					providedScope,
					providedExpression,
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(expectedErr))

			})
		})
		Context("interpolater.Interpolate doesn't error", func() {
			It("should return expected result", func() {
				/* arrange */
				scopeName := "dummyScopeName"
				scopeValue := "dummyScopeValue"

				interpolatedExpression := "dummyInterpolatedExpression"
				fakeInterpolater := new(interpolater.Fake)
				fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				expectedFileValue := filepath.Join(scopeValue, interpolatedExpression)

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					map[string]*model.Value{scopeName: {Dir: &scopeValue}},
					fmt.Sprintf("$(%v/path)", scopeName),
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(*actualValue).To(Equal(model.Value{Dir: &expectedFileValue}))
				Expect(actualErr).To(BeNil())
			})
		})
	})
	Context("expression is scope ref & deprecated path", func() {
		It("should call interpolater.Interpolate w/ expected args", func() {
			/* arrange */
			scopeName := "dummyScopeName"
			providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

			providedPath := "/dummyPath"
			providedExpression := fmt.Sprintf("$(%v)%v", scopeName, providedPath)
			providedPkgRef := new(pkg.FakeHandle)

			fakeInterpolater := new(interpolater.Fake)
			// err to trigger immediate return
			fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

			objectUnderTest := _evalDirer{
				interpolater: fakeInterpolater,
			}

			/* act */
			objectUnderTest.EvalToDir(
				providedScope,
				providedExpression,
				providedPkgRef,
			)

			/* assert */
			actualExpression,
				actualScope,
				actualPkgRef := fakeInterpolater.InterpolateArgsForCall(0)

			Expect(actualExpression).To(Equal(providedPath))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualPkgRef).To(Equal(providedPkgRef))
		})
		Context("interpolater.Interpolate errs", func() {
			It("should return expected result", func() {

				/* arrange */
				scopeName := "dummyScopeName"
				providedScope := map[string]*model.Value{scopeName: {Dir: new(string)}}

				providedPath := "/dummyPath"
				providedExpression := fmt.Sprintf("$(%v)%v", scopeName, providedPath)

				fakeInterpolater := new(interpolater.Fake)
				interpolateErr := errors.New("dummyError")
				fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

				expectedErr := fmt.Errorf(
					"unable to evaluate path %v; error was %v",
					providedPath,
					interpolateErr.Error(),
				)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					providedScope,
					providedExpression,
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(expectedErr))

			})
		})
		Context("interpolater.Interpolate doesn't error", func() {
			It("should return expected result", func() {
				/* arrange */
				scopeName := "dummyScopeName"
				scopeValue := "dummyDirValue"

				interpolatedExpression := "dummyInterpolatedExpression"
				fakeInterpolater := new(interpolater.Fake)
				fakeInterpolater.InterpolateReturns(interpolatedExpression, nil)

				objectUnderTest := _evalDirer{
					interpolater: fakeInterpolater,
				}

				expectedFileValue := filepath.Join(scopeValue, interpolatedExpression)

				/* act */
				actualValue, actualErr := objectUnderTest.EvalToDir(
					map[string]*model.Value{scopeName: {Dir: &scopeValue}},
					fmt.Sprintf("$(%v)/path", scopeName),
					new(pkg.FakeHandle),
				)

				/* assert */
				Expect(*actualValue).To(Equal(model.Value{Dir: &expectedFileValue}))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
