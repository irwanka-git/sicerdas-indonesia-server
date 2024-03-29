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

func newRefModelKejiwaanDewasa(db *gorm.DB, opts ...gen.DOOption) refModelKejiwaanDewasa {
	_refModelKejiwaanDewasa := refModelKejiwaanDewasa{}

	_refModelKejiwaanDewasa.refModelKejiwaanDewasaDo.UseDB(db, opts...)
	_refModelKejiwaanDewasa.refModelKejiwaanDewasaDo.UseModel(&entity.RefModelKejiwaanDewasa{})

	tableName := _refModelKejiwaanDewasa.refModelKejiwaanDewasaDo.TableName()
	_refModelKejiwaanDewasa.ALL = field.NewAsterisk(tableName)
	_refModelKejiwaanDewasa.ID = field.NewInt32(tableName, "id")
	_refModelKejiwaanDewasa.Nama = field.NewString(tableName, "nama")
	_refModelKejiwaanDewasa.FieldSkoring = field.NewString(tableName, "field_skoring")

	_refModelKejiwaanDewasa.fillFieldMap()

	return _refModelKejiwaanDewasa
}

type refModelKejiwaanDewasa struct {
	refModelKejiwaanDewasaDo refModelKejiwaanDewasaDo

	ALL          field.Asterisk
	ID           field.Int32
	Nama         field.String
	FieldSkoring field.String

	fieldMap map[string]field.Expr
}

func (r refModelKejiwaanDewasa) Table(newTableName string) *refModelKejiwaanDewasa {
	r.refModelKejiwaanDewasaDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refModelKejiwaanDewasa) As(alias string) *refModelKejiwaanDewasa {
	r.refModelKejiwaanDewasaDo.DO = *(r.refModelKejiwaanDewasaDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refModelKejiwaanDewasa) updateTableName(table string) *refModelKejiwaanDewasa {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Nama = field.NewString(table, "nama")
	r.FieldSkoring = field.NewString(table, "field_skoring")

	r.fillFieldMap()

	return r
}

func (r *refModelKejiwaanDewasa) WithContext(ctx context.Context) *refModelKejiwaanDewasaDo {
	return r.refModelKejiwaanDewasaDo.WithContext(ctx)
}

func (r refModelKejiwaanDewasa) TableName() string { return r.refModelKejiwaanDewasaDo.TableName() }

func (r refModelKejiwaanDewasa) Alias() string { return r.refModelKejiwaanDewasaDo.Alias() }

func (r refModelKejiwaanDewasa) Columns(cols ...field.Expr) gen.Columns {
	return r.refModelKejiwaanDewasaDo.Columns(cols...)
}

func (r *refModelKejiwaanDewasa) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refModelKejiwaanDewasa) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["id"] = r.ID
	r.fieldMap["nama"] = r.Nama
	r.fieldMap["field_skoring"] = r.FieldSkoring
}

func (r refModelKejiwaanDewasa) clone(db *gorm.DB) refModelKejiwaanDewasa {
	r.refModelKejiwaanDewasaDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refModelKejiwaanDewasa) replaceDB(db *gorm.DB) refModelKejiwaanDewasa {
	r.refModelKejiwaanDewasaDo.ReplaceDB(db)
	return r
}

type refModelKejiwaanDewasaDo struct{ gen.DO }

func (r refModelKejiwaanDewasaDo) Debug() *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Debug())
}

func (r refModelKejiwaanDewasaDo) WithContext(ctx context.Context) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refModelKejiwaanDewasaDo) ReadDB() *refModelKejiwaanDewasaDo {
	return r.Clauses(dbresolver.Read)
}

func (r refModelKejiwaanDewasaDo) WriteDB() *refModelKejiwaanDewasaDo {
	return r.Clauses(dbresolver.Write)
}

func (r refModelKejiwaanDewasaDo) Session(config *gorm.Session) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Session(config))
}

func (r refModelKejiwaanDewasaDo) Clauses(conds ...clause.Expression) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refModelKejiwaanDewasaDo) Returning(value interface{}, columns ...string) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refModelKejiwaanDewasaDo) Not(conds ...gen.Condition) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refModelKejiwaanDewasaDo) Or(conds ...gen.Condition) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refModelKejiwaanDewasaDo) Select(conds ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refModelKejiwaanDewasaDo) Where(conds ...gen.Condition) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refModelKejiwaanDewasaDo) Order(conds ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refModelKejiwaanDewasaDo) Distinct(cols ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refModelKejiwaanDewasaDo) Omit(cols ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refModelKejiwaanDewasaDo) Join(table schema.Tabler, on ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refModelKejiwaanDewasaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refModelKejiwaanDewasaDo) RightJoin(table schema.Tabler, on ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refModelKejiwaanDewasaDo) Group(cols ...field.Expr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refModelKejiwaanDewasaDo) Having(conds ...gen.Condition) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refModelKejiwaanDewasaDo) Limit(limit int) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refModelKejiwaanDewasaDo) Offset(offset int) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refModelKejiwaanDewasaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refModelKejiwaanDewasaDo) Unscoped() *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refModelKejiwaanDewasaDo) Create(values ...*entity.RefModelKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refModelKejiwaanDewasaDo) CreateInBatches(values []*entity.RefModelKejiwaanDewasa, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refModelKejiwaanDewasaDo) Save(values ...*entity.RefModelKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refModelKejiwaanDewasaDo) First() (*entity.RefModelKejiwaanDewasa, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefModelKejiwaanDewasa), nil
	}
}

func (r refModelKejiwaanDewasaDo) Take() (*entity.RefModelKejiwaanDewasa, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefModelKejiwaanDewasa), nil
	}
}

func (r refModelKejiwaanDewasaDo) Last() (*entity.RefModelKejiwaanDewasa, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefModelKejiwaanDewasa), nil
	}
}

func (r refModelKejiwaanDewasaDo) Find() ([]*entity.RefModelKejiwaanDewasa, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefModelKejiwaanDewasa), err
}

func (r refModelKejiwaanDewasaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefModelKejiwaanDewasa, err error) {
	buf := make([]*entity.RefModelKejiwaanDewasa, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refModelKejiwaanDewasaDo) FindInBatches(result *[]*entity.RefModelKejiwaanDewasa, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refModelKejiwaanDewasaDo) Attrs(attrs ...field.AssignExpr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refModelKejiwaanDewasaDo) Assign(attrs ...field.AssignExpr) *refModelKejiwaanDewasaDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refModelKejiwaanDewasaDo) Joins(fields ...field.RelationField) *refModelKejiwaanDewasaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refModelKejiwaanDewasaDo) Preload(fields ...field.RelationField) *refModelKejiwaanDewasaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refModelKejiwaanDewasaDo) FirstOrInit() (*entity.RefModelKejiwaanDewasa, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefModelKejiwaanDewasa), nil
	}
}

func (r refModelKejiwaanDewasaDo) FirstOrCreate() (*entity.RefModelKejiwaanDewasa, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefModelKejiwaanDewasa), nil
	}
}

func (r refModelKejiwaanDewasaDo) FindByPage(offset int, limit int) (result []*entity.RefModelKejiwaanDewasa, count int64, err error) {
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

func (r refModelKejiwaanDewasaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refModelKejiwaanDewasaDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refModelKejiwaanDewasaDo) Delete(models ...*entity.RefModelKejiwaanDewasa) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refModelKejiwaanDewasaDo) withDO(do gen.Dao) *refModelKejiwaanDewasaDo {
	r.DO = *do.(*gen.DO)
	return r
}
