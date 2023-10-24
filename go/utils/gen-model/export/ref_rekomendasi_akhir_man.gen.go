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

func newRefRekomendasiAkhirMan(db *gorm.DB, opts ...gen.DOOption) refRekomendasiAkhirMan {
	_refRekomendasiAkhirMan := refRekomendasiAkhirMan{}

	_refRekomendasiAkhirMan.refRekomendasiAkhirManDo.UseDB(db, opts...)
	_refRekomendasiAkhirMan.refRekomendasiAkhirManDo.UseModel(&entity.RefRekomendasiAkhirMan{})

	tableName := _refRekomendasiAkhirMan.refRekomendasiAkhirManDo.TableName()
	_refRekomendasiAkhirMan.ALL = field.NewAsterisk(tableName)
	_refRekomendasiAkhirMan.ID = field.NewInt32(tableName, "id")
	_refRekomendasiAkhirMan.RekomMinat = field.NewString(tableName, "rekom_minat")
	_refRekomendasiAkhirMan.RekomSikapPelajaran = field.NewString(tableName, "rekom_sikap_pelajaran")
	_refRekomendasiAkhirMan.RekomTmi = field.NewString(tableName, "rekom_tmi")
	_refRekomendasiAkhirMan.RekomAkhir = field.NewString(tableName, "rekom_akhir")

	_refRekomendasiAkhirMan.fillFieldMap()

	return _refRekomendasiAkhirMan
}

type refRekomendasiAkhirMan struct {
	refRekomendasiAkhirManDo refRekomendasiAkhirManDo

	ALL                 field.Asterisk
	ID                  field.Int32
	RekomMinat          field.String
	RekomSikapPelajaran field.String
	RekomTmi            field.String
	RekomAkhir          field.String

	fieldMap map[string]field.Expr
}

func (r refRekomendasiAkhirMan) Table(newTableName string) *refRekomendasiAkhirMan {
	r.refRekomendasiAkhirManDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refRekomendasiAkhirMan) As(alias string) *refRekomendasiAkhirMan {
	r.refRekomendasiAkhirManDo.DO = *(r.refRekomendasiAkhirManDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refRekomendasiAkhirMan) updateTableName(table string) *refRekomendasiAkhirMan {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.RekomMinat = field.NewString(table, "rekom_minat")
	r.RekomSikapPelajaran = field.NewString(table, "rekom_sikap_pelajaran")
	r.RekomTmi = field.NewString(table, "rekom_tmi")
	r.RekomAkhir = field.NewString(table, "rekom_akhir")

	r.fillFieldMap()

	return r
}

func (r *refRekomendasiAkhirMan) WithContext(ctx context.Context) *refRekomendasiAkhirManDo {
	return r.refRekomendasiAkhirManDo.WithContext(ctx)
}

func (r refRekomendasiAkhirMan) TableName() string { return r.refRekomendasiAkhirManDo.TableName() }

func (r refRekomendasiAkhirMan) Alias() string { return r.refRekomendasiAkhirManDo.Alias() }

func (r refRekomendasiAkhirMan) Columns(cols ...field.Expr) gen.Columns {
	return r.refRekomendasiAkhirManDo.Columns(cols...)
}

func (r *refRekomendasiAkhirMan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refRekomendasiAkhirMan) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["rekom_minat"] = r.RekomMinat
	r.fieldMap["rekom_sikap_pelajaran"] = r.RekomSikapPelajaran
	r.fieldMap["rekom_tmi"] = r.RekomTmi
	r.fieldMap["rekom_akhir"] = r.RekomAkhir
}

func (r refRekomendasiAkhirMan) clone(db *gorm.DB) refRekomendasiAkhirMan {
	r.refRekomendasiAkhirManDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refRekomendasiAkhirMan) replaceDB(db *gorm.DB) refRekomendasiAkhirMan {
	r.refRekomendasiAkhirManDo.ReplaceDB(db)
	return r
}

type refRekomendasiAkhirManDo struct{ gen.DO }

func (r refRekomendasiAkhirManDo) Debug() *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Debug())
}

