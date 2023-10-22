// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSkoringPenjurusanKuliahV3 = "skoring_penjurusan_kuliah_v3"

// SkoringPenjurusanKuliahV3 mapped from table <skoring_penjurusan_kuliah_v3>
type SkoringPenjurusanKuliahV3 struct {
	ID             int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	IDUser         int32   `gorm:"column:id_user" json:"id_user"`
	IDQuiz         int32   `gorm:"column:id_quiz" json:"id_quiz"`
	TpaIu          int32   `gorm:"column:tpa_iu" json:"tpa_iu"`
	TpaPv          int32   `gorm:"column:tpa_pv" json:"tpa_pv"`
	TpaPk          int32   `gorm:"column:tpa_pk" json:"tpa_pk"`
	TpaPa          int32   `gorm:"column:tpa_pa" json:"tpa_pa"`
	TpaPm          int32   `gorm:"column:tpa_pm" json:"tpa_pm"`
	TpaKt          int32   `gorm:"column:tpa_kt" json:"tpa_kt"`
	TpaIq          int32   `gorm:"column:tpa_iq" json:"tpa_iq"`
	SkorIq         float32 `gorm:"column:skor_iq" json:"skor_iq"`
	MinatIpa1      int32   `gorm:"column:minat_ipa1" json:"minat_ipa1"`
	MinatIpa2      int32   `gorm:"column:minat_ipa2" json:"minat_ipa2"`
	MinatIpa3      int32   `gorm:"column:minat_ipa3" json:"minat_ipa3"`
	MinatIpa4      int32   `gorm:"column:minat_ipa4" json:"minat_ipa4"`
	MinatIpa5      int32   `gorm:"column:minat_ipa5" json:"minat_ipa5"`
	MinatIps1      int32   `gorm:"column:minat_ips1" json:"minat_ips1"`
	MinatIps2      int32   `gorm:"column:minat_ips2" json:"minat_ips2"`
	MinatIps3      int32   `gorm:"column:minat_ips3" json:"minat_ips3"`
	MinatIps4      int32   `gorm:"column:minat_ips4" json:"minat_ips4"`
	MinatIps5      int32   `gorm:"column:minat_ips5" json:"minat_ips5"`
	SikapAgm       int32   `gorm:"column:sikap_agm" json:"sikap_agm"`
	SikapPkn       int32   `gorm:"column:sikap_pkn" json:"sikap_pkn"`
	SikapInd       int32   `gorm:"column:sikap_ind" json:"sikap_ind"`
	SikapEng       int32   `gorm:"column:sikap_eng" json:"sikap_eng"`
	SikapMat       int32   `gorm:"column:sikap_mat" json:"sikap_mat"`
	SikapFis       int32   `gorm:"column:sikap_fis" json:"sikap_fis"`
	SikapBio       int32   `gorm:"column:sikap_bio" json:"sikap_bio"`
	SikapKim       int32   `gorm:"column:sikap_kim" json:"sikap_kim"`
	SikapEko       int32   `gorm:"column:sikap_eko" json:"sikap_eko"`
	SikapSej       int32   `gorm:"column:sikap_sej" json:"sikap_sej"`
	SikapSos       int32   `gorm:"column:sikap_sos" json:"sikap_sos"`
	SikapGeo       int32   `gorm:"column:sikap_geo" json:"sikap_geo"`
	SikapSbd       int32   `gorm:"column:sikap_sbd" json:"sikap_sbd"`
	SikapOrg       int32   `gorm:"column:sikap_org" json:"sikap_org"`
	SikapMlk       int32   `gorm:"column:sikap_mlk" json:"sikap_mlk"`
	SikapTik       int32   `gorm:"column:sikap_tik" json:"sikap_tik"`
	MinatAgm1      int32   `gorm:"column:minat_agm1" json:"minat_agm1"`
	MinatAgm2      int32   `gorm:"column:minat_agm2" json:"minat_agm2"`
	MinatAgm3      int32   `gorm:"column:minat_agm3" json:"minat_agm3"`
	MinatAgm4      int32   `gorm:"column:minat_agm4" json:"minat_agm4"`
	MinatAgm5      int32   `gorm:"column:minat_agm5" json:"minat_agm5"`
	MinatDinas1    string  `gorm:"column:minat_dinas1" json:"minat_dinas1"`
	MinatDinas2    string  `gorm:"column:minat_dinas2" json:"minat_dinas2"`
	MinatDinas3    string  `gorm:"column:minat_dinas3" json:"minat_dinas3"`
	SelesaiSkoring int32   `gorm:"column:selesai_skoring" json:"selesai_skoring"`
	GpA            int32   `gorm:"column:gp_a" json:"gp_a"`
	GpB            int32   `gorm:"column:gp_b" json:"gp_b"`
	GpC            int32   `gorm:"column:gp_c" json:"gp_c"`
	GpD            int32   `gorm:"column:gp_d" json:"gp_d"`
	GpE            int32   `gorm:"column:gp_e" json:"gp_e"`
	GpF            int32   `gorm:"column:gp_f" json:"gp_f"`
	GpG            int32   `gorm:"column:gp_g" json:"gp_g"`
	GpH            int32   `gorm:"column:gp_h" json:"gp_h"`
	GpI            int32   `gorm:"column:gp_i" json:"gp_i"`
	GpJ            int32   `gorm:"column:gp_j" json:"gp_j"`
	GpK            int32   `gorm:"column:gp_k" json:"gp_k"`
	GpL            int32   `gorm:"column:gp_l" json:"gp_l"`
	RangkingGp1    string  `gorm:"column:rangking_gp1" json:"rangking_gp1"`
	RangkingGp2    string  `gorm:"column:rangking_gp2" json:"rangking_gp2"`
	RangkingGp3    string  `gorm:"column:rangking_gp3" json:"rangking_gp3"`
	TpaPs          int16   `gorm:"column:tpa_ps" json:"tpa_ps"`
	GayaAuditoris  int16   `gorm:"column:gaya_auditoris" json:"gaya_auditoris"`
	GayaVisual     int16   `gorm:"column:gaya_visual" json:"gaya_visual"`
	GayaKinestetik int16   `gorm:"column:gaya_kinestetik" json:"gaya_kinestetik"`
}

// TableName SkoringPenjurusanKuliahV3's table name
func (*SkoringPenjurusanKuliahV3) TableName() string {
	return TableNameSkoringPenjurusanKuliahV3
}
