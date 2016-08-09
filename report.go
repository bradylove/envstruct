package envstruct

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

var ReportWriter io.Writer = os.Stdout

func WriteReport(t interface{}) error {
	w := tabwriter.NewWriter(ReportWriter, 0, 8, 2, ' ', 0)

	fmt.Fprintln(w, "FIELD NAME:\tTYPE:\tENV:\tVALUE:")

	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		tagProperties := extractSliceInputs(tag.Get("env"))
		envVar := strings.ToUpper(tagProperties[indexEnvVar])

		fmt.Fprintln(w, fmt.Sprintf("%v\t%v\t%v\t%v", typeField.Name, valueField.Type(), envVar, valueField))
	}

	return w.Flush()
}
