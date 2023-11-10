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

func newSkorRekomPeminatanSma(db *gorm.DB, opts ...gen.DOOption) skorRekomPeminatanSma {
	_skorRekomPeminatanSma := skorRekomPeminatanSma{}

	_skorRekomPeminatanSma.skorRekomPeminatanSmaDo.UseDB(db, opts...)
	_skorRekomPeminatanSma.skorRekomPeminatanSmaDo.UseModel(&entity.SkorRekomPeminatanSma{})

	tableName := _skorRekomPeminatanSma.skorRekomPeminatanSmaDo.TableName()
	_skorRekomPeminatanSma.ALL = field.NewAsterisk(tableName)
	_skorRekomPeminatanSma.IDUser = field.NewInt32(tableName, "id_user")
	_skorRekomPeminatanSma.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorRekomPeminatanSma.RekomMinat = field.NewString(tableName, "rekom_minat")
	_skorRekomPeminatanSma.RekomSikapPelajaran = field.NewString(tableName, "rekom_sikap_pelajaran")
	_skorRekomPeminatanSma.RekomMapel = field.NewString(tableName, "rekom_mapel")

	_skorRekomPeminatanSma.fillFieldMap()

	return _skorRekomPeminatanSma
}

type skorRekomPeminatanSma struct {
	skorRekomPeminatanSmaDo skorRekomPeminatanSmaDo

	ALL                 field.Asterisk
	IDUser              field.Int32
	IDQuiz              field.Int32
	RekomMinat          field.String
	RekomSikapPelajaran field.String
	RekomMapel          field.String

	fieldMap map[string]field.Expr
}

func (s skorRekomPeminatanSma) Table(newTableName string) *skorRekomPeminatanSma {
	s.skorRekomPeminatanSmaDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorRekomPeminatanSma) As(alias string) *skorRekomPeminatanSma {
	s.skorRekomPeminatanSmaDo.DO = *(s.skorRekomPeminatanSmaDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorRekomPeminatanSma) updateTableName(table string) *skorRekomPeminatanSma {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.RekomMinat = field.NewString(table, "rekom_minat")
	s.RekomSikapPelajaran = field.NewString(table, "rekom_sikap_pelajaran")
	s.RekomMapel = field.NewString(table, "rekom_mapel")

	s.fillFieldMap()

	return s
}

func (s *skorRekomPeminatanSma) WithContext(ctx context.Context) *skorRekomPeminatanSmaDo {
	return s.skorRekomPeminatanSmaDo.WithContext(ctx)
}

func (s skorRekomPeminatanSma) TableName() string { return s.skorRekomPeminatanSmaDo.TableName() }

func (s skorRekomPeminatanSma) Alias() string { return s.skorRekomPeminatanSmaDo.Alias() }

func (s skorRekomPeminatanSma) Columns(cols ...field.Expr) gen.Columns {
	return s.skorRekomPeminatanSmaDo.Columns(cols...)
}

func (s *skorRekomPeminatanSma) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorRekomPeminatanSma) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["rekom_minat"] = s.RekomMinat
	s.fieldMap["rekom_sikap_pelajaran"] = s.RekomSikapPelajaran
	s.fieldMap["rekom_mapel"] = s.RekomMapel
}

func (s skorRekomPeminatanSma) clone(db *gorm.DB) skorRekomPeminatanSma {
	s.skorRekomPeminatanSmaDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorRekomPeminatanSma) replaceDB(db *gorm.DB) skorRekomPeminatanSma {
	s.skorRekomPeminatanSmaDo.ReplaceDB(db)
	return s
}

type skorRekomPeminatanSmaDo struct{ gen.DO }

func (s skorRekomPeminatanSmaDo) Debug() *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Debug())
}

func (s skorRekomPeminatanSmaDo) WithContext(ctx context.Context) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorRekomPeminatanSmaDo) ReadDB() *skorRekomPeminatanSmaDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorRekomPeminatanSmaDo) WriteDB() *skorRekomPeminatanSmaDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorRekomPeminatanSmaDo) Session(config *gorm.Session) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorRekomPeminatanSmaDo) Clauses(conds ...clause.Expression) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorRekomPeminatanSmaDo) Returning(value interface{}, columns ...string) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorRekomPeminatanSmaDo) Not(conds ...gen.Condition) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorRekomPeminatanSmaDo) Or(conds ...gen.Condition) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorRekomPeminatanSmaDo) Select(conds ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorRekomPeminatanSmaDo) Where(conds ...gen.Condition) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorRekomPeminatanSmaDo) Order(conds ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorRekomPeminatanSmaDo) Distinct(cols ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorRekomPeminatanSmaDo) Omit(cols ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorRekomPeminatanSmaDo) Join(table schema.Tabler, on ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorRekomPeminatanSmaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorRekomPeminatanSmaDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorRekomPeminatanSmaDo) Group(cols ...field.Expr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorRekomPeminatanSmaDo) Having(conds ...gen.Condition) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorRekomPeminatanSmaDo) Limit(limit int) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorRekomPeminatanSmaDo) Offset(offset int) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorRekomPeminatanSmaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorRekomPeminatanSmaDo) Unscoped() *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorRekomPeminatanSmaDo) Create(values ...*entity.SkorRekomPeminatanSma) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorRekomPeminatanSmaDo) CreateInBatches(values []*entity.SkorRekomPeminatanSma, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorRekomPeminatanSmaDo) Save(values ...*entity.SkorRekomPeminatanSma) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorRekomPeminatanSmaDo) First() (*entity.SkorRekomPeminatanSma, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomPeminatanSma), nil
	}
}

func (s skorRekomPeminatanSmaDo) Take() (*entity.SkorRekomPeminatanSma, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomPeminatanSma), nil
	}
}

func (s skorRekomPeminatanSmaDo) Last() (*entity.SkorRekomPeminatanSma, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomPeminatanSma), nil
	}
}

func (s skorRekomPeminatanSmaDo) Find() ([]*entity.SkorRekomPeminatanSma, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorRekomPeminatanSma), err
}

func (s skorRekomPeminatanSmaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorRekomPeminatanSma, err error) {
	buf := make([]*entity.SkorRekomPeminatanSma, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorRekomPeminatanSmaDo) FindInBatches(result *[]*entity.SkorRekomPeminatanSma, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorRekomPeminatanSmaDo) Attrs(attrs ...field.AssignExpr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorRekomPeminatanSmaDo) Assign(attrs ...field.AssignExpr) *skorRekomPeminatanSmaDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorRekomPeminatanSmaDo) Joins(fields ...field.RelationField) *skorRekomPeminatanSmaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorRekomPeminatanSmaDo) Preload(fields ...field.RelationField) *skorRekomPeminatanSmaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorRekomPeminatanSmaDo) FirstOrInit() (*entity.SkorRekomPeminatanSma, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomPeminatanSma), nil
	}
}

func (s skorRekomPeminatanSmaDo) FirstOrCreate() (*entity.SkorRekomPeminatanSma, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorRekomPeminatanSma), nil
	}
}

func (s skorRekomPeminatanSmaDo) FindByPage(offset int, limit int) (result []*entity.SkorRekomPeminatanSma, count int64, err error) {
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

func (s skorRekomPeminatanSmaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorRekomPeminatanSmaDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorRekomPeminatanSmaDo) Delete(models ...*entity.SkorRekomPeminatanSma) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorRekomPeminatanSmaDo) withDO(do gen.Dao) *skorRekomPeminatanSmaDo {
	s.DO = *do.(*gen.DO)
	return s
}