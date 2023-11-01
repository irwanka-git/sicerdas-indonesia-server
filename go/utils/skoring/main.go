package main

import (
	"fmt"
	"irwanka/sicerdas/repository"
)

func main() {
	skoringRepo := repository.NewSkoringRepository()
	cek, _ := skoringRepo.GetStatusRunningSkoring()
	if cek.Status == 1 {
		fmt.Println("Skoring sedang berjalan..")
		return
	}
	skoringRepo.StartRunningSkoring()
	list_belum_skoring, _ := skoringRepo.GetUserSesiBelumSkoring()
	for _, user := range list_belum_skoring {
		text := fmt.Sprintf("ID USER : %v \t ID QUIZ : %v \t Sesi: %v (%v) \t Kota: %v", user.IDUser, user.IDQuiz, user.NamaSesi, user.IDQuizTemplate, user.Kota)
		fmt.Println(text)
		kategori_tabel, _ := skoringRepo.GetTabelSkoring(user.IDQuiz)
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		skoringRepo.GenerateTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
		for _, kategori := range kategori_tabel {
			fmt.Println(kategori.Tabel)
			if kategori.Tabel == "skor_kognitif" {
				skoringRepo.SkoringKognitif(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "skor_kognitif_pmk" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "skor_gaya_pekerjaan" {
				skoringRepo.SkoringGayaPekerjaan(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "skor_sikap_pelajaran" {
				skoringRepo.SkoringSikapPelajaran(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_sikap_pelajaran_mk" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_sikap_pelajaran_man" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_peminatan_smk" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_peminatan_sma" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_peminatan_man" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kuliah_dinas" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kuliah_alam" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kuliah_sosial" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kuliah_agama" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_mbti" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_karakteristik_pribadi" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_minat_indonesia" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kecerdasan_majemuk" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_suasana_kerja" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_gaya_belajar" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kejiwaan_dewasa" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_kesehatan_mental" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_mode_belajar" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
			if kategori.Tabel == "=> skor_ssct" {
				skoringRepo.SkoringKognitifPMK(user.IDQuiz, user.IDUser)
			}
		}
		skoringRepo.ClearTabelTemporaryJawabanUser(user.IDQuiz, user.IDUser)
	}
	skoringRepo.StopRunningSkoring()
}
