package main

import (
	"fmt"
	"github.com/XxThunderBlastxX/config-env"
)

type Server struct {
	Name string `env:"NAME"`
}

func main() {
	fmt.Println(config_env.ParseEnv[Server]())
}
