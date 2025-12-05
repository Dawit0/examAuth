package service

import (
	"fmt"
	"net/smtp"
)

type GmailMailer struct {
	Email       string
	AppPassword string
}

func NewGmailMailer(email, appPassword string) *GmailMailer {
	return &GmailMailer{
		Email:       email,
		AppPassword: appPassword,
	}
}

func (m *GmailMailer) SendMail(toEmail, otp string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // TLS

	auth := smtp.PlainAuth("", m.Email, m.AppPassword, smtpHost)

	subject := "Your OTP Code"
	body := fmt.Sprintf("Your One-Time Password (OTP) is: %s\nThis code expires in 5 minutes.", otp)

	// SMTP Email Format
	message := []byte(
		"From: " + m.Email + "\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body +
			"\r\n",
	)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, m.Email, []string{toEmail}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully to:", toEmail)
	return nil
}
