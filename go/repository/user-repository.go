package repository

import (
	"errors"
	"irwanka/sicerdas/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	AuthLogin(credentials *entity.Credentials) (*entity.User, error)
	GetlUserByUuid(uuid string) (*entity.User, error)
	GetTopTenInfoCerdas() ([]*entity.InfoCerdas, error)
	SubmitPasswordBaru(user_id int32, password entity.PasswordChange) error
}

func NewUserRepository() UserRepository {
	return &repo{}
}

func (*repo) SubmitPasswordBaru(user_id int32, password entity.PasswordChange) error {

	newPassword, err := bcrypt.GenerateFromPassword([]byte(password.PasswordBaru), 14)
	if err != nil {
		return err
	}
	// print(string(newPassword))
	db.Table("users").Where("id = ?", user_id).UpdateColumn("password", string(newPassword))
	return nil
}

func (*repo) GetTopTenInfoCerdas() ([]*entity.InfoCerdas, error) {
	var list []*entity.InfoCerdas
	db.Table("info_cerdas").Order("id_info asc").Limit(10).Scan(&list)
	return list, nil
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
