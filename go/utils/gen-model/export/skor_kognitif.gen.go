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

func newSkorKognitif(db *gorm.DB, opts ...gen.DOOption) skorKognitif {
	_skorKognitif := skorKognitif{}

	_skorKognitif.skorKognitifDo.UseDB(db, opts...)
	_skorKognitif.skorKognitifDo.UseModel(&entity.SkorKognitif{})

	tableName := _skorKognitif.skorKognitifDo.TableName()
	_skorKognitif.ALL = field.NewAsterisk(tableName)
	_skorKognitif.ID = field.NewInt64(tableName, "id")
	_skorKognitif.IDUser = field.NewInt32(tableName, "id_user")
	_skorKognitif.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorKognitif.TpaIu = field.NewInt32(tableName, "tpa_iu")
	_skorKognitif.TpaPv = field.NewInt32(tableName, "tpa_pv")
	_skorKognitif.TpaPk = field.NewInt32(tableName, "tpa_pk")
	_skorKognitif.TpaPa = field.NewInt32(tableName, "tpa_pa")
	_skorKognitif.TpaPs = field.NewInt32(tableName, "tpa_ps")
	_skorKognitif.TpaPm = field.NewInt32(tableName, "tpa_pm")
	_skorKognitif.TpaKt = field.NewInt32(tableName, "tpa_kt")
	_skorKognitif.TpaIq = field.NewInt32(tableName, "tpa_iq")
	_skorKognitif.SkorIq = field.NewFloat32(tableName, "skor_iq")
	_skorKognitif.KlasifikasiIu = field.NewString(tableName, "klasifikasi_iu")
	_skorKognitif.KlasifikasiPv = field.NewString(tableName, "klasifikasi_pv")
	_skorKognitif.KlasifikasiPk = field.NewString(tableName, "klasifikasi_pk")
	_skorKognitif.KlasifikasiPa = field.NewString(tableName, "klasifikasi_pa")
	_skorKognitif.KlasifikasiPs = field.NewString(tableName, "klasifikasi_ps")
	_skorKognitif.KlasifikasiPm = field.NewString(tableName, "klasifikasi_pm")
	_skorKognitif.KlasifikasiKt = field.NewString(tableName, "klasifikasi_kt")
	_skorKognitif.KlasifikasiIq = field.NewString(tableName, "klasifikasi_iq")

	_skorKognitif.fillFieldMap()

	return _skorKognitif
}

type skorKognitif struct {
	skorKognitifDo skorKognitifDo

	ALL           field.Asterisk
	ID            field.Int64
	IDUser        field.Int32
	IDQuiz        field.Int32
	TpaIu         field.Int32
	TpaPv         field.Int32
	TpaPk         field.Int32
	TpaPa         field.Int32
	TpaPs         field.Int32
	TpaPm         field.Int32
	TpaKt         field.Int32
	TpaIq         field.Int32
	SkorIq        field.Float32
	KlasifikasiIu field.String
	KlasifikasiPv field.String
	KlasifikasiPk field.String
	KlasifikasiPa field.String
	KlasifikasiPs field.String
	KlasifikasiPm field.String
	KlasifikasiKt field.String
	KlasifikasiIq field.String

	fieldMap map[string]field.Expr
}

func (s skorKognitif) Table(newTableName string) *skorKognitif {
	s.skorKognitifDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorKognitif) As(alias string) *skorKognitif {
	s.skorKognitifDo.DO = *(s.skorKognitifDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorKognitif) updateTableName(table string) *skorKognitif {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
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
	s.KlasifikasiIu = field.NewString(table, "klasifikasi_iu")
	s.KlasifikasiPv = field.NewString(table, "klasifikasi_pv")
	s.KlasifikasiPk = field.NewString(table, "klasifikasi_pk")
	s.KlasifikasiPa = field.NewString(table, "klasifikasi_pa")
	s.KlasifikasiPs = field.NewString(table, "klasifikasi_ps")
	s.KlasifikasiPm = field.NewString(table, "klasifikasi_pm")
	s.KlasifikasiKt = field.NewString(table, "klasifikasi_kt")
	s.KlasifikasiIq = field.NewString(table, "klasifikasi_iq")

	s.fillFieldMap()

	return s
}

func (s *skorKognitif) WithContext(ctx context.Context) *skorKognitifDo {
	return s.skorKognitifDo.WithContext(ctx)
}

func (s skorKognitif) TableName() string { return s.skorKognitifDo.TableName() }

func (s skorKognitif) Alias() string { return s.skorKognitifDo.Alias() }

func (s skorKognitif) Columns(cols ...field.Expr) gen.Columns {
	return s.skorKognitifDo.Columns(cols...)
}

func (s *skorKognitif) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorKognitif) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
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
	s.fieldMap["klasifikasi_iu"] = s.KlasifikasiIu
	s.fieldMap["klasifikasi_pv"] = s.KlasifikasiPv
	s.fieldMap["klasifikasi_pk"] = s.KlasifikasiPk
	s.fieldMap["klasifikasi_pa"] = s.KlasifikasiPa
	s.fieldMap["klasifikasi_ps"] = s.KlasifikasiPs
	s.fieldMap["klasifikasi_pm"] = s.KlasifikasiPm
	s.fieldMap["klasifikasi_kt"] = s.KlasifikasiKt
	s.fieldMap["klasifikasi_iq"] = s.KlasifikasiIq
}

