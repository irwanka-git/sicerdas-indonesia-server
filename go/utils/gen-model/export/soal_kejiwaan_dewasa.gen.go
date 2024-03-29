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

func newSoalKejiwaanDewasa(db *gorm.DB, opts ...gen.DOOption) soalKejiwaanDewasa {
	_soalKejiwaanDewasa := soalKejiwaanDewasa{}

	_soalKejiwaanDewasa.soalKejiwaanDewasaDo.UseDB(db, opts...)
	_soalKejiwaanDewasa.soalKejiwaanDewasaDo.UseModel(&entity.SoalKejiwaanDewasa{})

	tableName := _soalKejiwaanDewasa.soalKejiwaanDewasaDo.TableName()
	_soalKejiwaanDewasa.ALL = field.NewAsterisk(tableName)
	_soalKejiwaanDewasa.IDSoal = field.NewInt16(tableName, "id_soal")
	_soalKejiwaanDewasa.IDModel = field.NewInt16(tableName, "id_model")
	_soalKejiwaanDewasa.Unsur = field.NewString(tableName, "unsur")
	_soalKejiwaanDewasa.Urutan = field.NewInt16(tableName, "urutan")

	_soalKejiwaanDewasa.fillFieldMap()

	return _soalKejiwaanDewasa
}

type soalKejiwaanDewasa struct {
	soalKejiwaanDewasaDo soalKejiwaanDewasaDo

	ALL     field.Asterisk
	IDSoal  field.Int16
	IDModel field.Int16
	Unsur   field.String
	Urutan  field.Int16

	fieldMap map[string]field.Expr
}

func (s soalKejiwaanDewasa) Table(newTableName string) *soalKejiwaanDewasa {
	s.soalKejiwaanDewasaDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s soalKejiwaanDewasa) As(alias string) *soalKejiwaanDewasa {
	s.soalKejiwaanDewasaDo.DO = *(s.soalKejiwaanDewasaDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *soalKejiwaanDewasa) updateTableName(table string) *soalKejiwaanDewasa {
	s.ALL = field.NewAsterisk(table)
	s.IDSoal = field.NewInt16(table, "id_soal")
	s.IDModel = field.NewInt16(table, "id_model")
	s.Unsur = field.NewString(table, "unsur")
	s.Urutan = field.NewInt16(table, "urutan")

	s.fillFieldMap()

	return s
}

func (s *soalKejiwaanDewasa) WithContext(ctx context.Context) *soalKejiwaanDewasaDo {
	return s.soalKejiwaanDewasaDo.WithContext(ctx)
}

func (s soalKejiwaanDewasa) TableName() string { return s.soalKejiwaanDewasaDo.TableName() }

func (s soalKejiwaanDewasa) Alias() string { return s.soalKejiwaanDewasaDo.Alias() }

func (s soalKejiwaanDewasa) Columns(cols ...field.Expr) gen.Columns {
	return s.soalKejiwaanDewasaDo.Columns(cols...)
}

func (s *soalKejiwaanDewasa) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *soalKejiwaanDewasa) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["id_soal"] = s.IDSoal
	s.fieldMap["id_model"] = s.IDModel
	s.fieldMap["unsur"] = s.Unsur
	s.fieldMap["urutan"] = s.Urutan
}

func (s soalKejiwaanDewasa) clone(db *gorm.DB) soalKejiwaanDewasa {
	s.soalKejiwaanDewasaDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s soalKejiwaanDewasa) replaceDB(db *gorm.DB) soalKejiwaanDewasa {
	s.soalKejiwaanDewasaDo.ReplaceDB(db)
	return s
}

type soalKejiwaanDewasaDo struct{ gen.DO }

func (s soalKejiwaanDewasaDo) Debug() *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Debug())
}

func (s soalKejiwaanDewasaDo) WithContext(ctx context.Context) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s soalKejiwaanDewasaDo) ReadDB() *soalKejiwaanDewasaDo {
	return s.Clauses(dbresolver.Read)
}

