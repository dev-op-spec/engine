package params

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/sdk-golang/model"
	"path/filepath"
)

var _ = Context("Defaulter", func() {
	Context("NewDefaulter", func() {
		It("shouldn't return nil", func() {
			/* arrange/act/assert */
			Expect(NewDefaulter()).To(Not(BeNil()))
		})
	})
	Context("Default", func() {
		Context("param is array", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := new([]interface{})

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {Array: &model.ArrayParam{Default: providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputs := map[string]*model.Value{
						providedOutputName: {Array: providedOutputDefault},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						"dummyOpPath",
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is boolean", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := true

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {Boolean: &model.BooleanParam{Default: &providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputs := map[string]*model.Value{
						providedOutputName: {Boolean: &providedOutputDefault},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						"dummyOpPath",
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is dir", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := "/pkgDirDefault"

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {Dir: &model.DirParam{Default: &providedOutputDefault}},
					}
					providedOpPath := "dummyOpPath"

					objectUnderTest := _defaulter{}

					expectedOutputValue := filepath.Join(providedOpPath, providedOutputDefault)
					expectedOutputs := map[string]*model.Value{
						providedOutputName: {Dir: &expectedOutputValue},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						providedOpPath,
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is file", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := "/pkgFileDefault"
					providedOpPath := "dummyOpPath"

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {File: &model.FileParam{Default: &providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputValue := filepath.Join(providedOpPath, providedOutputDefault)
					expectedOutputs := map[string]*model.Value{
						providedOutputName: {File: &expectedOutputValue},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						providedOpPath,
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is number", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := 2.2

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {Number: &model.NumberParam{Default: &providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputs := map[string]*model.Value{
						providedOutputName: {Number: &providedOutputDefault},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						"dummyOpPath",
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is object", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := new(map[string]interface{})

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {Object: &model.ObjectParam{Default: providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputs := map[string]*model.Value{
						providedOutputName: {Object: providedOutputDefault},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						"dummyOpPath",
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
		Context("param is string", func() {
			Context("default exists", func() {
				It("should set output to default", func() {
					/* arrange */
					providedOutputName := "outputName"
					providedOutputDefault := "outputDefault"

					providedOutputParams := map[string]*model.Param{
						providedOutputName: {String: &model.StringParam{Default: &providedOutputDefault}},
					}

					objectUnderTest := _defaulter{}

					expectedOutputs := map[string]*model.Value{
						providedOutputName: {String: &providedOutputDefault},
					}

					/* act */
					actualOutputs := objectUnderTest.Default(
						map[string]*model.Value{},
						providedOutputParams,
						"dummyOpPath",
					)

					/* assert */
					Expect(actualOutputs).To(Equal(expectedOutputs))
				})
			})
		})
	})
})
