package repository

import "irwanka/sicerdas/entity"

//SOAL KOGNITIF
//Minat Kuliah Eksakta
//Minat Kuliah Sosial

type SinkronRepository interface {
	//for table
	GetAllSoalKognitif() ([]*entity.SoalKognitif, error)
	UpdateGambarSoalKognitif(id int32, url string) error
	UpdateRefGambar(filename string, url string) error

	GetAllSoalMinatKuliahEksakta() ([]*entity.SoalMinatKuliahEksakta, error)
	UpdateGambarSoalMinatKuliahEksakta(id int32, url string) error
}

func NewSinkronRepository() SinkronRepository {
	return &repo{}
}

func (*repo) GetAllSoalKognitif() ([]*entity.SoalKognitif, error) {
	var list []*entity.SoalKognitif
	db.Table("soal_kognitif").Scan(&list)
	return list, nil
}

func (*repo) UpdateGambarSoalKognitif(id int32, url string) error {
	db.Table("soal_kognitif").Where("id_soal", id).UpdateColumn("pertanyaan_gambar", url)
	return nil
}

func (*repo) UpdateRefGambar(filename string, url string) error {
	db.Table("gambar").Where("filename", filename).UpdateColumn("filename", url)
	return nil
}

func (*repo) GetAllSoalMinatKuliahEksakta() ([]*entity.SoalMinatKuliahEksakta, error) {
	var list []*entity.SoalMinatKuliahEksakta
	db.Table("soal_minat_kuliah_eksakta").Scan(&list)
	return list, nil
}

func (*repo) UpdateGambarSoalMinatKuliahEksakta(id int32, url string) error {
	db.Table("soal_minat_kuliah_eksakta").Where("id_soal", id).UpdateColumn("gambar", url)
	return nil
}
