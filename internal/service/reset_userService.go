package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
)

type SendMailer interface {
	SendMail(email string, otp string) error
}

type ResetUserService struct {
	repo   *userRepo.ResetUserRepo
	mailer SendMailer
	otpTTl time.Duration
}

func NewResetUserService(repo *userRepo.ResetUserRepo, mailer SendMailer) *ResetUserService {
	return &ResetUserService{repo: repo, mailer: mailer, otpTTl: 15 * time.Minute}
}

func (r *ResetUserService) RequestResetPasswordEmail(email string) error {
	user, err := r.repo.GetByEmail(email)
	if err != nil {
		return err
	}

	otp, err := generateNumericOTP(6)
	if err != nil {
		return err
	}
	expiredAt := time.Now().Add(r.otpTTl)
	err = r.repo.SavePasswordReset(email, user.ID(), otp, expiredAt)
	if err != nil {
		return err
	}
	return r.mailer.SendMail(email, otp)
}

func (r *ResetUserService) ResetPassword(email string, otp string, newPassword string) error {
	Token, err := r.repo.FindValidResetByEmailAndOTP(email, otp)
	if err != nil {
		return err
	}
	if Token == nil || Token.Used() {
		return errors.New("invalid reset token")
	}
	err = r.repo.MarkPasswordResetUsed(Token.Id())
	if err != nil {
		return err
	}
	user, err := r.repo.GetByEmail(Token.Email())
	if err != nil {
		return err
	}

	user.SetPassword(newPassword)
	err = r.repo.UpdateUserPassword(user.ID(), newPassword)
	if err != nil {
		return err
	}
	return nil
}

func generateNumericOTP(n int) (string, error) {
	otp := ""
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10)) // random number 0-9
		if err != nil {
			return "", fmt.Errorf("failed to generate OTP: %w", err)
		}
		otp += num.String()
	}
	return otp, nil
}
