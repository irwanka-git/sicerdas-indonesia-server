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

func newRefSkalaKdPenalaranMekanika(db *gorm.DB, opts ...gen.DOOption) refSkalaKdPenalaranMekanika {
	_refSkalaKdPenalaranMekanika := refSkalaKdPenalaranMekanika{}

	_refSkalaKdPenalaranMekanika.refSkalaKdPenalaranMekanikaDo.UseDB(db, opts...)
	_refSkalaKdPenalaranMekanika.refSkalaKdPenalaranMekanikaDo.UseModel(&entity.RefSkalaKdPenalaranMekanika{})

	tableName := _refSkalaKdPenalaranMekanika.refSkalaKdPenalaranMekanikaDo.TableName()
	_refSkalaKdPenalaranMekanika.ALL = field.NewAsterisk(tableName)
	_refSkalaKdPenalaranMekanika.ID = field.NewInt32(tableName, "id")
	_refSkalaKdPenalaranMekanika.Skor = field.NewInt32(tableName, "skor")
	_refSkalaKdPenalaranMekanika.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_refSkalaKdPenalaranMekanika.fillFieldMap()

	return _refSkalaKdPenalaranMekanika
}

type refSkalaKdPenalaranMekanika struct {
	refSkalaKdPenalaranMekanikaDo refSkalaKdPenalaranMekanikaDo

	ALL         field.Asterisk
	ID          field.Int32
	Skor        field.Int32
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refSkalaKdPenalaranMekanika) Table(newTableName string) *refSkalaKdPenalaranMekanika {
	r.refSkalaKdPenalaranMekanikaDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkalaKdPenalaranMekanika) As(alias string) *refSkalaKdPenalaranMekanika {
	r.refSkalaKdPenalaranMekanikaDo.DO = *(r.refSkalaKdPenalaranMekanikaDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkalaKdPenalaranMekanika) updateTableName(table string) *refSkalaKdPenalaranMekanika {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Skor = field.NewInt32(table, "skor")
	r.Klasifikasi = field.NewString(table, "klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refSkalaKdPenalaranMekanika) WithContext(ctx context.Context) *refSkalaKdPenalaranMekanikaDo {
	return r.refSkalaKdPenalaranMekanikaDo.WithContext(ctx)
}

func (r refSkalaKdPenalaranMekanika) TableName() string {
	return r.refSkalaKdPenalaranMekanikaDo.TableName()
}

func (r refSkalaKdPenalaranMekanika) Alias() string { return r.refSkalaKdPenalaranMekanikaDo.Alias() }

func (r refSkalaKdPenalaranMekanika) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkalaKdPenalaranMekanikaDo.Columns(cols...)
}

func (r *refSkalaKdPenalaranMekanika) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkalaKdPenalaranMekanika) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["id"] = r.ID
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["klasifikasi"] = r.Klasifikasi
}

func (r refSkalaKdPenalaranMekanika) clone(db *gorm.DB) refSkalaKdPenalaranMekanika {
	r.refSkalaKdPenalaranMekanikaDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkalaKdPenalaranMekanika) replaceDB(db *gorm.DB) refSkalaKdPenalaranMekanika {
	r.refSkalaKdPenalaranMekanikaDo.ReplaceDB(db)
	return r
}

type refSkalaKdPenalaranMekanikaDo struct{ gen.DO }

func (r refSkalaKdPenalaranMekanikaDo) Debug() *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkalaKdPenalaranMekanikaDo) WithContext(ctx context.Context) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkalaKdPenalaranMekanikaDo) ReadDB() *refSkalaKdPenalaranMekanikaDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkalaKdPenalaranMekanikaDo) WriteDB() *refSkalaKdPenalaranMekanikaDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkalaKdPenalaranMekanikaDo) Session(config *gorm.Session) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkalaKdPenalaranMekanikaDo) Clauses(conds ...clause.Expression) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Returning(value interface{}, columns ...string) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkalaKdPenalaranMekanikaDo) Not(conds ...gen.Condition) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Or(conds ...gen.Condition) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Select(conds ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Where(conds ...gen.Condition) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Order(conds ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Distinct(cols ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkalaKdPenalaranMekanikaDo) Omit(cols ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkalaKdPenalaranMekanikaDo) Join(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkalaKdPenalaranMekanikaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkalaKdPenalaranMekanikaDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkalaKdPenalaranMekanikaDo) Group(cols ...field.Expr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkalaKdPenalaranMekanikaDo) Having(conds ...gen.Condition) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkalaKdPenalaranMekanikaDo) Limit(limit int) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkalaKdPenalaranMekanikaDo) Offset(offset int) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkalaKdPenalaranMekanikaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkalaKdPenalaranMekanikaDo) Unscoped() *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkalaKdPenalaranMekanikaDo) Create(values ...*entity.RefSkalaKdPenalaranMekanika) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkalaKdPenalaranMekanikaDo) CreateInBatches(values []*entity.RefSkalaKdPenalaranMekanika, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkalaKdPenalaranMekanikaDo) Save(values ...*entity.RefSkalaKdPenalaranMekanika) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkalaKdPenalaranMekanikaDo) First() (*entity.RefSkalaKdPenalaranMekanika, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranMekanika), nil
	}
}

func (r refSkalaKdPenalaranMekanikaDo) Take() (*entity.RefSkalaKdPenalaranMekanika, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranMekanika), nil
	}
}

func (r refSkalaKdPenalaranMekanikaDo) Last() (*entity.RefSkalaKdPenalaranMekanika, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranMekanika), nil
	}
}

func (r refSkalaKdPenalaranMekanikaDo) Find() ([]*entity.RefSkalaKdPenalaranMekanika, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkalaKdPenalaranMekanika), err
}

func (r refSkalaKdPenalaranMekanikaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkalaKdPenalaranMekanika, err error) {
	buf := make([]*entity.RefSkalaKdPenalaranMekanika, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkalaKdPenalaranMekanikaDo) FindInBatches(result *[]*entity.RefSkalaKdPenalaranMekanika, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkalaKdPenalaranMekanikaDo) Attrs(attrs ...field.AssignExpr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkalaKdPenalaranMekanikaDo) Assign(attrs ...field.AssignExpr) *refSkalaKdPenalaranMekanikaDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkalaKdPenalaranMekanikaDo) Joins(fields ...field.RelationField) *refSkalaKdPenalaranMekanikaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkalaKdPenalaranMekanikaDo) Preload(fields ...field.RelationField) *refSkalaKdPenalaranMekanikaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkalaKdPenalaranMekanikaDo) FirstOrInit() (*entity.RefSkalaKdPenalaranMekanika, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranMekanika), nil
	}
}

func (r refSkalaKdPenalaranMekanikaDo) FirstOrCreate() (*entity.RefSkalaKdPenalaranMekanika, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkalaKdPenalaranMekanika), nil
	}
}

func (r refSkalaKdPenalaranMekanikaDo) FindByPage(offset int, limit int) (result []*entity.RefSkalaKdPenalaranMekanika, count int64, err error) {
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

func (r refSkalaKdPenalaranMekanikaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkalaKdPenalaranMekanikaDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkalaKdPenalaranMekanikaDo) Delete(models ...*entity.RefSkalaKdPenalaranMekanika) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkalaKdPenalaranMekanikaDo) withDO(do gen.Dao) *refSkalaKdPenalaranMekanikaDo {
	r.DO = *do.(*gen.DO)
	return r
}