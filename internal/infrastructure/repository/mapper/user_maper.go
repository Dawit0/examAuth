package mapper

import (
	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/repository/model"
)

func MapDomainToModel(user domain.User) (*model.UserModel, error) {
	email := user.Email()
	badge := user.Badge()
	isActive := user.IsActive()
	score := user.Score()
	return &model.UserModel{
		ID:        user.ID(),
		Username:  user.Username(),
		Phone:     user.Phone(),
		Email:     &email,
		Password:  user.Password(),
		CreatedAt: user.CreatedAt(),
		IsActive:  &isActive,
		Badge:     &badge,
		Score:     &score,
	}, nil
}

func MapModelToDomain(model model.UserModel) (*domain.User, error) {
	domain_val, err := domain.NewUser(model.Email, model.Password, model.Badge, model.Username, model.Phone, model.IsActive, model.Score)
	if err != nil {
		return nil, err
	}

	domain_val.Id_Set(model.ID)
	return domain_val, nil
}
