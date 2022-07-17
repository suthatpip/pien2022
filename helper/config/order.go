package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type OrderStatus struct {
	INITIAL_ORDER     string `env-default:"S-000"`
	APPROVED          string `env-default:"S-001"`
	PENDING_PAYMENT   string `env-default:"S-002"`
	PAYMENT_COMPLETED string `env-default:"S-003"`
	ON_PROCESS        string `env-default:"S-004"`
	PUBLISH           string `env-default:"S-005"`
	FAILED            string `env-default:"S-006"`
	CANCELED          string `env-default:"S-007"`
	VALIDATE_PAYMENT  string `env-default:"S-008"`
	COMPLETE          string `env-default:"S-009"`
}

var orderStatus OrderStatus

func init() {
	cleanenv.ReadEnv(&orderStatus)

}

func GetOrderStatus() OrderStatus {
	return orderStatus
}
