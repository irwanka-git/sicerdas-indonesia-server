// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package export

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"irwanka/sicerdas/utils/gen-model/entity"
)

func newSkoringMinatMan(db *gorm.DB, opts ...gen.DOOption) skoringMinatMan {
	_skoringMinatMan := skoringMinatMan{}

	_skoringMinatMan.skoringMinatManDo.UseDB(db, opts...)
	_skoringMinatMan.skoringMinatManDo.UseModel(&entity.SkoringMinatMan{})

	tableName := _skoringMinatMan.skoringMinatManDo.TableName()
	_skoringMinatMan.ALL = field.NewAsterisk(tableName)
	_skoringMinatMan.ID = field.NewInt32(tableName, "id")
	_skoringMinatMan.IDUser = field.NewInt32(tableName, "id_user")
	_skoringMinatMan.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skoringMinatMan.TpaIu = field.NewInt32(tableName, "tpa_iu")
	_skoringMinatMan.TpaPv = field.NewInt32(tableName, "tpa_pv")
	_skoringMinatMan.TpaPk = field.NewInt32(tableName, "tpa_pk")
	_skoringMinatMan.TpaPa = field.NewInt32(tableName, "tpa_pa")
	_skoringMinatMan.TpaPs = field.NewInt32(tableName, "tpa_ps")
	_skoringMinatMan.TpaPm = field.NewInt32(tableName, "tpa_pm")
	_skoringMinatMan.TpaKt = field.NewInt32(tableName, "tpa_kt")
	_skoringMinatMan.TpaIq = field.NewInt32(tableName, "tpa_iq")
	_skoringMinatMan.SkorIq = field.NewFloat32(tableName, "skor_iq")
	_skoringMinatMan.MinatSains = field.NewInt32(tableName, "minat_sains")
	_skoringMinatMan.MinatHumaniora = field.NewInt32(tableName, "minat_humaniora")
	_skoringMinatMan.MinatBahasa = field.NewInt32(tableName, "minat_bahasa")
	_skoringMinatMan.MinatAgama = field.NewInt32(tableName, "minat_agama")
	_skoringMinatMan.MinatRentang = field.NewInt32(tableName, "minat_rentang")
	_skoringMinatMan.SikapAgm = field.NewInt32(tableName, "sikap_agm")
	_skoringMinatMan.SikapPkn = field.NewInt32(tableName, "sikap_pkn")
	_skoringMinatMan.SikapInd = field.NewInt32(tableName, "sikap_ind")
	_skoringMinatMan.SikapEng = field.NewInt32(tableName, "sikap_eng")
	_skoringMinatMan.SikapMat = field.NewInt32(tableName, "sikap_mat")
	_skoringMinatMan.SikapFis = field.NewInt32(tableName, "sikap_fis")
	_skoringMinatMan.SikapBio = field.NewInt32(tableName, "sikap_bio")
	_skoringMinatMan.SikapEko = field.NewInt32(tableName, "sikap_eko")
	_skoringMinatMan.SikapSej = field.NewInt32(tableName, "sikap_sej")
	_skoringMinatMan.SikapGeo = field.NewInt32(tableName, "sikap_geo")
	_skoringMinatMan.SikapIlmuAlam = field.NewInt32(tableName, "sikap_ilmu_alam")
	_skoringMinatMan.SikapIlmuSosial = field.NewInt32(tableName, "sikap_ilmu_sosial")
	_skoringMinatMan.SikapRentang = field.NewInt32(tableName, "sikap_rentang")
	_skoringMinatMan.TmiIlmuAlam = field.NewInt32(tableName, "tmi_ilmu_alam")
	_skoringMinatMan.TmiIlmuSosial = field.NewInt32(tableName, "tmi_ilmu_sosial")
	_skoringMinatMan.TmiRentang = field.NewInt32(tableName, "tmi_rentang")
	_skoringMinatMan.TipojungE = field.NewInt32(tableName, "tipojung_e")
	_skoringMinatMan.TipojungI = field.NewInt32(tableName, "tipojung_i")
	_skoringMinatMan.TipojungS = field.NewInt32(tableName, "tipojung_s")
	_skoringMinatMan.TipojungN = field.NewInt32(tableName, "tipojung_n")
	_skoringMinatMan.TipojungT = field.NewInt32(tableName, "tipojung_t")
	_skoringMinatMan.TipojungF = field.NewInt32(tableName, "tipojung_f")
	_skoringMinatMan.TipojungJ = field.NewInt32(tableName, "tipojung_j")
	_skoringMinatMan.TipojungP = field.NewInt32(tableName, "tipojung_p")
	_skoringMinatMan.TipojungKode = field.NewString(tableName, "tipojung_kode")
	_skoringMinatMan.PribadiMotivasi = field.NewInt32(tableName, "pribadi_motivasi")
	_skoringMinatMan.PribadiJuang = field.NewInt32(tableName, "pribadi_juang")
	_skoringMinatMan.PribadiYakin = field.NewInt32(tableName, "pribadi_yakin")
	_skoringMinatMan.PribadiPercaya = field.NewInt32(tableName, "pribadi_percaya")
	_skoringMinatMan.PribadiKonsep = field.NewInt32(tableName, "pribadi_konsep")
	_skoringMinatMan.PribadiKreativitas = field.NewInt32(tableName, "pribadi_kreativitas")
	_skoringMinatMan.PribadiMimpin = field.NewInt32(tableName, "pribadi_mimpin")
	_skoringMinatMan.PribadiEntrepreneur = field.NewInt32(tableName, "pribadi_entrepreneur")
	_skoringMinatMan.PribadiStress = field.NewInt32(tableName, "pribadi_stress")
	_skoringMinatMan.PribadiEmosi = field.NewInt32(tableName, "pribadi_emosi")
	_skoringMinatMan.PribadiSosial = field.NewInt32(tableName, "pribadi_sosial")
	_skoringMinatMan.PribadiEmpati = field.NewInt32(tableName, "pribadi_empati")
	_skoringMinatMan.RekomMinat = field.NewString(tableName, "rekom_minat")
	_skoringMinatMan.RekomSikapPelajaran = field.NewString(tableName, "rekom_sikap_pelajaran")
	_skoringMinatMan.RekomTmi = field.NewString(tableName, "rekom_tmi")
	_skoringMinatMan.RekomAkhir = field.NewString(tableName, "rekom_akhir")
	_skoringMinatMan.SelesaiSkoring = field.NewInt32(tableName, "selesai_skoring")

	_skoringMinatMan.fillFieldMap()

	return _skoringMinatMan
}

