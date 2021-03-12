package dirs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Interpret", func() {
	Context("dir.Interpret errs", func() {
		It("should return expected error", func() {
			/* arrange */
			identifier := "identifier"
			dataDir, err := ioutil.TempDir("", "")
			Expect(err).To(BeNil())

			/* act */
			_, actualErr := Interpret(
				map[string]*model.Value{
					identifier: {
						Socket: new(string),
					},
				},
				map[string]interface{}{
					"/something": fmt.Sprintf("$(%s)", identifier),
				},
				dataDir,
				"dataDirPath",
			)

			/* assert */
			Expect(actualErr).To(Equal(errors.New("unable to bind /something to $(identifier); error was unable to interpret $(identifier) to dir; error was unable to coerce socket to dir; incompatible types")))
		})
	})
	Context("dir.Interpret doesn't err", func() {
		Context("value.Dir not prefixed by dataCachePath", func() {
			It("should return expected results", func() {
				/* arrange */
				identifier := "identifier"

				dataDir, err := ioutil.TempDir("", "")
				if nil != err {
					panic(err)
				}
				dirPath, err := ioutil.TempDir("", "")
				if nil != err {
					panic(err)
				}

				expectedDirs := map[string]string{
					"/something": filepath.Join(dataDir, "/something"),
				}

				/* act */
				actualContainerCallDirs, actualErr := Interpret(
					map[string]*model.Value{
						identifier: {Dir: &dirPath},
					},
					map[string]interface{}{
						// implicitly bound
						"/something": fmt.Sprintf("$(%s)", identifier),
					},
					dataDir,
					filepath.Dir(dirPath),
				)

				/* assert */
				Expect(actualErr).To(BeNil())
				Expect(actualContainerCallDirs).To(Equal(expectedDirs))

			})
		})
		Context("value.Dir prefixed by dataCachePath", func() {
			Context("dircopier.OS errs", func() {
				It("should return expected result", func() {
					/* arrange */
					identifier := "identifier"
					containerPath := "/something"
					dirValue, err := ioutil.TempDir("", "")
					if nil != err {
						panic(err)
					}

					scratchDirPath, err := ioutil.TempDir("", "")
					if nil != err {
						panic(err)
					}

					expectedResult := map[string]string{
						containerPath: filepath.Join(scratchDirPath, containerPath),
					}

					/* act */
					actualResult, actualErr := Interpret(
						map[string]*model.Value{
							identifier: &model.Value{Dir: &dirValue},
						},
						map[string]interface{}{
							// implicitly bound
							containerPath: fmt.Sprintf("$(%s)", identifier),
						},
						scratchDirPath,
						filepath.Dir(dirValue),
					)

					/* assert */
					Expect(actualErr).To(BeNil())
					Expect(actualResult).To(Equal(expectedResult))
				})
			})
		})
	})
})
