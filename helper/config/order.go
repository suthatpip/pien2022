package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type OrderStatus struct {
	INITIAL_ORDER     string `env-default:"0"`
	APPROVED          string `env-default:"1"`
	PENDING_PAYMENT   string `env-default:"2"`
	PAYMENT_COMPLETED string `env-default:"3"`
	ON_PROCESS        string `env-default:"4"`
	PUBLISH           string `env-default:"5"`
	FAILED            string `env-default:"6"`
	CANCELED          string `env-default:"7"`
}

var orderStatus OrderStatus

func init() {
	cleanenv.ReadEnv(&orderStatus)

}

func GetOrderStatus() OrderStatus {
	return orderStatus
}
