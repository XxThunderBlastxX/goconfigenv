package configenv

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

// Tags : env
// Example : `env:"NAME"`
// Validations :-
// default - sets the default value if the environment variable if not present
// omitempty - ignores the field if the environment variable is not present

func ParseEnv[T any]() (T, error) {
	if err := godotenv.Load(); err != nil {
		return reflect.Zero(reflect.TypeFor[T]()).Interface().(T), err
	}

	// Creating a zero instance of the struct
	var zero T
	// Creating a new instance of the struct
	instance := reflect.New(reflect.TypeFor[T]()).Elem()

	// Checking if the type is a struct
	if err := checkForTypeStruct[T](); err != nil {
		return zero, err
	}

	// Looping through the fields of the struct
	for i := 0; i < reflect.TypeFor[T]().NumField(); i++ {
		tag := reflect.TypeFor[T]().Field(i).Tag.Get("env")
		fieldType := reflect.TypeFor[T]().Field(i).Type

		params := getAllParameters(tag)
		envFieldName := params[0]

		// Checking for "required" param
		if slices.Contains(params, "required") {
			if ok := checkEnvFound(envFieldName); !ok {
				return zero, fmt.Errorf("environment variable %s is not set", envFieldName)
			}
		}

		envValue := os.Getenv(envFieldName)

		if err := setFieldValue(&instance, i, fieldType, envValue); err != nil {
			return zero, err
		}
	}

	return instance.Interface().(T), nil
}

func checkForTypeStruct[T any]() error {
	if reflect.TypeFor[T]().Kind() != reflect.Struct {
		return fmt.Errorf("type is not a struct")
	}

	return nil
}

func getAllParameters(str string) []string {
	return strings.Split(str, ",")
}

func checkEnvFound(env string) bool {
	if _, ok := os.LookupEnv(env); !ok {
		return false
	}

	return true
}

func setFieldValue(instance *reflect.Value, i int, fieldType reflect.Type, envValue string) error {
	switch fieldType.Kind() {
	case reflect.String:
		instance.Field(i).SetString(envValue)
	case reflect.Int:
		intVal, err := strconv.Atoi(envValue)
		if err != nil {
			return err
		}
		instance.Field(i).SetInt(int64(intVal))
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(envValue)
		if err != nil {
			return err
		}
		instance.Field(i).SetBool(boolVal)
	default:
		return fmt.Errorf("unsupported type %s", fieldType.Kind())
	}
	return nil
}
