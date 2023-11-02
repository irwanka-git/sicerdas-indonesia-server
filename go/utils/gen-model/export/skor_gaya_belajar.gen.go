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

func newSkorGayaBelajar(db *gorm.DB, opts ...gen.DOOption) skorGayaBelajar {
	_skorGayaBelajar := skorGayaBelajar{}

	_skorGayaBelajar.skorGayaBelajarDo.UseDB(db, opts...)
	_skorGayaBelajar.skorGayaBelajarDo.UseModel(&entity.SkorGayaBelajar{})

	tableName := _skorGayaBelajar.skorGayaBelajarDo.TableName()
	_skorGayaBelajar.ALL = field.NewAsterisk(tableName)
	_skorGayaBelajar.IDUser = field.NewInt32(tableName, "id_user")
	_skorGayaBelajar.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorGayaBelajar.GayaAuditoris = field.NewInt16(tableName, "gaya_auditoris")
	_skorGayaBelajar.GayaVisual = field.NewInt16(tableName, "gaya_visual")
	_skorGayaBelajar.GayaKinestetik = field.NewInt16(tableName, "gaya_kinestetik")
	_skorGayaBelajar.KlasifikasiAuditoris = field.NewString(tableName, "klasifikasi_auditoris")
	_skorGayaBelajar.KlasifikasiVisual = field.NewString(tableName, "klasifikasi_visual")
	_skorGayaBelajar.KlasifikasiKinestetik = field.NewString(tableName, "klasifikasi_kinestetik")

	_skorGayaBelajar.fillFieldMap()

	return _skorGayaBelajar
}

type skorGayaBelajar struct {
	skorGayaBelajarDo skorGayaBelajarDo

	ALL                   field.Asterisk
	IDUser                field.Int32
	IDQuiz                field.Int32
	GayaAuditoris         field.Int16
	GayaVisual            field.Int16
	GayaKinestetik        field.Int16
	KlasifikasiAuditoris  field.String
	KlasifikasiVisual     field.String
	KlasifikasiKinestetik field.String

	fieldMap map[string]field.Expr
}

func (s skorGayaBelajar) Table(newTableName string) *skorGayaBelajar {
	s.skorGayaBelajarDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorGayaBelajar) As(alias string) *skorGayaBelajar {
	s.skorGayaBelajarDo.DO = *(s.skorGayaBelajarDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorGayaBelajar) updateTableName(table string) *skorGayaBelajar {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.GayaAuditoris = field.NewInt16(table, "gaya_auditoris")
	s.GayaVisual = field.NewInt16(table, "gaya_visual")
	s.GayaKinestetik = field.NewInt16(table, "gaya_kinestetik")
	s.KlasifikasiAuditoris = field.NewString(table, "klasifikasi_auditoris")
	s.KlasifikasiVisual = field.NewString(table, "klasifikasi_visual")
	s.KlasifikasiKinestetik = field.NewString(table, "klasifikasi_kinestetik")

	s.fillFieldMap()

	return s
}

func (s *skorGayaBelajar) WithContext(ctx context.Context) *skorGayaBelajarDo {
	return s.skorGayaBelajarDo.WithContext(ctx)
}

func (s skorGayaBelajar) TableName() string { return s.skorGayaBelajarDo.TableName() }

func (s skorGayaBelajar) Alias() string { return s.skorGayaBelajarDo.Alias() }

func (s skorGayaBelajar) Columns(cols ...field.Expr) gen.Columns {
	return s.skorGayaBelajarDo.Columns(cols...)
}

func (s *skorGayaBelajar) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorGayaBelajar) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["gaya_auditoris"] = s.GayaAuditoris
	s.fieldMap["gaya_visual"] = s.GayaVisual
	s.fieldMap["gaya_kinestetik"] = s.GayaKinestetik
	s.fieldMap["klasifikasi_auditoris"] = s.KlasifikasiAuditoris
	s.fieldMap["klasifikasi_visual"] = s.KlasifikasiVisual
	s.fieldMap["klasifikasi_kinestetik"] = s.KlasifikasiKinestetik
}

