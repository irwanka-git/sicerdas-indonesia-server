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

func newSkorSikapPelajaran(db *gorm.DB, opts ...gen.DOOption) skorSikapPelajaran {
	_skorSikapPelajaran := skorSikapPelajaran{}

	_skorSikapPelajaran.skorSikapPelajaranDo.UseDB(db, opts...)
	_skorSikapPelajaran.skorSikapPelajaranDo.UseModel(&entity.SkorSikapPelajaran{})

	tableName := _skorSikapPelajaran.skorSikapPelajaranDo.TableName()
	_skorSikapPelajaran.ALL = field.NewAsterisk(tableName)
	_skorSikapPelajaran.IDUser = field.NewInt32(tableName, "id_user")
	_skorSikapPelajaran.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorSikapPelajaran.SikapAgm = field.NewInt32(tableName, "sikap_agm")
	_skorSikapPelajaran.SikapPkn = field.NewInt32(tableName, "sikap_pkn")
	_skorSikapPelajaran.SikapInd = field.NewInt32(tableName, "sikap_ind")
	_skorSikapPelajaran.SikapEng = field.NewInt32(tableName, "sikap_eng")
	_skorSikapPelajaran.SikapMat = field.NewInt32(tableName, "sikap_mat")
	_skorSikapPelajaran.SikapFis = field.NewInt32(tableName, "sikap_fis")
	_skorSikapPelajaran.SikapBio = field.NewInt32(tableName, "sikap_bio")
	_skorSikapPelajaran.SikapEko = field.NewInt32(tableName, "sikap_eko")
	_skorSikapPelajaran.SikapSej = field.NewInt32(tableName, "sikap_sej")
	_skorSikapPelajaran.SikapGeo = field.NewInt32(tableName, "sikap_geo")
	_skorSikapPelajaran.KlasifikasiAgm = field.NewString(tableName, "klasifikasi_agm")
	_skorSikapPelajaran.KlasifikasiPkn = field.NewString(tableName, "klasifikasi_pkn")
	_skorSikapPelajaran.KlasifikasiInd = field.NewString(tableName, "klasifikasi_ind")
	_skorSikapPelajaran.KlasifikasiEng = field.NewString(tableName, "klasifikasi_eng")
	_skorSikapPelajaran.KlasifikasiMat = field.NewString(tableName, "klasifikasi_mat")
	_skorSikapPelajaran.KlasifikasiFis = field.NewString(tableName, "klasifikasi_fis")
	_skorSikapPelajaran.KlasifikasiBio = field.NewString(tableName, "klasifikasi_bio")
	_skorSikapPelajaran.KlasifikasiEko = field.NewString(tableName, "klasifikasi_eko")
	_skorSikapPelajaran.KlasifikasiSej = field.NewString(tableName, "klasifikasi_sej")
	_skorSikapPelajaran.KlasifikasiGeo = field.NewString(tableName, "klasifikasi_geo")
	_skorSikapPelajaran.SikapRentang = field.NewInt32(tableName, "sikap_rentang")
	_skorSikapPelajaran.SikapIlmuAlam = field.NewInt32(tableName, "sikap_ilmu_alam")
	_skorSikapPelajaran.SikapIlmuSosial = field.NewInt32(tableName, "sikap_ilmu_sosial")
	_skorSikapPelajaran.RekomSikapPelajaran = field.NewString(tableName, "rekom_sikap_pelajaran")

	_skorSikapPelajaran.fillFieldMap()

	return _skorSikapPelajaran
}

type skorSikapPelajaran struct {
	skorSikapPelajaranDo skorSikapPelajaranDo

	ALL                 field.Asterisk
	IDUser              field.Int32
	IDQuiz              field.Int32
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
	KlasifikasiAgm      field.String
	KlasifikasiPkn      field.String
	KlasifikasiInd      field.String
	KlasifikasiEng      field.String
	KlasifikasiMat      field.String
	KlasifikasiFis      field.String
	KlasifikasiBio      field.String
	KlasifikasiEko      field.String
	KlasifikasiSej      field.String
	KlasifikasiGeo      field.String
	SikapRentang        field.Int32
	SikapIlmuAlam       field.Int32
	SikapIlmuSosial     field.Int32
	RekomSikapPelajaran field.String

	fieldMap map[string]field.Expr
}

