package param

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/types"
)

var _ = Context("Validate", func() {
	objectUnderTest := NewValidator()
	Context("invoked w/ nil param", func() {
		It("should return expected error", func() {
			/* arrange */
			providedValue := &types.Value{}

			expectedErrs := []error{errors.New("param required")}

			/* act */
			actualErrs := objectUnderTest.Validate(providedValue, nil)
			Expect(actualErrs).To(Equal(expectedErrs))
		})
	})
})
