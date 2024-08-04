package main

import (
	"fmt"
	configenv "github.com/XxThunderBlastxX/goconfigenv"
)

type AppConfig struct {
	Host        string `env:"HOST,default=localhost"`
	Port        int    `env:"PORT,default=8080"`
	InnerStruct InnerConfig
}

type InnerConfig struct {
	Username string `env:"USERNAME,default=John"`
	Password string `env:"PASSWORD"`
}

func main() {
	config, err := configenv.ParseEnv[AppConfig]()
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Host)
	fmt.Println(config.Port)
	fmt.Println(config.InnerStruct.Username)
	fmt.Println(config.InnerStruct.Password)
}
