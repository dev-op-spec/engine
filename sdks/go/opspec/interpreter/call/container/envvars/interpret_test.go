package envvars

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Interpret", func() {
	Context("object.Interpret errs", func() {
		It("should return expected result", func() {
			/* arrange */
			/* act */
			_, actualErr := Interpret(
				map[string]*model.Value{},
				"$()",
			)

			/* assert */
			Expect(actualErr).To(Equal(errors.New("unable to interpret '$()' as envVars; error was unable to interpret $() to object; error was unable to interpret '' as reference; '' not in scope")))
		})
	})
	Context("object.Interpret doesn't err", func() {
		Context("str.Interpret errs", func() {

			It("should return expected result", func() {
				/* arrange */
				identifier := "identifier"
				objectValue := map[string]interface{}{
					"key": "$()",
				}

				providedScope := map[string]*model.Value{
					identifier: &model.Value{
						Object: &objectValue,
					},
				}

				/* act */
				_, actualErr := Interpret(
					providedScope,
					fmt.Sprintf("$(%s)", identifier),
				)

				/* assert */
				Expect(actualErr).To(Equal(errors.New("unable to interpret $() as value of env var 'key'; error was unable to interpret $() to string; error was unable to interpret '' as reference; '' not in scope")))
			})
		})
		Context("str.Interpret doesn't err", func() {
			It("should return expected result", func() {
				/* arrange */
				identifier := "identifier"
				objectValue := map[string]interface{}{}

				providedScope := map[string]*model.Value{
					identifier: &model.Value{
						Object: &objectValue,
					},
				}

				/* act */
				actualResult, actualErr := Interpret(
					providedScope,
					fmt.Sprintf("$(%s)", identifier),
				)

				/* assert */
				Expect(actualErr).To(BeNil())
				Expect(actualResult).To(Equal(map[string]string{}))
			})
		})
	})
})
