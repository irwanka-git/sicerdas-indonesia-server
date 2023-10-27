package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type QuisService interface {
	GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error)
	GetlDetilQuizByToken(token string) (*entity.Quiz, error)
	GetAllSoalSessionQuiz(token string) ([]*entity.SoalSession, error)
	UpdateURLFirebaseSoalQuiz(token string, url string) error
}

var (
	quizrepo repository.QuizRepository
)

func NewQuizService(repo repository.QuizRepository) QuisService {
	quizrepo = repo
	return &service{}
}

func (*service) GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error) {
	return quizrepo.GetListInfoSessionQuiz(token)
}

func (*service) GetlDetilQuizByToken(token string) (*entity.Quiz, error) {
	return quizrepo.GetlDetilQuizByToken(token)
}

func (*service) GetAllSoalSessionQuiz(token string) ([]*entity.SoalSession, error) {
	return quizrepo.GetAllSoalSessionQuiz(token)
}

func (*service) UpdateURLFirebaseSoalQuiz(token string, url string) error {
	return quizrepo.UpdateURLFirebaseSoqalQuiz(token, url)
}
