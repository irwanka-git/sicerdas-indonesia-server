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

func newRefKlasifikasiPribadiManajerial(db *gorm.DB, opts ...gen.DOOption) refKlasifikasiPribadiManajerial {
	_refKlasifikasiPribadiManajerial := refKlasifikasiPribadiManajerial{}

	_refKlasifikasiPribadiManajerial.refKlasifikasiPribadiManajerialDo.UseDB(db, opts...)
	_refKlasifikasiPribadiManajerial.refKlasifikasiPribadiManajerialDo.UseModel(&entity.RefKlasifikasiPribadiManajerial{})

	tableName := _refKlasifikasiPribadiManajerial.refKlasifikasiPribadiManajerialDo.TableName()
	_refKlasifikasiPribadiManajerial.ALL = field.NewAsterisk(tableName)
	_refKlasifikasiPribadiManajerial.Skor = field.NewInt32(tableName, "skor")
	_refKlasifikasiPribadiManajerial.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_refKlasifikasiPribadiManajerial.fillFieldMap()

	return _refKlasifikasiPribadiManajerial
}

type refKlasifikasiPribadiManajerial struct {
	refKlasifikasiPribadiManajerialDo refKlasifikasiPribadiManajerialDo

	ALL         field.Asterisk
	Skor        field.Int32
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refKlasifikasiPribadiManajerial) Table(newTableName string) *refKlasifikasiPribadiManajerial {
	r.refKlasifikasiPribadiManajerialDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refKlasifikasiPribadiManajerial) As(alias string) *refKlasifikasiPribadiManajerial {
	r.refKlasifikasiPribadiManajerialDo.DO = *(r.refKlasifikasiPribadiManajerialDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refKlasifikasiPribadiManajerial) updateTableName(table string) *refKlasifikasiPribadiManajerial {
	r.ALL = field.NewAsterisk(table)
	r.Skor = field.NewInt32(table, "skor")
	r.Klasifikasi = field.NewString(table, "klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refKlasifikasiPribadiManajerial) WithContext(ctx context.Context) *refKlasifikasiPribadiManajerialDo {
	return r.refKlasifikasiPribadiManajerialDo.WithContext(ctx)
}

func (r refKlasifikasiPribadiManajerial) TableName() string {
	return r.refKlasifikasiPribadiManajerialDo.TableName()
}

func (r refKlasifikasiPribadiManajerial) Alias() string {
	return r.refKlasifikasiPribadiManajerialDo.Alias()
}

func (r refKlasifikasiPribadiManajerial) Columns(cols ...field.Expr) gen.Columns {
	return r.refKlasifikasiPribadiManajerialDo.Columns(cols...)
}

func (r *refKlasifikasiPribadiManajerial) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refKlasifikasiPribadiManajerial) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["klasifikasi"] = r.Klasifikasi
}

func (r refKlasifikasiPribadiManajerial) clone(db *gorm.DB) refKlasifikasiPribadiManajerial {
	r.refKlasifikasiPribadiManajerialDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refKlasifikasiPribadiManajerial) replaceDB(db *gorm.DB) refKlasifikasiPribadiManajerial {
	r.refKlasifikasiPribadiManajerialDo.ReplaceDB(db)
	return r
}

type refKlasifikasiPribadiManajerialDo struct{ gen.DO }

func (r refKlasifikasiPribadiManajerialDo) Debug() *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Debug())
}

func (r refKlasifikasiPribadiManajerialDo) WithContext(ctx context.Context) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refKlasifikasiPribadiManajerialDo) ReadDB() *refKlasifikasiPribadiManajerialDo {
	return r.Clauses(dbresolver.Read)
}

func (r refKlasifikasiPribadiManajerialDo) WriteDB() *refKlasifikasiPribadiManajerialDo {
	return r.Clauses(dbresolver.Write)
}

func (r refKlasifikasiPribadiManajerialDo) Session(config *gorm.Session) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Session(config))
}

