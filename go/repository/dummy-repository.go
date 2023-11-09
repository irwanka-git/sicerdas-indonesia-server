package repository

import (
	"errors"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	ID_USER_DUMMY = 100
)

type DummyQuizUserRepository interface {
	CekDummyQuizUser(id_quiz_template int) (*entity.QuizSesiUser, error)
	CreateDummySesiTemplate(id_quiz_template int) entity.QuizSesi
	GenerateJawaban(id_quiz int) ([]*entity.JawabanQuiz, error)
	SubmitJawabanUser(id_quiz int, jawaban string) error
	DeleteOldDummy(id_quiz int) error
}

var (
	// soalRepoDummy SoalRepository = NewSoalRepository()
	quizRepoDummy QuizRepository = NewQuizRepository()
)

func NewDummyQuizUserRepository() DummyQuizUserRepository {
	return &repo{}
}

func (*repo) DeleteOldDummy(id_quiz int) error {

	var quisDummy = entity.QuizSesi{}
	db.Table("quiz_sesi").Where("id_quiz = ?", id_quiz).Delete(&quisDummy)
	var temp = entity.QuizSesiDetil{}
	db.Table("quiz_sesi_detil").Where("id_quiz = ?", id_quiz).Delete(&temp)
	var quiz_user = entity.QuizSesiUser{}
	db.Table("quiz_sesi_user").Where("id_quiz = ?", id_quiz).Delete(&quiz_user)

	return nil
}

func (*repo) CreateDummySesiTemplate(id_quiz_template int) entity.QuizSesi {
	var quiz_sesi_template entity.QuizSesiTemplate
	db.Table("quiz_sesi_template").Where("id_quiz_template", id_quiz_template).First(&quiz_sesi_template)
	var quisDummy = entity.QuizSesi{}
	token := generateNewToken()
	quisDummy.Arsip = 1
	quisDummy.CoverTemplate = "default.pdf"
	quisDummy.IDQuizTemplate = int32(id_quiz_template)
	quisDummy.Jenis = "quiz"
	quisDummy.Token = token
	quisDummy.Tanggal = time.Now()
	quisDummy.NamaSesi = "DUMMY - " + quiz_sesi_template.NamaSesi
	quisDummy.UUID = uuid.NewString()
	quisDummy.Kota = "Jambi"
	quisDummy.NamaAsesor = "JELPA PERIANTALO, M.Psi. Psi."
	quisDummy.NomorSipp = "0012-12-2-1"
	quisDummy.IDUserBiro = 239
	quisDummy.IDLokasi = 17
	quisDummy.ModelReport = "-"
	quisDummy.Gambar = "https://storage.googleapis.com/sicerdas-indonesia-service-repository/gambar/6091699310030.png"
	db.Table("quiz_sesi").Create(&quisDummy)

	var quizDummy entity.QuizSesi
	db.Table("quiz_sesi").Where("uuid", quisDummy.UUID).First(&quizDummy)

	id_quiz := quisDummy.IDQuiz
	var sesi_template []*entity.QuizSesiDetilTemplate

	db.Table("quiz_sesi_detil_template").Where("id_quiz_template = ?", id_quiz_template).Order("urutan asc").Scan(&sesi_template)
	for i := 0; i < len(sesi_template); i++ {
		var temp = entity.QuizSesiDetil{}
		temp.IDQuiz = id_quiz
		temp.Urutan = sesi_template[i].Urutan
		temp.Durasi = sesi_template[i].Durasi
		temp.KunciWaktu = sesi_template[i].KunciWaktu
		temp.IDSesiMaster = sesi_template[i].IDSesiMaster
		db.Table("quiz_sesi_detil").Create(&temp)

		//KHUSUS SKALA_PEMINATAN_SMK
		if sesi_template[i].IDSesiMaster == 13 {
			//genearte random mappping pilihan smk
			var random []*entity.ResultPeminatanSMK
			db.Raw(`select nomor , max(id_kegiatan) as urutan   from soal_peminatan_smk 
					group by nomor 
					order by random() limit 5 `).Scan(&random)
			for n := 0; n < len(random); n++ {
				var temp = entity.QuizSesiMappingSmk{}
				temp.IDQuiz = id_quiz
				temp.IDKegiatan = int32(random[n].Urutan)
				temp.UUID = uuid.NewString()
				db.Table("quiz_sesi_mapping_smk").Create(&temp)
			}
		}
	}

	var quiz_user = entity.QuizSesiUserCreate{}
	quiz_user.IDQuiz = id_quiz
	quiz_user.IDUser = ID_USER_DUMMY
	quiz_user.Skoring = 0
	quiz_user.Submit = 0
	quiz_user.UUID = uuid.NewString()

	db.Table("quiz_sesi_user").Create(&quiz_user)
	fmt.Println(quiz_user)
	fmt.Println(quizDummy)
	return quizDummy
}

func generateNewToken() string {
	tokenInt := rand.Intn(99999-11111) + 11111
	token := strconv.Itoa(tokenInt)
	var keepRunning = true
	var exist = int64(0)
	for keepRunning {
		db.Table("quiz_sesi").Where("token = ?", token).Count(&exist)
		if exist == 0 {
			break
		}
	}
	return token
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