func (r refRekomendasiAkhirManDo) WithContext(ctx context.Context) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refRekomendasiAkhirManDo) ReadDB() *refRekomendasiAkhirManDo {
	return r.Clauses(dbresolver.Read)
}

func (r refRekomendasiAkhirManDo) WriteDB() *refRekomendasiAkhirManDo {
	return r.Clauses(dbresolver.Write)
}

func (r refRekomendasiAkhirManDo) Session(config *gorm.Session) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Session(config))
}

func (r refRekomendasiAkhirManDo) Clauses(conds ...clause.Expression) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refRekomendasiAkhirManDo) Returning(value interface{}, columns ...string) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refRekomendasiAkhirManDo) Not(conds ...gen.Condition) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refRekomendasiAkhirManDo) Or(conds ...gen.Condition) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refRekomendasiAkhirManDo) Select(conds ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refRekomendasiAkhirManDo) Where(conds ...gen.Condition) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refRekomendasiAkhirManDo) Order(conds ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refRekomendasiAkhirManDo) Distinct(cols ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refRekomendasiAkhirManDo) Omit(cols ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refRekomendasiAkhirManDo) Join(table schema.Tabler, on ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refRekomendasiAkhirManDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refRekomendasiAkhirManDo) RightJoin(table schema.Tabler, on ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refRekomendasiAkhirManDo) Group(cols ...field.Expr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refRekomendasiAkhirManDo) Having(conds ...gen.Condition) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refRekomendasiAkhirManDo) Limit(limit int) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refRekomendasiAkhirManDo) Offset(offset int) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refRekomendasiAkhirManDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refRekomendasiAkhirManDo) Unscoped() *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refRekomendasiAkhirManDo) Create(values ...*entity.RefRekomendasiAkhirMan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refRekomendasiAkhirManDo) CreateInBatches(values []*entity.RefRekomendasiAkhirMan, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refRekomendasiAkhirManDo) Save(values ...*entity.RefRekomendasiAkhirMan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refRekomendasiAkhirManDo) First() (*entity.RefRekomendasiAkhirMan, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiAkhirMan), nil
	}
}

func (r refRekomendasiAkhirManDo) Take() (*entity.RefRekomendasiAkhirMan, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiAkhirMan), nil
	}
}

func (r refRekomendasiAkhirManDo) Last() (*entity.RefRekomendasiAkhirMan, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiAkhirMan), nil
	}
}

func (r refRekomendasiAkhirManDo) Find() ([]*entity.RefRekomendasiAkhirMan, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefRekomendasiAkhirMan), err
}

func (r refRekomendasiAkhirManDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefRekomendasiAkhirMan, err error) {
	buf := make([]*entity.RefRekomendasiAkhirMan, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refRekomendasiAkhirManDo) FindInBatches(result *[]*entity.RefRekomendasiAkhirMan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refRekomendasiAkhirManDo) Attrs(attrs ...field.AssignExpr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refRekomendasiAkhirManDo) Assign(attrs ...field.AssignExpr) *refRekomendasiAkhirManDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refRekomendasiAkhirManDo) Joins(fields ...field.RelationField) *refRekomendasiAkhirManDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refRekomendasiAkhirManDo) Preload(fields ...field.RelationField) *refRekomendasiAkhirManDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refRekomendasiAkhirManDo) FirstOrInit() (*entity.RefRekomendasiAkhirMan, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiAkhirMan), nil
	}
}

func (r refRekomendasiAkhirManDo) FirstOrCreate() (*entity.RefRekomendasiAkhirMan, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefRekomendasiAkhirMan), nil
	}
}

func (r refRekomendasiAkhirManDo) FindByPage(offset int, limit int) (result []*entity.RefRekomendasiAkhirMan, count int64, err error) {
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

func (r refRekomendasiAkhirManDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refRekomendasiAkhirManDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refRekomendasiAkhirManDo) Delete(models ...*entity.RefRekomendasiAkhirMan) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refRekomendasiAkhirManDo) withDO(do gen.Dao) *refRekomendasiAkhirManDo {
	r.DO = *do.(*gen.DO)
	return r
}