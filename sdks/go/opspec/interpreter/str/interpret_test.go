package str

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/data/coerce"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Interpret", func() {
	Context("value.Interpret errs", func() {
		It("should return expected result", func() {

			/* arrange */
			/* act */
			_, actualErr := Interpret(
				map[string]*model.Value{},
				"$()",
			)

			/* assert */
			Expect(actualErr).To(Equal(errors.New("unable to interpret $() to string; error was unable to interpret '' as reference; '' not in scope")))
		})
	})
	Context("value.Interpret doesn't err", func() {
		It("should return expected result", func() {
			/* arrange */
			providedExpression := "providedExpression"
			providedValue := model.Value{String: &providedExpression}

			expectedValue, err := coerce.ToString(&providedValue)
			if nil != err {
				panic(err)
			}

			/* act */
			actualValue, actualErr := Interpret(
				map[string]*model.Value{},
				providedExpression,
			)

			/* assert */
			Expect(actualErr).To(BeNil())
			Expect(*actualValue).To(Equal(*expectedValue))
		})
	})
})
