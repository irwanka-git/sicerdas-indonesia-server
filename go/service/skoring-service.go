package service

import (
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/repository"
)

type SkoringService interface {
	SkoringAllKategori(kategori_tabel []*entity.KategoriTabel, id_quiz int32, id_user int32) error
}

var (
	skoringRepoService repository.SkoringRepository
)

func NewSkoringService(repo repository.SkoringRepository) SkoringService {
	skoringRepoService = repo
	return &service{}
}

func (*service) SkoringAllKategori(kategori_tabel []*entity.KategoriTabel, id_quiz int32, id_user int32) error {
	for _, kategori := range kategori_tabel {
		//1
		if kategori.Tabel == "skor_kognitif" {
			skoringRepoService.SkoringKognitif(id_quiz, id_user)
		}
		//2
		if kategori.Tabel == "skor_kognitif_pmk" {
			skoringRepoService.SkoringKognitifPMK(id_quiz, id_user)
		}
		//3
		if kategori.Tabel == "skor_gaya_pekerjaan" {
			skoringRepoService.SkoringGayaPekerjaan(id_quiz, id_user)
		}
		//4
		if kategori.Tabel == "skor_sikap_pelajaran" {
			skoringRepoService.SkoringSikapPelajaran(id_quiz, id_user)
		}
		//5
		if kategori.Tabel == "skor_sikap_pelajaran_mk" {
			skoringRepoService.SkoringSikapPelajaranMK(id_quiz, id_user)
		}
		//6
		if kategori.Tabel == "skor_sikap_pelajaran_man" {
			skoringRepoService.SkoringSikapPelajaran(id_quiz, id_user)
		}
		//7
		if kategori.Tabel == "skor_peminatan_smk" {
			skoringRepoService.SkoringPeminatanSMK(id_quiz, id_user)
		}
		//8
		if kategori.Tabel == "skor_peminatan_sma" {
			skoringRepoService.SkoringPeminatanSMA(id_quiz, id_user)
		}
		//9
		if kategori.Tabel == "skor_peminatan_man" {
			skoringRepoService.SkoringPeminatanMAN(id_quiz, id_user)
		}
		//10
		if kategori.Tabel == "skor_kuliah_dinas" {
			skoringRepoService.SkoringMinatKuliahKedinasan(id_quiz, id_user)
		}
		//11
		if kategori.Tabel == "skor_kuliah_alam" {
			skoringRepoService.SkoringKuliahAlam(id_quiz, id_user)
		}
		//12
		if kategori.Tabel == "skor_kuliah_sosial" {
			skoringRepoService.SkoringKuliahSosial(id_quiz, id_user)
		}
		//13
		if kategori.Tabel == "skor_kuliah_agama" {
			skoringRepoService.SkoringKuliahAgama(id_quiz, id_user)
		}
		//14
		if kategori.Tabel == "skor_mbti" {
			skoringRepoService.SkoringMBTI(id_quiz, id_user)
		}
		//15
		if kategori.Tabel == "skor_karakteristik_pribadi" {
			skoringRepoService.SkoringKarakteristikPribadi(id_quiz, id_user)
		}
		//16
		if kategori.Tabel == "skor_minat_indonesia" {
			skoringRepoService.SkoringTesMinatIndonesia(id_quiz, id_user)
		}
		//17
		if kategori.Tabel == "skor_kecerdasan_majemuk" {
			skoringRepoService.SkoringKecerdasanMajemuk(id_quiz, id_user)
		}
		//18
		if kategori.Tabel == "skor_suasana_kerja" {
			skoringRepoService.SkoringSuasanaKerja(id_quiz, id_user)
		}
		//19
		if kategori.Tabel == "skor_gaya_belajar" {
			skoringRepoService.SkoringGayaBelajar(id_quiz, id_user)
		}
		//20
		if kategori.Tabel == "skor_kejiwaan_dewasa" {
			skoringRepoService.SkoringKejiwaanDewasa(id_quiz, id_user)
		}
		//21
		if kategori.Tabel == "skor_kesehatan_mental" {
			skoringRepoService.SkoringKesehatanMental(id_quiz, id_user)
		}
		//22
		if kategori.Tabel == "skor_mode_belajar" {
			skoringRepoService.SkoringModeBelajar(id_quiz, id_user)
		}
		//23
		if kategori.Tabel == "skor_ssct" {
			skoringRepoService.SkoringSSCT(id_quiz, id_user)
		}
	}
	return nil
}
