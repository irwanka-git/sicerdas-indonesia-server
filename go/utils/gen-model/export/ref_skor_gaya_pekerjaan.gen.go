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

func newRefSkorGayaPekerjaan(db *gorm.DB, opts ...gen.DOOption) refSkorGayaPekerjaan {
	_refSkorGayaPekerjaan := refSkorGayaPekerjaan{}

	_refSkorGayaPekerjaan.refSkorGayaPekerjaanDo.UseDB(db, opts...)
	_refSkorGayaPekerjaan.refSkorGayaPekerjaanDo.UseModel(&entity.RefSkorGayaPekerjaan{})

	tableName := _refSkorGayaPekerjaan.refSkorGayaPekerjaanDo.TableName()
	_refSkorGayaPekerjaan.ALL = field.NewAsterisk(tableName)
	_refSkorGayaPekerjaan.ID = field.NewInt16(tableName, "id")
	_refSkorGayaPekerjaan.Respon = field.NewString(tableName, "respon")
	_refSkorGayaPekerjaan.U = field.NewInt16(tableName, "u")
	_refSkorGayaPekerjaan.T = field.NewInt16(tableName, "t")
	_refSkorGayaPekerjaan.C = field.NewInt16(tableName, "c")
	_refSkorGayaPekerjaan.Jawaban = field.NewString(tableName, "jawaban")

	_refSkorGayaPekerjaan.fillFieldMap()

	return _refSkorGayaPekerjaan
}

type refSkorGayaPekerjaan struct {
	refSkorGayaPekerjaanDo refSkorGayaPekerjaanDo

	ALL     field.Asterisk
	ID      field.Int16
	Respon  field.String
	U       field.Int16
	T       field.Int16
	C       field.Int16
	Jawaban field.String

	fieldMap map[string]field.Expr
}

func (r refSkorGayaPekerjaan) Table(newTableName string) *refSkorGayaPekerjaan {
	r.refSkorGayaPekerjaanDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkorGayaPekerjaan) As(alias string) *refSkorGayaPekerjaan {
	r.refSkorGayaPekerjaanDo.DO = *(r.refSkorGayaPekerjaanDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkorGayaPekerjaan) updateTableName(table string) *refSkorGayaPekerjaan {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt16(table, "id")
	r.Respon = field.NewString(table, "respon")
	r.U = field.NewInt16(table, "u")
	r.T = field.NewInt16(table, "t")
	r.C = field.NewInt16(table, "c")
	r.Jawaban = field.NewString(table, "jawaban")

	r.fillFieldMap()

	return r
}

func (r *refSkorGayaPekerjaan) WithContext(ctx context.Context) *refSkorGayaPekerjaanDo {
	return r.refSkorGayaPekerjaanDo.WithContext(ctx)
}

func (r refSkorGayaPekerjaan) TableName() string { return r.refSkorGayaPekerjaanDo.TableName() }

func (r refSkorGayaPekerjaan) Alias() string { return r.refSkorGayaPekerjaanDo.Alias() }

func (r refSkorGayaPekerjaan) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkorGayaPekerjaanDo.Columns(cols...)
}

func (r *refSkorGayaPekerjaan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkorGayaPekerjaan) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id"] = r.ID
	r.fieldMap["respon"] = r.Respon
	r.fieldMap["u"] = r.U
	r.fieldMap["t"] = r.T
	r.fieldMap["c"] = r.C
	r.fieldMap["jawaban"] = r.Jawaban
}

func (r refSkorGayaPekerjaan) clone(db *gorm.DB) refSkorGayaPekerjaan {
	r.refSkorGayaPekerjaanDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkorGayaPekerjaan) replaceDB(db *gorm.DB) refSkorGayaPekerjaan {
	r.refSkorGayaPekerjaanDo.ReplaceDB(db)
	return r
}

type refSkorGayaPekerjaanDo struct{ gen.DO }

