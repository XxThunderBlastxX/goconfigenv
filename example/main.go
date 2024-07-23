package main

import (
	"fmt"
	configenv "github.com/XxThunderBlastxX/config-env"
)

// Example usage
type Config struct {
	Port int `env:"APP_PORT"`
}

func main() {
	config, err := configenv.ParseEnv[Config]()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	fmt.Printf("Loaded config: %+v\n", config)

}
