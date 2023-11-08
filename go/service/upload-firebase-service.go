package service

import (
	"irwanka/sicerdas/repository"
)

type UploadFirebaseService interface {
	UploadSoalQuizToFirebase(src string, directory string) (string, error)
	UploadGambarToFirebase(src string, directory string) (string, error)
	UploadReportPDFToFirebase(src string, directory string) (string, error)
}

var (
	uploadfirebaserepo repository.UploadFirebaseRepository
)

func NewUploadFirebaseRepository(repo repository.UploadFirebaseRepository) UploadFirebaseService {
	uploadfirebaserepo = repo
	return &service{}
}

func (*service) UploadReportPDFToFirebase(src string, directory string) (string, error) {
	return uploadfirebaserepo.UploadReportPDFToFirebase(src, directory)
}

func (*service) UploadGambarToFirebase(src string, directory string) (string, error) {
	return uploadfirebaserepo.UploadGambarToFirebase(src, directory)
}

func (*service) UploadSoalQuizToFirebase(src string, directory string) (string, error) {
	return uploadfirebaserepo.UploadSoalQuizToFirebase(src, directory)
}
