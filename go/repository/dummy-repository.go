package repository

import (
	"errors"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"math/rand"
	"strconv"
	"time"
)

const (
	ID_USER_DUMMY = 100
)

type DummyQuizUserRepository interface {
	CekDummyQuizUser(id_quiz_template int) (*entity.QuizSesiUser, error)
	GenerateJawaban(id_quiz int) ([]*entity.JawabanQuiz, error)
	SubmitJawabanUser(id_quiz int, jawaban string) error
}

var (
	// soalRepoDummy SoalRepository = NewSoalRepository()
	quizRepoDummy QuizRepository = NewQuizRepository()
)

func NewDummyQuizUserRepository() DummyQuizUserRepository {
	return &repo{}
}

func (*repo) CekDummyQuizUser(id_quiz_template int) (*entity.QuizSesiUser, error) {
	var quizUser *entity.QuizSesiUser
	db.Raw(`select a.* from 
			quiz_sesi_user as a, 
			quiz_sesi as b, quiz_sesi_template as c 
				where a.id_quiz = b.id_quiz 
				and b.id_quiz_template = c.id_quiz_template 
				and c.id_quiz_template = ? and a.id_user = ? order by b.id_quiz desc limit 1`,
		id_quiz_template, ID_USER_DUMMY).First(&quizUser)
	if quizUser.IDQuiz == 0 {
		return nil, errors.New("belum ada data dummy untuk template ini")
	}
	return quizUser, nil
}

func (*repo) GenerateJawaban(id_quiz int) ([]*entity.JawabanQuiz, error) {
	var quiz *entity.Quiz
	db.Table("quiz_sesi").Where("id_quiz = ?", id_quiz).First(&quiz)
	fmt.Println(quiz)

	token := quiz.Token
	listInfoSesi, err := quizRepoDummy.GetListInfoSessionQuiz(token)
	if err != nil {
		return []*entity.JawabanQuiz{}, errors.New("not found list session")
	}
	listSoalSesi, err := quizRepoDummy.GetAllSoalSessionQuiz(token)
	if err != nil {
		return []*entity.JawabanQuiz{}, errors.New("not found list soal")
	}

	var jawaban = []*entity.JawabanQuiz{}

	for i := 0; i < len(listInfoSesi); i++ {
		sesi := listInfoSesi[i]
		var randomJawaban = entity.JawabanQuiz{}
		randomJawaban.Kategori = sesi.Kategori
		var jawabanArray = []string{}
		jawabanArray = append(jawabanArray, "0")
		if sesi.Mode == "PG" {
			for s := 0; s < len(listSoalSesi); s++ {
				if listSoalSesi[s].Kategori == sesi.Kategori {
					var pilihanPG = []string{}
					for p := 0; p < len(listSoalSesi[s].Pilihan); p++ {
						pilihanPG = append(pilihanPG, listSoalSesi[s].Pilihan[p].Value)
					}
					rand.Shuffle(len(pilihanPG), func(x, y int) { pilihanPG[x], pilihanPG[y] = pilihanPG[y], pilihanPG[x] })
					jawabanArray = append(jawabanArray, pilihanPG[0])
				}
			}
		}

		if sesi.Mode == "RT" {
			for s := 0; s < len(listSoalSesi); s++ {
				if listSoalSesi[s].Kategori == sesi.Kategori {
					min := 1
					max := 7
					var jawaban = ""
					for p := 0; p < sesi.PanjangJawaban; p++ {
						random := rand.Intn(max-min) + min
						jawaban = jawaban + strconv.Itoa(random)
					}
					jawabanArray = append(jawabanArray, jawaban)
				}
			}
		}

		if sesi.Mode == "TOP" {
			var pilihanTOP = []string{}
			for s := 0; s < len(listSoalSesi); s++ {
				if listSoalSesi[s].Kategori == sesi.Kategori {
					pilihanTOP = append(pilihanTOP, listSoalSesi[s].Nomor)
				}
			}
			rand.Shuffle(len(pilihanTOP), func(x, y int) { pilihanTOP[x], pilihanTOP[y] = pilihanTOP[y], pilihanTOP[x] })
			for p := 0; p < sesi.Jawaban; p++ {
				jawabanArray = append(jawabanArray, pilihanTOP[p])
			}
		}

		if sesi.Mode == "PP" {
			for s := 0; s < len(listSoalSesi); s++ {
				if listSoalSesi[s].Kategori == sesi.Kategori {
					var jawaban = ""
					var pilihanPP = []string{}
					for p := 0; p < len(listSoalSesi[s].Pilihan); p++ {
						pilihanPP = append(pilihanPP, listSoalSesi[s].Pilihan[p].Value)
					}
					rand.Shuffle(len(pilihanPP), func(x, y int) { pilihanPP[x], pilihanPP[y] = pilihanPP[y], pilihanPP[x] })
					for p := 0; p < len(pilihanPP); p++ {
						jawaban = jawaban + pilihanPP[p]
					}
					jawabanArray = append(jawabanArray, jawaban)
				}
			}
		}

		if sesi.Mode == "PGS" {
			for s := 0; s < len(listSoalSesi); s++ {
				if listSoalSesi[s].Kategori == sesi.Kategori {
					var jawaban = ""
					var pilihanPGS = []string{}
					for p := 0; p < len(listSoalSesi[s].Pilihan); p++ {
						pilihanPGS = append(pilihanPGS, listSoalSesi[s].Pilihan[p].Value)
					}
					for p := 0; p < len(listSoalSesi[s].PernyataanMulti); p++ {
						rand.Shuffle(len(pilihanPGS), func(x, y int) { pilihanPGS[x], pilihanPGS[y] = pilihanPGS[y], pilihanPGS[x] })
						jawaban = jawaban + pilihanPGS[0]
					}
					jawabanArray = append(jawabanArray, jawaban)
				}
			}
		}

		// fmt.Println(jawabanArray)
		randomJawaban.Jawaban = jawabanArray
		if sesi.Kategori != "BREAK" {
			jawaban = append(jawaban, &randomJawaban)
			// fmt.Println(sesi.Mode)
		}
	}
	// fmt.Println(jawaban)
	return jawaban, nil
}

func (*repo) SubmitJawabanUser(id_quiz int, jawaban string) error {
	now := helper.StringTimeYMDHIS(time.Now())
	db.Table("quiz_sesi_user").
		Where("id_quiz = ?", id_quiz).
		Where("id_user = ?", ID_USER_DUMMY).
		UpdateColumns(map[string]interface{}{
			"submit":    1,
			"skoring":   0,
			"jawaban":   jawaban,
			"submit_at": now})
	return nil
}
