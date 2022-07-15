package config

// const (
// 	merchant_secret_key = "Jr9G4UV2iB9btz8V"
// 	merchant_api_key    = "DfQhMg0L"
// 	merchantId          = "33015"
// )
import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Paysolution struct {
	MERCHANT_SECRET_KEY string `env-default:"Jr9G4UV2iB9btz8V"`
	MERCHANT_API_KEY    string `env-default:"DfQhMg0L"`
	MERCHANTID          string `env-default:"33015"`
}

var paysolution Paysolution

func init() {
	cleanenv.ReadEnv(&paysolution)

}

func GetPaysolutionEnv() Paysolution {
	return paysolution
}
