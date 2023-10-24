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

func newSeriCetakanDoc(db *gorm.DB, opts ...gen.DOOption) seriCetakanDoc {
	_seriCetakanDoc := seriCetakanDoc{}

	_seriCetakanDoc.seriCetakanDocDo.UseDB(db, opts...)
	_seriCetakanDoc.seriCetakanDocDo.UseModel(&entity.SeriCetakanDoc{})

	tableName := _seriCetakanDoc.seriCetakanDocDo.TableName()
	_seriCetakanDoc.ALL = field.NewAsterisk(tableName)
	_seriCetakanDoc.ID = field.NewInt32(tableName, "id")
	_seriCetakanDoc.NoSeri = field.NewString(tableName, "no_seri")
	_seriCetakanDoc.Token = field.NewString(tableName, "token")
	_seriCetakanDoc.CreatedAt = field.NewTime(tableName, "created_at")
	_seriCetakanDoc.IDUser = field.NewInt32(tableName, "id_user")
	_seriCetakanDoc.URL = field.NewString(tableName, "url")
	_seriCetakanDoc.JenisTes = field.NewString(tableName, "jenis_tes")
	_seriCetakanDoc.Filename = field.NewString(tableName, "filename")
	_seriCetakanDoc.Pathname = field.NewString(tableName, "pathname")

	_seriCetakanDoc.fillFieldMap()

	return _seriCetakanDoc
}

type seriCetakanDoc struct {
	seriCetakanDocDo seriCetakanDocDo

	ALL       field.Asterisk
	ID        field.Int32
	NoSeri    field.String
	Token     field.String
	CreatedAt field.Time
	IDUser    field.Int32
	URL       field.String
	JenisTes  field.String
	Filename  field.String
	Pathname  field.String

	fieldMap map[string]field.Expr
}

func (s seriCetakanDoc) Table(newTableName string) *seriCetakanDoc {
	s.seriCetakanDocDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seriCetakanDoc) As(alias string) *seriCetakanDoc {
	s.seriCetakanDocDo.DO = *(s.seriCetakanDocDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seriCetakanDoc) updateTableName(table string) *seriCetakanDoc {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.NoSeri = field.NewString(table, "no_seri")
	s.Token = field.NewString(table, "token")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.IDUser = field.NewInt32(table, "id_user")
	s.URL = field.NewString(table, "url")
	s.JenisTes = field.NewString(table, "jenis_tes")
	s.Filename = field.NewString(table, "filename")
	s.Pathname = field.NewString(table, "pathname")

	s.fillFieldMap()

	return s
}

func (s *seriCetakanDoc) WithContext(ctx context.Context) *seriCetakanDocDo {
	return s.seriCetakanDocDo.WithContext(ctx)
}

func (s seriCetakanDoc) TableName() string { return s.seriCetakanDocDo.TableName() }

func (s seriCetakanDoc) Alias() string { return s.seriCetakanDocDo.Alias() }

func (s seriCetakanDoc) Columns(cols ...field.Expr) gen.Columns {
	return s.seriCetakanDocDo.Columns(cols...)
}

func (s *seriCetakanDoc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seriCetakanDoc) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["no_seri"] = s.NoSeri
	s.fieldMap["token"] = s.Token
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["url"] = s.URL
	s.fieldMap["jenis_tes"] = s.JenisTes
	s.fieldMap["filename"] = s.Filename
	s.fieldMap["pathname"] = s.Pathname
}

func (s seriCetakanDoc) clone(db *gorm.DB) seriCetakanDoc {
	s.seriCetakanDocDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seriCetakanDoc) replaceDB(db *gorm.DB) seriCetakanDoc {
	s.seriCetakanDocDo.ReplaceDB(db)
	return s
}

type seriCetakanDocDo struct{ gen.DO }

func (s seriCetakanDocDo) Debug() *seriCetakanDocDo {
	return s.withDO(s.DO.Debug())
}

