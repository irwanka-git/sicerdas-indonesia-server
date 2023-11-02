package main

import (
	"fmt"
	"irwanka/sicerdas/helper"
	"irwanka/sicerdas/repository"
	"os"
	"time"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	skoringRepo := repository.NewSkoringRepository()
	cek, _ := skoringRepo.GetStatusRunningSkoring()
	if cek.Status == 1 {
		fmt.Println("Skoring sedang berjalan..")
		return
	}
	skoringRepo.StartRunningSkoring()
	list_belum_skoring, _ := skoringRepo.GetUserSesiBelumSkoring()
	for _, user := range list_belum_skoring {

		kategori_tabel, _ := skoringRepo.GetTabelSkoring(user.IDQuiz)
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		skoringRepo.GenerateTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		for _, kategori := range kategori_tabel {
			//1
			if kategori.Tabel == "skor_kognitif" {
				skoringRepo.SkoringKognitif(user.IDQuiz, user.IDUser)
			}
			//2
			if kategori.Tabel == "skor_kognitif_pmk" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			//3
			if kategori.Tabel == "skor_gaya_pekerjaan" {
				skoringRepo.SkoringGayaPekerjaan(user.IDQuiz, user.IDUser)
			}
			//4
			if kategori.Tabel == "skor_sikap_pelajaran" {
				skoringRepo.SkoringSikapPelajaran(user.IDQuiz, user.IDUser)
			}
			//5
			if kategori.Tabel == "skor_sikap_pelajaran_mk" {
				skoringRepo.SkoringSikapPelajaranMK(user.IDQuiz, user.IDUser)
			}
			//6
			if kategori.Tabel == "skor_sikap_pelajaran_man" {
				skoringRepo.SkoringSikapPelajaran(user.IDQuiz, user.IDUser)
			}
			//7
			if kategori.Tabel == "skor_peminatan_smk" {
				skoringRepo.SkoringPeminatanSMK(user.IDQuiz, user.IDUser)
			}
			//8
			if kategori.Tabel == "skor_peminatan_sma" {
				skoringRepo.SkoringPeminatanSMA(user.IDQuiz, user.IDUser)
			}
			//9
			if kategori.Tabel == "skor_peminatan_man" {
				skoringRepo.SkoringPeminatanMAN(user.IDQuiz, user.IDUser)
			}
			//10
			if kategori.Tabel == "skor_kuliah_dinas" {
				skoringRepo.SkoringMinatKuliahKedinasan(user.IDQuiz, user.IDUser)
			}
			//11
			if kategori.Tabel == "skor_kuliah_alam" {
				skoringRepo.SkoringKuliahAlam(user.IDQuiz, user.IDUser)
			}
			//12
			if kategori.Tabel == "skor_kuliah_sosial" {
				skoringRepo.SkoringKuliahSosial(user.IDQuiz, user.IDUser)
			}
			//13
			if kategori.Tabel == "skor_kuliah_agama" {
				skoringRepo.SkoringKuliahAgama(user.IDQuiz, user.IDUser)
			}
			//14
			if kategori.Tabel == "skor_mbti" {
				skoringRepo.SkoringMBTI(user.IDQuiz, user.IDUser)
			}
			//15
			if kategori.Tabel == "skor_karakteristik_pribadi" {
				skoringRepo.SkoringKarakteristikPribadi(user.IDQuiz, user.IDUser)
			}
			//16
			if kategori.Tabel == "skor_minat_indonesia" {
				skoringRepo.SkoringTesMinatIndonesia(user.IDQuiz, user.IDUser)
			}
			//17
			if kategori.Tabel == "skor_kecerdasan_majemuk" {
				skoringRepo.SkoringKecerdasanMajemuk(user.IDQuiz, user.IDUser)
			}
			//18
			if kategori.Tabel == "skor_suasana_kerja" {
				skoringRepo.SkoringSuasanaKerja(user.IDQuiz, user.IDUser)
			}
			//19
			if kategori.Tabel == "skor_gaya_belajar" {
				skoringRepo.SkoringGayaBelajar(user.IDQuiz, user.IDUser)
			}
			//20
			if kategori.Tabel == "skor_kejiwaan_dewasa" {
				skoringRepo.SkoringKejiwaanDewasa(user.IDQuiz, user.IDUser)
			}
			//21
			if kategori.Tabel == "skor_kesehatan_mental" {
				skoringRepo.SkoringKesehatanMental(user.IDQuiz, user.IDUser)
			}
			//22
			if kategori.Tabel == "skor_mode_belajar" {
				skoringRepo.SkoringModeBelajar(user.IDQuiz, user.IDUser)
			}
			//23
			if kategori.Tabel == "skor_ssct" {
				skoringRepo.SkoringSSCT(user.IDQuiz, user.IDUser)
			}
		}
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		now := helper.StringTimeYMDHIS(time.Now())
		skoringRepo.FinishSkoring(user.IDQuiz, user.IDUser, now)
		text := fmt.Sprintf("%v  \t ID USER : %v   ID QUIZ : %v  Sesi: %v (%v)  Kota: %v", now, user.IDUser, user.IDQuiz, user.NamaSesi, user.IDQuizTemplate, user.Kota)
		fmt.Println(text)
	}
	skoringRepo.StopRunningSkoring()
}
