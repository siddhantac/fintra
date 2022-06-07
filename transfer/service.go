package transfer

type Service struct {
	handler Handler
}

func NewService(handler Handler) Service {
	return Service{handler: handler}
}
