package entity

type SoalKognitif struct {
	IDSoal           int32  `gorm:"column:id_soal;primaryKey;autoIncrement:true" json:"id_soal"`
	Bidang           string `gorm:"column:bidang" json:"bidang"`
	Urutan           int32  `gorm:"column:urutan" json:"urutan"`
	IDPetunjuk       int32  `gorm:"column:id_petunjuk" json:"id_petunjuk"`
	Pertanyaan       string `gorm:"column:pertanyaan" json:"pertanyaan"`
	PertanyaanGambar string `gorm:"column:pertanyaan_gambar" json:"pertanyaan_gambar"`
	PilihanA         string `gorm:"column:pilihan_a" json:"pilihan_a"`
	PilihanAGambar   string `gorm:"column:pilihan_a_gambar" json:"pilihan_a_gambar"`
	PilihanB         string `gorm:"column:pilihan_b" json:"pilihan_b"`
	PilihanBGambar   string `gorm:"column:pilihan_b_gambar" json:"pilihan_b_gambar"`
	PilihanC         string `gorm:"column:pilihan_c" json:"pilihan_c"`
	PilihanCGambar   string `gorm:"column:pilihan_c_gambar" json:"pilihan_c_gambar"`
	PilihanD         string `gorm:"column:pilihan_d" json:"pilihan_d"`
	PilihanDGambar   string `gorm:"column:pilihan_d_gambar" json:"pilihan_d_gambar"`
	PilihanE         string `gorm:"column:pilihan_e" json:"pilihan_e"`
	PilihanEGambar   string `gorm:"column:pilihan_e_gambar" json:"pilihan_e_gambar"`
	PilihanJawaban   string `gorm:"column:pilihan_jawaban" json:"pilihan_jawaban"`
	UUID             string `gorm:"column:uuid" json:"uuid"`
	Paket            string `gorm:"column:paket" json:"paket"`
}
