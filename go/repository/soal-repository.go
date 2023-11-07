package repository

import (
	"fmt"
	"html"
	"irwanka/sicerdas/entity"
	"irwanka/sicerdas/helper"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type SoalRepository interface {
	//for api
	GetSoalBreak(token string) ([]*entity.SoalSession, error)
	GetSoalKognitif(paket string, bidang string, token string) ([]*entity.SoalSession, error)
	GetSoalPeminatanSMK(id_quiz int32, paket string, demo bool, token string) ([]*entity.SoalSession, error)
	GetSoalSikapPelajaran(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalSikapPelajaranKuliah(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalTesMinatIndonesia(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalTipologiJung(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalKarakteristikPribadi(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalSkalaPeminatanSMA(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalSkalaPeminatanMAN(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalMinatKuliahEksakta(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalMinatKuliahSosial(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalMinatKuliahDinas(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalMinatKuliahAgama(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalSuasanaKerja(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalKecerdasanMajemuk(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalGayaPekerjaan(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalGayaBelajar(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalTesModeBelajar(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalSSCTRemaja(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalKesehatanMentalID(token string, demo bool) ([]*entity.SoalSession, error)
	GetSoalKejiwaanDewasaID(token string, demo bool) ([]*entity.SoalSession, error)
}

func NewSoalRepository() SoalRepository {
	return &repo{}
}

func (*repo) GetSoalBreak(token string) ([]*entity.SoalSession, error) {
	var soal = entity.SoalSession{}
	soal.Token = token
	soal.Kategori = "BREAK"
	soal.MaxSikap = 0
	soal.MinSikap = 0

	soal.Sn1 = ""
	soal.Sn2 = ""
	soal.Sn3 = ""
	soal.Sp1 = ""
	soal.Sp2 = ""
	soal.Sp3 = ""

	soal.Pernyataan = ""
	soal.Pilihan = []*entity.PilihanJawaban{}
	soal.Mode = "BREAK"

	soal.Gambar = ""
	soal.Nomor = ""
	soal.Uuid = ""
	var listSoal = []*entity.SoalSession{}
	listSoal = append(listSoal, &soal)
	return listSoal, nil
}

func (*repo) GetSoalKognitif(paket string, bidang string, token string) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan      int    `json:"urutan"`
		Uuid        string `json:"uuid"`
		Bidang      string `json:"bidang"`
		Pertanyaan  string `json:"pertanyaan"`
		ImageBase64 string `json:"image_base64"`
		TypeImage   string `json:"type_image"`
		PilihanA    string `json:"pilihan_a"`
		PilihanB    string `json:"pilihan_b"`
		PilihanC    string `json:"pilihan_c"`
		PilihanD    string `json:"pilihan_d"`
		PilihanE    string `json:"pilihan_e"`
		IsiPetunjuk string `json:"isi_petunjuk"`
	}
	db.Raw(`SELECT
				a.urutan,
				a.uuid,
				a.bidang,
				a.pertanyaan,
				c.image_base64,
				c.type as type_image,
				a.pilihan_a,
				a.pilihan_b,
				a.pilihan_c,
				a.pilihan_d,
				a.pilihan_e,
				b.isi_petunjuk 
			FROM
				soal_kognitif AS a
				LEFT JOIN petunjuk_soal AS b ON a.id_petunjuk = b.id_petunjuk 
				LEFT JOIN gambar AS c ON a.pertanyaan_gambar = c.filename 
				where a.bidang = ?
				and a.paket =  ?
			ORDER BY
				bidang, urutan`, bidang, paket).Scan(&listResultSoal)

	for i := 0; i < len(listResultSoal); i++ {
		var prefix_kategori = "KOGNITIF_"
		if paket != "NON" {
			prefix_kategori = prefix_kategori + paket + "_"
		}
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pertanyaan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		if listResultSoal[i].ImageBase64 != "" {
			var typeImage = "data:image/jpg;base64"
			if listResultSoal[i].TypeImage == "image/png" {
				typeImage = "data:image/png;base64"
			}
			var src = fmt.Sprintf("%s,%s", typeImage, listResultSoal[i].ImageBase64)
			var gambar = "<br><img width=\"100%\"" + fmt.Sprintf(" src=\"%v\"", src) + ">"
			gambar = html.UnescapeString(gambar)
			pertanyaan = pertanyaan + gambar
		}

		var pilihan = []*entity.PilihanJawaban{}
		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanD != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanD
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanE != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanE
			tmp.Value = "E"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = prefix_kategori + bidang
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = ""
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)

		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = "<p>" + listResultSoal[i].IsiPetunjuk + "</p>"
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}

	return listSoal, nil
}

func (*repo) GetSoalPeminatanSMK(id_quiz int32, paket string, demo bool, token string) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}

	var listResultSoal []struct {
		Urutan      string `json:"urutan"`
		Uuid        string `json:"uuid"`
		Kegiatan    string `json:"kegiatan"`
		Gambar      string `json:"gambar"`
		Pertanyaan  string `json:"pertanyaan"`
		ImageBase64 string `json:"image_base64"`
		TypeImage   string `json:"type_image"`
	}

	if demo {
		db.Raw(`select x.*, 
				c.image_base64,
				c.type as type_image from (select 
				a.nomor as urutan,
				a.uuid,
				a.kegiatan,					
				a.gambar
				from soal_peminatan_smk as a   
				) as x 
				LEFT JOIN gambar AS c ON x.gambar = c.filename
				order by x.urutan limit 3`).Scan(&listResultSoal)
	} else {
		if paket == "" {
			db.Raw(`select x.*, 
				c.image_base64,
				c.type as type_image from (select 
				a.nomor as urutan,
				a.uuid,
				a.kegiatan,					
				a.gambar
				from soal_peminatan_smk as a, quiz_sesi_mapping_smk as b   
				where a.id_kegiatan = b.id_kegiatan
				and b.id_quiz = ? ) as x 
				LEFT JOIN gambar AS c ON x.gambar = c.filename
				order by x.urutan`, id_quiz).Scan(&listResultSoal)
		} else {
			db.Raw(`select x.*, 
				c.image_base64,
				c.type as type_image from (select 
				a.nomor as urutan,
				a.uuid,
				a.kegiatan,					
				a.gambar
				from soal_peminatan_smk as a, quiz_sesi_mapping_smk as b   
				where a.id_kegiatan = b.id_kegiatan
				and b.id_quiz = ? and a.paket=?  ) as x 
				LEFT JOIN gambar AS c ON x.gambar = c.filename
				order by x.urutan`, id_quiz, paket).Scan(&listResultSoal)
		}
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		var gambar = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Kegiatan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "class=\"ql-align-justify\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "style=\"color: black;\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span >", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span >", "")

		// if listResultSoal[i].ImageBase64 != "" {
		// 	var typeImage = "data:image/jpeg;base64"
		// 	if listResultSoal[i].TypeImage == "image/png" {
		// 		typeImage = "data:image/png;base64"
		// 	}
		// 	gambar = fmt.Sprintf("%s,%s", typeImage, listResultSoal[i].ImageBase64)
		// }

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PEMINATAN_SMK"
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = gambar
		soal.Nomor = listResultSoal[i].Urutan

		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSikapPelajaran(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan        string `json:"urutan"`
		Uuid          string `json:"uuid"`
		Pelajaran     string `json:"pelajaran"`
		SikapNegatif1 string `json:"sikap_negatif1"`
		SikapNegatif2 string `json:"sikap_negatif2"`
		SikapNegatif3 string `json:"sikap_negatif3"`
		SikapPositif1 string `json:"sikap_positif1"`
		SikapPositif2 string `json:"sikap_positif2"`
		SikapPositif3 string `json:"sikap_positif3"`
		Kelompok      string `json:"kelompok"`
	}

	if demo {
		db.Raw(`SELECT
					urutan,
					uuid,
					pelajaran,
					sikap_negatif1,
					sikap_negatif2,
					sikap_negatif3,
					sikap_positif1,
					sikap_positif2,
					sikap_positif3,
					kelompok
				FROM
					soal_sikap_pelajaran
					order by urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					urutan,
					uuid,
					pelajaran,
					sikap_negatif1,
					sikap_negatif2,
					sikap_negatif3,
					sikap_positif1,
					sikap_positif2,
					sikap_positif3,
					kelompok
				FROM
					soal_sikap_pelajaran
					order by urutan`).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pelajaran)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = "<strong>" + pertanyaan + "</strong> adalah pelajaran yang ..."

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SIKAP_TERHADAP_PELAJARAN"
		soal.MaxSikap = 7
		soal.MinSikap = 0

		soal.Gambar = ""
		urutan, _ := strconv.Atoi(listResultSoal[i].Urutan)
		soal.Nomor = fmt.Sprintf("%02d", urutan)
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = listResultSoal[i].SikapNegatif1
		soal.Sn2 = listResultSoal[i].SikapNegatif2
		soal.Sn3 = listResultSoal[i].SikapNegatif3
		soal.Sp1 = listResultSoal[i].SikapPositif1
		soal.Sp2 = listResultSoal[i].SikapPositif2
		soal.Sp3 = listResultSoal[i].SikapPositif3

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "RT"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSikapPelajaranKuliah(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan        string `json:"urutan"`
		Uuid          string `json:"uuid"`
		Pelajaran     string `json:"pelajaran"`
		SikapNegatif1 string `json:"sikap_negatif1"`
		SikapNegatif2 string `json:"sikap_negatif2"`
		SikapNegatif3 string `json:"sikap_negatif3"`
		SikapPositif1 string `json:"sikap_positif1"`
		SikapPositif2 string `json:"sikap_positif2"`
		SikapPositif3 string `json:"sikap_positif3"`
		Kelompok      string `json:"kelompok"`
	}

	if demo {
		db.Raw(`SELECT
					urutan,
					uuid,
					pelajaran,
					sikap_negatif1,
					sikap_negatif2,
					sikap_negatif3,
					sikap_positif1,
					sikap_positif2,
					sikap_positif3,
					kelompok
				FROM
					soal_sikap_pelajaran_kuliah
					order by urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					urutan,
					uuid,
					pelajaran,
					sikap_negatif1,
					sikap_negatif2,
					sikap_negatif3,
					sikap_positif1,
					sikap_positif2,
					sikap_positif3,
					kelompok
				FROM
					soal_sikap_pelajaran_kuliah
					order by kelompok, urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pelajaran)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = "<strong>" + pertanyaan + "</strong> adalah pelajaran yang ..."

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PMK_SIKAP_PELAJARAN"
		soal.MaxSikap = 7
		soal.MinSikap = 0

		soal.Gambar = ""
		urutan, _ := strconv.Atoi(listResultSoal[i].Urutan)
		soal.Nomor = fmt.Sprintf("%02d", urutan)
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = listResultSoal[i].SikapNegatif1
		soal.Sn2 = listResultSoal[i].SikapNegatif2
		soal.Sn3 = listResultSoal[i].SikapNegatif3
		soal.Sp1 = listResultSoal[i].SikapPositif1
		soal.Sp2 = listResultSoal[i].SikapPositif2
		soal.Sp3 = listResultSoal[i].SikapPositif3

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "RT"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalTesMinatIndonesia(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     string `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
	}

	if demo {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.pernyataan
				FROM
					soal_tmi as a order by a.urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
				a.urutan,
				a.uuid,
				a.pernyataan
				FROM soal_tmi as a order by a.urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var soal = entity.SoalSession{}
		soal.Token = token
		urutan, _ := strconv.Atoi(listResultSoal[i].Urutan)
		soal.Nomor = fmt.Sprintf("%02d", urutan)
		soal.Kategori = "SKALA_TES_MINAT_INDONESIA"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}

	return listSoal, nil
}

func (*repo) GetSoalTipologiJung(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     int32  `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
		PilihanA   string `json:"pilihan_a"`
		PilihanB   string `json:"pilihan_b"`
	}

	if demo {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.pernyataan,
					a.pilihan_a,			
					a.pilihan_b
				FROM
					soal_tipologi_jung as a order by a.urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.pernyataan,
					a.pilihan_a,			
					a.pilihan_b
				FROM
					soal_tipologi_jung as a
					order by a.urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}
		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token

		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "SKALA_TES_TIPOLOGI_JUNG"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalKarakteristikPribadi(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     int32  `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
		Pilihan_1  string `json:"pilihan_1"`
		Pilihan_2  string `json:"pilihan_2"`
		Pilihan_3  string `json:"pilihan_3"`
		Pilihan_4  string `json:"pilihan_4"`
	}

	if demo {
		db.Raw(`SELECT a.urutan,
				a.uuid,
				a.pernyataan,
				a.pilihan_1,
				a.pilihan_2,
				a.pilihan_3,
				a.pilihan_4
				FROM
					soal_karakteristik_pribadi as a order by a.urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT a.urutan,
					a.uuid,
					a.pernyataan,
					a.pilihan_1,
					a.pilihan_2,
					a.pilihan_3,
					a.pilihan_4
				FROM
					soal_karakteristik_pribadi as a
					order by a.urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].Pilihan_1 != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].Pilihan_1
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].Pilihan_2 != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].Pilihan_2
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].Pilihan_3 != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].Pilihan_3
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].Pilihan_4 != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].Pilihan_4
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "SKALA_TES_KARAKTERISTIK_PRIBADI"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSkalaPeminatanSMA(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     int32  `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
		PilihanA   string `json:"pilihan_a"`
		PilihanB   string `json:"pilihan_b"`
		PilihanC   string `json:"pilihan_c"`
	}

	if demo {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c
				FROM
					soal_peminatan_sma 
					order by urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c
				FROM
					soal_peminatan_sma 
					order by urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "SKALA_PEMINATAN_SMA"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSkalaPeminatanMAN(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     int32  `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
		PilihanA   string `json:"pilihan_a"`
		PilihanB   string `json:"pilihan_b"`
		PilihanC   string `json:"pilihan_c"`
		PilihanD   string `json:"pilihan_d"`
	}

	if demo {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c,
					pilihan_d
				FROM
					soal_peminatan_man 
					order by urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c,
					pilihan_d
				FROM
					soal_peminatan_man 
					order by urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanD != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanD
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "SKALA_PEMINATAN_MAN"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalMinatKuliahEksakta(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}

	var listResultSoal []struct {
		Urutan    int32  `json:"urutan"`
		Uuid      string `json:"uuid"`
		Indikator string `json:"indikator"`
		Minat     string `json:"minat"`
	}

	if demo {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.minat
				FROM
					soal_minat_kuliah_eksakta as a 
							order by a.urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.minat
				FROM
					soal_minat_kuliah_eksakta as a 
							order by a.urutan asc`).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		var gambar = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Indikator)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "class=\"ql-align-justify\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "style=\"color: black;\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span >", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span >", "")
		// pertanyaan = strip.StripTags(pertanyaan)

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PMK_MINAT_ILMU_ALAM"
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = gambar
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)

		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalMinatKuliahSosial(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}

	var listResultSoal []struct {
		Urutan    int32  `json:"urutan"`
		Uuid      string `json:"uuid"`
		Indikator string `json:"indikator"`
		Minat     string `json:"minat"`
	}

	if demo {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.minat
				FROM
				soal_minat_kuliah_sosial as a 
							order by a.urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.minat
				FROM
				soal_minat_kuliah_sosial as a 
							order by a.urutan asc`).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		var gambar = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Indikator)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "class=\"ql-align-justify\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "style=\"color: black;\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span >", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span >", "")
		// pertanyaan = strip.StripTags(pertanyaan)

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PMK_MINAT_ILMU_SOSIAL"
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = gambar
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)

		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalMinatKuliahDinas(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Nomor       string `json:"nomor"`
		Uuid        string `json:"uuid"`
		Deskripsi   string `json:"deskripsi"`
		PernyataanA string `json:"pernyataan_a"`
		PernyataanB string `json:"pernyataan_b"`
		PernyataanC string `json:"pernyataan_c"`
		PernyataanD string `json:"pernyataan_d"`
		PernyataanE string `json:"pernyataan_e"`
		PernyataanF string `json:"pernyataan_f"`
		PernyataanG string `json:"pernyataan_g"`
		PernyataanH string `json:"pernyataan_h"`
		PernyataanI string `json:"pernyataan_i"`
		PernyataanJ string `json:"pernyataan_j"`
		PernyataanK string `json:"pernyataan_k"`
		PernyataanL string `json:"pernyataan_l"`
	}

	if demo {
		db.Raw(`SELECT
					*
				FROM
					soal_minat_kuliah_dinas 
					order by nomor limit 1 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					*
				FROM
					soal_minat_kuliah_dinas 
					order by nomor`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = listResultSoal[i].Deskripsi
		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PernyataanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanD != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanD
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanE != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanE
			tmp.Value = "E"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanF != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanF
			tmp.Value = "F"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanG != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanG
			tmp.Value = "G"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanH != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanH
			tmp.Value = "H"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanI != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanI
			tmp.Value = "I"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanJ != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanJ
			tmp.Value = "J"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanK != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanK
			tmp.Value = "K"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanL != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanL
			tmp.Value = "L"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		urutan, _ := strconv.Atoi(listResultSoal[i].Nomor)
		soal.Nomor = fmt.Sprintf("%02d", urutan)
		soal.Kategori = "SKALA_PMK_SEKOLAH_KEDINASAN"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalMinatKuliahAgama(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}

	var listResultSoal []struct {
		Urutan    int32  `json:"urutan"`
		Uuid      string `json:"uuid"`
		Indikator string `json:"indikator"`
		Jurusan   string `json:"jurusan"`
	}

	if demo {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.jurusan
				FROM
					soal_minat_kuliah_agama as a
							order by urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					a.urutan,
					a.uuid,
					a.indikator, 
					a.jurusan
				FROM
					soal_minat_kuliah_agama as a
					order by a.urutan`).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		var gambar = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Indikator)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "class=\"ql-align-justify\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "style=\"color: black;\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span >", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span >", "")
		// pertanyaan = strip.StripTags(pertanyaan)

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PMK_ILMU_AGAMA"
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = gambar
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)

		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSuasanaKerja(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}

	var listResultSoal []struct {
		Urutan   string `json:"urutan"`
		Uuid     string `json:"uuid"`
		Kegiatan string `json:"kegiatan"`
	}

	if demo {
		db.Raw(`select x.*, 
				c.image_base64,
				c.type as type_image from (select 
				a.nomor as urutan,
				a.uuid,
				a.kegiatan,					
				a.gambar
				from soal_minat_kuliah_suasana_kerja as a
				) as x 
				LEFT JOIN gambar AS c ON x.gambar = c.filename
				order by x.urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`select x.*, 
				c.image_base64,
				c.type as type_image from (select 
				a.nomor as urutan,
				a.uuid,
				a.kegiatan,					
				a.gambar
				from soal_minat_kuliah_suasana_kerja as a
				) as x 
				LEFT JOIN gambar AS c ON x.gambar = c.filename
				order by x.urutan `).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		var gambar = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Kegiatan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "class=\"ql-align-justify\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "style=\"color: black;\"", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</p>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "</span >", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span>", "")
		pertanyaan = strings.ReplaceAll(pertanyaan, "<span >", "")
		// pertanyaan = strip.StripTags(pertanyaan)

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SKALA_PMK_SUASANA_KERJA"
		soal.MaxSikap = 0
		soal.MinSikap = 0

		soal.Gambar = gambar
		soal.Nomor = listResultSoal[i].Urutan
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "TOP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalKecerdasanMajemuk(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Nomor       string `json:"nomor"`
		Uuid        string `json:"uuid"`
		Deskripsi   string `json:"deskripsi"`
		PernyataanA string `json:"pernyataan_a"`
		PernyataanB string `json:"pernyataan_b"`
		PernyataanC string `json:"pernyataan_c"`
		PernyataanD string `json:"pernyataan_d"`
		PernyataanE string `json:"pernyataan_e"`
		PernyataanF string `json:"pernyataan_f"`
		PernyataanG string `json:"pernyataan_g"`
		PernyataanH string `json:"pernyataan_h"`
		PernyataanI string `json:"pernyataan_i"`
		PernyataanJ string `json:"pernyataan_j"`
		PernyataanK string `json:"pernyataan_k"`
		PernyataanL string `json:"pernyataan_l"`
	}

	if demo {
		db.Raw(`SELECT
					*
				FROM
					soal_kecerdasan_majemuk 
					order by nomor limit 1 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					*
				FROM
					soal_kecerdasan_majemuk 
					order by nomor`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = listResultSoal[i].Deskripsi
		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PernyataanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanD != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanD
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanE != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanE
			tmp.Value = "E"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanF != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanF
			tmp.Value = "F"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanG != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanG
			tmp.Value = "G"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanH != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanH
			tmp.Value = "H"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanI != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanI
			tmp.Value = "I"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanJ != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanJ
			tmp.Value = "J"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanK != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanK
			tmp.Value = "K"
			pilihan = append(pilihan, &tmp)
		}

		if listResultSoal[i].PernyataanL != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PernyataanL
			tmp.Value = "L"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		urutan, _ := strconv.Atoi(listResultSoal[i].Nomor)
		soal.Nomor = fmt.Sprintf("%02d", urutan)
		soal.Kategori = "SKALA_KECERDASAN_MAJEMUK"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalGayaPekerjaan(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Nomor     int32  `json:"urutan"`
		Uuid      string `json:"uuid"`
		Deskripsi string `json:"pernyataan"`
	}

	if demo {
		db.Raw(`SELECT
					a.nomor,
					a.deskripsi, 
					a.uuid
				FROM
					soal_gaya_pekerjaan as a order by a.nomor limit 2`).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
				a.nomor,
				a.deskripsi, 
				a.uuid
			FROM
				soal_gaya_pekerjaan as a
				order by a.nomor asc`).Scan(&listResultSoal)
	}
	var listPilihanJawaban []struct {
		Jawaban string `json:"jawaban"`
		Respon  string `json:"respon"`
	}

	db.Raw(`SELECT
				a.jawaban,
				a.respon
			FROM
			ref_skor_gaya_pekerjaan as a
				order by a.jawaban asc`).Scan(&listPilihanJawaban)

	var pilihan = []*entity.PilihanJawaban{}
	for i := 0; i < len(listPilihanJawaban); i++ {
		// pilihan = append(pilihan)
		var tmp = entity.PilihanJawaban{}
		tmp.Text = listPilihanJawaban[i].Respon
		tmp.Value = listPilihanJawaban[i].Jawaban
		pilihan = append(pilihan, &tmp)
	}
	// fmt.Println(pilihan)

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Deskripsi)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Nomor)
		soal.Kategori = "SKALA_GAYA_PEKERJAAN"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalGayaBelajar(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan     int32  `json:"urutan"`
		Uuid       string `json:"uuid"`
		Pernyataan string `json:"pernyataan"`
		PilihanA   string `json:"pilihan_a"`
		PilihanB   string `json:"pilihan_b"`
		PilihanC   string `json:"pilihan_c"`
	}

	if demo {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c
				FROM
				soal_gaya_belajar 
					order by urutan limit 3 `).Scan(&listResultSoal)
	} else {
		db.Raw(`SELECT
					urutan,
					uuid,
					pernyataan,					
					pilihan_a,
					pilihan_b,
					pilihan_c
				FROM
				soal_gaya_belajar 
					order by urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Pernyataan)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanA
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanB
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = listResultSoal[i].PilihanC
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "SKALA_GAYA_BELAJAR"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PG"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalTesModeBelajar(token string, demo bool) ([]*entity.SoalSession, error) {

	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan   int32  `json:"urutan"`
		Uuid     string `json:"uuid"`
		Soal     string `json:"soal"`
		PilihanA string `json:"pilihan_a"`
		PilihanB string `json:"pilihan_b"`
		PilihanC string `json:"pilihan_c"`
		PilihanD string `json:"pilihan_d"`
		PilihanE string `json:"pilihan_e"`
	}

	if demo {
		db.Raw(`select 
				a.urutan,
				a.soal, 
				a.uuid,
				a.pilihan_a,
				a.pilihan_b, 
				a.pilihan_c, 
				a.pilihan_d, 
				a.pilihan_e
				from soal_mode_belajar as a 
				order by urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`select 
					a.urutan,
					a.soal, 
					a.uuid,
					a.pilihan_a,
					a.pilihan_b, 
					a.pilihan_c, 
					a.pilihan_d, 
					a.pilihan_e
				from soal_mode_belajar as a 
					order by urutan`).Scan(&listResultSoal)
	}
	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].Soal)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")

		var pilihan = []*entity.PilihanJawaban{}

		if listResultSoal[i].PilihanA != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = helper.GetPilihanFromTagDot(listResultSoal[i].PilihanA, 1)
			tmp.Value = "A"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanB != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = helper.GetPilihanFromTagDot(listResultSoal[i].PilihanB, 1)
			tmp.Value = "B"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanC != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = helper.GetPilihanFromTagDot(listResultSoal[i].PilihanC, 1)
			tmp.Value = "C"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanD != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = helper.GetPilihanFromTagDot(listResultSoal[i].PilihanD, 1)
			tmp.Value = "D"
			pilihan = append(pilihan, &tmp)
		}
		if listResultSoal[i].PilihanE != "" {
			var tmp = entity.PilihanJawaban{}
			tmp.Text = helper.GetPilihanFromTagDot(listResultSoal[i].PilihanE, 1)
			tmp.Value = "E"
			pilihan = append(pilihan, &tmp)
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Kategori = "TES_MODE_BELAJAR"

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Section = ""
		soal.Sp3 = ""

		soal.Pernyataan = pertanyaan
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PP"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalSSCTRemaja(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listResultSoal []struct {
		Urutan          int32  `json:"urutan"`
		Uuid            string `json:"uuid"`
		SubjekPenilaian string `json:"subjek_penilaian"`
		SikapNegatif1   string `json:"sikap_negatif1"`
		SikapNegatif2   string `json:"sikap_negatif2"`
		SikapNegatif3   string `json:"sikap_negatif3"`
		SikapPositif1   string `json:"sikap_positif1"`
		SikapPositif2   string `json:"sikap_positif2"`
		SikapPositif3   string `json:"sikap_positif3"`
	}

	if demo {
		db.Raw(`select 
				urutan, 
				subjek_penilaian,
				sikap_negatif1, 
				sikap_negatif2, 
				sikap_negatif3,
				sikap_positif1, 
				sikap_positif2, 
				sikap_positif3,
				uuid
		 		from soal_ssct_remaja as a  
					order by urutan limit 3`).Scan(&listResultSoal)
	} else {
		db.Raw(`select 
				urutan, 
				subjek_penilaian,
				sikap_negatif1, 
				sikap_negatif2, 
				sikap_negatif3,
				sikap_positif1, 
				sikap_positif2, 
				sikap_positif3,
				uuid
				from soal_ssct_remaja as a 
					order by urutan`).Scan(&listResultSoal)
	}

	for i := 0; i < len(listResultSoal); i++ {
		var pertanyaan = ""
		pertanyaan = html.UnescapeString(listResultSoal[i].SubjekPenilaian)
		pertanyaan = strings.ReplaceAll(pertanyaan, "&hellip;", "")
		pertanyaan = "<strong>" + pertanyaan + "</strong>"

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Kategori = "SSCT_REMAJA"
		soal.MaxSikap = 7
		soal.MinSikap = 0

		soal.Gambar = ""
		soal.Nomor = fmt.Sprintf("%02d", listResultSoal[i].Urutan)
		soal.Uuid = listResultSoal[i].Uuid

		soal.Sn1 = listResultSoal[i].SikapNegatif1
		soal.Sn2 = listResultSoal[i].SikapNegatif2
		soal.Sn3 = listResultSoal[i].SikapNegatif3
		soal.Sp1 = listResultSoal[i].SikapPositif1
		soal.Sp2 = listResultSoal[i].SikapPositif2
		soal.Sp3 = listResultSoal[i].SikapPositif3

		soal.Pernyataan = pertanyaan
		soal.Pilihan = []*entity.PilihanJawaban{}
		soal.Petunjuk = ""
		soal.Mode = "RT"
		soal.PernyataanMulti = []*entity.ItemSoalMulti{}
		listSoal = append(listSoal, &soal)
	}

	return listSoal, nil
}

func (*repo) GetSoalKesehatanMentalID(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listModel []struct {
		IdModel int32 `json:"id_model"`
	}
	if demo {
		db.Raw(`select a.id_model from 
			soal_kesehatan_mental as a, ref_model_kesehatan_mental as b 
			where a.id_model = b.id
			GROUP BY a.id_model, b.nama , b.id 
			order by a.id_model limit 1`).Scan(&listModel)
	} else {
		db.Raw(`select a.id_model from 
			soal_kesehatan_mental as a, ref_model_kesehatan_mental as b 
			where a.id_model = b.id
			GROUP BY a.id_model, b.nama , b.id 
			order by a.id_model `).Scan(&listModel)
	}
	var pilihan = []*entity.PilihanJawaban{}
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Ya",
		Value: "3",
	})
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Kadang",
		Value: "2",
	})
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Tidak",
		Value: "1",
	})

	for i := 0; i < len(listModel); i++ {
		section := "Section " + fmt.Sprintf("%02d", i+1)
		id_model := listModel[i].IdModel
		var listMultiSoal []struct {
			Unsur string `json:"unsur"`
		}
		var pertanyaan_multi = []*entity.ItemSoalMulti{}
		if demo {
			db.Raw(`select a.unsur from soal_kesehatan_mental as a 
					where a.id_model = ? order by a.urutan `, id_model).Scan(&listMultiSoal)
		} else {
			db.Raw(`select a.unsur from soal_kesehatan_mental as a 
					where a.id_model = ? order by a.urutan`, id_model).Scan(&listMultiSoal)
		}
		for n := 0; n < len(listMultiSoal); n++ {

			pertanyaan_multi = append(pertanyaan_multi, &entity.ItemSoalMulti{
				Nomor:      n + 1,
				Pernyataan: listMultiSoal[n].Unsur,
			})
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", i+1)
		soal.Kategori = "TES_KESEHATAN_MENTAL_ID"
		soal.PernyataanMulti = pertanyaan_multi

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = uuid.NewString()

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = ""
		soal.PernyataanMulti = pertanyaan_multi
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PGS"
		soal.Section = section
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}

func (*repo) GetSoalKejiwaanDewasaID(token string, demo bool) ([]*entity.SoalSession, error) {
	var listSoal = []*entity.SoalSession{}
	var listModel []struct {
		IdModel int32 `json:"id_model"`
	}
	if demo {
		db.Raw(`select a.id_model from 
				soal_kejiwaan_dewasa as a, ref_model_kejiwaan_dewasa as b 
				where a.id_model = b.id
				GROUP BY a.id_model, b.nama , b.id 
			order by a.id_model limit 1`).Scan(&listModel)
	} else {
		db.Raw(`select a.id_model from 
		soal_kejiwaan_dewasa as a, ref_model_kejiwaan_dewasa as b 
		where a.id_model = b.id
		GROUP BY a.id_model, b.nama , b.id 
		order by a.id_model `).Scan(&listModel)
	}
	var pilihan = []*entity.PilihanJawaban{}
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Ya",
		Value: "3",
	})
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Kadang",
		Value: "2",
	})
	pilihan = append(pilihan, &entity.PilihanJawaban{
		Text:  "Tidak",
		Value: "1",
	})

	for i := 0; i < len(listModel); i++ {
		section := "Section " + fmt.Sprintf("%02d", i+1)
		id_model := listModel[i].IdModel
		var listMultiSoal []struct {
			Unsur string `json:"unsur"`
		}
		var pertanyaan_multi = []*entity.ItemSoalMulti{}
		if demo {
			db.Raw(`select a.unsur from soal_kejiwaan_dewasa as a 
					where a.id_model = ? order by a.urutan`, id_model).Scan(&listMultiSoal)
		} else {
			db.Raw(`select a.unsur from soal_kejiwaan_dewasa as a 
					where a.id_model = ? order by a.urutan`, id_model).Scan(&listMultiSoal)
		}
		for n := 0; n < len(listMultiSoal); n++ {

			pertanyaan_multi = append(pertanyaan_multi, &entity.ItemSoalMulti{
				Nomor:      n + 1,
				Pernyataan: listMultiSoal[n].Unsur,
			})
		}

		var soal = entity.SoalSession{}
		soal.Token = token
		soal.Nomor = fmt.Sprintf("%02d", i+1)
		soal.Kategori = "TES_KEJIWAAN_DEWASA_ID"
		soal.PernyataanMulti = pertanyaan_multi

		soal.MaxSikap = 0
		soal.MinSikap = 0
		soal.Gambar = ""
		soal.Uuid = uuid.NewString()

		soal.Sn1 = ""
		soal.Sn2 = ""
		soal.Sn3 = ""
		soal.Sp1 = ""
		soal.Sp2 = ""
		soal.Sp3 = ""

		soal.Pernyataan = ""
		soal.PernyataanMulti = pertanyaan_multi
		soal.Pilihan = pilihan
		soal.Petunjuk = ""
		soal.Mode = "PGS"
		soal.Section = section
		listSoal = append(listSoal, &soal)
	}
	return listSoal, nil
}