func (s skorSikapPelajaran) Table(newTableName string) *skorSikapPelajaran {
	s.skorSikapPelajaranDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorSikapPelajaran) As(alias string) *skorSikapPelajaran {
	s.skorSikapPelajaranDo.DO = *(s.skorSikapPelajaranDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorSikapPelajaran) updateTableName(table string) *skorSikapPelajaran {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
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
	s.KlasifikasiAgm = field.NewString(table, "klasifikasi_agm")
	s.KlasifikasiPkn = field.NewString(table, "klasifikasi_pkn")
	s.KlasifikasiInd = field.NewString(table, "klasifikasi_ind")
	s.KlasifikasiEng = field.NewString(table, "klasifikasi_eng")
	s.KlasifikasiMat = field.NewString(table, "klasifikasi_mat")
	s.KlasifikasiFis = field.NewString(table, "klasifikasi_fis")
	s.KlasifikasiBio = field.NewString(table, "klasifikasi_bio")
	s.KlasifikasiEko = field.NewString(table, "klasifikasi_eko")
	s.KlasifikasiSej = field.NewString(table, "klasifikasi_sej")
	s.KlasifikasiGeo = field.NewString(table, "klasifikasi_geo")
	s.SikapRentang = field.NewInt32(table, "sikap_rentang")
	s.SikapIlmuAlam = field.NewInt32(table, "sikap_ilmu_alam")
	s.SikapIlmuSosial = field.NewInt32(table, "sikap_ilmu_sosial")
	s.RekomSikapPelajaran = field.NewString(table, "rekom_sikap_pelajaran")

	s.fillFieldMap()

	return s
}

func (s *skorSikapPelajaran) WithContext(ctx context.Context) *skorSikapPelajaranDo {
	return s.skorSikapPelajaranDo.WithContext(ctx)
}

func (s skorSikapPelajaran) TableName() string { return s.skorSikapPelajaranDo.TableName() }

func (s skorSikapPelajaran) Alias() string { return s.skorSikapPelajaranDo.Alias() }

func (s skorSikapPelajaran) Columns(cols ...field.Expr) gen.Columns {
	return s.skorSikapPelajaranDo.Columns(cols...)
}

func (s *skorSikapPelajaran) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorSikapPelajaran) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 26)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
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
	s.fieldMap["klasifikasi_agm"] = s.KlasifikasiAgm
	s.fieldMap["klasifikasi_pkn"] = s.KlasifikasiPkn
	s.fieldMap["klasifikasi_ind"] = s.KlasifikasiInd
	s.fieldMap["klasifikasi_eng"] = s.KlasifikasiEng
	s.fieldMap["klasifikasi_mat"] = s.KlasifikasiMat
	s.fieldMap["klasifikasi_fis"] = s.KlasifikasiFis
	s.fieldMap["klasifikasi_bio"] = s.KlasifikasiBio
	s.fieldMap["klasifikasi_eko"] = s.KlasifikasiEko
	s.fieldMap["klasifikasi_sej"] = s.KlasifikasiSej
	s.fieldMap["klasifikasi_geo"] = s.KlasifikasiGeo
	s.fieldMap["sikap_rentang"] = s.SikapRentang
	s.fieldMap["sikap_ilmu_alam"] = s.SikapIlmuAlam
	s.fieldMap["sikap_ilmu_sosial"] = s.SikapIlmuSosial
	s.fieldMap["rekom_sikap_pelajaran"] = s.RekomSikapPelajaran
}

