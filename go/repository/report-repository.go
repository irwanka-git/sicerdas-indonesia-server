package repository

type ReportRepository interface {
}

func NewReportRepository() ReportRepository {
	return &repo{}
}
