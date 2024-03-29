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

func newRefSkalaSkorSsct(db *gorm.DB, opts ...gen.DOOption) refSkalaSkorSsct {
	_refSkalaSkorSsct := refSkalaSkorSsct{}

	_refSkalaSkorSsct.refSkalaSkorSsctDo.UseDB(db, opts...)
	_refSkalaSkorSsct.refSkalaSkorSsctDo.UseModel(&entity.RefSkalaSkorSsct{})

	tableName := _refSkalaSkorSsct.refSkalaSkorSsctDo.TableName()
	_refSkalaSkorSsct.ALL = field.NewAsterisk(tableName)
	_refSkalaSkorSsct.ID = field.NewInt32(tableName, "id")
	_refSkalaSkorSsct.Skor = field.NewInt32(tableName, "skor")
	_refSkalaSkorSsct.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_refSkalaSkorSsct.fillFieldMap()

	return _refSkalaSkorSsct
}

type refSkalaSkorSsct struct {
	refSkalaSkorSsctDo refSkalaSkorSsctDo

	ALL         field.Asterisk
	ID          field.Int32
	Skor        field.Int32
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refSkalaSkorSsct) Table(newTableName string) *refSkalaSkorSsct {
	r.refSkalaSkorSsctDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkalaSkorSsct) As(alias string) *refSkalaSkorSsct {
	r.refSkalaSkorSsctDo.DO = *(r.refSkalaSkorSsctDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkalaSkorSsct) updateTableName(table string) *refSkalaSkorSsct {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Skor = field.NewInt32(table, "skor")
	r.Klasifikasi = field.NewString(table, "klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refSkalaSkorSsct) WithContext(ctx context.Context) *refSkalaSkorSsctDo {
	return r.refSkalaSkorSsctDo.WithContext(ctx)
}

func (r refSkalaSkorSsct) TableName() string { return r.refSkalaSkorSsctDo.TableName() }

func (r refSkalaSkorSsct) Alias() string { return r.refSkalaSkorSsctDo.Alias() }

func (r refSkalaSkorSsct) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkalaSkorSsctDo.Columns(cols...)
}

func (r *refSkalaSkorSsct) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkalaSkorSsct) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["id"] = r.ID
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["klasifikasi"] = r.Klasifikasi
}

func (r refSkalaSkorSsct) clone(db *gorm.DB) refSkalaSkorSsct {
	r.refSkalaSkorSsctDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkalaSkorSsct) replaceDB(db *gorm.DB) refSkalaSkorSsct {
	r.refSkalaSkorSsctDo.ReplaceDB(db)
	return r
}

type refSkalaSkorSsctDo struct{ gen.DO }

func (r refSkalaSkorSsctDo) Debug() *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkalaSkorSsctDo) WithContext(ctx context.Context) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkalaSkorSsctDo) ReadDB() *refSkalaSkorSsctDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkalaSkorSsctDo) WriteDB() *refSkalaSkorSsctDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkalaSkorSsctDo) Session(config *gorm.Session) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkalaSkorSsctDo) Clauses(conds ...clause.Expression) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkalaSkorSsctDo) Returning(value interface{}, columns ...string) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkalaSkorSsctDo) Not(conds ...gen.Condition) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkalaSkorSsctDo) Or(conds ...gen.Condition) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkalaSkorSsctDo) Select(conds ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkalaSkorSsctDo) Where(conds ...gen.Condition) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkalaSkorSsctDo) Order(conds ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkalaSkorSsctDo) Distinct(cols ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkalaSkorSsctDo) Omit(cols ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkalaSkorSsctDo) Join(table schema.Tabler, on ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkalaSkorSsctDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkalaSkorSsctDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkalaSkorSsctDo) Group(cols ...field.Expr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkalaSkorSsctDo) Having(conds ...gen.Condition) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkalaSkorSsctDo) Limit(limit int) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkalaSkorSsctDo) Offset(offset int) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkalaSkorSsctDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkalaSkorSsctDo) Unscoped() *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkalaSkorSsctDo) Create(values ...*entity.RefSkalaSkorSsct) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkalaSkorSsctDo) CreateInBatches(values []*entity.RefSkalaSkorSsct, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkalaSkorSsctDo) Save(values ...*entity.RefSkalaSkorSsct) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkalaSkorSsctDo) First() (*entity.RefSkalaSkorSsct, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaSkorSsct), nil
	}
}

func (r refSkalaSkorSsctDo) Take() (*entity.RefSkalaSkorSsct, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaSkorSsct), nil
	}
}

func (r refSkalaSkorSsctDo) Last() (*entity.RefSkalaSkorSsct, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaSkorSsct), nil
	}
}

func (r refSkalaSkorSsctDo) Find() ([]*entity.RefSkalaSkorSsct, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkalaSkorSsct), err
}

func (r refSkalaSkorSsctDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkalaSkorSsct, err error) {
	buf := make([]*entity.RefSkalaSkorSsct, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkalaSkorSsctDo) FindInBatches(result *[]*entity.RefSkalaSkorSsct, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkalaSkorSsctDo) Attrs(attrs ...field.AssignExpr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkalaSkorSsctDo) Assign(attrs ...field.AssignExpr) *refSkalaSkorSsctDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkalaSkorSsctDo) Joins(fields ...field.RelationField) *refSkalaSkorSsctDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkalaSkorSsctDo) Preload(fields ...field.RelationField) *refSkalaSkorSsctDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkalaSkorSsctDo) FirstOrInit() (*entity.RefSkalaSkorSsct, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaSkorSsct), nil
	}
}

func (r refSkalaSkorSsctDo) FirstOrCreate() (*entity.RefSkalaSkorSsct, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaSkorSsct), nil
	}
}

func (r refSkalaSkorSsctDo) FindByPage(offset int, limit int) (result []*entity.RefSkalaSkorSsct, count int64, err error) {
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

func (r refSkalaSkorSsctDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkalaSkorSsctDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkalaSkorSsctDo) Delete(models ...*entity.RefSkalaSkorSsct) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkalaSkorSsctDo) withDO(do gen.Dao) *refSkalaSkorSsctDo {
	r.DO = *do.(*gen.DO)
	return r
}
