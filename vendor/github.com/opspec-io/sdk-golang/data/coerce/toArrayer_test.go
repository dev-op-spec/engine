package coerce

import (
	"errors"
	"fmt"
	"github.com/golang-interfaces/encoding-ijson"
	"github.com/golang-interfaces/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"reflect"
)

var _ = Context("toArrayer", func() {
	Context("ToArray", func() {
		Context("Value is nil", func() {
			It("should return expected result", func() {
				/* arrange */
				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(nil)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(BeNil())
			})
		})
		Context("Value.Array isn't nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedValue := &model.Value{
					Array: []interface{}{"dummyItem"},
				}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(Equal(providedValue))
				Expect(actualErr).To(BeNil())
			})
		})
		Context("Value.Boolean isn't nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedBoolean := true
				providedValue := &model.Value{
					Boolean: &providedBoolean,
				}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce boolean '%v' to array; incompatible types", providedBoolean)))
			})
		})
		Context("Value.Dir isn't nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedDir := "dummyValue"
				providedValue := &model.Value{
					Dir: &providedDir,
				}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce dir '%v' to array; incompatible types", providedDir)))
			})
		})
		Context("Value.File isn't nil", func() {
			It("should call ioutil.ReadFile w/ expected args", func() {
				/* arrange */
				providedFile := "dummyFile"

				providedValue := &model.Value{
					File: &providedFile,
				}

				fakeIOUtil := new(iioutil.Fake)
				// err to trigger immediate return
				fakeIOUtil.ReadFileReturns(nil, errors.New("dummyError"))

				fileUnderTest := _toArrayer{
					ioUtil: fakeIOUtil,
				}

				/* act */
				fileUnderTest.ToArray(providedValue)

				/* assert */
				Expect(fakeIOUtil.ReadFileArgsForCall(0)).To(Equal(providedFile))
			})
			Context("ioutil.ReadFile errs", func() {
				It("should return expected result", func() {
					/* arrange */
					fakeIOUtil := new(iioutil.Fake)

					marshalErr := errors.New("dummyError")
					fakeIOUtil.ReadFileReturns(nil, marshalErr)

					fileUnderTest := _toArrayer{
						ioUtil: fakeIOUtil,
					}

					/* act */
					actualValue, actualErr := fileUnderTest.ToArray(
						&model.Value{File: new(string)},
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce file to array; error was %v", marshalErr.Error())))
				})
			})
			Context("ioutil.ReadFile doesn't err", func() {
				It("should call json.Unmarshal w/ expected args", func() {
					/* arrange */
					fakeIOUtil := new(iioutil.Fake)

					marshaledBytes := []byte{2, 3, 5}
					fakeIOUtil.ReadFileReturns(marshaledBytes, nil)

					fakeJSON := new(ijson.Fake)
					// err to trigger immediate return
					fakeJSON.UnmarshalReturns(errors.New("dummyError"))

					arrayUnderTest := _toArrayer{
						ioUtil: fakeIOUtil,
						json:   fakeJSON,
					}

					/* act */
					arrayUnderTest.ToArray(
						&model.Value{File: new(string)},
					)

					/* assert */
					actualBytes,
						_ := fakeJSON.UnmarshalArgsForCall(0)
					Expect(actualBytes).To(Equal(marshaledBytes))
				})
				Context("json.Unmarshal errs", func() {
					It("should return expected result", func() {
						/* arrange */

						fakeJSON := new(ijson.Fake)

						unmarshalError := errors.New("dummyError")
						fakeJSON.UnmarshalReturns(unmarshalError)

						arrayUnderTest := _toArrayer{
							ioUtil: new(iioutil.Fake),
							json:   fakeJSON,
						}

						/* act */
						actualValue, actualErr := arrayUnderTest.ToArray(
							&model.Value{File: new(string)},
						)

						/* assert */
						Expect(actualValue).To(BeNil())
						Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce file to array; error was %v", unmarshalError.Error())))
					})
				})
				Context("json.Unmarshal doesn't err", func() {
					It("should return expected result", func() {
						/* arrange */
						fakeJSON := new(ijson.Fake)

						arrayItem := "dummyMapValue"
						expectedValue := model.Value{Array: []interface{}{arrayItem}}

						fakeJSON.UnmarshalStub = func(data []byte, v interface{}) error {
							reflect.ValueOf(v).Elem().Set(reflect.ValueOf([]interface{}{arrayItem}))
							return nil
						}

						arrayUnderTest := _toArrayer{
							ioUtil: new(iioutil.Fake),
							json:   fakeJSON,
						}

						/* act */
						actualValue, actualErr := arrayUnderTest.ToArray(
							&model.Value{File: new(string)},
						)

						/* assert */
						Expect(*actualValue).To(Equal(expectedValue))
						Expect(actualErr).To(BeNil())
					})
				})
			})
		})
		Context("Value.Number isn't nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedNumber := 2.2
				providedValue := &model.Value{
					Number: &providedNumber,
				}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce number '%v' to array; incompatible types", providedNumber)))
			})
		})
		Context("Value.Socket isn't nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedSocket := "dummySocket"
				providedValue := &model.Value{
					Socket: &providedSocket,
				}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce socket '%v' to array; incompatible types", providedSocket)))
			})
		})
		Context("Value.String isn't nil", func() {
			It("should call json.Unmarshal w/ expected args", func() {
				/* arrange */
				providedString := "{}"

				providedValue := &model.Value{
					String: &providedString,
				}

				fakeJSON := new(ijson.Fake)
				// err to trigger immediate return
				fakeJSON.UnmarshalReturns(errors.New("dummyError"))

				arrayUnderTest := _toArrayer{
					json: fakeJSON,
				}

				/* act */
				arrayUnderTest.ToArray(providedValue)

				/* assert */
				actualBytes,
					_ := fakeJSON.UnmarshalArgsForCall(0)
				Expect(actualBytes).To(Equal([]byte(providedString)))
			})
			Context("json.Unmarshal errs", func() {
				It("should return expected result", func() {
					/* arrange */

					fakeJSON := new(ijson.Fake)

					unmarshalError := errors.New("dummyError")
					fakeJSON.UnmarshalReturns(unmarshalError)

					arrayUnderTest := _toArrayer{
						json: fakeJSON,
					}

					/* act */
					actualValue, actualErr := arrayUnderTest.ToArray(
						&model.Value{String: new(string)},
					)

					/* assert */
					Expect(actualValue).To(BeNil())
					Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce string to array; error was %v", unmarshalError.Error())))
				})
			})
			Context("json.Unmarshal doesn't err", func() {
				It("should return expected result", func() {
					/* arrange */
					fakeJSON := new(ijson.Fake)

					arrayItem := "dummyMapValue"
					expectedValue := model.Value{Array: []interface{}{arrayItem}}

					fakeJSON.UnmarshalStub = func(data []byte, v interface{}) error {
						reflect.ValueOf(v).Elem().Set(reflect.ValueOf([]interface{}{arrayItem}))
						return nil
					}

					arrayUnderTest := _toArrayer{
						json: fakeJSON,
					}

					/* act */
					actualValue, actualErr := arrayUnderTest.ToArray(
						&model.Value{String: new(string)},
					)

					/* assert */
					Expect(*actualValue).To(Equal(expectedValue))
					Expect(actualErr).To(BeNil())
				})
			})
		})
		Context("Value.Dir,File,Number,Array,String nil", func() {
			It("should return expected result", func() {
				/* arrange */
				providedValue := &model.Value{}

				arrayUnderTest := _toArrayer{}

				/* act */
				actualValue, actualErr := arrayUnderTest.ToArray(providedValue)

				/* assert */
				Expect(actualValue).To(BeNil())
				Expect(actualErr).To(Equal(fmt.Errorf("unable to coerce '%+v' to array", providedValue)))
			})
		})
	})
})
