package godotenvstruct

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Bind(prefix string, s interface{}) error {
	var missing []string
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("cfg must be a non-nil pointer to struct")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("env")
		if tag == "" {
			tag = t.Name() + "_" + field.Name
		}
		str, err := GetEnv(prefix + tag)
		if err != nil {
			missing = append(missing, err.Error())
			continue
		}

		err = setFieldValue(value, str)
		if err != nil {
			missing = append(missing, err.Error())
		}
	}
	m := strings.Join(missing, "\n")
	if m != "" {
		return errors.New("some error occured: \n" + m)
	}
	return nil
}

func GetEnv(key string) (string, error) {
	str := os.Getenv(key)
	if str == "" {
		return "", errors.New("- " + strings.Replace(key, "_", ".", -1) + " can't be null or empty")
	}
	return str, nil
}

func setFieldValue(value reflect.Value, str string) error {
	if value.Type() == reflect.TypeOf(time.Duration(0)) {
		dur, err := time.ParseDuration(str)
		if err != nil {
			return err
		}

		value.SetInt(int64(dur))
		return nil
	}

	switch value.Kind() {
	case reflect.String:
		value.SetString(str)
	case reflect.Int, reflect.Int64:
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		value.SetBool(b)
	case reflect.Float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		value.SetFloat(f)
	default:
		return fmt.Errorf("- unsupported field type: %s", value.Kind())
	}

	return nil
}
