package main

import (
	"fmt"
	configenv "github.com/XxThunderBlastxX/config-env"
)

type Config struct {
	Port int `env:"PORT"`
}

func main() {
	//var config Config
	config, err := configenv.ParseEnv[Config]()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	fmt.Printf("Loaded config: %+v\n", config)
}
