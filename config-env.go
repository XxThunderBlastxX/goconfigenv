package config_env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
)

// Tags : env
// Example : `env:"NAME"`
// Validations :-
// required - checks if the environment variable is present and compulsory
// default - sets the default value if the environment variable if not present
// omitempty - ignores the field if the environment variable is not present

func ParseEnv[T any]() T {
	_ = godotenv.Load()

	// Checking if the type is a struct
	if reflect.TypeFor[T]().Kind() != reflect.Struct {
		panic("type is not a struct")
	}

	// Creating a new instance of the struct
	instance := reflect.New(reflect.TypeFor[T]()).Elem()

	// Looping through the fields of the struct
	for i := 0; i < reflect.TypeFor[T]().NumField(); i++ {
		tag := reflect.TypeFor[T]().Field(i).Tag.Get("env")

		// Validating the environment variable
		_ = ValidateEnv[T]()

		envValue := os.Getenv(tag)

		// Setting the value of the field
		instance.Field(i).SetString(envValue)
	}

	return instance.Interface().(T)
}

func ValidateEnv[T any]() error {
	for i := 0; i < reflect.TypeFor[T]().NumField(); i++ {
		tag := reflect.TypeFor[T]().Field(i).Tag.Get("env")

		// Checking if the environment variable exists
		if _, ok := os.LookupEnv(tag); !ok {
			return fmt.Errorf("environment variable %s not found", tag)
		}
	}

	return nil
}

// func validateType[T any]() error {}
