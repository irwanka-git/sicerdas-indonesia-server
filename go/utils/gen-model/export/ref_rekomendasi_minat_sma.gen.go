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

func newRefRekomendasiMinatSma(db *gorm.DB, opts ...gen.DOOption) refRekomendasiMinatSma {
	_refRekomendasiMinatSma := refRekomendasiMinatSma{}

	_refRekomendasiMinatSma.refRekomendasiMinatSmaDo.UseDB(db, opts...)
	_refRekomendasiMinatSma.refRekomendasiMinatSmaDo.UseModel(&entity.RefRekomendasiMinatSma{})

	tableName := _refRekomendasiMinatSma.refRekomendasiMinatSmaDo.TableName()
	_refRekomendasiMinatSma.ALL = field.NewAsterisk(tableName)
	_refRekomendasiMinatSma.ID = field.NewInt32(tableName, "id")
	_refRekomendasiMinatSma.IlmuAlam = field.NewInt32(tableName, "ilmu_alam")
	_refRekomendasiMinatSma.IlmuSosial = field.NewInt32(tableName, "ilmu_sosial")
	_refRekomendasiMinatSma.Perbedaan = field.NewInt32(tableName, "perbedaan")
	_refRekomendasiMinatSma.Rekomendasi = field.NewString(tableName, "rekomendasi")

	_refRekomendasiMinatSma.fillFieldMap()

	return _refRekomendasiMinatSma
}

type refRekomendasiMinatSma struct {
	refRekomendasiMinatSmaDo refRekomendasiMinatSmaDo

	ALL         field.Asterisk
	ID          field.Int32
	IlmuAlam    field.Int32
	IlmuSosial  field.Int32
	Perbedaan   field.Int32
	Rekomendasi field.String

	fieldMap map[string]field.Expr
}

func (r refRekomendasiMinatSma) Table(newTableName string) *refRekomendasiMinatSma {
	r.refRekomendasiMinatSmaDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refRekomendasiMinatSma) As(alias string) *refRekomendasiMinatSma {
	r.refRekomendasiMinatSmaDo.DO = *(r.refRekomendasiMinatSmaDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refRekomendasiMinatSma) updateTableName(table string) *refRekomendasiMinatSma {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.IlmuAlam = field.NewInt32(table, "ilmu_alam")
	r.IlmuSosial = field.NewInt32(table, "ilmu_sosial")
	r.Perbedaan = field.NewInt32(table, "perbedaan")
	r.Rekomendasi = field.NewString(table, "rekomendasi")

	r.fillFieldMap()

	return r
}

func (r *refRekomendasiMinatSma) WithContext(ctx context.Context) *refRekomendasiMinatSmaDo {
	return r.refRekomendasiMinatSmaDo.WithContext(ctx)
}

func (r refRekomendasiMinatSma) TableName() string { return r.refRekomendasiMinatSmaDo.TableName() }

func (r refRekomendasiMinatSma) Alias() string { return r.refRekomendasiMinatSmaDo.Alias() }

func (r refRekomendasiMinatSma) Columns(cols ...field.Expr) gen.Columns {
	return r.refRekomendasiMinatSmaDo.Columns(cols...)
}

func (r *refRekomendasiMinatSma) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refRekomendasiMinatSma) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["ilmu_alam"] = r.IlmuAlam
	r.fieldMap["ilmu_sosial"] = r.IlmuSosial
	r.fieldMap["perbedaan"] = r.Perbedaan
	r.fieldMap["rekomendasi"] = r.Rekomendasi
}

func (r refRekomendasiMinatSma) clone(db *gorm.DB) refRekomendasiMinatSma {
	r.refRekomendasiMinatSmaDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refRekomendasiMinatSma) replaceDB(db *gorm.DB) refRekomendasiMinatSma {
	r.refRekomendasiMinatSmaDo.ReplaceDB(db)
	return r
}

type refRekomendasiMinatSmaDo struct{ gen.DO }

func (r refRekomendasiMinatSmaDo) Debug() *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Debug())
}

