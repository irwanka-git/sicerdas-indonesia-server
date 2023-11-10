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

func newSkorRekomKuliahB(db *gorm.DB, opts ...gen.DOOption) skorRekomKuliahB {
	_skorRekomKuliahB := skorRekomKuliahB{}

	_skorRekomKuliahB.skorRekomKuliahBDo.UseDB(db, opts...)
	_skorRekomKuliahB.skorRekomKuliahBDo.UseModel(&entity.SkorRekomKuliahB{})

	tableName := _skorRekomKuliahB.skorRekomKuliahBDo.TableName()
	_skorRekomKuliahB.ALL = field.NewAsterisk(tableName)
	_skorRekomKuliahB.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorRekomKuliahB.IDUser = field.NewInt32(tableName, "id_user")
	_skorRekomKuliahB.RekomKuliahAlam = field.NewString(tableName, "rekom_kuliah_alam")
	_skorRekomKuliahB.RekomKuliahSosial = field.NewString(tableName, "rekom_kuliah_sosial")
	_skorRekomKuliahB.RekomKuliahDinas = field.NewString(tableName, "rekom_kuliah_dinas")
	_skorRekomKuliahB.RekomKuliahAgama = field.NewString(tableName, "rekom_kuliah_agama")

	_skorRekomKuliahB.fillFieldMap()

	return _skorRekomKuliahB
}

type skorRekomKuliahB struct {
	skorRekomKuliahBDo skorRekomKuliahBDo

	ALL               field.Asterisk
	IDQuiz            field.Int32
	IDUser            field.Int32
	RekomKuliahAlam   field.String
	RekomKuliahSosial field.String
	RekomKuliahDinas  field.String
	RekomKuliahAgama  field.String

	fieldMap map[string]field.Expr
}

func (s skorRekomKuliahB) Table(newTableName string) *skorRekomKuliahB {
	s.skorRekomKuliahBDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorRekomKuliahB) As(alias string) *skorRekomKuliahB {
	s.skorRekomKuliahBDo.DO = *(s.skorRekomKuliahBDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorRekomKuliahB) updateTableName(table string) *skorRekomKuliahB {
	s.ALL = field.NewAsterisk(table)
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.IDUser = field.NewInt32(table, "id_user")
	s.RekomKuliahAlam = field.NewString(table, "rekom_kuliah_alam")
	s.RekomKuliahSosial = field.NewString(table, "rekom_kuliah_sosial")
	s.RekomKuliahDinas = field.NewString(table, "rekom_kuliah_dinas")
	s.RekomKuliahAgama = field.NewString(table, "rekom_kuliah_agama")

	s.fillFieldMap()

	return s
}

func (s *skorRekomKuliahB) WithContext(ctx context.Context) *skorRekomKuliahBDo {
	return s.skorRekomKuliahBDo.WithContext(ctx)
}

func (s skorRekomKuliahB) TableName() string { return s.skorRekomKuliahBDo.TableName() }

func (s skorRekomKuliahB) Alias() string { return s.skorRekomKuliahBDo.Alias() }

func (s skorRekomKuliahB) Columns(cols ...field.Expr) gen.Columns {
	return s.skorRekomKuliahBDo.Columns(cols...)
}

func (s *skorRekomKuliahB) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorRekomKuliahB) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["rekom_kuliah_alam"] = s.RekomKuliahAlam
	s.fieldMap["rekom_kuliah_sosial"] = s.RekomKuliahSosial
	s.fieldMap["rekom_kuliah_dinas"] = s.RekomKuliahDinas
	s.fieldMap["rekom_kuliah_agama"] = s.RekomKuliahAgama
}

func (s skorRekomKuliahB) clone(db *gorm.DB) skorRekomKuliahB {
	s.skorRekomKuliahBDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorRekomKuliahB) replaceDB(db *gorm.DB) skorRekomKuliahB {
	s.skorRekomKuliahBDo.ReplaceDB(db)
	return s
}

type skorRekomKuliahBDo struct{ gen.DO }

func (s skorRekomKuliahBDo) Debug() *skorRekomKuliahBDo {
	return s.withDO(s.DO.Debug())
}

func (s skorRekomKuliahBDo) WithContext(ctx context.Context) *skorRekomKuliahBDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorRekomKuliahBDo) ReadDB() *skorRekomKuliahBDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorRekomKuliahBDo) WriteDB() *skorRekomKuliahBDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorRekomKuliahBDo) Session(config *gorm.Session) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorRekomKuliahBDo) Clauses(conds ...clause.Expression) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorRekomKuliahBDo) Returning(value interface{}, columns ...string) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorRekomKuliahBDo) Not(conds ...gen.Condition) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorRekomKuliahBDo) Or(conds ...gen.Condition) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorRekomKuliahBDo) Select(conds ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorRekomKuliahBDo) Where(conds ...gen.Condition) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorRekomKuliahBDo) Order(conds ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorRekomKuliahBDo) Distinct(cols ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorRekomKuliahBDo) Omit(cols ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorRekomKuliahBDo) Join(table schema.Tabler, on ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorRekomKuliahBDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorRekomKuliahBDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorRekomKuliahBDo) Group(cols ...field.Expr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorRekomKuliahBDo) Having(conds ...gen.Condition) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorRekomKuliahBDo) Limit(limit int) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorRekomKuliahBDo) Offset(offset int) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorRekomKuliahBDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorRekomKuliahBDo) Unscoped() *skorRekomKuliahBDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorRekomKuliahBDo) Create(values ...*entity.SkorRekomKuliahB) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorRekomKuliahBDo) CreateInBatches(values []*entity.SkorRekomKuliahB, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorRekomKuliahBDo) Save(values ...*entity.SkorRekomKuliahB) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorRekomKuliahBDo) First() (*entity.SkorRekomKuliahB, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomKuliahB), nil
	}
}

func (s skorRekomKuliahBDo) Take() (*entity.SkorRekomKuliahB, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomKuliahB), nil
	}
}

func (s skorRekomKuliahBDo) Last() (*entity.SkorRekomKuliahB, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomKuliahB), nil
	}
}

func (s skorRekomKuliahBDo) Find() ([]*entity.SkorRekomKuliahB, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorRekomKuliahB), err
}

func (s skorRekomKuliahBDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorRekomKuliahB, err error) {
	buf := make([]*entity.SkorRekomKuliahB, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorRekomKuliahBDo) FindInBatches(result *[]*entity.SkorRekomKuliahB, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorRekomKuliahBDo) Attrs(attrs ...field.AssignExpr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorRekomKuliahBDo) Assign(attrs ...field.AssignExpr) *skorRekomKuliahBDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorRekomKuliahBDo) Joins(fields ...field.RelationField) *skorRekomKuliahBDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorRekomKuliahBDo) Preload(fields ...field.RelationField) *skorRekomKuliahBDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorRekomKuliahBDo) FirstOrInit() (*entity.SkorRekomKuliahB, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomKuliahB), nil
	}
}

func (s skorRekomKuliahBDo) FirstOrCreate() (*entity.SkorRekomKuliahB, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomKuliahB), nil
	}
}

func (s skorRekomKuliahBDo) FindByPage(offset int, limit int) (result []*entity.SkorRekomKuliahB, count int64, err error) {
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

func (s skorRekomKuliahBDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorRekomKuliahBDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorRekomKuliahBDo) Delete(models ...*entity.SkorRekomKuliahB) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorRekomKuliahBDo) withDO(do gen.Dao) *skorRekomKuliahBDo {
	s.DO = *do.(*gen.DO)
	return s
}