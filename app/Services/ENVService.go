package UserServices

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"strings"
)

type EnvValues struct {
	AppKey  string `json:"APP_KEY"`
	AppMode string `json:"APP_MODE"`
}

func GetEnvValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(strings.ToUpper(key))

	return value
}

func GetEnvVars() EnvValues {

	var envValues EnvValues

	var env = reflect.ValueOf(envValues)

	for i := 0; i < env.NumField(); i++ {
		field := env.Field(i)
		fieldName := env.Type().Field(i).Name
		fieldValue := env.Interface()

		fmt.Printf("%s: %v %v %v\n", fieldName, field, fieldValue, env.Elem())
	}

	return EnvValues{
		AppKey:  "",
		AppMode: "",
	}
}
