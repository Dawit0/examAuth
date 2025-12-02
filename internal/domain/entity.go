package domain

import (
	"regexp"
	"time"
)

type User struct {
	id        uint
	username  string
	phone     string
	email     string
	password  string
	createdAT time.Time
	isActive  bool
	badge     string

	score float64
}

func NewUser(email, password, badge, username, phone string, isactive bool, score int64) (*User, error) {
	if len(email) != 0 {
		regex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

		if !regex.MatchString(email) {
			return nil, ErrInvalidEmail
		}
	}

	if len(username) == 0 {
		return nil, ErrInvalidUsername
	}

	phonregex := regexp.MustCompile(`^(?:\+?251|0)?9\d{8}$`)

	if !phonregex.MatchString(phone) {
		return nil, ErrInvalidPhone
	}

	if len(password) < 4 {
		return nil, ErrInvalidPassword
	}

	if score == 0 {
		score = 0
	}
	nows := time.Now()

	return &User{
		email:     email,
		username:  username,
		phone:     phone,
		password:  password,
		badge:     badge,
		isActive:  isactive,
		score:     float64(score),
		createdAT: nows,
	}, nil
}

func WithoutValidation(email, password, badge, username, phone string, isactive bool, score float64, times time.Time) (*User, error) {
	return &User{
		email:     email,
		username:  username,
		phone:     phone,
		password:  password,
		badge:     badge,
		isActive:  isactive,
		score:     score,
		createdAT: times,
	}, nil
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Badge() string {
	return u.badge
}

func (u User) IsActive() bool {
	return u.isActive
}

func (u User) Score() float64 {
	return u.score
}

func (u User) CreatedAt() time.Time {
	return u.createdAT
}

func (u User) ID() uint {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Phone() string {
	return u.phone
}

func (u *User) Id_Set(id uint) {
	u.id = id
}
