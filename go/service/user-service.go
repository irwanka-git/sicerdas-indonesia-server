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
}

var (
	userrepo repository.UserRepository
)

func NewUserService(repo repository.UserRepository) UserService {
	userrepo = repo
	return &service{}
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
