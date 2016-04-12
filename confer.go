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

		switch valueField.Kind() {
		case reflect.String:
			valueField.SetString(envVal)
		case reflect.Bool:
			valueField.SetBool(envVal == "true" || envVal == "1")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(envVal, 10, 64)
			_ = err
			valueField.SetInt(int64(n))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n, err := strconv.ParseUint(envVal, 10, 64)
			_ = err
			valueField.SetUint(uint64(n))
		}
	}

	return nil
}
