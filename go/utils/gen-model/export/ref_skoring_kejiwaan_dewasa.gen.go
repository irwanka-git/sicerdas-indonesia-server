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

func newRefSkoringKejiwaanDewasa(db *gorm.DB, opts ...gen.DOOption) refSkoringKejiwaanDewasa {
	_refSkoringKejiwaanDewasa := refSkoringKejiwaanDewasa{}

	_refSkoringKejiwaanDewasa.refSkoringKejiwaanDewasaDo.UseDB(db, opts...)
	_refSkoringKejiwaanDewasa.refSkoringKejiwaanDewasaDo.UseModel(&entity.RefSkoringKejiwaanDewasa{})

	tableName := _refSkoringKejiwaanDewasa.refSkoringKejiwaanDewasaDo.TableName()
	_refSkoringKejiwaanDewasa.ALL = field.NewAsterisk(tableName)
	_refSkoringKejiwaanDewasa.Skor = field.NewInt16(tableName, "skor")
	_refSkoringKejiwaanDewasa.Nilai = field.NewInt16(tableName, "nilai")

	_refSkoringKejiwaanDewasa.fillFieldMap()

	return _refSkoringKejiwaanDewasa
}

type refSkoringKejiwaanDewasa struct {
	refSkoringKejiwaanDewasaDo refSkoringKejiwaanDewasaDo

	ALL   field.Asterisk
	Skor  field.Int16
	Nilai field.Int16

	fieldMap map[string]field.Expr
}

func (r refSkoringKejiwaanDewasa) Table(newTableName string) *refSkoringKejiwaanDewasa {
	r.refSkoringKejiwaanDewasaDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkoringKejiwaanDewasa) As(alias string) *refSkoringKejiwaanDewasa {
	r.refSkoringKejiwaanDewasaDo.DO = *(r.refSkoringKejiwaanDewasaDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkoringKejiwaanDewasa) updateTableName(table string) *refSkoringKejiwaanDewasa {
	r.ALL = field.NewAsterisk(table)
	r.Skor = field.NewInt16(table, "skor")
	r.Nilai = field.NewInt16(table, "nilai")

	r.fillFieldMap()

	return r
}

func (r *refSkoringKejiwaanDewasa) WithContext(ctx context.Context) *refSkoringKejiwaanDewasaDo {
	return r.refSkoringKejiwaanDewasaDo.WithContext(ctx)
}

func (r refSkoringKejiwaanDewasa) TableName() string { return r.refSkoringKejiwaanDewasaDo.TableName() }

func (r refSkoringKejiwaanDewasa) Alias() string { return r.refSkoringKejiwaanDewasaDo.Alias() }

func (r refSkoringKejiwaanDewasa) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkoringKejiwaanDewasaDo.Columns(cols...)
}

func (r *refSkoringKejiwaanDewasa) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkoringKejiwaanDewasa) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["nilai"] = r.Nilai
}

func (r refSkoringKejiwaanDewasa) clone(db *gorm.DB) refSkoringKejiwaanDewasa {
	r.refSkoringKejiwaanDewasaDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkoringKejiwaanDewasa) replaceDB(db *gorm.DB) refSkoringKejiwaanDewasa {
	r.refSkoringKejiwaanDewasaDo.ReplaceDB(db)
	return r
}

type refSkoringKejiwaanDewasaDo struct{ gen.DO }

func (r refSkoringKejiwaanDewasaDo) Debug() *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkoringKejiwaanDewasaDo) WithContext(ctx context.Context) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkoringKejiwaanDewasaDo) ReadDB() *refSkoringKejiwaanDewasaDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkoringKejiwaanDewasaDo) WriteDB() *refSkoringKejiwaanDewasaDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkoringKejiwaanDewasaDo) Session(config *gorm.Session) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkoringKejiwaanDewasaDo) Clauses(conds ...clause.Expression) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Returning(value interface{}, columns ...string) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkoringKejiwaanDewasaDo) Not(conds ...gen.Condition) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Or(conds ...gen.Condition) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Select(conds ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Where(conds ...gen.Condition) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Order(conds ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Distinct(cols ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkoringKejiwaanDewasaDo) Omit(cols ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkoringKejiwaanDewasaDo) Join(table schema.Tabler, on ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkoringKejiwaanDewasaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkoringKejiwaanDewasaDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkoringKejiwaanDewasaDo) Group(cols ...field.Expr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkoringKejiwaanDewasaDo) Having(conds ...gen.Condition) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkoringKejiwaanDewasaDo) Limit(limit int) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkoringKejiwaanDewasaDo) Offset(offset int) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkoringKejiwaanDewasaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkoringKejiwaanDewasaDo) Unscoped() *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkoringKejiwaanDewasaDo) Create(values ...*entity.RefSkoringKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkoringKejiwaanDewasaDo) CreateInBatches(values []*entity.RefSkoringKejiwaanDewasa, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkoringKejiwaanDewasaDo) Save(values ...*entity.RefSkoringKejiwaanDewasa) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkoringKejiwaanDewasaDo) First() (*entity.RefSkoringKejiwaanDewasa, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringKejiwaanDewasa), nil
	}
}

func (r refSkoringKejiwaanDewasaDo) Take() (*entity.RefSkoringKejiwaanDewasa, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringKejiwaanDewasa), nil
	}
}

func (r refSkoringKejiwaanDewasaDo) Last() (*entity.RefSkoringKejiwaanDewasa, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringKejiwaanDewasa), nil
	}
}

func (r refSkoringKejiwaanDewasaDo) Find() ([]*entity.RefSkoringKejiwaanDewasa, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkoringKejiwaanDewasa), err
}

func (r refSkoringKejiwaanDewasaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkoringKejiwaanDewasa, err error) {
	buf := make([]*entity.RefSkoringKejiwaanDewasa, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkoringKejiwaanDewasaDo) FindInBatches(result *[]*entity.RefSkoringKejiwaanDewasa, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkoringKejiwaanDewasaDo) Attrs(attrs ...field.AssignExpr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkoringKejiwaanDewasaDo) Assign(attrs ...field.AssignExpr) *refSkoringKejiwaanDewasaDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkoringKejiwaanDewasaDo) Joins(fields ...field.RelationField) *refSkoringKejiwaanDewasaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkoringKejiwaanDewasaDo) Preload(fields ...field.RelationField) *refSkoringKejiwaanDewasaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkoringKejiwaanDewasaDo) FirstOrInit() (*entity.RefSkoringKejiwaanDewasa, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringKejiwaanDewasa), nil
	}
}

func (r refSkoringKejiwaanDewasaDo) FirstOrCreate() (*entity.RefSkoringKejiwaanDewasa, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringKejiwaanDewasa), nil
	}
}

func (r refSkoringKejiwaanDewasaDo) FindByPage(offset int, limit int) (result []*entity.RefSkoringKejiwaanDewasa, count int64, err error) {
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

func (r refSkoringKejiwaanDewasaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkoringKejiwaanDewasaDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkoringKejiwaanDewasaDo) Delete(models ...*entity.RefSkoringKejiwaanDewasa) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkoringKejiwaanDewasaDo) withDO(do gen.Dao) *refSkoringKejiwaanDewasaDo {
	r.DO = *do.(*gen.DO)
	return r
}
