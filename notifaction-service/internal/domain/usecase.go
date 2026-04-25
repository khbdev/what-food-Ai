package domain

type SMSUsecase interface {
	Handle(body []byte) error
}