func (r refSkorGayaPekerjaanDo) Debug() *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkorGayaPekerjaanDo) WithContext(ctx context.Context) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkorGayaPekerjaanDo) ReadDB() *refSkorGayaPekerjaanDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkorGayaPekerjaanDo) WriteDB() *refSkorGayaPekerjaanDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkorGayaPekerjaanDo) Session(config *gorm.Session) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkorGayaPekerjaanDo) Clauses(conds ...clause.Expression) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkorGayaPekerjaanDo) Returning(value interface{}, columns ...string) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkorGayaPekerjaanDo) Not(conds ...gen.Condition) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkorGayaPekerjaanDo) Or(conds ...gen.Condition) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkorGayaPekerjaanDo) Select(conds ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkorGayaPekerjaanDo) Where(conds ...gen.Condition) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkorGayaPekerjaanDo) Order(conds ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkorGayaPekerjaanDo) Distinct(cols ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkorGayaPekerjaanDo) Omit(cols ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkorGayaPekerjaanDo) Join(table schema.Tabler, on ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkorGayaPekerjaanDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkorGayaPekerjaanDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkorGayaPekerjaanDo) Group(cols ...field.Expr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkorGayaPekerjaanDo) Having(conds ...gen.Condition) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkorGayaPekerjaanDo) Limit(limit int) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkorGayaPekerjaanDo) Offset(offset int) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkorGayaPekerjaanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkorGayaPekerjaanDo) Unscoped() *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkorGayaPekerjaanDo) Create(values ...*entity.RefSkorGayaPekerjaan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkorGayaPekerjaanDo) CreateInBatches(values []*entity.RefSkorGayaPekerjaan, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkorGayaPekerjaanDo) Save(values ...*entity.RefSkorGayaPekerjaan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkorGayaPekerjaanDo) First() (*entity.RefSkorGayaPekerjaan, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorGayaPekerjaan), nil
	}
}

func (r refSkorGayaPekerjaanDo) Take() (*entity.RefSkorGayaPekerjaan, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorGayaPekerjaan), nil
	}
}

func (r refSkorGayaPekerjaanDo) Last() (*entity.RefSkorGayaPekerjaan, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorGayaPekerjaan), nil
	}
}

func (r refSkorGayaPekerjaanDo) Find() ([]*entity.RefSkorGayaPekerjaan, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkorGayaPekerjaan), err
}

func (r refSkorGayaPekerjaanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkorGayaPekerjaan, err error) {
	buf := make([]*entity.RefSkorGayaPekerjaan, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkorGayaPekerjaanDo) FindInBatches(result *[]*entity.RefSkorGayaPekerjaan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkorGayaPekerjaanDo) Attrs(attrs ...field.AssignExpr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkorGayaPekerjaanDo) Assign(attrs ...field.AssignExpr) *refSkorGayaPekerjaanDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkorGayaPekerjaanDo) Joins(fields ...field.RelationField) *refSkorGayaPekerjaanDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkorGayaPekerjaanDo) Preload(fields ...field.RelationField) *refSkorGayaPekerjaanDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkorGayaPekerjaanDo) FirstOrInit() (*entity.RefSkorGayaPekerjaan, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorGayaPekerjaan), nil
	}
}

func (r refSkorGayaPekerjaanDo) FirstOrCreate() (*entity.RefSkorGayaPekerjaan, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorGayaPekerjaan), nil
	}
}

func (r refSkorGayaPekerjaanDo) FindByPage(offset int, limit int) (result []*entity.RefSkorGayaPekerjaan, count int64, err error) {
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

func (r refSkorGayaPekerjaanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkorGayaPekerjaanDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkorGayaPekerjaanDo) Delete(models ...*entity.RefSkorGayaPekerjaan) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkorGayaPekerjaanDo) withDO(do gen.Dao) *refSkorGayaPekerjaanDo {
	r.DO = *do.(*gen.DO)
	return r
}
