package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type ReportService interface {
	GetKomponenReportTemplate(uuid string) (*entity.QuizSesiReportAndTemplate, error)
	GetQuizUserDummyFromTemplate(id_quiz_template int32) (*entity.QuizSesiUser, error)
	//get data skoring
	GetSkoringKognitif(id_quiz int, id_user int) (*entity.SkorKognitif, error)
	GetSkoringKognitifPMK(id_quiz int, id_user int) (*entity.SkorKognitif, error)
}

var (
	reportRepo repository.ReportRepository
)

func NewReporService(repo repository.ReportRepository) ReportService {
	reportRepo = repo
	return &service{}
}

func (*service) GetKomponenReportTemplate(uuid string) (*entity.QuizSesiReportAndTemplate, error) {
	return reportRepo.GetKomponenReportTemplate(uuid)
}

func (*service) GetQuizUserDummyFromTemplate(id_quiz_template int32) (*entity.QuizSesiUser, error) {
	return reportRepo.GetQuizUserDummyFromTemplate(id_quiz_template)
}

func (*service) GetSkoringKognitif(id_quiz int, id_user int) (*entity.SkorKognitif, error) {
	return reportRepo.GetSkoringKognitif(id_quiz, id_user)
}

func (*service) GetSkoringKognitifPMK(id_quiz int, id_user int) (*entity.SkorKognitif, error) {
	return reportRepo.GetSkoringKognitifPMK(id_quiz, id_user)
}
