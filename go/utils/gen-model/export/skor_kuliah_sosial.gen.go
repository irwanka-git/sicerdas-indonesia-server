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

func newSkorKuliahSosial(db *gorm.DB, opts ...gen.DOOption) skorKuliahSosial {
	_skorKuliahSosial := skorKuliahSosial{}

	_skorKuliahSosial.skorKuliahSosialDo.UseDB(db, opts...)
	_skorKuliahSosial.skorKuliahSosialDo.UseModel(&entity.SkorKuliahSosial{})

	tableName := _skorKuliahSosial.skorKuliahSosialDo.TableName()
	_skorKuliahSosial.ALL = field.NewAsterisk(tableName)
	_skorKuliahSosial.IDUser = field.NewInt32(tableName, "id_user")
	_skorKuliahSosial.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorKuliahSosial.MinatIps1 = field.NewInt32(tableName, "minat_ips1")
	_skorKuliahSosial.MinatIps2 = field.NewInt32(tableName, "minat_ips2")
	_skorKuliahSosial.MinatIps3 = field.NewInt32(tableName, "minat_ips3")
	_skorKuliahSosial.MinatIps4 = field.NewInt32(tableName, "minat_ips4")
	_skorKuliahSosial.MinatIps5 = field.NewInt32(tableName, "minat_ips5")

	_skorKuliahSosial.fillFieldMap()

	return _skorKuliahSosial
}

type skorKuliahSosial struct {
	skorKuliahSosialDo skorKuliahSosialDo

	ALL       field.Asterisk
	IDUser    field.Int32
	IDQuiz    field.Int32
	MinatIps1 field.Int32
	MinatIps2 field.Int32
	MinatIps3 field.Int32
	MinatIps4 field.Int32
	MinatIps5 field.Int32

	fieldMap map[string]field.Expr
}

func (s skorKuliahSosial) Table(newTableName string) *skorKuliahSosial {
	s.skorKuliahSosialDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorKuliahSosial) As(alias string) *skorKuliahSosial {
	s.skorKuliahSosialDo.DO = *(s.skorKuliahSosialDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorKuliahSosial) updateTableName(table string) *skorKuliahSosial {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.MinatIps1 = field.NewInt32(table, "minat_ips1")
	s.MinatIps2 = field.NewInt32(table, "minat_ips2")
	s.MinatIps3 = field.NewInt32(table, "minat_ips3")
	s.MinatIps4 = field.NewInt32(table, "minat_ips4")
	s.MinatIps5 = field.NewInt32(table, "minat_ips5")

	s.fillFieldMap()

	return s
}

func (s *skorKuliahSosial) WithContext(ctx context.Context) *skorKuliahSosialDo {
	return s.skorKuliahSosialDo.WithContext(ctx)
}

func (s skorKuliahSosial) TableName() string { return s.skorKuliahSosialDo.TableName() }

func (s skorKuliahSosial) Alias() string { return s.skorKuliahSosialDo.Alias() }

func (s skorKuliahSosial) Columns(cols ...field.Expr) gen.Columns {
	return s.skorKuliahSosialDo.Columns(cols...)
}

func (s *skorKuliahSosial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorKuliahSosial) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["minat_ips1"] = s.MinatIps1
	s.fieldMap["minat_ips2"] = s.MinatIps2
	s.fieldMap["minat_ips3"] = s.MinatIps3
	s.fieldMap["minat_ips4"] = s.MinatIps4
	s.fieldMap["minat_ips5"] = s.MinatIps5
}

func (s skorKuliahSosial) clone(db *gorm.DB) skorKuliahSosial {
	s.skorKuliahSosialDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorKuliahSosial) replaceDB(db *gorm.DB) skorKuliahSosial {
	s.skorKuliahSosialDo.ReplaceDB(db)
	return s
}

type skorKuliahSosialDo struct{ gen.DO }

func (s skorKuliahSosialDo) Debug() *skorKuliahSosialDo {
	return s.withDO(s.DO.Debug())
}

func (s skorKuliahSosialDo) WithContext(ctx context.Context) *skorKuliahSosialDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorKuliahSosialDo) ReadDB() *skorKuliahSosialDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorKuliahSosialDo) WriteDB() *skorKuliahSosialDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorKuliahSosialDo) Session(config *gorm.Session) *skorKuliahSosialDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorKuliahSosialDo) Clauses(conds ...clause.Expression) *skorKuliahSosialDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorKuliahSosialDo) Returning(value interface{}, columns ...string) *skorKuliahSosialDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorKuliahSosialDo) Not(conds ...gen.Condition) *skorKuliahSosialDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorKuliahSosialDo) Or(conds ...gen.Condition) *skorKuliahSosialDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorKuliahSosialDo) Select(conds ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorKuliahSosialDo) Where(conds ...gen.Condition) *skorKuliahSosialDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorKuliahSosialDo) Order(conds ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorKuliahSosialDo) Distinct(cols ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorKuliahSosialDo) Omit(cols ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorKuliahSosialDo) Join(table schema.Tabler, on ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorKuliahSosialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorKuliahSosialDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorKuliahSosialDo) Group(cols ...field.Expr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorKuliahSosialDo) Having(conds ...gen.Condition) *skorKuliahSosialDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorKuliahSosialDo) Limit(limit int) *skorKuliahSosialDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorKuliahSosialDo) Offset(offset int) *skorKuliahSosialDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorKuliahSosialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorKuliahSosialDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorKuliahSosialDo) Unscoped() *skorKuliahSosialDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorKuliahSosialDo) Create(values ...*entity.SkorKuliahSosial) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorKuliahSosialDo) CreateInBatches(values []*entity.SkorKuliahSosial, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorKuliahSosialDo) Save(values ...*entity.SkorKuliahSosial) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorKuliahSosialDo) First() (*entity.SkorKuliahSosial, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKuliahSosial), nil
	}
}

func (s skorKuliahSosialDo) Take() (*entity.SkorKuliahSosial, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKuliahSosial), nil
	}
}

func (s skorKuliahSosialDo) Last() (*entity.SkorKuliahSosial, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKuliahSosial), nil
	}
}

func (s skorKuliahSosialDo) Find() ([]*entity.SkorKuliahSosial, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorKuliahSosial), err
}

func (s skorKuliahSosialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorKuliahSosial, err error) {
	buf := make([]*entity.SkorKuliahSosial, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorKuliahSosialDo) FindInBatches(result *[]*entity.SkorKuliahSosial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorKuliahSosialDo) Attrs(attrs ...field.AssignExpr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorKuliahSosialDo) Assign(attrs ...field.AssignExpr) *skorKuliahSosialDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorKuliahSosialDo) Joins(fields ...field.RelationField) *skorKuliahSosialDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorKuliahSosialDo) Preload(fields ...field.RelationField) *skorKuliahSosialDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorKuliahSosialDo) FirstOrInit() (*entity.SkorKuliahSosial, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKuliahSosial), nil
	}
}

func (s skorKuliahSosialDo) FirstOrCreate() (*entity.SkorKuliahSosial, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorKuliahSosial), nil
	}
}

func (s skorKuliahSosialDo) FindByPage(offset int, limit int) (result []*entity.SkorKuliahSosial, count int64, err error) {
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

func (s skorKuliahSosialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorKuliahSosialDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorKuliahSosialDo) Delete(models ...*entity.SkorKuliahSosial) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorKuliahSosialDo) withDO(do gen.Dao) *skorKuliahSosialDo {
	s.DO = *do.(*gen.DO)
	return s
}