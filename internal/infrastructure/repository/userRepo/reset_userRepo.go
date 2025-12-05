package userRepo

import (
	"errors"
	"time"

	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/mapper"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ResetUserRepo struct {
	DB *gorm.DB
}

func NewResetUserRepo(db *gorm.DB) *ResetUserRepo {
	db.AutoMigrate(&model.PasswordResetModel{}, &model.UserModel{})
	return &ResetUserRepo{DB: db}
}

func (r *ResetUserRepo) GetByEmail(email string) (*domain.User, error) {
    var u model.UserModel
    if err := r.DB.Where("email = ?", email).First(&u).Error; err != nil {
        return nil, err
    }
    return mapper.MapModelToDomain(u)
}

func (ur *ResetUserRepo) UpdateUserPassword(id uint, hashpass string) error {
	pass, _ := bcrypt.GenerateFromPassword([]byte(hashpass), bcrypt.DefaultCost)
	err := ur.DB.Model(&model.UserModel{}).Where("id = ?", id).Update("password", pass).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *ResetUserRepo) SavePasswordReset(email string, UserID uint, otp string, expiredAt time.Time) error {
	models := model.PasswordResetModel{
		Email:     email,
		UserID:    UserID,
		OTP:       otp,
		ExpiresAt: expiredAt,
		Used:      false,
	}

	err := ur.DB.Model(&model.PasswordResetModel{}).Create(&models).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *ResetUserRepo) FindValidResetByEmailAndOTP(email string, otp string) (*domain.ForgetPassword, error) {
	var models model.PasswordResetModel
	err := ur.DB.Model(&model.PasswordResetModel{}).Where("email = ? AND otp = ? AND used = false", email, otp).First(&models).Error
	if err != nil {
		return nil, err
	}

	if models.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("password reset expired")
	}

	val, err := domain.NewForgetPassword(models.UserID, models.Email, models.OTP, models.ExpiresAt, models.Used)
	if err != nil {
		return nil, errors.New("failed to create forget password")
	}

	val.Set_Id(models.ID)

	return val, nil
}

func (r *ResetUserRepo) MarkPasswordResetUsed(id uint) error {
	return r.DB.Model(&model.PasswordResetModel{}).Where("id = ?", id).Update("used", true).Error
}

// optional: invalidate all previous resets for a phone (on new request or after used)
func (r *ResetUserRepo) InvalidatePasswordResetsByEmail(email string) error {
	return r.DB.Model(&model.PasswordResetModel{}).Where("email = ?", email).Update("used", true).Error
}
