package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const Price float64 = 49

type Env struct {
	Owner       string `env-default:"piennews"`
	Environment string `env:"ENVIRONMENT" env-default:"http://piennews001.thddns.net:3030"`
}

var cfg Env

func init() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("Invalid env")
	}
}

func Get() Env {
	return cfg
}

func GetEnv() string {
	return fmt.Sprintf("%+v", cfg)
}
