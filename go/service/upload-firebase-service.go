package service

import (
	"irwanka/sicerdas/repository"
)

type UploadFirebaseService interface {
	UploadFileToFirebase(src string, bucket string) (string, error)
}

var (
	uploadfirebaserepo repository.UploadFirebaseRepository
)

func NewUploadFirebaseRepository(repo repository.UploadFirebaseRepository) UploadFirebaseService {
	uploadfirebaserepo = repo
	return &service{}
}

func (*service) UploadFileToFirebase(src string, directory string) (string, error) {
	return uploadfirebaserepo.UploadFileToFirebase(src, directory)
}
