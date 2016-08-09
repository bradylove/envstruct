package envstruct_test

import (
	"bytes"

	"github.com/bradylove/envstruct"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Report", func() {
	var (
		ts         TestStruct
		outputText string
	)

	Describe("Report()", func() {
		BeforeEach(func() {
			for k, v := range baseEnvVars {
				os.Setenv(k, v)
			}

			err := envstruct.Load(&ts)
			Expect(err).ToNot(HaveOccurred())

			outputBuffer := bytes.NewBuffer(nil)
			envstruct.ReportWriter = outputBuffer

			err = envstruct.WriteReport(&ts)
			Expect(err).ToNot(HaveOccurred())

			outputText = string(outputBuffer.Bytes())
		})

		It("prints a report of the given envstruct struct", func() {
			Expect(outputText).To(Equal(expectedReportOutput))
		})
	})
})

const (
	expectedReportOutput = `FIELD NAME:       TYPE:          ENV:                VALUE:
NonEnvThing       string
DefaultThing      string         DEFAULT_THING
StringThing       string         STRING_THING        stringy thingy
RequiredThing     string         REQUIRED_THING      im so required
BoolThing         bool           BOOL_THING          true
IntThing          int            INT_THING           100
Int8Thing         int8           INT8_THING          20
Int16Thing        int16          INT16_THING         2000
Int32Thing        int32          INT32_THING         200000
Int64Thing        int64          INT64_THING         200000000
UintThing         uint           UINT_THING          100
Uint8Thing        uint8          UINT8_THING         20
Uint16Thing       uint16         UINT16_THING        2000
Uint32Thing       uint32         UINT32_THING        200000
Uint64Thing       uint64         UINT64_THING        200000000
StringSliceThing  []string       STRING_SLICE_THING  [one two three]
IntSliceThing     []int          INT_SLICE_THING     [1 2 3]
DurationThing     time.Duration  DURATION_THING      2000000000
URLThing          *url.URL       URL_THING           &{http  <nil> github.com /some/path   }
`
)
