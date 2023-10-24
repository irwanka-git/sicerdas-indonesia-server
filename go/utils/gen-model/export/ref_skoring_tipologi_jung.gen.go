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

func newRefSkoringTipologiJung(db *gorm.DB, opts ...gen.DOOption) refSkoringTipologiJung {
	_refSkoringTipologiJung := refSkoringTipologiJung{}

	_refSkoringTipologiJung.refSkoringTipologiJungDo.UseDB(db, opts...)
	_refSkoringTipologiJung.refSkoringTipologiJungDo.UseModel(&entity.RefSkoringTipologiJung{})

	tableName := _refSkoringTipologiJung.refSkoringTipologiJungDo.TableName()
	_refSkoringTipologiJung.ALL = field.NewAsterisk(tableName)
	_refSkoringTipologiJung.ID = field.NewInt32(tableName, "id")
	_refSkoringTipologiJung.Kolom = field.NewString(tableName, "kolom")
	_refSkoringTipologiJung.SkorA = field.NewInt32(tableName, "skor_a")
	_refSkoringTipologiJung.KodeKlasifikasi = field.NewString(tableName, "kode_klasifikasi")

	_refSkoringTipologiJung.fillFieldMap()

	return _refSkoringTipologiJung
}

type refSkoringTipologiJung struct {
	refSkoringTipologiJungDo refSkoringTipologiJungDo

	ALL             field.Asterisk
	ID              field.Int32
	Kolom           field.String
	SkorA           field.Int32
	KodeKlasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refSkoringTipologiJung) Table(newTableName string) *refSkoringTipologiJung {
	r.refSkoringTipologiJungDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkoringTipologiJung) As(alias string) *refSkoringTipologiJung {
	r.refSkoringTipologiJungDo.DO = *(r.refSkoringTipologiJungDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkoringTipologiJung) updateTableName(table string) *refSkoringTipologiJung {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Kolom = field.NewString(table, "kolom")
	r.SkorA = field.NewInt32(table, "skor_a")
	r.KodeKlasifikasi = field.NewString(table, "kode_klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refSkoringTipologiJung) WithContext(ctx context.Context) *refSkoringTipologiJungDo {
	return r.refSkoringTipologiJungDo.WithContext(ctx)
}

func (r refSkoringTipologiJung) TableName() string { return r.refSkoringTipologiJungDo.TableName() }

func (r refSkoringTipologiJung) Alias() string { return r.refSkoringTipologiJungDo.Alias() }

func (r refSkoringTipologiJung) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkoringTipologiJungDo.Columns(cols...)
}

func (r *refSkoringTipologiJung) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkoringTipologiJung) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 4)
	r.fieldMap["id"] = r.ID
	r.fieldMap["kolom"] = r.Kolom
	r.fieldMap["skor_a"] = r.SkorA
	r.fieldMap["kode_klasifikasi"] = r.KodeKlasifikasi
}

func (r refSkoringTipologiJung) clone(db *gorm.DB) refSkoringTipologiJung {
	r.refSkoringTipologiJungDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkoringTipologiJung) replaceDB(db *gorm.DB) refSkoringTipologiJung {
	r.refSkoringTipologiJungDo.ReplaceDB(db)
	return r
}

type refSkoringTipologiJungDo struct{ gen.DO }

func (r refSkoringTipologiJungDo) Debug() *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkoringTipologiJungDo) WithContext(ctx context.Context) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkoringTipologiJungDo) ReadDB() *refSkoringTipologiJungDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkoringTipologiJungDo) WriteDB() *refSkoringTipologiJungDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkoringTipologiJungDo) Session(config *gorm.Session) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkoringTipologiJungDo) Clauses(conds ...clause.Expression) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkoringTipologiJungDo) Returning(value interface{}, columns ...string) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkoringTipologiJungDo) Not(conds ...gen.Condition) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkoringTipologiJungDo) Or(conds ...gen.Condition) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkoringTipologiJungDo) Select(conds ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkoringTipologiJungDo) Where(conds ...gen.Condition) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkoringTipologiJungDo) Order(conds ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkoringTipologiJungDo) Distinct(cols ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkoringTipologiJungDo) Omit(cols ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkoringTipologiJungDo) Join(table schema.Tabler, on ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkoringTipologiJungDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkoringTipologiJungDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkoringTipologiJungDo) Group(cols ...field.Expr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkoringTipologiJungDo) Having(conds ...gen.Condition) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkoringTipologiJungDo) Limit(limit int) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkoringTipologiJungDo) Offset(offset int) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkoringTipologiJungDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkoringTipologiJungDo) Unscoped() *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkoringTipologiJungDo) Create(values ...*entity.RefSkoringTipologiJung) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkoringTipologiJungDo) CreateInBatches(values []*entity.RefSkoringTipologiJung, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkoringTipologiJungDo) Save(values ...*entity.RefSkoringTipologiJung) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkoringTipologiJungDo) First() (*entity.RefSkoringTipologiJung, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringTipologiJung), nil
	}
}

func (r refSkoringTipologiJungDo) Take() (*entity.RefSkoringTipologiJung, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringTipologiJung), nil
	}
}

func (r refSkoringTipologiJungDo) Last() (*entity.RefSkoringTipologiJung, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringTipologiJung), nil
	}
}

func (r refSkoringTipologiJungDo) Find() ([]*entity.RefSkoringTipologiJung, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkoringTipologiJung), err
}

func (r refSkoringTipologiJungDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkoringTipologiJung, err error) {
	buf := make([]*entity.RefSkoringTipologiJung, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkoringTipologiJungDo) FindInBatches(result *[]*entity.RefSkoringTipologiJung, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkoringTipologiJungDo) Attrs(attrs ...field.AssignExpr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkoringTipologiJungDo) Assign(attrs ...field.AssignExpr) *refSkoringTipologiJungDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkoringTipologiJungDo) Joins(fields ...field.RelationField) *refSkoringTipologiJungDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkoringTipologiJungDo) Preload(fields ...field.RelationField) *refSkoringTipologiJungDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkoringTipologiJungDo) FirstOrInit() (*entity.RefSkoringTipologiJung, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringTipologiJung), nil
	}
}

func (r refSkoringTipologiJungDo) FirstOrCreate() (*entity.RefSkoringTipologiJung, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringTipologiJung), nil
	}
}

func (r refSkoringTipologiJungDo) FindByPage(offset int, limit int) (result []*entity.RefSkoringTipologiJung, count int64, err error) {
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

func (r refSkoringTipologiJungDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkoringTipologiJungDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkoringTipologiJungDo) Delete(models ...*entity.RefSkoringTipologiJung) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkoringTipologiJungDo) withDO(do gen.Dao) *refSkoringTipologiJungDo {
	r.DO = *do.(*gen.DO)
	return r
}