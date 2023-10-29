package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type UserService interface {
	AuthLogin(credentials *entity.Credentials) (*entity.User, error)
	GetlUserByUuid(uuid string) (*entity.User, error)
	GetTopTenInfoCerdas() ([]*entity.InfoCerdas, error)
	SubmitPasswordBaru(user_id int32, password entity.PasswordChange) error

	UpdateGambarInfo(id int32, gambar string) error
	GetAllInfoCerdas() ([]*entity.InfoCerdas, error)
}

var (
	userrepo repository.UserRepository
)

func NewUserService(repo repository.UserRepository) UserService {
	userrepo = repo
	return &service{}
}

func (*service) GetAllInfoCerdas() ([]*entity.InfoCerdas, error) {
	return userrepo.GetAllInfoCerdas()
}

func (*service) UpdateGambarInfo(id int32, gambar string) error {
	return userrepo.UpdateGambarInfo(id, gambar)
}

func (*service) SubmitPasswordBaru(user_id int32, password entity.PasswordChange) error {
	return userrepo.SubmitPasswordBaru(user_id, password)
}

func (*service) AuthLogin(credentials *entity.Credentials) (*entity.User, error) {
	return userrepo.AuthLogin(credentials)
}

func (*service) GetTopTenInfoCerdas() ([]*entity.InfoCerdas, error) {
	return userrepo.GetTopTenInfoCerdas()
}

func (*service) GetlUserByUuid(uuid string) (*entity.User, error) {
	return userrepo.GetlUserByUuid(uuid)
}