func (s soalKejiwaanDewasaDo) WriteDB() *soalKejiwaanDewasaDo {
	return s.Clauses(dbresolver.Write)
}

func (s soalKejiwaanDewasaDo) Session(config *gorm.Session) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Session(config))
}

func (s soalKejiwaanDewasaDo) Clauses(conds ...clause.Expression) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s soalKejiwaanDewasaDo) Returning(value interface{}, columns ...string) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s soalKejiwaanDewasaDo) Not(conds ...gen.Condition) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s soalKejiwaanDewasaDo) Or(conds ...gen.Condition) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s soalKejiwaanDewasaDo) Select(conds ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s soalKejiwaanDewasaDo) Where(conds ...gen.Condition) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s soalKejiwaanDewasaDo) Order(conds ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s soalKejiwaanDewasaDo) Distinct(cols ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s soalKejiwaanDewasaDo) Omit(cols ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s soalKejiwaanDewasaDo) Join(table schema.Tabler, on ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s soalKejiwaanDewasaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s soalKejiwaanDewasaDo) RightJoin(table schema.Tabler, on ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s soalKejiwaanDewasaDo) Group(cols ...field.Expr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s soalKejiwaanDewasaDo) Having(conds ...gen.Condition) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s soalKejiwaanDewasaDo) Limit(limit int) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s soalKejiwaanDewasaDo) Offset(offset int) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s soalKejiwaanDewasaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s soalKejiwaanDewasaDo) Unscoped() *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Unscoped())
}

func (s soalKejiwaanDewasaDo) Create(values ...*entity.SoalKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s soalKejiwaanDewasaDo) CreateInBatches(values []*entity.SoalKejiwaanDewasa, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s soalKejiwaanDewasaDo) Save(values ...*entity.SoalKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s soalKejiwaanDewasaDo) First() (*entity.SoalKejiwaanDewasa, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKejiwaanDewasa), nil
	}
}

func (s soalKejiwaanDewasaDo) Take() (*entity.SoalKejiwaanDewasa, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKejiwaanDewasa), nil
	}
}

func (s soalKejiwaanDewasaDo) Last() (*entity.SoalKejiwaanDewasa, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKejiwaanDewasa), nil
	}
}

func (s soalKejiwaanDewasaDo) Find() ([]*entity.SoalKejiwaanDewasa, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SoalKejiwaanDewasa), err
}

func (s soalKejiwaanDewasaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SoalKejiwaanDewasa, err error) {
	buf := make([]*entity.SoalKejiwaanDewasa, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s soalKejiwaanDewasaDo) FindInBatches(result *[]*entity.SoalKejiwaanDewasa, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s soalKejiwaanDewasaDo) Attrs(attrs ...field.AssignExpr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s soalKejiwaanDewasaDo) Assign(attrs ...field.AssignExpr) *soalKejiwaanDewasaDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s soalKejiwaanDewasaDo) Joins(fields ...field.RelationField) *soalKejiwaanDewasaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s soalKejiwaanDewasaDo) Preload(fields ...field.RelationField) *soalKejiwaanDewasaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s soalKejiwaanDewasaDo) FirstOrInit() (*entity.SoalKejiwaanDewasa, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKejiwaanDewasa), nil
	}
}

func (s soalKejiwaanDewasaDo) FirstOrCreate() (*entity.SoalKejiwaanDewasa, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKejiwaanDewasa), nil
	}
}

func (s soalKejiwaanDewasaDo) FindByPage(offset int, limit int) (result []*entity.SoalKejiwaanDewasa, count int64, err error) {
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

func (s soalKejiwaanDewasaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s soalKejiwaanDewasaDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s soalKejiwaanDewasaDo) Delete(models ...*entity.SoalKejiwaanDewasa) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *soalKejiwaanDewasaDo) withDO(do gen.Dao) *soalKejiwaanDewasaDo {
	s.DO = *do.(*gen.DO)
	return s
}