func (s skorGayaBelajar) clone(db *gorm.DB) skorGayaBelajar {
	s.skorGayaBelajarDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorGayaBelajar) replaceDB(db *gorm.DB) skorGayaBelajar {
	s.skorGayaBelajarDo.ReplaceDB(db)
	return s
}

type skorGayaBelajarDo struct{ gen.DO }

func (s skorGayaBelajarDo) Debug() *skorGayaBelajarDo {
	return s.withDO(s.DO.Debug())
}

func (s skorGayaBelajarDo) WithContext(ctx context.Context) *skorGayaBelajarDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorGayaBelajarDo) ReadDB() *skorGayaBelajarDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorGayaBelajarDo) WriteDB() *skorGayaBelajarDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorGayaBelajarDo) Session(config *gorm.Session) *skorGayaBelajarDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorGayaBelajarDo) Clauses(conds ...clause.Expression) *skorGayaBelajarDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorGayaBelajarDo) Returning(value interface{}, columns ...string) *skorGayaBelajarDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorGayaBelajarDo) Not(conds ...gen.Condition) *skorGayaBelajarDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorGayaBelajarDo) Or(conds ...gen.Condition) *skorGayaBelajarDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorGayaBelajarDo) Select(conds ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorGayaBelajarDo) Where(conds ...gen.Condition) *skorGayaBelajarDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorGayaBelajarDo) Order(conds ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorGayaBelajarDo) Distinct(cols ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorGayaBelajarDo) Omit(cols ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorGayaBelajarDo) Join(table schema.Tabler, on ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorGayaBelajarDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorGayaBelajarDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorGayaBelajarDo) Group(cols ...field.Expr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorGayaBelajarDo) Having(conds ...gen.Condition) *skorGayaBelajarDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorGayaBelajarDo) Limit(limit int) *skorGayaBelajarDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorGayaBelajarDo) Offset(offset int) *skorGayaBelajarDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorGayaBelajarDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorGayaBelajarDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorGayaBelajarDo) Unscoped() *skorGayaBelajarDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorGayaBelajarDo) Create(values ...*entity.SkorGayaBelajar) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorGayaBelajarDo) CreateInBatches(values []*entity.SkorGayaBelajar, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorGayaBelajarDo) Save(values ...*entity.SkorGayaBelajar) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorGayaBelajarDo) First() (*entity.SkorGayaBelajar, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorGayaBelajar), nil
	}
}

func (s skorGayaBelajarDo) Take() (*entity.SkorGayaBelajar, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorGayaBelajar), nil
	}
}

func (s skorGayaBelajarDo) Last() (*entity.SkorGayaBelajar, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorGayaBelajar), nil
	}
}

func (s skorGayaBelajarDo) Find() ([]*entity.SkorGayaBelajar, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorGayaBelajar), err
}

func (s skorGayaBelajarDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorGayaBelajar, err error) {
	buf := make([]*entity.SkorGayaBelajar, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorGayaBelajarDo) FindInBatches(result *[]*entity.SkorGayaBelajar, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorGayaBelajarDo) Attrs(attrs ...field.AssignExpr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorGayaBelajarDo) Assign(attrs ...field.AssignExpr) *skorGayaBelajarDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorGayaBelajarDo) Joins(fields ...field.RelationField) *skorGayaBelajarDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorGayaBelajarDo) Preload(fields ...field.RelationField) *skorGayaBelajarDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorGayaBelajarDo) FirstOrInit() (*entity.SkorGayaBelajar, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorGayaBelajar), nil
	}
}

func (s skorGayaBelajarDo) FirstOrCreate() (*entity.SkorGayaBelajar, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorGayaBelajar), nil
	}
}

func (s skorGayaBelajarDo) FindByPage(offset int, limit int) (result []*entity.SkorGayaBelajar, count int64, err error) {
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

func (s skorGayaBelajarDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorGayaBelajarDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorGayaBelajarDo) Delete(models ...*entity.SkorGayaBelajar) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorGayaBelajarDo) withDO(do gen.Dao) *skorGayaBelajarDo {
	s.DO = *do.(*gen.DO)
	return s
}
