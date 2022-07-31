package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const Price float64 = 49
const Token_expire int64 = 86400 * 30

type Env struct {
	OWNER       string `env-default:"piennews"`
	URL         string `env:"URL" env-default:"http://piennews001.thddns.net:3030"`
	SECRET      string `env:"SECRET" env-default:"secert"`
	ENVIRONMENT string `env:"ENVIRONMENT" env-default:"dev"`
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