func (r refRekomendasiMinatSmaDo) WithContext(ctx context.Context) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refRekomendasiMinatSmaDo) ReadDB() *refRekomendasiMinatSmaDo {
	return r.Clauses(dbresolver.Read)
}

func (r refRekomendasiMinatSmaDo) WriteDB() *refRekomendasiMinatSmaDo {
	return r.Clauses(dbresolver.Write)
}

func (r refRekomendasiMinatSmaDo) Session(config *gorm.Session) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Session(config))
}

func (r refRekomendasiMinatSmaDo) Clauses(conds ...clause.Expression) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refRekomendasiMinatSmaDo) Returning(value interface{}, columns ...string) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refRekomendasiMinatSmaDo) Not(conds ...gen.Condition) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refRekomendasiMinatSmaDo) Or(conds ...gen.Condition) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refRekomendasiMinatSmaDo) Select(conds ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refRekomendasiMinatSmaDo) Where(conds ...gen.Condition) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refRekomendasiMinatSmaDo) Order(conds ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refRekomendasiMinatSmaDo) Distinct(cols ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refRekomendasiMinatSmaDo) Omit(cols ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refRekomendasiMinatSmaDo) Join(table schema.Tabler, on ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refRekomendasiMinatSmaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refRekomendasiMinatSmaDo) RightJoin(table schema.Tabler, on ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refRekomendasiMinatSmaDo) Group(cols ...field.Expr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refRekomendasiMinatSmaDo) Having(conds ...gen.Condition) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refRekomendasiMinatSmaDo) Limit(limit int) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refRekomendasiMinatSmaDo) Offset(offset int) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refRekomendasiMinatSmaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refRekomendasiMinatSmaDo) Unscoped() *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refRekomendasiMinatSmaDo) Create(values ...*entity.RefRekomendasiMinatSma) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refRekomendasiMinatSmaDo) CreateInBatches(values []*entity.RefRekomendasiMinatSma, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refRekomendasiMinatSmaDo) Save(values ...*entity.RefRekomendasiMinatSma) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refRekomendasiMinatSmaDo) First() (*entity.RefRekomendasiMinatSma, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiMinatSma), nil
	}
}

func (r refRekomendasiMinatSmaDo) Take() (*entity.RefRekomendasiMinatSma, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiMinatSma), nil
	}
}

func (r refRekomendasiMinatSmaDo) Last() (*entity.RefRekomendasiMinatSma, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiMinatSma), nil
	}
}

func (r refRekomendasiMinatSmaDo) Find() ([]*entity.RefRekomendasiMinatSma, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefRekomendasiMinatSma), err
}

func (r refRekomendasiMinatSmaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefRekomendasiMinatSma, err error) {
	buf := make([]*entity.RefRekomendasiMinatSma, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refRekomendasiMinatSmaDo) FindInBatches(result *[]*entity.RefRekomendasiMinatSma, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refRekomendasiMinatSmaDo) Attrs(attrs ...field.AssignExpr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refRekomendasiMinatSmaDo) Assign(attrs ...field.AssignExpr) *refRekomendasiMinatSmaDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refRekomendasiMinatSmaDo) Joins(fields ...field.RelationField) *refRekomendasiMinatSmaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refRekomendasiMinatSmaDo) Preload(fields ...field.RelationField) *refRekomendasiMinatSmaDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refRekomendasiMinatSmaDo) FirstOrInit() (*entity.RefRekomendasiMinatSma, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiMinatSma), nil
	}
}

func (r refRekomendasiMinatSmaDo) FirstOrCreate() (*entity.RefRekomendasiMinatSma, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiMinatSma), nil
	}
}

func (r refRekomendasiMinatSmaDo) FindByPage(offset int, limit int) (result []*entity.RefRekomendasiMinatSma, count int64, err error) {
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

func (r refRekomendasiMinatSmaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refRekomendasiMinatSmaDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refRekomendasiMinatSmaDo) Delete(models ...*entity.RefRekomendasiMinatSma) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refRekomendasiMinatSmaDo) withDO(do gen.Dao) *refRekomendasiMinatSmaDo {
	r.DO = *do.(*gen.DO)
	return r
}
