package controller

import (
	"irwanka/sicerdas/repository"
)

var (
	soalRepository repository.SoalRepository = repository.NewSoalRepository()
)

type SoalController interface {
	// SinkronGambarKognitif(w http.ResponseWriter, r *http.Request)
}

func NewSoalController() SoalController {
	return &controller{}
}
