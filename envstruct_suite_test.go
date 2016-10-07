//go:generate hel

package envstruct_test

import (
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	baseEnvVars = map[string]string{
		"STRING_THING":         "stringy thingy",
		"REQUIRED_THING":       "im so required",
		"BOOL_THING":           "true",
		"INT_THING":            "100",
		"INT8_THING":           "20",
		"INT16_THING":          "2000",
		"INT32_THING":          "200000",
		"INT64_THING":          "200000000",
		"UINT_THING":           "100",
		"UINT8_THING":          "20",
		"UINT16_THING":         "2000",
		"UINT32_THING":         "200000",
		"UINT64_THING":         "200000000",
		"STRING_SLICE_THING":   "one,two,three",
		"INT_SLICE_THING":      "1,2,3",
		"DURATION_THING":       "2s",
		"URL_THING":            "http://github.com/some/path",
		"UNMARSHALLER_POINTER": "pointer",
		"UNMARSHALLER_VALUE":   "value",
		"CaSe_SeNsItIvE_ThInG": "case sensitive",
	}
)

type LargeTestStruct struct {
	NonEnvThing        string
	DefaultThing       string `env:"DEFAULT_THING"`
	StringThing        string `env:"STRING_THING"`
	RequiredThing      string `env:"REQUIRED_THING,noreport,required"`
	CaseSensitiveThing string `env:"CaSe_SeNsItIvE_ThInG"`

	BoolThing bool `env:"BOOL_THING"`

	IntThing    int    `env:"INT_THING"`
	Int8Thing   int8   `env:"INT8_THING"`
	Int16Thing  int16  `env:"INT16_THING"`
	Int32Thing  int32  `env:"INT32_THING"`
	Int64Thing  int64  `env:"INT64_THING"`
	UintThing   uint   `env:"UINT_THING"`
	Uint8Thing  uint8  `env:"UINT8_THING"`
	Uint16Thing uint16 `env:"UINT16_THING"`
	Uint32Thing uint32 `env:"UINT32_THING"`
	Uint64Thing uint64 `env:"UINT64_THING"`

	StringSliceThing []string `env:"STRING_SLICE_THING"`
	IntSliceThing    []int    `env:"INT_SLICE_THING"`

	DurationThing time.Duration `env:"DURATION_THING"`
	URLThing      *url.URL      `env:"URL_THING"`

	UnmarshallerPointer *mockUnmarshaller `env:"UNMARSHALLER_POINTER"`
	UnmarshallerValue   mockUnmarshaller  `env:"UNMARSHALLER_VALUE"`
}

type SmallTestStruct struct {
	HiddenThing        string   `env:"HIDDEN_THING,noreport"`
	StringThing        string   `env:"STRING_THING"`
	BoolThing          bool     `env:"BOOL_THING"`
	IntThing           int      `env:"INT_THING"`
	URLThing           *url.URL `env:"URL_THING"`
	StringSliceThing   []string `env:"STRING_SLICE_THING"`
	CaseSensitiveThing string   `env:"CaSe_SeNsItIvE_ThInG"`
}

func TestEnvstruct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Envstruct Suite")
}