func (s skorKognitif) clone(db *gorm.DB) skorKognitif {
	s.skorKognitifDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorKognitif) replaceDB(db *gorm.DB) skorKognitif {
	s.skorKognitifDo.ReplaceDB(db)
	return s
}

type skorKognitifDo struct{ gen.DO }

func (s skorKognitifDo) Debug() *skorKognitifDo {
	return s.withDO(s.DO.Debug())
}

func (s skorKognitifDo) WithContext(ctx context.Context) *skorKognitifDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorKognitifDo) ReadDB() *skorKognitifDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorKognitifDo) WriteDB() *skorKognitifDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorKognitifDo) Session(config *gorm.Session) *skorKognitifDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorKognitifDo) Clauses(conds ...clause.Expression) *skorKognitifDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorKognitifDo) Returning(value interface{}, columns ...string) *skorKognitifDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorKognitifDo) Not(conds ...gen.Condition) *skorKognitifDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorKognitifDo) Or(conds ...gen.Condition) *skorKognitifDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorKognitifDo) Select(conds ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorKognitifDo) Where(conds ...gen.Condition) *skorKognitifDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorKognitifDo) Order(conds ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorKognitifDo) Distinct(cols ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorKognitifDo) Omit(cols ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorKognitifDo) Join(table schema.Tabler, on ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorKognitifDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorKognitifDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorKognitifDo) Group(cols ...field.Expr) *skorKognitifDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorKognitifDo) Having(conds ...gen.Condition) *skorKognitifDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorKognitifDo) Limit(limit int) *skorKognitifDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorKognitifDo) Offset(offset int) *skorKognitifDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorKognitifDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorKognitifDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorKognitifDo) Unscoped() *skorKognitifDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorKognitifDo) Create(values ...*entity.SkorKognitif) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorKognitifDo) CreateInBatches(values []*entity.SkorKognitif, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorKognitifDo) Save(values ...*entity.SkorKognitif) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorKognitifDo) First() (*entity.SkorKognitif, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKognitif), nil
	}
}

func (s skorKognitifDo) Take() (*entity.SkorKognitif, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKognitif), nil
	}
}

func (s skorKognitifDo) Last() (*entity.SkorKognitif, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKognitif), nil
	}
}

func (s skorKognitifDo) Find() ([]*entity.SkorKognitif, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorKognitif), err
}

func (s skorKognitifDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorKognitif, err error) {
	buf := make([]*entity.SkorKognitif, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorKognitifDo) FindInBatches(result *[]*entity.SkorKognitif, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorKognitifDo) Attrs(attrs ...field.AssignExpr) *skorKognitifDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorKognitifDo) Assign(attrs ...field.AssignExpr) *skorKognitifDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorKognitifDo) Joins(fields ...field.RelationField) *skorKognitifDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorKognitifDo) Preload(fields ...field.RelationField) *skorKognitifDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorKognitifDo) FirstOrInit() (*entity.SkorKognitif, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKognitif), nil
	}
}

func (s skorKognitifDo) FirstOrCreate() (*entity.SkorKognitif, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKognitif), nil
	}
}

func (s skorKognitifDo) FindByPage(offset int, limit int) (result []*entity.SkorKognitif, count int64, err error) {
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

func (s skorKognitifDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorKognitifDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorKognitifDo) Delete(models ...*entity.SkorKognitif) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorKognitifDo) withDO(do gen.Dao) *skorKognitifDo {
	s.DO = *do.(*gen.DO)
	return s
}