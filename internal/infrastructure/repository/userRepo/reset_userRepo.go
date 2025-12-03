package userRepo

import (
	"time"

	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/mapper"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
	"gorm.io/gorm"
)

type ResetUserRepo struct {
	DB *gorm.DB
}

func NewResetUserRepo(db *gorm.DB) *ResetUserRepo {
	db.AutoMigrate(&model.PasswordResetModel{}, &model.UserModel{})
	return &ResetUserRepo{DB: db}
}

func (ur *ResetUserRepo) GetByPhone(phone string) (*domain.User, error) {
	var user model.UserModel
	err := ur.DB.Model(&model.UserModel{}).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}

	val, err := mapper.MapModelToDomain(user)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (ur *ResetUserRepo) UpdateUserPassword(id uint, hashpass string) error {
	err := ur.DB.Model(&model.UserModel{}).Where("id = ?", id).Update("password", hashpass).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *ResetUserRepo) CreateUserResetPassword(phone string, UserID uint, otp string, expiredAt time.Time) error {
	models := model.PasswordResetModel{
		Phone:     phone,
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

func (ur *ResetUserRepo) GetPasswordResetByPhoneAndOTP(phone string, otp string) (*domain.ForgetPassword, error) {
	var models model.PasswordResetModel
	err := ur.DB.Model(&model.PasswordResetModel{}).Where("phone = ? AND otp = ? AND used = false", phone, otp, false).First(&models).Error
	if err != nil {
		return nil, err
	}

	if models.ExpiresAt.Before(time.Now()) {
		return nil, nil
	}

	val, err := domain.NewForgetPassword(models.UserID, models.Phone, models.OTP, models.ExpiresAt, models.Used)
	if err != nil {
		return nil, err
	}

	val.Set_Id(models.ID)

	return val, nil
}

func (r *ResetUserRepo) MarkPasswordResetUsed(id uint) error {
	return r.DB.Model(&model.PasswordResetModel{}).Where("id = ?", id).Update("used", true).Error
}

// optional: invalidate all previous resets for a phone (on new request or after used)
func (r *ResetUserRepo) InvalidatePasswordResetsByPhone(phone string) error {
	return r.DB.Model(&model.PasswordResetModel{}).Where("phone = ?", phone).Update("used", true).Error
}
