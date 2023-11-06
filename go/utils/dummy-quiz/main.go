package main

import (
	"encoding/json"
	"fmt"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	dummyRepo := repository.NewDummyQuizUserRepository()
	skoringRepo := repository.NewSkoringRepository()
	quiz, err := dummyRepo.CekDummyQuizUser(100)
	if err != nil {
		fmt.Println("Belum ada sesi dummy untuk template ini..")
		return
	}
	id_quiz := quiz.IDQuiz
	fmt.Println(id_quiz)
	randomJawaban, errJawaban := dummyRepo.GenerateJawaban(int(id_quiz))
	if errJawaban != nil {
		fmt.Println("Error gagal buat jawaban sesi dummy untuk template ini..")
		return
	}
	var jawabanString = []entity.JawabanString{}
	for i := 0; i < len(randomJawaban); i++ {
		fmt.Println(randomJawaban[i].Jawaban)
		fmt.Println(randomJawaban[i].Kategori)
		tmp := entity.JawabanString{}
		tmp.Kategori = randomJawaban[i].Kategori
		jws, _ := json.Marshal(randomJawaban[i].Jawaban)
		tmp.Jawaban = string(jws)
		jawabanString = append(jawabanString, tmp)
	}
	j, _ := json.Marshal(jawabanString)
	jsonString := string(j)
	// log.Println(string(jsonString))

	dummyRepo.SubmitJawabanUser(int(id_quiz), jsonString)

	kategori_tabel, _ := skoringRepo.GetTabelSkoring(id_quiz)
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	skoringRepo.GenerateTabelTemporaryJawabanUser(id_quiz, 100)
	for _, kategori := range kategori_tabel {
		// print(kategori.Tabel)
		//1
		if kategori.Tabel == "skor_kognitif" {
			skoringRepo.SkoringKognitif(id_quiz, 100)
		}
		//2
		if kategori.Tabel == "skor_kognitif_pmk" {
			skoringRepo.SkoringKognitifPMK(id_quiz, 100)
		}
		//3
		if kategori.Tabel == "skor_gaya_pekerjaan" {
			skoringRepo.SkoringGayaPekerjaan(id_quiz, 100)
		}
		//4
		if kategori.Tabel == "skor_sikap_pelajaran" {
			skoringRepo.SkoringSikapPelajaran(id_quiz, 100)
		}
		//5
		if kategori.Tabel == "skor_sikap_pelajaran_mk" {
			skoringRepo.SkoringSikapPelajaranMK(id_quiz, 100)
		}
		//6
		if kategori.Tabel == "skor_sikap_pelajaran_man" {
			skoringRepo.SkoringSikapPelajaran(id_quiz, 100)
		}
		//7
		if kategori.Tabel == "skor_peminatan_smk" {
			skoringRepo.SkoringPeminatanSMK(id_quiz, 100)
		}
		//8
		if kategori.Tabel == "skor_peminatan_sma" {
			skoringRepo.SkoringPeminatanSMA(id_quiz, 100)
		}
		//9
		if kategori.Tabel == "skor_peminatan_man" {
			skoringRepo.SkoringPeminatanMAN(id_quiz, 100)
		}
		//10
		if kategori.Tabel == "skor_kuliah_dinas" {
			skoringRepo.SkoringMinatKuliahKedinasan(id_quiz, 100)
		}
		//11
		if kategori.Tabel == "skor_kuliah_alam" {
			skoringRepo.SkoringKuliahAlam(id_quiz, 100)
		}
		//12
		if kategori.Tabel == "skor_kuliah_sosial" {
			skoringRepo.SkoringKuliahSosial(id_quiz, 100)
		}
		//13
		if kategori.Tabel == "skor_kuliah_agama" {
			skoringRepo.SkoringKuliahAgama(id_quiz, 100)
		}
		//14
		if kategori.Tabel == "skor_mbti" {
			skoringRepo.SkoringMBTI(id_quiz, 100)
		}
		//15
		if kategori.Tabel == "skor_karakteristik_pribadi" {
			skoringRepo.SkoringKarakteristikPribadi(id_quiz, 100)
		}
		//16
		if kategori.Tabel == "skor_minat_indonesia" {
			skoringRepo.SkoringTesMinatIndonesia(id_quiz, 100)
		}
		//17
		if kategori.Tabel == "skor_kecerdasan_majemuk" {
			skoringRepo.SkoringKecerdasanMajemuk(id_quiz, 100)
		}
		//18
		if kategori.Tabel == "skor_suasana_kerja" {
			skoringRepo.SkoringSuasanaKerja(id_quiz, 100)
		}
		//19
		if kategori.Tabel == "skor_gaya_belajar" {
			skoringRepo.SkoringGayaBelajar(id_quiz, 100)
		}
		//20
		if kategori.Tabel == "skor_kejiwaan_dewasa" {
			skoringRepo.SkoringKejiwaanDewasa(id_quiz, 100)
		}
		//21
		if kategori.Tabel == "skor_kesehatan_mental" {
			skoringRepo.SkoringKesehatanMental(id_quiz, 100)
		}
		//22
		if kategori.Tabel == "skor_mode_belajar" {
			skoringRepo.SkoringModeBelajar(id_quiz, 100)
		}
		//23
		if kategori.Tabel == "skor_ssct" {
			skoringRepo.SkoringSSCT(id_quiz, 100)
		}
	}
	skoringRepo.ClearTabelTemporaryJawabanUser(id_quiz, 100)
	now := helper.StringTimeYMDHIS(time.Now())
	skoringRepo.FinishSkoring(id_quiz, 100, now)
}
