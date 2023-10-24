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

func newRefSkalaKdPenalaranSpasial(db *gorm.DB, opts ...gen.DOOption) refSkalaKdPenalaranSpasial {
	_refSkalaKdPenalaranSpasial := refSkalaKdPenalaranSpasial{}

	_refSkalaKdPenalaranSpasial.refSkalaKdPenalaranSpasialDo.UseDB(db, opts...)
	_refSkalaKdPenalaranSpasial.refSkalaKdPenalaranSpasialDo.UseModel(&entity.RefSkalaKdPenalaranSpasial{})

	tableName := _refSkalaKdPenalaranSpasial.refSkalaKdPenalaranSpasialDo.TableName()
	_refSkalaKdPenalaranSpasial.ALL = field.NewAsterisk(tableName)
	_refSkalaKdPenalaranSpasial.ID = field.NewInt32(tableName, "id")
	_refSkalaKdPenalaranSpasial.Skor = field.NewInt32(tableName, "skor")
	_refSkalaKdPenalaranSpasial.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_refSkalaKdPenalaranSpasial.fillFieldMap()

	return _refSkalaKdPenalaranSpasial
}

type refSkalaKdPenalaranSpasial struct {
	refSkalaKdPenalaranSpasialDo refSkalaKdPenalaranSpasialDo

	ALL         field.Asterisk
	ID          field.Int32
	Skor        field.Int32
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refSkalaKdPenalaranSpasial) Table(newTableName string) *refSkalaKdPenalaranSpasial {
	r.refSkalaKdPenalaranSpasialDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkalaKdPenalaranSpasial) As(alias string) *refSkalaKdPenalaranSpasial {
	r.refSkalaKdPenalaranSpasialDo.DO = *(r.refSkalaKdPenalaranSpasialDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkalaKdPenalaranSpasial) updateTableName(table string) *refSkalaKdPenalaranSpasial {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Skor = field.NewInt32(table, "skor")
	r.Klasifikasi = field.NewString(table, "klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refSkalaKdPenalaranSpasial) WithContext(ctx context.Context) *refSkalaKdPenalaranSpasialDo {
	return r.refSkalaKdPenalaranSpasialDo.WithContext(ctx)
}

func (r refSkalaKdPenalaranSpasial) TableName() string {
	return r.refSkalaKdPenalaranSpasialDo.TableName()
}

func (r refSkalaKdPenalaranSpasial) Alias() string { return r.refSkalaKdPenalaranSpasialDo.Alias() }

func (r refSkalaKdPenalaranSpasial) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkalaKdPenalaranSpasialDo.Columns(cols...)
}

func (r *refSkalaKdPenalaranSpasial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkalaKdPenalaranSpasial) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["id"] = r.ID
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["klasifikasi"] = r.Klasifikasi
}

func (r refSkalaKdPenalaranSpasial) clone(db *gorm.DB) refSkalaKdPenalaranSpasial {
	r.refSkalaKdPenalaranSpasialDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkalaKdPenalaranSpasial) replaceDB(db *gorm.DB) refSkalaKdPenalaranSpasial {
	r.refSkalaKdPenalaranSpasialDo.ReplaceDB(db)
	return r
}

type refSkalaKdPenalaranSpasialDo struct{ gen.DO }

func (r refSkalaKdPenalaranSpasialDo) Debug() *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkalaKdPenalaranSpasialDo) WithContext(ctx context.Context) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkalaKdPenalaranSpasialDo) ReadDB() *refSkalaKdPenalaranSpasialDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkalaKdPenalaranSpasialDo) WriteDB() *refSkalaKdPenalaranSpasialDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkalaKdPenalaranSpasialDo) Session(config *gorm.Session) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkalaKdPenalaranSpasialDo) Clauses(conds ...clause.Expression) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Returning(value interface{}, columns ...string) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkalaKdPenalaranSpasialDo) Not(conds ...gen.Condition) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Or(conds ...gen.Condition) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Select(conds ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Where(conds ...gen.Condition) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Order(conds ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Distinct(cols ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkalaKdPenalaranSpasialDo) Omit(cols ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkalaKdPenalaranSpasialDo) Join(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkalaKdPenalaranSpasialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkalaKdPenalaranSpasialDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkalaKdPenalaranSpasialDo) Group(cols ...field.Expr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkalaKdPenalaranSpasialDo) Having(conds ...gen.Condition) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkalaKdPenalaranSpasialDo) Limit(limit int) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkalaKdPenalaranSpasialDo) Offset(offset int) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkalaKdPenalaranSpasialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkalaKdPenalaranSpasialDo) Unscoped() *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkalaKdPenalaranSpasialDo) Create(values ...*entity.RefSkalaKdPenalaranSpasial) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkalaKdPenalaranSpasialDo) CreateInBatches(values []*entity.RefSkalaKdPenalaranSpasial, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkalaKdPenalaranSpasialDo) Save(values ...*entity.RefSkalaKdPenalaranSpasial) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkalaKdPenalaranSpasialDo) First() (*entity.RefSkalaKdPenalaranSpasial, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranSpasial), nil
	}
}

func (r refSkalaKdPenalaranSpasialDo) Take() (*entity.RefSkalaKdPenalaranSpasial, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranSpasial), nil
	}
}

func (r refSkalaKdPenalaranSpasialDo) Last() (*entity.RefSkalaKdPenalaranSpasial, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranSpasial), nil
	}
}

func (r refSkalaKdPenalaranSpasialDo) Find() ([]*entity.RefSkalaKdPenalaranSpasial, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkalaKdPenalaranSpasial), err
}

func (r refSkalaKdPenalaranSpasialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkalaKdPenalaranSpasial, err error) {
	buf := make([]*entity.RefSkalaKdPenalaranSpasial, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkalaKdPenalaranSpasialDo) FindInBatches(result *[]*entity.RefSkalaKdPenalaranSpasial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkalaKdPenalaranSpasialDo) Attrs(attrs ...field.AssignExpr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkalaKdPenalaranSpasialDo) Assign(attrs ...field.AssignExpr) *refSkalaKdPenalaranSpasialDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkalaKdPenalaranSpasialDo) Joins(fields ...field.RelationField) *refSkalaKdPenalaranSpasialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkalaKdPenalaranSpasialDo) Preload(fields ...field.RelationField) *refSkalaKdPenalaranSpasialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkalaKdPenalaranSpasialDo) FirstOrInit() (*entity.RefSkalaKdPenalaranSpasial, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranSpasial), nil
	}
}

func (r refSkalaKdPenalaranSpasialDo) FirstOrCreate() (*entity.RefSkalaKdPenalaranSpasial, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranSpasial), nil
	}
}

func (r refSkalaKdPenalaranSpasialDo) FindByPage(offset int, limit int) (result []*entity.RefSkalaKdPenalaranSpasial, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r refSkalaKdPenalaranSpasialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkalaKdPenalaranSpasialDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkalaKdPenalaranSpasialDo) Delete(models ...*entity.RefSkalaKdPenalaranSpasial) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkalaKdPenalaranSpasialDo) withDO(do gen.Dao) *refSkalaKdPenalaranSpasialDo {
	r.DO = *do.(*gen.DO)
	return r
}