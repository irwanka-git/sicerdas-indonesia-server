// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkorWlb = "skor_wlb"

// SkorWlb mapped from table <skor_wlb>
type SkorWlb struct {
	IDUser                      int32  `gorm:"column:id_user;primaryKey" json:"id_user"`
	IDQuiz                      int32  `gorm:"column:id_quiz;primaryKey" json:"id_quiz"`
	KedamaianHati               int32  `gorm:"column:kedamaian_hati" json:"kedamaian_hati"`
	PengembanganDiri            int32  `gorm:"column:pengembangan_diri" json:"pengembangan_diri"`
	Ibadah                      int32  `gorm:"column:ibadah" json:"ibadah"`
	Pendapatan                  int32  `gorm:"column:pendapatan" json:"pendapatan"`
	HubunganSosial              int32  `gorm:"column:hubungan_sosial" json:"hubungan_sosial"`
	Kesehatan                   int32  `gorm:"column:kesehatan" json:"kesehatan"`
	RekanKerja                  int32  `gorm:"column:rekan_kerja" json:"rekan_kerja"`
	NilaiKedamaianHati          int32  `gorm:"column:nilai_kedamaian_hati" json:"nilai_kedamaian_hati"`
	NilaiPengembanganDiri       int32  `gorm:"column:nilai_pengembangan_diri" json:"nilai_pengembangan_diri"`
	NilaiIbadah                 int32  `gorm:"column:nilai_ibadah" json:"nilai_ibadah"`
	NilaiPendapatan             int32  `gorm:"column:nilai_pendapatan" json:"nilai_pendapatan"`
	NilaiHubunganSosial         int32  `gorm:"column:nilai_hubungan_sosial" json:"nilai_hubungan_sosial"`
	NilaiKesehatan              int32  `gorm:"column:nilai_kesehatan" json:"nilai_kesehatan"`
	NilaiRekanKerja             int32  `gorm:"column:nilai_rekan_kerja" json:"nilai_rekan_kerja"`
	KlasifikasiKedamaianHati    string `gorm:"column:klasifikasi_kedamaian_hati" json:"klasifikasi_kedamaian_hati"`
	KlasifikasiPengembanganDiri string `gorm:"column:klasifikasi_pengembangan_diri" json:"klasifikasi_pengembangan_diri"`
	KlasifikasiIbadah           string `gorm:"column:klasifikasi_ibadah" json:"klasifikasi_ibadah"`
	KlasifikasiPendapatan       string `gorm:"column:klasifikasi_pendapatan" json:"klasifikasi_pendapatan"`
	KlasifikasiHubunganSosial   string `gorm:"column:klasifikasi_hubungan_sosial" json:"klasifikasi_hubungan_sosial"`
	KlasifikasiKesehatan        string `gorm:"column:klasifikasi_kesehatan" json:"klasifikasi_kesehatan"`
	KlasifikasiRekanKerja       string `gorm:"column:klasifikasi_rekan_kerja" json:"klasifikasi_rekan_kerja"`
}

// TableName SkorWlb's table name
func (*SkorWlb) TableName() string {
	return TableNameSkorWlb
}