func (r refKlasifikasiPribadiManajerialDo) Clauses(conds ...clause.Expression) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Returning(value interface{}, columns ...string) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refKlasifikasiPribadiManajerialDo) Not(conds ...gen.Condition) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Or(conds ...gen.Condition) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Select(conds ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Where(conds ...gen.Condition) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Order(conds ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Distinct(cols ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refKlasifikasiPribadiManajerialDo) Omit(cols ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refKlasifikasiPribadiManajerialDo) Join(table schema.Tabler, on ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refKlasifikasiPribadiManajerialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refKlasifikasiPribadiManajerialDo) RightJoin(table schema.Tabler, on ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refKlasifikasiPribadiManajerialDo) Group(cols ...field.Expr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refKlasifikasiPribadiManajerialDo) Having(conds ...gen.Condition) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refKlasifikasiPribadiManajerialDo) Limit(limit int) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refKlasifikasiPribadiManajerialDo) Offset(offset int) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refKlasifikasiPribadiManajerialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refKlasifikasiPribadiManajerialDo) Unscoped() *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refKlasifikasiPribadiManajerialDo) Create(values ...*entity.RefKlasifikasiPribadiManajerial) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refKlasifikasiPribadiManajerialDo) CreateInBatches(values []*entity.RefKlasifikasiPribadiManajerial, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refKlasifikasiPribadiManajerialDo) Save(values ...*entity.RefKlasifikasiPribadiManajerial) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refKlasifikasiPribadiManajerialDo) First() (*entity.RefKlasifikasiPribadiManajerial, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKlasifikasiPribadiManajerial), nil
	}
}

func (r refKlasifikasiPribadiManajerialDo) Take() (*entity.RefKlasifikasiPribadiManajerial, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKlasifikasiPribadiManajerial), nil
	}
}

func (r refKlasifikasiPribadiManajerialDo) Last() (*entity.RefKlasifikasiPribadiManajerial, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKlasifikasiPribadiManajerial), nil
	}
}

func (r refKlasifikasiPribadiManajerialDo) Find() ([]*entity.RefKlasifikasiPribadiManajerial, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefKlasifikasiPribadiManajerial), err
}

func (r refKlasifikasiPribadiManajerialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefKlasifikasiPribadiManajerial, err error) {
	buf := make([]*entity.RefKlasifikasiPribadiManajerial, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refKlasifikasiPribadiManajerialDo) FindInBatches(result *[]*entity.RefKlasifikasiPribadiManajerial, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refKlasifikasiPribadiManajerialDo) Attrs(attrs ...field.AssignExpr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refKlasifikasiPribadiManajerialDo) Assign(attrs ...field.AssignExpr) *refKlasifikasiPribadiManajerialDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refKlasifikasiPribadiManajerialDo) Joins(fields ...field.RelationField) *refKlasifikasiPribadiManajerialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refKlasifikasiPribadiManajerialDo) Preload(fields ...field.RelationField) *refKlasifikasiPribadiManajerialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refKlasifikasiPribadiManajerialDo) FirstOrInit() (*entity.RefKlasifikasiPribadiManajerial, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKlasifikasiPribadiManajerial), nil
	}
}

func (r refKlasifikasiPribadiManajerialDo) FirstOrCreate() (*entity.RefKlasifikasiPribadiManajerial, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKlasifikasiPribadiManajerial), nil
	}
}

func (r refKlasifikasiPribadiManajerialDo) FindByPage(offset int, limit int) (result []*entity.RefKlasifikasiPribadiManajerial, count int64, err error) {
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

func (r refKlasifikasiPribadiManajerialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refKlasifikasiPribadiManajerialDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refKlasifikasiPribadiManajerialDo) Delete(models ...*entity.RefKlasifikasiPribadiManajerial) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refKlasifikasiPribadiManajerialDo) withDO(do gen.Dao) *refKlasifikasiPribadiManajerialDo {
	r.DO = *do.(*gen.DO)
	return r
}
