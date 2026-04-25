package sms

import "log"

func Send(phone string, otp string) error {
	log.Println("[SMS MOCK]")
	log.Println("To:", phone)
	log.Println("OTP:", otp)
	return nil
}