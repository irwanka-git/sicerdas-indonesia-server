// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSoalMinatKuliahDina = "soal_minat_kuliah_dinas"

// SoalMinatKuliahDina mapped from table <soal_minat_kuliah_dinas>
type SoalMinatKuliahDina struct {
	Nomor       string `gorm:"column:nomor;primaryKey" json:"nomor"`
	Deskripsi   string `gorm:"column:deskripsi" json:"deskripsi"`
	PernyataanA string `gorm:"column:pernyataan_a" json:"pernyataan_a"`
	PernyataanB string `gorm:"column:pernyataan_b" json:"pernyataan_b"`
	PernyataanC string `gorm:"column:pernyataan_c" json:"pernyataan_c"`
	PernyataanD string `gorm:"column:pernyataan_d" json:"pernyataan_d"`
	PernyataanE string `gorm:"column:pernyataan_e" json:"pernyataan_e"`
	PernyataanF string `gorm:"column:pernyataan_f" json:"pernyataan_f"`
	PernyataanG string `gorm:"column:pernyataan_g" json:"pernyataan_g"`
	PernyataanH string `gorm:"column:pernyataan_h" json:"pernyataan_h"`
	PernyataanI string `gorm:"column:pernyataan_i" json:"pernyataan_i"`
	PernyataanJ string `gorm:"column:pernyataan_j" json:"pernyataan_j"`
	PernyataanK string `gorm:"column:pernyataan_k" json:"pernyataan_k"`
	PernyataanL string `gorm:"column:pernyataan_l" json:"pernyataan_l"`
	UUID        string `gorm:"column:uuid" json:"uuid"`
}

// TableName SoalMinatKuliahDina's table name
func (*SoalMinatKuliahDina) TableName() string {
	return TableNameSoalMinatKuliahDina
}