type skoringMinatMan struct {
	skoringMinatManDo skoringMinatManDo

	ALL                 field.Asterisk
	ID                  field.Int32
	IDUser              field.Int32
	IDQuiz              field.Int32
	TpaIu               field.Int32
	TpaPv               field.Int32
	TpaPk               field.Int32
	TpaPa               field.Int32
	TpaPs               field.Int32
	TpaPm               field.Int32
	TpaKt               field.Int32
	TpaIq               field.Int32
	SkorIq              field.Float32
	MinatSains          field.Int32
	MinatHumaniora      field.Int32
	MinatBahasa         field.Int32
	MinatAgama          field.Int32
	MinatRentang        field.Int32
	SikapAgm            field.Int32
	SikapPkn            field.Int32
	SikapInd            field.Int32
	SikapEng            field.Int32
	SikapMat            field.Int32
	SikapFis            field.Int32
	SikapBio            field.Int32
	SikapEko            field.Int32
	SikapSej            field.Int32
	SikapGeo            field.Int32
	SikapIlmuAlam       field.Int32
	SikapIlmuSosial     field.Int32
	SikapRentang        field.Int32
	TmiIlmuAlam         field.Int32
	TmiIlmuSosial       field.Int32
	TmiRentang          field.Int32
	TipojungE           field.Int32
	TipojungI           field.Int32
	TipojungS           field.Int32
	TipojungN           field.Int32
	TipojungT           field.Int32
	TipojungF           field.Int32
	TipojungJ           field.Int32
	TipojungP           field.Int32
	TipojungKode        field.String
	PribadiMotivasi     field.Int32
	PribadiJuang        field.Int32
	PribadiYakin        field.Int32
	PribadiPercaya      field.Int32
	PribadiKonsep       field.Int32
	PribadiKreativitas  field.Int32
	PribadiMimpin       field.Int32
	PribadiEntrepreneur field.Int32
	PribadiStress       field.Int32
	PribadiEmosi        field.Int32
	PribadiSosial       field.Int32
	PribadiEmpati       field.Int32
	RekomMinat          field.String
	RekomSikapPelajaran field.String
	RekomTmi            field.String
	RekomAkhir          field.String
	SelesaiSkoring      field.Int32

	fieldMap map[string]field.Expr
}

