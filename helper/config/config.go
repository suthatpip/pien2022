package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const Price float64 = 49

type Env struct {
	OWNER       string `env-default:"piennews"`
	ENVIRONMENT string `env:"ENVIRONMENT" env-default:"http://piennews001.thddns.net:3030"`
	SECRET      string `env:"SECRET" env-default:"secert"`
}

var cfg Env

func init() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("Invalid env")
	}
}

func GetENV() Env {
	return cfg
}
