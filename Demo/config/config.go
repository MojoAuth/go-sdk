package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Port    int    `env:"PORT" envDefault:"8081"`
	BaseURL string `env:"BASEURL" envDefault:"http://localhost:8081"`
	APIKey  string `env:"APIKEY" envDefault:""`
	ENV     string `env:"ENV" envDefault:"Development"`
}

var App config

func init() {
	// godotenv.Load("./public.env")
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	App = cfg
	fmt.Printf("%+v\n", cfg)
}
