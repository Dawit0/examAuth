package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
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

	otp := generateOTP(6)
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

func generateOTP(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)[:n]
}
