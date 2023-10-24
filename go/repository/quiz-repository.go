package repository

import (
	"errors"
	"fmt"
	"irwanka/sicerdas/entity"
	"strings"
)

type QuizRepository interface {
	GetListInfoSessionQuiz(token string) ([]*entity.QuizSesiInfo, error)
	GetlDetilQuizByToken(token string) (*entity.Quiz, error)
	GetAllSoalSessionQuiz(token string) ([]*entity.SoalSession, error)
}

var (
	soalSessionRepo SoalRepository = NewSoalRepository()
)

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
		temp.PanjangJawaban = int(listSesi[i].PanjangJawaban)

		temp.Mode = listSesi[i].Mode
		temp.Play = 0
		temp.Token = token
		temp.PetunjukSesi = listSesi[i].PetunjukSesi
		temp.NamaSesiUjian = listSesi[i].NamaSesiUjian
		temp.Soal = listSesi[i].Soal
		temp.Sections = []string{}

		listSesiInfo = append(listSesiInfo, &temp)
	}

	return listSesiInfo, nil
}

func (*repo) GetAllSoalSessionQuiz(token string) ([]*entity.SoalSession, error) {
	var quiz *entity.Quiz
	resultQuiz := db.Table("quiz_sesi").Where("token = ?", token).First(&quiz)
	if resultQuiz.RowsAffected == 0 {
		return []*entity.SoalSession{}, errors.New("Data not found")
	}

	var listSoal = []*entity.SoalSession{}
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

	for i := 0; i < len(listSesi); i++ {
		splitSoal := strings.Split(listSesi[i].Soal, "/")

		// //BAKAT DAN KARIR

		// //SOAL GAYA BELAJAR

		//pola : /soal-kognitif/{paket}/{bidang}
		if splitSoal[1] == "soal-kognitif" && len(splitSoal) == 4 {
			paket := splitSoal[2]
			bidang := splitSoal[3]
			currentList, _ := soalSessionRepo.GetSoalKognitif(paket, bidang, token)
			listSoal = append(listSoal, currentList...)
		}

		//pola : /soal-skala-peminatan-smk
		if splitSoal[1] == "soal-skala-peminatan-smk" && len(splitSoal) == 2 {
			currentList, _ := soalSessionRepo.GetSoalPeminatanSMK(quiz.IDQuiz, "", false, token)
			listSoal = append(listSoal, currentList...)
		}
		//pola : /soal-skala-peminatan-smk-demo
		if splitSoal[1] == "soal-skala-peminatan-smk-demo" && len(splitSoal) == 2 {
			currentList, _ := soalSessionRepo.GetSoalPeminatanSMK(quiz.IDQuiz, "", true, token)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-skala-peminatan-smk/{paket}
		if splitSoal[1] == "soal-skala-peminatan-smk" && len(splitSoal) == 3 {
			paket := splitSoal[2]
			currentList, _ := soalSessionRepo.GetSoalPeminatanSMK(quiz.IDQuiz, paket, false, token)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-sikap-pelajaran
		if splitSoal[1] == "soal-sikap-pelajaran" {
			currentList, _ := soalSessionRepo.GetSoalSikapPelajaran(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-sikap-pelajaran-demo
		if splitSoal[1] == "soal-sikap-pelajaran-demo" {
			currentList, _ := soalSessionRepo.GetSoalSikapPelajaran(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-pmk-sikap-pelajaran
		if splitSoal[1] == "soal-pmk-sikap-pelajaran" {
			currentList, _ := soalSessionRepo.GetSoalSikapPelajaranKuliah(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-pmk-sikap-pelajaran-demo
		if splitSoal[1] == "soal-pmk-sikap-pelajaran-demo" {
			currentList, _ := soalSessionRepo.GetSoalSikapPelajaranKuliah(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-tmi
		if splitSoal[1] == "soal-tmi" {
			currentList, _ := soalSessionRepo.GetSoalTesMinatIndonesia(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-tmi-demo
		if splitSoal[1] == "soal-tmi-demo" {
			currentList, _ := soalSessionRepo.GetSoalTesMinatIndonesia(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-tipologi-jung
		if splitSoal[1] == "soal-tipologi-jung" {
			currentList, _ := soalSessionRepo.GetSoalTipologiJung(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-tipologi-jung-demo
		if splitSoal[1] == "soal-tipologi-jung-demo" {
			currentList, _ := soalSessionRepo.GetSoalTipologiJung(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-tes-karakteristik-pribadi
		if splitSoal[1] == "soal-tes-karakteristik-pribadi" {
			currentList, _ := soalSessionRepo.GetSoalKarakteristikPribadi(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-tes-karakteristik-pribadi-demo
		if splitSoal[1] == "soal-tes-karakteristik-pribadi-demo" {
			currentList, _ := soalSessionRepo.GetSoalKarakteristikPribadi(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-skala-peminatan-sma
		if splitSoal[1] == "soal-skala-peminatan-sma" {
			currentList, _ := soalSessionRepo.GetSoalSkalaPeminatanSMA(token, false)
			listSoal = append(listSoal, currentList...)
		}
		if splitSoal[1] == "soal-skala-peminatan-sma-demo" {
			currentList, _ := soalSessionRepo.GetSoalSkalaPeminatanSMA(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-skala-peminatan-man
		if splitSoal[1] == "soal-skala-peminatan-man" {
			currentList, _ := soalSessionRepo.GetSoalSkalaPeminatanMAN(token, false)
			listSoal = append(listSoal, currentList...)
		}
		if splitSoal[1] == "soal-skala-peminatan-man-demo" {
			currentList, _ := soalSessionRepo.GetSoalSkalaPeminatanMAN(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//Pola: '/soal-minat-kuliah-eksakta'
		if splitSoal[1] == "soal-minat-kuliah-eksakta" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahEksakta(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//Pola: '/soal-minat-kuliah-eksakta-demo'
		if splitSoal[1] == "soal-minat-kuliah-eksakta-demo" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahEksakta(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: '/soal-minat-kuliah-sosial'
		if splitSoal[1] == "soal-minat-kuliah-sosial" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahSosial(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: '/soal-minat-kuliah-sosial-demo'
		if splitSoal[1] == "soal-minat-kuliah-sosial" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahSosial(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: '/soal-minat-kuliah-kedinasan'
		if splitSoal[1] == "soal-minat-kuliah-kedinasan" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahDinas(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: '/soal-minat-kuliah-kedinasan-demo'
		if splitSoal[1] == "soal-minat-kuliah-kedinasan-demo" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahDinas(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-minat-kuliah-agama
		if splitSoal[1] == "soal-minat-kuliah-agama" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahAgama(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-minat-kuliah-agama-demo
		if splitSoal[1] == "soal-minat-kuliah-agama-demo" {
			currentList, _ := soalSessionRepo.GetSoalMinatKuliahAgama(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola:  '/soal-minat-kuliah-suasana-kerja'
		if splitSoal[1] == "soal-minat-kuliah-suasana-kerja" {
			currentList, _ := soalSessionRepo.GetSoalSuasanaKerja(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-minat-kuliah-suasana-kerja-demo
		if splitSoal[1] == "soal-minat-kuliah-suasana-kerja-demo" {
			currentList, _ := soalSessionRepo.GetSoalSuasanaKerja(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola:  '/soal-kecerdasan-majemuk'
		if splitSoal[1] == "soal-kecerdasan-majemuk" {
			currentList, _ := soalSessionRepo.GetSoalKecerdasanMajemuk(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-kecerdasan-majemuk-demo
		if splitSoal[1] == "soal-kecerdasan-majemuk-demo" {
			currentList, _ := soalSessionRepo.GetSoalKecerdasanMajemuk(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola:  '/soal-gaya-pekerjaan'
		if splitSoal[1] == "soal-gaya-pekerjaan" {
			currentList, _ := soalSessionRepo.GetSoalGayaPekerjaan(token, false)
			listSoal = append(listSoal, currentList...)
		}
		//pola: /soal-gaya-pekerjaan-demo
		if splitSoal[1] == "soal-gaya-pekerjaan-demo" {
			currentList, _ := soalSessionRepo.GetSoalGayaPekerjaan(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola: /soal-gaya-belajar
		//pola: /soal-gaya-belajar-demo
		if splitSoal[1] == "soal-gaya-belajar" {
			currentList, _ := soalSessionRepo.GetSoalGayaBelajar(token, false)
			listSoal = append(listSoal, currentList...)
		}
		if splitSoal[1] == "soal-gaya-belajar-demo" {
			currentList, _ := soalSessionRepo.GetSoalGayaBelajar(token, true)
			listSoal = append(listSoal, currentList...)
		}

		//pola : /soal-break
		if splitSoal[1] == "soal-break" && len(splitSoal) == 2 {
			currentList, _ := soalSessionRepo.GetSoalBreak(token)
			listSoal = append(listSoal, currentList...)
		}

	}
	return listSoal, nil
}