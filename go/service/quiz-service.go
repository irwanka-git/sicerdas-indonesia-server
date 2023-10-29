package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type QuisService interface {
	GetlListQuizByUser(id int) ([]*entity.QuizUserApi, error)
	GetStatusQuizUser(id int, token string) (*entity.QuizUserApi, error)
	GetSalamPembuka(token string) (string, error)
	SubmitJawabanQuiz(jawaban string, id_quiz int32, user entity.User) error

	GetlDetilQuizByToken(token string) (*entity.Quiz, error)
	GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error)
	GetAllSoalSessionQuiz(token string) ([]*entity.SoalSession, error)
	UpdateURLFirebaseSoalQuiz(token string, url string) error

	GetAllQuizTemplate() ([]*entity.QuizSesiTemplate, error)
	GetAllQuizSesi() ([]*entity.Quiz, error)
	UpdateGambarQuizTemplate(id int32, gambar string) error
	UpdateGambarQuizSesi(id int32, gambar string) error
}

var (
	quizrepo repository.QuizRepository
)

func NewQuizService(repo repository.QuizRepository) QuisService {
	quizrepo = repo
	return &service{}
}
func (*service) SubmitJawabanQuiz(jawaban string, id_quiz int32, user entity.User) error {
	return quizrepo.SubmitJawabanQuiz(jawaban, id_quiz, user)
}

func (*service) GetSalamPembuka(token string) (string, error) {
	return quizrepo.GetSalamPembuka(token)
}

func (*service) GetStatusQuizUser(id int, token string) (*entity.QuizUserApi, error) {
	return quizrepo.GetStatusQuizUser(id, token)
}

func (*service) GetlListQuizByUser(id int) ([]*entity.QuizUserApi, error) {
	return quizrepo.GetlListQuizByUser(id)
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

func (*service) GetAllQuizTemplate() ([]*entity.QuizSesiTemplate, error) {
	return quizrepo.GetAllQuizTemplate()
}

func (*service) UpdateGambarQuizTemplate(id int32, gambar string) error {
	return quizrepo.UpdateGambarQuizTemplate(id, gambar)
}

func (*service) UpdateGambarQuizSesi(id int32, gambar string) error {
	return quizrepo.UpdateGambarQuizSesi(id, gambar)
}

func (*service) GetAllQuizSesi() ([]*entity.Quiz, error) {
	return quizrepo.GetAllQuizSesi()
}
