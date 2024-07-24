# GoConfigEnv
[![GoDoc](https://godoc.org/github.com/alexliesenfeld/go-config-env?status.svg)](https://godoc.org/github.com/alexliesenfeld/go-config-env)

GoConfigEnv is a Go library that allows you to easily load configuration values from environment variables and map them to a struct.

## 💉 Installation

```bash
  go get -u github.com/XxThunderBlastxX/goconfigenv
```

## 🛠️ Usage

Add all your application configuration in your `.env` file:

```bash
  PORT=8080
  HOST=localhost
  DEBUG=true
```

Create a struct that represents your configuration and just simply call the function `ParseEnv` and pass the struct type as its generic type. The library will automatically map the values from the environment variables and return the struct. :

```go
  package main

  import (
    "fmt"
    "github.com/XxThunderBlastxX/goconfigenv"
  )

  type Config struct {
    Port  int    `env:"PORT" `
    Host  string `env:"HOST"`
    Debug bool   `env:"DEBUG"`
  }

  func main() {
    config, err := goconfigenv.ParseEnv[Config]()
    if err != nil {
      fmt.Println(err)
    }

	fmt.Println(config.Port)
	fmt.Println(config.Host)
	fmt.Println(config.Debug)
  }
```

You can also configure your tags by setting the default value to the field. It takes the default value if the environment variable is not set:

```go
  type Config struct {
    Port  int    `env:"PORT,default=8080"`
    Host  string `env:"HOST,default=localhost"`
    Debug bool   `env:"DEBUG,default=true"`
  }
```

By default it takes the values from `.env` file at root of the dir. If you want to change this behaviour you can set the custom path and file location using [godotenv](https://github.com/joho/godotenv)

```go
    err := godotenv.Load("path/to/your/.env")
```

## 📄 License
This project is licensed under the MIT License - see the [LICENCE](LICENCE) file for details.

## 🤝 Show your support
Give a ⭐️ if this project helped you!

## 🤝 Contribution
Contributions, issues and feature requests are always welcome!

Feel free to check [issues page](https://github.com/XxThunderBlastxX/goconfigenv).