func (s skoringMinatMan) Table(newTableName string) *skoringMinatMan {
	s.skoringMinatManDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skoringMinatMan) As(alias string) *skoringMinatMan {
	s.skoringMinatManDo.DO = *(s.skoringMinatManDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skoringMinatMan) updateTableName(table string) *skoringMinatMan {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.TpaIu = field.NewInt32(table, "tpa_iu")
	s.TpaPv = field.NewInt32(table, "tpa_pv")
	s.TpaPk = field.NewInt32(table, "tpa_pk")
	s.TpaPa = field.NewInt32(table, "tpa_pa")
	s.TpaPs = field.NewInt32(table, "tpa_ps")
	s.TpaPm = field.NewInt32(table, "tpa_pm")
	s.TpaKt = field.NewInt32(table, "tpa_kt")
	s.TpaIq = field.NewInt32(table, "tpa_iq")
	s.SkorIq = field.NewFloat32(table, "skor_iq")
	s.MinatSains = field.NewInt32(table, "minat_sains")
	s.MinatHumaniora = field.NewInt32(table, "minat_humaniora")
	s.MinatBahasa = field.NewInt32(table, "minat_bahasa")
	s.MinatAgama = field.NewInt32(table, "minat_agama")
	s.MinatRentang = field.NewInt32(table, "minat_rentang")
	s.SikapAgm = field.NewInt32(table, "sikap_agm")
	s.SikapPkn = field.NewInt32(table, "sikap_pkn")
	s.SikapInd = field.NewInt32(table, "sikap_ind")
	s.SikapEng = field.NewInt32(table, "sikap_eng")
	s.SikapMat = field.NewInt32(table, "sikap_mat")
	s.SikapFis = field.NewInt32(table, "sikap_fis")
	s.SikapBio = field.NewInt32(table, "sikap_bio")
	s.SikapEko = field.NewInt32(table, "sikap_eko")
	s.SikapSej = field.NewInt32(table, "sikap_sej")
	s.SikapGeo = field.NewInt32(table, "sikap_geo")
	s.SikapIlmuAlam = field.NewInt32(table, "sikap_ilmu_alam")
	s.SikapIlmuSosial = field.NewInt32(table, "sikap_ilmu_sosial")
	s.SikapRentang = field.NewInt32(table, "sikap_rentang")
	s.TmiIlmuAlam = field.NewInt32(table, "tmi_ilmu_alam")
	s.TmiIlmuSosial = field.NewInt32(table, "tmi_ilmu_sosial")
	s.TmiRentang = field.NewInt32(table, "tmi_rentang")
	s.TipojungE = field.NewInt32(table, "tipojung_e")
	s.TipojungI = field.NewInt32(table, "tipojung_i")
	s.TipojungS = field.NewInt32(table, "tipojung_s")
	s.TipojungN = field.NewInt32(table, "tipojung_n")
	s.TipojungT = field.NewInt32(table, "tipojung_t")
	s.TipojungF = field.NewInt32(table, "tipojung_f")
	s.TipojungJ = field.NewInt32(table, "tipojung_j")
	s.TipojungP = field.NewInt32(table, "tipojung_p")
	s.TipojungKode = field.NewString(table, "tipojung_kode")
	s.PribadiMotivasi = field.NewInt32(table, "pribadi_motivasi")
	s.PribadiJuang = field.NewInt32(table, "pribadi_juang")
	s.PribadiYakin = field.NewInt32(table, "pribadi_yakin")
	s.PribadiPercaya = field.NewInt32(table, "pribadi_percaya")
	s.PribadiKonsep = field.NewInt32(table, "pribadi_konsep")
	s.PribadiKreativitas = field.NewInt32(table, "pribadi_kreativitas")
	s.PribadiMimpin = field.NewInt32(table, "pribadi_mimpin")
	s.PribadiEntrepreneur = field.NewInt32(table, "pribadi_entrepreneur")
	s.PribadiStress = field.NewInt32(table, "pribadi_stress")
	s.PribadiEmosi = field.NewInt32(table, "pribadi_emosi")
	s.PribadiSosial = field.NewInt32(table, "pribadi_sosial")
	s.PribadiEmpati = field.NewInt32(table, "pribadi_empati")
	s.RekomMinat = field.NewString(table, "rekom_minat")
	s.RekomSikapPelajaran = field.NewString(table, "rekom_sikap_pelajaran")
	s.RekomTmi = field.NewString(table, "rekom_tmi")
	s.RekomAkhir = field.NewString(table, "rekom_akhir")
	s.SelesaiSkoring = field.NewInt32(table, "selesai_skoring")

	s.fillFieldMap()

	return s
}

func (s *skoringMinatMan) WithContext(ctx context.Context) *skoringMinatManDo {
	return s.skoringMinatManDo.WithContext(ctx)
}

func (s skoringMinatMan) TableName() string { return s.skoringMinatManDo.TableName() }

func (s skoringMinatMan) Alias() string { return s.skoringMinatManDo.Alias() }

func (s skoringMinatMan) Columns(cols ...field.Expr) gen.Columns {
	return s.skoringMinatManDo.Columns(cols...)
}

func (s *skoringMinatMan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skoringMinatMan) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 59)
	s.fieldMap["id"] = s.ID
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["tpa_iu"] = s.TpaIu
	s.fieldMap["tpa_pv"] = s.TpaPv
	s.fieldMap["tpa_pk"] = s.TpaPk
	s.fieldMap["tpa_pa"] = s.TpaPa
	s.fieldMap["tpa_ps"] = s.TpaPs
	s.fieldMap["tpa_pm"] = s.TpaPm
	s.fieldMap["tpa_kt"] = s.TpaKt
	s.fieldMap["tpa_iq"] = s.TpaIq
	s.fieldMap["skor_iq"] = s.SkorIq
	s.fieldMap["minat_sains"] = s.MinatSains
	s.fieldMap["minat_humaniora"] = s.MinatHumaniora
	s.fieldMap["minat_bahasa"] = s.MinatBahasa
	s.fieldMap["minat_agama"] = s.MinatAgama
	s.fieldMap["minat_rentang"] = s.MinatRentang
	s.fieldMap["sikap_agm"] = s.SikapAgm
	s.fieldMap["sikap_pkn"] = s.SikapPkn
	s.fieldMap["sikap_ind"] = s.SikapInd
	s.fieldMap["sikap_eng"] = s.SikapEng
	s.fieldMap["sikap_mat"] = s.SikapMat
	s.fieldMap["sikap_fis"] = s.SikapFis
	s.fieldMap["sikap_bio"] = s.SikapBio
	s.fieldMap["sikap_eko"] = s.SikapEko
	s.fieldMap["sikap_sej"] = s.SikapSej
	s.fieldMap["sikap_geo"] = s.SikapGeo
	s.fieldMap["sikap_ilmu_alam"] = s.SikapIlmuAlam
	s.fieldMap["sikap_ilmu_sosial"] = s.SikapIlmuSosial
	s.fieldMap["sikap_rentang"] = s.SikapRentang
	s.fieldMap["tmi_ilmu_alam"] = s.TmiIlmuAlam
	s.fieldMap["tmi_ilmu_sosial"] = s.TmiIlmuSosial
	s.fieldMap["tmi_rentang"] = s.TmiRentang
	s.fieldMap["tipojung_e"] = s.TipojungE
	s.fieldMap["tipojung_i"] = s.TipojungI
	s.fieldMap["tipojung_s"] = s.TipojungS
	s.fieldMap["tipojung_n"] = s.TipojungN
	s.fieldMap["tipojung_t"] = s.TipojungT
	s.fieldMap["tipojung_f"] = s.TipojungF
	s.fieldMap["tipojung_j"] = s.TipojungJ
	s.fieldMap["tipojung_p"] = s.TipojungP
	s.fieldMap["tipojung_kode"] = s.TipojungKode
	s.fieldMap["pribadi_motivasi"] = s.PribadiMotivasi
	s.fieldMap["pribadi_juang"] = s.PribadiJuang
	s.fieldMap["pribadi_yakin"] = s.PribadiYakin
	s.fieldMap["pribadi_percaya"] = s.PribadiPercaya
	s.fieldMap["pribadi_konsep"] = s.PribadiKonsep
	s.fieldMap["pribadi_kreativitas"] = s.PribadiKreativitas
	s.fieldMap["pribadi_mimpin"] = s.PribadiMimpin
	s.fieldMap["pribadi_entrepreneur"] = s.PribadiEntrepreneur
	s.fieldMap["pribadi_stress"] = s.PribadiStress
	s.fieldMap["pribadi_emosi"] = s.PribadiEmosi
	s.fieldMap["pribadi_sosial"] = s.PribadiSosial
	s.fieldMap["pribadi_empati"] = s.PribadiEmpati
	s.fieldMap["rekom_minat"] = s.RekomMinat
	s.fieldMap["rekom_sikap_pelajaran"] = s.RekomSikapPelajaran
	s.fieldMap["rekom_tmi"] = s.RekomTmi
	s.fieldMap["rekom_akhir"] = s.RekomAkhir
	s.fieldMap["selesai_skoring"] = s.SelesaiSkoring
}

