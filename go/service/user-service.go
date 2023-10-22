package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type UserService interface {
	AuthLogin(credentials *entity.Credentials) (*entity.User, error)
	GetlUserByUuid(uuid string) (*entity.User, error)
}

var (
	userrepo repository.UserRepository
)

func NewUserService(repo repository.UserRepository) UserService {
	userrepo = repo
	return &service{}
}

func (*service) AuthLogin(credentials *entity.Credentials) (*entity.User, error) {
	return userrepo.AuthLogin(credentials)
}

func (*service) GetlUserByUuid(uuid string) (*entity.User, error) {
	return userrepo.GetlUserByUuid(uuid)
}
