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

func newRefBidangKognitif(db *gorm.DB, opts ...gen.DOOption) refBidangKognitif {
	_refBidangKognitif := refBidangKognitif{}

	_refBidangKognitif.refBidangKognitifDo.UseDB(db, opts...)
	_refBidangKognitif.refBidangKognitifDo.UseModel(&entity.RefBidangKognitif{})

	tableName := _refBidangKognitif.refBidangKognitifDo.TableName()
	_refBidangKognitif.ALL = field.NewAsterisk(tableName)
	_refBidangKognitif.NamaBidang = field.NewString(tableName, "nama_bidang")
	_refBidangKognitif.FieldSkoring = field.NewString(tableName, "field_skoring")
	_refBidangKognitif.Deskripsi = field.NewString(tableName, "deskripsi")

	_refBidangKognitif.fillFieldMap()

	return _refBidangKognitif
}

type refBidangKognitif struct {
	refBidangKognitifDo refBidangKognitifDo

	ALL          field.Asterisk
	NamaBidang   field.String
	FieldSkoring field.String
	Deskripsi    field.String

	fieldMap map[string]field.Expr
}

func (r refBidangKognitif) Table(newTableName string) *refBidangKognitif {
	r.refBidangKognitifDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refBidangKognitif) As(alias string) *refBidangKognitif {
	r.refBidangKognitifDo.DO = *(r.refBidangKognitifDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refBidangKognitif) updateTableName(table string) *refBidangKognitif {
	r.ALL = field.NewAsterisk(table)
	r.NamaBidang = field.NewString(table, "nama_bidang")
	r.FieldSkoring = field.NewString(table, "field_skoring")
	r.Deskripsi = field.NewString(table, "deskripsi")

	r.fillFieldMap()

	return r
}

func (r *refBidangKognitif) WithContext(ctx context.Context) *refBidangKognitifDo {
	return r.refBidangKognitifDo.WithContext(ctx)
}

func (r refBidangKognitif) TableName() string { return r.refBidangKognitifDo.TableName() }

func (r refBidangKognitif) Alias() string { return r.refBidangKognitifDo.Alias() }

func (r refBidangKognitif) Columns(cols ...field.Expr) gen.Columns {
	return r.refBidangKognitifDo.Columns(cols...)
}

func (r *refBidangKognitif) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refBidangKognitif) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["nama_bidang"] = r.NamaBidang
	r.fieldMap["field_skoring"] = r.FieldSkoring
	r.fieldMap["deskripsi"] = r.Deskripsi
}

func (r refBidangKognitif) clone(db *gorm.DB) refBidangKognitif {
	r.refBidangKognitifDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refBidangKognitif) replaceDB(db *gorm.DB) refBidangKognitif {
	r.refBidangKognitifDo.ReplaceDB(db)
	return r
}

type refBidangKognitifDo struct{ gen.DO }

func (r refBidangKognitifDo) Debug() *refBidangKognitifDo {
	return r.withDO(r.DO.Debug())
}

func (r refBidangKognitifDo) WithContext(ctx context.Context) *refBidangKognitifDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refBidangKognitifDo) ReadDB() *refBidangKognitifDo {
	return r.Clauses(dbresolver.Read)
}

func (r refBidangKognitifDo) WriteDB() *refBidangKognitifDo {
	return r.Clauses(dbresolver.Write)
}

func (r refBidangKognitifDo) Session(config *gorm.Session) *refBidangKognitifDo {
	return r.withDO(r.DO.Session(config))
}

func (r refBidangKognitifDo) Clauses(conds ...clause.Expression) *refBidangKognitifDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refBidangKognitifDo) Returning(value interface{}, columns ...string) *refBidangKognitifDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refBidangKognitifDo) Not(conds ...gen.Condition) *refBidangKognitifDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refBidangKognitifDo) Or(conds ...gen.Condition) *refBidangKognitifDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refBidangKognitifDo) Select(conds ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refBidangKognitifDo) Where(conds ...gen.Condition) *refBidangKognitifDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refBidangKognitifDo) Order(conds ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refBidangKognitifDo) Distinct(cols ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refBidangKognitifDo) Omit(cols ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refBidangKognitifDo) Join(table schema.Tabler, on ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refBidangKognitifDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refBidangKognitifDo) RightJoin(table schema.Tabler, on ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refBidangKognitifDo) Group(cols ...field.Expr) *refBidangKognitifDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refBidangKognitifDo) Having(conds ...gen.Condition) *refBidangKognitifDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refBidangKognitifDo) Limit(limit int) *refBidangKognitifDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refBidangKognitifDo) Offset(offset int) *refBidangKognitifDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refBidangKognitifDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refBidangKognitifDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refBidangKognitifDo) Unscoped() *refBidangKognitifDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refBidangKognitifDo) Create(values ...*entity.RefBidangKognitif) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refBidangKognitifDo) CreateInBatches(values []*entity.RefBidangKognitif, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refBidangKognitifDo) Save(values ...*entity.RefBidangKognitif) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refBidangKognitifDo) First() (*entity.RefBidangKognitif, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefBidangKognitif), nil
	}
}

func (r refBidangKognitifDo) Take() (*entity.RefBidangKognitif, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefBidangKognitif), nil
	}
}

func (r refBidangKognitifDo) Last() (*entity.RefBidangKognitif, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefBidangKognitif), nil
	}
}

func (r refBidangKognitifDo) Find() ([]*entity.RefBidangKognitif, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefBidangKognitif), err
}

func (r refBidangKognitifDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefBidangKognitif, err error) {
	buf := make([]*entity.RefBidangKognitif, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refBidangKognitifDo) FindInBatches(result *[]*entity.RefBidangKognitif, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refBidangKognitifDo) Attrs(attrs ...field.AssignExpr) *refBidangKognitifDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refBidangKognitifDo) Assign(attrs ...field.AssignExpr) *refBidangKognitifDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refBidangKognitifDo) Joins(fields ...field.RelationField) *refBidangKognitifDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refBidangKognitifDo) Preload(fields ...field.RelationField) *refBidangKognitifDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refBidangKognitifDo) FirstOrInit() (*entity.RefBidangKognitif, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefBidangKognitif), nil
	}
}

func (r refBidangKognitifDo) FirstOrCreate() (*entity.RefBidangKognitif, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefBidangKognitif), nil
	}
}

func (r refBidangKognitifDo) FindByPage(offset int, limit int) (result []*entity.RefBidangKognitif, count int64, err error) {
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

func (r refBidangKognitifDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refBidangKognitifDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refBidangKognitifDo) Delete(models ...*entity.RefBidangKognitif) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refBidangKognitifDo) withDO(do gen.Dao) *refBidangKognitifDo {
	r.DO = *do.(*gen.DO)
	return r
}