package pkg

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path/filepath"
)

var _ = Context("refParser", func() {
	Context("Ref", func() {
		Context("ToPath", func() {
			It("should return expected path", func() {
				/* arrange */
				providedBasePath := "/dummy/path"
				objectUnderTest := &Ref{
					Name:    "test.com/org/pkg-name",
					Version: "0.0.0",
				}

				expectedPath := filepath.Join(
					providedBasePath,
					filepath.FromSlash(fmt.Sprintf("%v#%v", objectUnderTest.Name, objectUnderTest.Version)),
				)

				/* act */
				actualPath := objectUnderTest.ToPath(providedBasePath)

				/* assert */
				Expect(actualPath).To(Equal(expectedPath))
			})
		})
	})
	Context("Parse", func() {
		Context("url.Parse errors", func() {
			It("should error", func() {
				/* arrange */
				providedPkgRef := "::"

				objectUnderTest := _Pkg{
					refParser: newRefParser(),
				}

				/* act */
				_, actualErr := objectUnderTest.Parse(providedPkgRef)

				/* assert */
				Expect(actualErr).To(Not(BeNil()))
			})
		})
		Context("url.Parse doesn't error", func() {
			It("should return expected Ref", func() {
				/* arrange */
				providedFullyQualifiedPkgName := "somehost.com/path/pkgName"
				providedPkgVersion := "0.0.0"
				providedOpRef := fmt.Sprintf("%v#%v/some/op/path", providedFullyQualifiedPkgName, providedPkgVersion)
				expectedPkgRef := &Ref{
					Name:    providedFullyQualifiedPkgName,
					Version: providedPkgVersion,
				}
				objectUnderTest := _Pkg{
					refParser: newRefParser(),
				}

				/* act */
				actualPkgRef, actualErr := objectUnderTest.Parse(providedOpRef)

				/* assert */
				Expect(actualPkgRef).To(Equal(expectedPkgRef))
				Expect(actualErr).To(BeNil())

			})
		})
	})
})