func (s seriCetakanDocDo) WithContext(ctx context.Context) *seriCetakanDocDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seriCetakanDocDo) ReadDB() *seriCetakanDocDo {
	return s.Clauses(dbresolver.Read)
}

func (s seriCetakanDocDo) WriteDB() *seriCetakanDocDo {
	return s.Clauses(dbresolver.Write)
}

func (s seriCetakanDocDo) Session(config *gorm.Session) *seriCetakanDocDo {
	return s.withDO(s.DO.Session(config))
}

func (s seriCetakanDocDo) Clauses(conds ...clause.Expression) *seriCetakanDocDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seriCetakanDocDo) Returning(value interface{}, columns ...string) *seriCetakanDocDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seriCetakanDocDo) Not(conds ...gen.Condition) *seriCetakanDocDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seriCetakanDocDo) Or(conds ...gen.Condition) *seriCetakanDocDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seriCetakanDocDo) Select(conds ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seriCetakanDocDo) Where(conds ...gen.Condition) *seriCetakanDocDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seriCetakanDocDo) Order(conds ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seriCetakanDocDo) Distinct(cols ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seriCetakanDocDo) Omit(cols ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seriCetakanDocDo) Join(table schema.Tabler, on ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seriCetakanDocDo) LeftJoin(table schema.Tabler, on ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seriCetakanDocDo) RightJoin(table schema.Tabler, on ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seriCetakanDocDo) Group(cols ...field.Expr) *seriCetakanDocDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seriCetakanDocDo) Having(conds ...gen.Condition) *seriCetakanDocDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seriCetakanDocDo) Limit(limit int) *seriCetakanDocDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seriCetakanDocDo) Offset(offset int) *seriCetakanDocDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seriCetakanDocDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *seriCetakanDocDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seriCetakanDocDo) Unscoped() *seriCetakanDocDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seriCetakanDocDo) Create(values ...*entity.SeriCetakanDoc) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seriCetakanDocDo) CreateInBatches(values []*entity.SeriCetakanDoc, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seriCetakanDocDo) Save(values ...*entity.SeriCetakanDoc) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seriCetakanDocDo) First() (*entity.SeriCetakanDoc, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SeriCetakanDoc), nil
	}
}

func (s seriCetakanDocDo) Take() (*entity.SeriCetakanDoc, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SeriCetakanDoc), nil
	}
}

func (s seriCetakanDocDo) Last() (*entity.SeriCetakanDoc, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SeriCetakanDoc), nil
	}
}

func (s seriCetakanDocDo) Find() ([]*entity.SeriCetakanDoc, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SeriCetakanDoc), err
}

func (s seriCetakanDocDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SeriCetakanDoc, err error) {
	buf := make([]*entity.SeriCetakanDoc, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seriCetakanDocDo) FindInBatches(result *[]*entity.SeriCetakanDoc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seriCetakanDocDo) Attrs(attrs ...field.AssignExpr) *seriCetakanDocDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seriCetakanDocDo) Assign(attrs ...field.AssignExpr) *seriCetakanDocDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seriCetakanDocDo) Joins(fields ...field.RelationField) *seriCetakanDocDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seriCetakanDocDo) Preload(fields ...field.RelationField) *seriCetakanDocDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seriCetakanDocDo) FirstOrInit() (*entity.SeriCetakanDoc, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SeriCetakanDoc), nil
	}
}

func (s seriCetakanDocDo) FirstOrCreate() (*entity.SeriCetakanDoc, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SeriCetakanDoc), nil
	}
}

func (s seriCetakanDocDo) FindByPage(offset int, limit int) (result []*entity.SeriCetakanDoc, count int64, err error) {
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

func (s seriCetakanDocDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seriCetakanDocDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seriCetakanDocDo) Delete(models ...*entity.SeriCetakanDoc) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seriCetakanDocDo) withDO(do gen.Dao) *seriCetakanDocDo {
	s.DO = *do.(*gen.DO)
	return s
}