func (s skorSikapPelajaran) clone(db *gorm.DB) skorSikapPelajaran {
	s.skorSikapPelajaranDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorSikapPelajaran) replaceDB(db *gorm.DB) skorSikapPelajaran {
	s.skorSikapPelajaranDo.ReplaceDB(db)
	return s
}

type skorSikapPelajaranDo struct{ gen.DO }

func (s skorSikapPelajaranDo) Debug() *skorSikapPelajaranDo {
	return s.withDO(s.DO.Debug())
}

func (s skorSikapPelajaranDo) WithContext(ctx context.Context) *skorSikapPelajaranDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorSikapPelajaranDo) ReadDB() *skorSikapPelajaranDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorSikapPelajaranDo) WriteDB() *skorSikapPelajaranDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorSikapPelajaranDo) Session(config *gorm.Session) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorSikapPelajaranDo) Clauses(conds ...clause.Expression) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorSikapPelajaranDo) Returning(value interface{}, columns ...string) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorSikapPelajaranDo) Not(conds ...gen.Condition) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorSikapPelajaranDo) Or(conds ...gen.Condition) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorSikapPelajaranDo) Select(conds ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorSikapPelajaranDo) Where(conds ...gen.Condition) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorSikapPelajaranDo) Order(conds ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorSikapPelajaranDo) Distinct(cols ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorSikapPelajaranDo) Omit(cols ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorSikapPelajaranDo) Join(table schema.Tabler, on ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorSikapPelajaranDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorSikapPelajaranDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorSikapPelajaranDo) Group(cols ...field.Expr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorSikapPelajaranDo) Having(conds ...gen.Condition) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorSikapPelajaranDo) Limit(limit int) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorSikapPelajaranDo) Offset(offset int) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorSikapPelajaranDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorSikapPelajaranDo) Unscoped() *skorSikapPelajaranDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorSikapPelajaranDo) Create(values ...*entity.SkorSikapPelajaran) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorSikapPelajaranDo) CreateInBatches(values []*entity.SkorSikapPelajaran, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorSikapPelajaranDo) Save(values ...*entity.SkorSikapPelajaran) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorSikapPelajaranDo) First() (*entity.SkorSikapPelajaran, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSikapPelajaran), nil
	}
}

func (s skorSikapPelajaranDo) Take() (*entity.SkorSikapPelajaran, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSikapPelajaran), nil
	}
}

func (s skorSikapPelajaranDo) Last() (*entity.SkorSikapPelajaran, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSikapPelajaran), nil
	}
}

func (s skorSikapPelajaranDo) Find() ([]*entity.SkorSikapPelajaran, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorSikapPelajaran), err
}

func (s skorSikapPelajaranDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorSikapPelajaran, err error) {
	buf := make([]*entity.SkorSikapPelajaran, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorSikapPelajaranDo) FindInBatches(result *[]*entity.SkorSikapPelajaran, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorSikapPelajaranDo) Attrs(attrs ...field.AssignExpr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorSikapPelajaranDo) Assign(attrs ...field.AssignExpr) *skorSikapPelajaranDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorSikapPelajaranDo) Joins(fields ...field.RelationField) *skorSikapPelajaranDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorSikapPelajaranDo) Preload(fields ...field.RelationField) *skorSikapPelajaranDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorSikapPelajaranDo) FirstOrInit() (*entity.SkorSikapPelajaran, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSikapPelajaran), nil
	}
}

func (s skorSikapPelajaranDo) FirstOrCreate() (*entity.SkorSikapPelajaran, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSikapPelajaran), nil
	}
}

func (s skorSikapPelajaranDo) FindByPage(offset int, limit int) (result []*entity.SkorSikapPelajaran, count int64, err error) {
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

func (s skorSikapPelajaranDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorSikapPelajaranDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorSikapPelajaranDo) Delete(models ...*entity.SkorSikapPelajaran) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorSikapPelajaranDo) withDO(do gen.Dao) *skorSikapPelajaranDo {
	s.DO = *do.(*gen.DO)
	return s
}