func (s skoringMinatMan) clone(db *gorm.DB) skoringMinatMan {
	s.skoringMinatManDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skoringMinatMan) replaceDB(db *gorm.DB) skoringMinatMan {
	s.skoringMinatManDo.ReplaceDB(db)
	return s
}

type skoringMinatManDo struct{ gen.DO }

func (s skoringMinatManDo) Debug() *skoringMinatManDo {
	return s.withDO(s.DO.Debug())
}

func (s skoringMinatManDo) WithContext(ctx context.Context) *skoringMinatManDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skoringMinatManDo) ReadDB() *skoringMinatManDo {
	return s.Clauses(dbresolver.Read)
}

func (s skoringMinatManDo) WriteDB() *skoringMinatManDo {
	return s.Clauses(dbresolver.Write)
}

func (s skoringMinatManDo) Session(config *gorm.Session) *skoringMinatManDo {
	return s.withDO(s.DO.Session(config))
}

func (s skoringMinatManDo) Clauses(conds ...clause.Expression) *skoringMinatManDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skoringMinatManDo) Returning(value interface{}, columns ...string) *skoringMinatManDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skoringMinatManDo) Not(conds ...gen.Condition) *skoringMinatManDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skoringMinatManDo) Or(conds ...gen.Condition) *skoringMinatManDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skoringMinatManDo) Select(conds ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skoringMinatManDo) Where(conds ...gen.Condition) *skoringMinatManDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skoringMinatManDo) Order(conds ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skoringMinatManDo) Distinct(cols ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skoringMinatManDo) Omit(cols ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skoringMinatManDo) Join(table schema.Tabler, on ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skoringMinatManDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skoringMinatManDo) RightJoin(table schema.Tabler, on ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skoringMinatManDo) Group(cols ...field.Expr) *skoringMinatManDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skoringMinatManDo) Having(conds ...gen.Condition) *skoringMinatManDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skoringMinatManDo) Limit(limit int) *skoringMinatManDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skoringMinatManDo) Offset(offset int) *skoringMinatManDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skoringMinatManDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skoringMinatManDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skoringMinatManDo) Unscoped() *skoringMinatManDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skoringMinatManDo) Create(values ...*entity.SkoringMinatMan) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skoringMinatManDo) CreateInBatches(values []*entity.SkoringMinatMan, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skoringMinatManDo) Save(values ...*entity.SkoringMinatMan) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skoringMinatManDo) First() (*entity.SkoringMinatMan, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkoringMinatMan), nil
	}
}

