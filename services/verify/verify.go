package verify

type serviceInterface interface {
}

type service struct {
}

func NewService() serviceInterface {
	return &service{}
}
