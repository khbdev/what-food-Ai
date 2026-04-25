package usecase

import (
	"encoding/json"

)

// =====================
// USECASE (PURE BUSINESS)
// =====================

type SMSUsecase struct{}

// DI constructor
func NewSMSUsecase() *SMSUsecase {
	return &SMSUsecase{}
}

// faqat process
func (u *SMSUsecase) Handle(body []byte) error {

	var data mod.SMSOTP

	// parse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// business
	return sms.Send(data.Phone, data.OTP)
}