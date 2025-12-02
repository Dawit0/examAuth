package service

import (
	"github.com/Dawit0/examAuth/internal/domain"
	repo "github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService(userrp *repo.UserRepo) *UserService {
	return &UserService{UserRepo: userrp}
}

func (uc *UserService) CreateUser(user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, nil
	}
	return uc.UserRepo.CreateUser(user)
}
