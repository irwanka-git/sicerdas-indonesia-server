package repository

import (
	"errors"
	"irwanka/sicerdas/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	AuthLogin(credentials *entity.Credentials) (*entity.User, error)
	GetlUserByUuid(uuid string) (*entity.User, error)
}

func NewUserRepository() UserRepository {
	return &repo{}
}

func (*repo) AuthLogin(credentials *entity.Credentials) (*entity.User, error) {
	var userCek *entity.User
	result := db.Table("users").Where("username = ?", credentials.Username).First(&userCek)
	if result.RowsAffected == 0 {
		return nil, errors.New("User ID tidak ditemukan, silahkan hubungi admin")
	}
	errPassword := bcrypt.CompareHashAndPassword([]byte(userCek.Password), []byte(credentials.Password))
	if errPassword != nil {
		return nil, errors.New("Password tidak valid")
	}
	return userCek, nil
}

func (*repo) GetlUserByUuid(uuid string) (*entity.User, error) {
	var userCek *entity.User
	result := db.Table("users").Where("uuid = ?", uuid).First(&userCek)
	if result.RowsAffected == 0 {
		return nil, errors.New("User ID tidak ditemukan, silahkan hubungi admin")
	}
	return userCek, nil
}