func (s skoringMinatManDo) Take() (*entity.SkoringMinatMan, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkoringMinatMan), nil
	}
}

func (s skoringMinatManDo) Last() (*entity.SkoringMinatMan, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkoringMinatMan), nil
	}
}

func (s skoringMinatManDo) Find() ([]*entity.SkoringMinatMan, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkoringMinatMan), err
}

func (s skoringMinatManDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkoringMinatMan, err error) {
	buf := make([]*entity.SkoringMinatMan, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skoringMinatManDo) FindInBatches(result *[]*entity.SkoringMinatMan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skoringMinatManDo) Attrs(attrs ...field.AssignExpr) *skoringMinatManDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skoringMinatManDo) Assign(attrs ...field.AssignExpr) *skoringMinatManDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skoringMinatManDo) Joins(fields ...field.RelationField) *skoringMinatManDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skoringMinatManDo) Preload(fields ...field.RelationField) *skoringMinatManDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skoringMinatManDo) FirstOrInit() (*entity.SkoringMinatMan, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkoringMinatMan), nil
	}
}

func (s skoringMinatManDo) FirstOrCreate() (*entity.SkoringMinatMan, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkoringMinatMan), nil
	}
}

func (s skoringMinatManDo) FindByPage(offset int, limit int) (result []*entity.SkoringMinatMan, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s skoringMinatManDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skoringMinatManDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skoringMinatManDo) Delete(models ...*entity.SkoringMinatMan) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skoringMinatManDo) withDO(do gen.Dao) *skoringMinatManDo {
	s.DO = *do.(*gen.DO)
	return s
}
