package envstruct_test

import (
	"envstruct"
	"os"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestStruct struct {
	StringThing   string `env:"string_thing"`
	RequiredThing string `env:"required_thing,required"`

	BoolThing bool `env:"bool_thing"`

	IntThing    int    `env:"int_thing"`
	Int8Thing   int8   `env:"int8_thing"`
	Int16Thing  int16  `env:"int16_thing"`
	Int32Thing  int32  `env:"int32_thing"`
	Int64Thing  int64  `env:"int64_thing"`
	UintThing   uint   `env:"uint_thing"`
	Uint8Thing  uint8  `env:"uint8_thing"`
	Uint16Thing uint16 `env:"uint16_thing"`
	Uint32Thing uint32 `env:"uint32_thing"`
	Uint64Thing uint64 `env:"uint64_thing"`

	StringSliceThing []string `env:"string_slice_thing"`
	IntSliceThing    []int    `env:"int_slice_thing"`
}

var _ = Describe("envstruct", func() {
	Describe("Load()", func() {
		var (
			ts        TestStruct
			loadError error
			envVars   map[string]string
		)

		BeforeEach(func() {
			envVars = map[string]string{
				"STRING_THING":       "stringy thingy",
				"REQUIRED_THING":     "im so required",
				"BOOL_THING":         "true",
				"INT_THING":          "100",
				"INT8_THING":         "20",
				"INT16_THING":        "2000",
				"INT32_THING":        "200000",
				"INT64_THING":        "200000000",
				"UINT_THING":         "100",
				"UINT8_THING":        "20",
				"UINT16_THING":       "2000",
				"UINT32_THING":       "200000",
				"UINT64_THING":       "200000000",
				"STRING_SLICE_THING": "one,two,three",
				"INT_SLICE_THING":    "1,2,3",
			}
		})

		JustBeforeEach(func() {
			for k, v := range envVars {
				os.Setenv(k, v)
			}
		})

		Context("when load is successfull", func() {
			JustBeforeEach(func() {
				loadError = envstruct.Load(&ts)
			})

			AfterEach(func() {
				for k := range envVars {
					os.Setenv(k, "")
				}
			})

			It("does not return an error", func() {
				Expect(loadError).ToNot(HaveOccurred())
			})

			Context("with strings", func() {
				It("populates the string thing", func() {
					Expect(ts.StringThing).To(Equal("stringy thingy"))
				})
			})

			Context("with bools", func() {
				Context("with 'true'", func() {
					It("is true", func() {
						Expect(ts.BoolThing).To(BeTrue())
					})
				})

				Context("with 'false'", func() {
					BeforeEach(func() {
						envVars["BOOL_THING"] = "false"
					})

					It("is true", func() {
						Expect(ts.BoolThing).To(BeFalse())
					})
				})

				Context("with '1'", func() {
					BeforeEach(func() {
						envVars["BOOL_THING"] = "1"
					})

					It("is true", func() {
						Expect(ts.BoolThing).To(BeTrue())
					})
				})

				Context("with '0'", func() {
					BeforeEach(func() {
						envVars["BOOL_THING"] = "0"
					})

					It("is false", func() {
						Expect(ts.BoolThing).To(BeFalse())
					})
				})
			})

			Context("with ints", func() {
				It("populates the int thing", func() {
					Expect(ts.IntThing).To(Equal(100))
				})

				It("populates the int 8 thing", func() {
					Expect(ts.Int8Thing).To(Equal(int8(20)))
				})

				It("populates the int 16 thing", func() {
					Expect(ts.Int16Thing).To(Equal(int16(2000)))
				})

				It("populates the int 32 thing", func() {
					Expect(ts.Int32Thing).To(Equal(int32(200000)))
				})

				It("populates the int 64 thing", func() {
					Expect(ts.Int64Thing).To(Equal(int64(200000000)))
				})
			})

			Context("with uints", func() {
				It("populates the uint thing", func() {
					Expect(ts.UintThing).To(Equal(uint(100)))
				})

				It("populates the uint 8 thing", func() {
					Expect(ts.Uint8Thing).To(Equal(uint8(20)))
				})

				It("populates the uint 16 thing", func() {
					Expect(ts.Uint16Thing).To(Equal(uint16(2000)))
				})

				It("populates the uint 32 thing", func() {
					Expect(ts.Uint32Thing).To(Equal(uint32(200000)))
				})

				It("populates the uint 64 thing", func() {
					Expect(ts.Uint64Thing).To(Equal(uint64(200000000)))
				})
			})

			Context("with comma separated strings", func() {
				Context("slice of strings", func() {
					It("populates a slice of strings", func() {
						Expect(ts.StringSliceThing).To(Equal([]string{"one", "two", "three"}))
					})

					Context("with leading and trailing spaces", func() {
						BeforeEach(func() {
							envVars["STRING_SLICE_THING"] = "one , two , three"
						})

						It("populates a slice of strings", func() {
							Expect(ts.StringSliceThing).To(Equal([]string{"one", "two", "three"}))
						})
					})
				})

				Context("slice of ints", func() {
					It("populates a slice of ints", func() {
						Expect(ts.IntSliceThing).To(Equal([]int{1, 2, 3}))
					})
				})
			})
		})

		Context("when load is unsuccessfull", func() {
			Context("when a required environment variable is not given", func() {
				BeforeEach(func() {
					envVars["REQUIRED_THING"] = ""
				})

				It("returns a validation error", func() {
					loadError = envstruct.Load(&ts)

					Expect(loadError).To(MatchError(fmt.Errorf("REQUIRED_THING is required but was empty")))
				})
			})

			Context("with an invalid int", func() {
				BeforeEach(func() {
					envVars["INT_THING"] = "Hello!"
				})

				It("returns an error", func() {
					Expect(envstruct.Load(&ts)).ToNot(Succeed())
				})
			})

			Context("with an invalid uint", func() {
				BeforeEach(func() {
					envVars["UINT_THING"] = "Hello!"
				})

				It("returns an error", func() {
					Expect(envstruct.Load(&ts)).ToNot(Succeed())
				})
			})
		})
	})
})
