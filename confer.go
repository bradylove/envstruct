package confer

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Load(t interface{}) error {
	val := reflect.ValueOf(t).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		envVal := os.Getenv(strings.ToUpper(tag.Get("env")))

		setField(valueField, envVal)
	}

	return nil
}

func setField(value reflect.Value, input string) error {
	switch value.Kind() {
	case reflect.String:
		value.SetString(input)
	case reflect.Bool:
		value.SetBool(input == "true" || input == "1")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(input, 10, 64)
		_ = err
		value.SetInt(int64(n))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(input, 10, 64)
		_ = err
		value.SetUint(uint64(n))
	case reflect.Slice:
		inputs := extractSliceInputs(input)

		rs := reflect.MakeSlice(value.Type(), len(inputs), len(inputs))
		for i, val := range inputs {
			setField(rs.Index(i), val)
		}

		value.Set(rs)
	}

	return nil
}

func extractSliceInputs(input string) []string {
	inputs := strings.Split(input, ",")

	for i, v := range inputs {
		inputs[i] = strings.TrimSpace(v)
	}

	return inputs
}
