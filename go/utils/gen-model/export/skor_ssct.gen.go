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

func newSkorSsct(db *gorm.DB, opts ...gen.DOOption) skorSsct {
	_skorSsct := skorSsct{}

	_skorSsct.skorSsctDo.UseDB(db, opts...)
	_skorSsct.skorSsctDo.UseModel(&entity.SkorSsct{})

	tableName := _skorSsct.skorSsctDo.TableName()
	_skorSsct.ALL = field.NewAsterisk(tableName)
	_skorSsct.IDUser = field.NewInt32(tableName, "id_user")
	_skorSsct.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorSsct.Urutan = field.NewInt32(tableName, "urutan")
	_skorSsct.Skor = field.NewInt32(tableName, "skor")
	_skorSsct.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_skorSsct.fillFieldMap()

	return _skorSsct
}

type skorSsct struct {
	skorSsctDo skorSsctDo

	ALL         field.Asterisk
	IDUser      field.Int32
	IDQuiz      field.Int32
	Urutan      field.Int32
	Skor        field.Int32
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (s skorSsct) Table(newTableName string) *skorSsct {
	s.skorSsctDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorSsct) As(alias string) *skorSsct {
	s.skorSsctDo.DO = *(s.skorSsctDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorSsct) updateTableName(table string) *skorSsct {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.Urutan = field.NewInt32(table, "urutan")
	s.Skor = field.NewInt32(table, "skor")
	s.Klasifikasi = field.NewString(table, "klasifikasi")

	s.fillFieldMap()

	return s
}

func (s *skorSsct) WithContext(ctx context.Context) *skorSsctDo { return s.skorSsctDo.WithContext(ctx) }

func (s skorSsct) TableName() string { return s.skorSsctDo.TableName() }

func (s skorSsct) Alias() string { return s.skorSsctDo.Alias() }

func (s skorSsct) Columns(cols ...field.Expr) gen.Columns { return s.skorSsctDo.Columns(cols...) }

func (s *skorSsct) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorSsct) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["urutan"] = s.Urutan
	s.fieldMap["skor"] = s.Skor
	s.fieldMap["klasifikasi"] = s.Klasifikasi
}

func (s skorSsct) clone(db *gorm.DB) skorSsct {
	s.skorSsctDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorSsct) replaceDB(db *gorm.DB) skorSsct {
	s.skorSsctDo.ReplaceDB(db)
	return s
}

type skorSsctDo struct{ gen.DO }

func (s skorSsctDo) Debug() *skorSsctDo {
	return s.withDO(s.DO.Debug())
}

func (s skorSsctDo) WithContext(ctx context.Context) *skorSsctDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorSsctDo) ReadDB() *skorSsctDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorSsctDo) WriteDB() *skorSsctDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorSsctDo) Session(config *gorm.Session) *skorSsctDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorSsctDo) Clauses(conds ...clause.Expression) *skorSsctDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorSsctDo) Returning(value interface{}, columns ...string) *skorSsctDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorSsctDo) Not(conds ...gen.Condition) *skorSsctDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorSsctDo) Or(conds ...gen.Condition) *skorSsctDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorSsctDo) Select(conds ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorSsctDo) Where(conds ...gen.Condition) *skorSsctDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorSsctDo) Order(conds ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorSsctDo) Distinct(cols ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorSsctDo) Omit(cols ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorSsctDo) Join(table schema.Tabler, on ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorSsctDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorSsctDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorSsctDo) Group(cols ...field.Expr) *skorSsctDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorSsctDo) Having(conds ...gen.Condition) *skorSsctDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorSsctDo) Limit(limit int) *skorSsctDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorSsctDo) Offset(offset int) *skorSsctDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorSsctDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorSsctDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorSsctDo) Unscoped() *skorSsctDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorSsctDo) Create(values ...*entity.SkorSsct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorSsctDo) CreateInBatches(values []*entity.SkorSsct, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorSsctDo) Save(values ...*entity.SkorSsct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorSsctDo) First() (*entity.SkorSsct, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSsct), nil
	}
}

func (s skorSsctDo) Take() (*entity.SkorSsct, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSsct), nil
	}
}

func (s skorSsctDo) Last() (*entity.SkorSsct, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSsct), nil
	}
}

func (s skorSsctDo) Find() ([]*entity.SkorSsct, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorSsct), err
}

func (s skorSsctDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorSsct, err error) {
	buf := make([]*entity.SkorSsct, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorSsctDo) FindInBatches(result *[]*entity.SkorSsct, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorSsctDo) Attrs(attrs ...field.AssignExpr) *skorSsctDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorSsctDo) Assign(attrs ...field.AssignExpr) *skorSsctDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorSsctDo) Joins(fields ...field.RelationField) *skorSsctDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorSsctDo) Preload(fields ...field.RelationField) *skorSsctDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorSsctDo) FirstOrInit() (*entity.SkorSsct, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSsct), nil
	}
}

func (s skorSsctDo) FirstOrCreate() (*entity.SkorSsct, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorSsct), nil
	}
}

func (s skorSsctDo) FindByPage(offset int, limit int) (result []*entity.SkorSsct, count int64, err error) {
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

func (s skorSsctDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorSsctDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorSsctDo) Delete(models ...*entity.SkorSsct) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorSsctDo) withDO(do gen.Dao) *skorSsctDo {
	s.DO = *do.(*gen.DO)
	return s
}