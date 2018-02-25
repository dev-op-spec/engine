package pkg

import (
	"errors"
	"fmt"
	"github.com/golang-interfaces/gopkg.in-src-d-go-git.v4"
	"github.com/golang-interfaces/ios"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg/manifest"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
	"path/filepath"
	"time"
)

var _ = Context("puller", func() {
	Context("Pull", func() {
		It("should call refParser.Parse w/ expected args", func() {
			/* arrange */
			providedPkgRef := "dummyPkgRef"

			fakeRefParser := new(fakeRefParser)
			// error to trigger immediate return
			fakeRefParser.ParseReturns(nil, errors.New("dummyError"))

			objectUnderTest := _puller{
				refParser: fakeRefParser,
			}

			/* act */
			objectUnderTest.Pull(
				"dummyPath",
				providedPkgRef,
				nil,
			)

			/* assert */
			Expect(fakeRefParser.ParseArgsForCall(0)).To(Equal(providedPkgRef))
		})
		Context("refParser.Parse errs", func() {
			It("should return error", func() {
				/* arrange */
				expectedError := errors.New("dummyError")

				fakeRefParser := new(fakeRefParser)
				fakeRefParser.ParseReturns(nil, expectedError)

				objectUnderTest := _puller{
					refParser: fakeRefParser,
				}

				/* act */
				actualError := objectUnderTest.Pull(
					"dummyPath",
					"dummyPkgRef",
					nil,
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
		Context("refParser.Parse doesn't err", func() {
			It("should call git.PlainClone w/ expected args", func() {

				/* arrange */
				providedPath := "dummyPath"
				providedPullCreds := &model.PullCreds{
					Username: "dummyUsername",
					Password: "dummyPassword",
				}

				ref := &Ref{
					Name:    "dummyPkgRef",
					Version: "0.0.0",
				}

				fakeRefParser := new(fakeRefParser)
				fakeRefParser.ParseReturns(ref, nil)

				expectedCloneOptions := &git.CloneOptions{
					Auth: &http.BasicAuth{
						Username: providedPullCreds.Username,
						Password: providedPullCreds.Password,
					},
					URL:           fmt.Sprintf("https://%v", ref.Name),
					ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%v", ref.Version)),
					Depth:         1,
					Progress:      os.Stdout,
				}

				fakeGit := new(igit.Fake)

				objectUnderTest := _puller{
					git:       fakeGit,
					os:        new(ios.Fake),
					refParser: fakeRefParser,
				}

				/* act */
				objectUnderTest.Pull(
					providedPath,
					"dummyPkgRef",
					providedPullCreds,
				)

				/* assert */
				actualPath,
					actualIsBare,
					actualCloneOptions := fakeGit.PlainCloneArgsForCall(0)

				Expect(actualPath).To(Equal(ref.ToPath(providedPath)))
				Expect(actualIsBare).To(BeFalse())
				Expect(actualCloneOptions).To(Equal(expectedCloneOptions))
			})
			Context("git.PlainClone errors", func() {
				Context("err.Error() returns git.ErrRepositoryAlreadyExists", func() {
					It("shouldn't call fs.RemoveAll or error", func() {

						/* arrange */
						fakeRefParser := new(fakeRefParser)
						fakeRefParser.ParseReturns(&Ref{}, nil)

						fakeOS := new(ios.Fake)

						fakeGit := new(igit.Fake)
						fakeGit.PlainCloneReturns(nil, git.ErrRepositoryAlreadyExists)

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        fakeOS,
							refParser: fakeRefParser,
						}

						/* act */
						actualError := objectUnderTest.Pull(
							"dummyPath",
							"dummyPkgRef",
							nil,
						)

						/* assert */
						Expect(actualError).To(BeNil())
						Expect(fakeOS.RemoveAllCallCount()).To(Equal(0))
					})
				})
				Context("err.Error() returns transport.ErrAuthenticationRequired error", func() {
					It("should call fs.RemoveAll w/ expected args & return expected error", func() {

						/* arrange */
						providedPath := "dummyPath"

						ref := &Ref{
							Name:    "dummyPkgRef",
							Version: "0.0.0",
						}

						fakeRefParser := new(fakeRefParser)
						fakeRefParser.ParseReturns(ref, nil)

						expectedPath := filepath.Join(
							providedPath,
							fmt.Sprintf("%v#%v", ref.Name, ref.Version),
						)

						fakeOS := new(ios.Fake)
						expectedError := model.ErrPkgPullAuthentication{}

						fakeGit := new(igit.Fake)
						fakeGit.PlainCloneReturns(nil, transport.ErrAuthenticationRequired)

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        fakeOS,
							refParser: fakeRefParser,
						}

						/* act */
						actualError := objectUnderTest.Pull(
							providedPath,
							"dummyPkgRef",
							nil,
						)

						/* assert */
						Expect(fakeOS.RemoveAllArgsForCall(0)).To(Equal(expectedPath))
						Expect(actualError).To(Equal(expectedError))
					})
				})
				Context("err.Error() returns transport.ErrAuthorizationFailed error", func() {
					It("should call fs.RemoveAll w/ expected args & return expected error", func() {

						/* arrange */
						providedPath := "dummyPath"

						ref := &Ref{
							Name:    "dummyPkgRef",
							Version: "0.0.0",
						}

						fakeRefParser := new(fakeRefParser)
						fakeRefParser.ParseReturns(ref, nil)

						expectedPath := filepath.Join(
							providedPath,
							fmt.Sprintf("%v#%v", ref.Name, ref.Version),
						)

						fakeOS := new(ios.Fake)
						expectedError := model.ErrPkgPullAuthorization{}

						fakeGit := new(igit.Fake)
						fakeGit.PlainCloneReturns(nil, transport.ErrAuthorizationFailed)

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        fakeOS,
							refParser: fakeRefParser,
						}

						/* act */
						actualError := objectUnderTest.Pull(
							providedPath,
							"dummyPkgRef",
							nil,
						)

						/* assert */
						Expect(fakeOS.RemoveAllArgsForCall(0)).To(Equal(expectedPath))
						Expect(actualError).To(Equal(expectedError))
					})
				})
				Context("err.Error() returns other error", func() {
					It("should call fs.RemoveAll w/ expected args & return error", func() {

						/* arrange */
						providedPath := "dummypath"
						ref := &Ref{
							Name:    "dummyPkgRef",
							Version: "0.0.0",
						}

						fakeRefParser := new(fakeRefParser)
						fakeRefParser.ParseReturns(ref, nil)

						fakeOS := new(ios.Fake)

						expectedError := errors.New("dummyError")

						fakeGit := new(igit.Fake)
						fakeGit.PlainCloneReturns(nil, expectedError)

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        fakeOS,
							refParser: fakeRefParser,
						}

						/* act */
						actualError := objectUnderTest.Pull(
							providedPath,
							"dummyPkgRef",
							nil,
						)

						/* assert */
						Expect(fakeOS.RemoveAllArgsForCall(0)).To(Equal(ref.ToPath(providedPath)))
						Expect(actualError).To(Equal(expectedError))
					})
				})
			})
			Context("git.PlainClone doesn't error", func() {
				It("shouldn't err", func() {
					/* arrange */
					fakeRefParser := new(fakeRefParser)
					fakeRefParser.ParseReturns(&Ref{}, nil)

					objectUnderTest := _puller{
						git:       new(igit.Fake),
						os:        new(ios.Fake),
						refParser: fakeRefParser,
					}

					/* act */
					actualErr := objectUnderTest.Pull(
						"dummyPath",
						"dummyPkgRef",
						nil,
					)

					/* assert */
					Expect(actualErr).To(BeNil())
				})

				It("should remove pkg '.git' sub dir & return errors", func() {

					/* arrange */
					providedPath := "dummypath"

					ref := &Ref{
						Name:    "dummyPkgRef",
						Version: "0.0.0",
					}

					fakeRefParser := new(fakeRefParser)
					fakeRefParser.ParseReturns(ref, nil)

					expectedPath := filepath.Join(ref.ToPath(providedPath), ".git")

					fakeOS := new(ios.Fake)
					expectedError := errors.New("dummyError")
					fakeOS.RemoveAllReturns(expectedError)

					objectUnderTest := _puller{
						git:       new(igit.Fake),
						os:        fakeOS,
						refParser: fakeRefParser,
					}

					/* act */
					actualError := objectUnderTest.Pull(
						providedPath,
						"dummyPkgRef",
						nil,
					)

					/* assert */
					Expect(fakeOS.RemoveAllArgsForCall(0)).To(Equal(expectedPath))
					Expect(actualError).To(Equal(expectedError))
				})
				Context("puller.Pull called in parallel w/ same args", func() {
					It("should not call git.PlainClone & return same result", func() {
						/* arrange */

						fakeGit := new(igit.Fake)
						fakeGit.PlainCloneStub = func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
							// ensure go routing has time to overlap
							<-time.After(100 * time.Millisecond)
							return nil, nil
						}

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        new(ios.Fake),
							refParser: _refParser{},
						}

						pkgRef := "dummyPkgName0#1.0.0"

						go objectUnderTest.Pull(
							"dummyPath",
							pkgRef,
							nil,
						)

						/* act */
						actualErr := objectUnderTest.Pull(
							"dummyPath",
							pkgRef,
							nil,
						)

						/* assert */
						Expect(fakeGit.PlainCloneCallCount()).To(Equal(1))
						Expect(actualErr).To(BeNil())
					})
				})
				Context("puller.Pull called in parallel w/ different args", func() {
					It("should call git.PlainClone", func() {
						/* arrange */
						expectedView := &model.PkgManifest{Name: "dummyName"}
						expectedErr := errors.New("dummyError")

						fakeGit := new(igit.Fake)

						fakeManifest := new(manifest.Fake)
						fakeManifest.UnmarshalReturns(expectedView, expectedErr)

						objectUnderTest := _puller{
							git:       fakeGit,
							os:        new(ios.Fake),
							refParser: _refParser{},
						}

						objectUnderTest.Pull(
							"dummyPath",
							"dummyPkgName0#1.0.0",
							nil,
						)

						/* act */
						actualErr := objectUnderTest.Pull(
							"dummyPath",
							"dummyPkgName1#1.0.0",
							nil,
						)

						/* assert */
						Expect(fakeGit.PlainCloneCallCount()).To(Equal(2))
						Expect(actualErr).To(BeNil())
					})
				})
			})
		})
	})
})
