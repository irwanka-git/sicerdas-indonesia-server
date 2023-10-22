package repository

import (
	"errors"
	"fmt"
	"irwanka/sicerdas/entity"
)

type QuizRepository interface {
	GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error)
	GetlDetilQuizByToken(token string) (*entity.Quiz, error)
}

func NewQuizRepository() QuizRepository {
	return &repo{}
}

func (*repo) GetlDetilQuizByToken(token string) (*entity.Quiz, error) {
	var quiz *entity.Quiz
	result := db.Table("quiz_sesi").Where("token = ?", token).First(&quiz)
	if result.RowsAffected == 0 {
		return nil, errors.New("Data not found")
	}
	return quiz, nil
}

func (*repo) GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error) {
	var listSesiInfo []*entity.QuizSesiInfo

	var listSesi []struct {
		*entity.QuizSesiMaster
		Durasi     int    `json:"message"`
		KunciWaktu int    `json:"kunci_waktu"`
		Token      string `json:"token"`
		Urutan     int    `json:"urutan"`
		IdQuiz     int    `json:"id_quiz"`
	}

	db.Raw(`select c.*, b.durasi, b.kunci_waktu, a.token , b.urutan, a.id_quiz 
			from quiz_sesi as a, quiz_sesi_detil as b, quiz_sesi_master as c 
			where  a.id_quiz = b.id_quiz
			and c.id_sesi_master = b.id_sesi_master and  a.token = ? order by b.urutan asc`, token).Scan(&listSesi)

	if len(listSesi) == 0 {
		return []*entity.QuizSesiInfo{}, errors.New("token tidak valid")
	}

	for i := 0; i < len(listSesi); i++ {
		//khusus peminatan smk jumlah jawaban tergantung kondisi SMK;
		var temp = entity.QuizSesiInfo{}
		if listSesi[i].Kategori == "SKALA_PEMINATAN_SMK" {
			var jumlahJawaban int64
			if listSesi[i].NamaSesiUjian == "TES PEMINATAN SMK - DEMO" {
				jumlahJawaban = 3
			} else {
				db.Table("quiz_sesi_mapping_smk").Where("id_quiz = ?", listSesi[i].IdQuiz).Count(&jumlahJawaban)
			}
			listSesi[i].Soal = fmt.Sprintf("%v?token=%v", listSesi[i].Soal, token)
		}

		temp.Urutan = listSesi[i].Urutan
		temp.Durasi = listSesi[i].Durasi
		temp.Finish = 0
		temp.KunciWaktu = listSesi[i].KunciWaktu
		temp.Kategori = listSesi[i].Kategori
		temp.Jawaban = int(listSesi[i].Jawaban)
		temp.Mode = listSesi[i].Mode
		temp.Play = 0
		temp.PetunjukSesi = listSesi[i].PetunjukSesi
		temp.NamaSesiUjian = listSesi[i].NamaSesiUjian
		temp.Soal = listSesi[i].Soal

		listSesiInfo = append(listSesiInfo, &temp)
	}

	return listSesiInfo, nil
}
