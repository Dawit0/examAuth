package service

import "fmt"

type TestMailer struct {
	FixedOTP string
}

func NewTestMailer(otp string) *TestMailer {
	return &TestMailer{
		FixedOTP: otp,
	}
}

func (m *TestMailer) SendMail(email, otp string) error {
	fmt.Println("========== TEST MAILER ==========")
	fmt.Println("To:      ", email)
	fmt.Println("OTP:     ", otp)
	fmt.Println("Message: This is a test OTP email (not sent).")
	fmt.Println("=================================")
	return nil
}
