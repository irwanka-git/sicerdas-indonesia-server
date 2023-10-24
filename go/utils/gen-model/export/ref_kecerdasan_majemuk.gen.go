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

func newRefKecerdasanMajemuk(db *gorm.DB, opts ...gen.DOOption) refKecerdasanMajemuk {
	_refKecerdasanMajemuk := refKecerdasanMajemuk{}

	_refKecerdasanMajemuk.refKecerdasanMajemukDo.UseDB(db, opts...)
	_refKecerdasanMajemuk.refKecerdasanMajemukDo.UseModel(&entity.RefKecerdasanMajemuk{})

	tableName := _refKecerdasanMajemuk.refKecerdasanMajemukDo.TableName()
	_refKecerdasanMajemuk.ALL = field.NewAsterisk(tableName)
	_refKecerdasanMajemuk.No = field.NewString(tableName, "no")
	_refKecerdasanMajemuk.NamaKecerdasan = field.NewString(tableName, "nama_kecerdasan")
	_refKecerdasanMajemuk.NamaKecil = field.NewString(tableName, "nama_kecil")
	_refKecerdasanMajemuk.Icon = field.NewString(tableName, "icon")

	_refKecerdasanMajemuk.fillFieldMap()

	return _refKecerdasanMajemuk
}

type refKecerdasanMajemuk struct {
	refKecerdasanMajemukDo refKecerdasanMajemukDo

	ALL            field.Asterisk
	No             field.String
	NamaKecerdasan field.String
	NamaKecil      field.String
	Icon           field.String

	fieldMap map[string]field.Expr
}

func (r refKecerdasanMajemuk) Table(newTableName string) *refKecerdasanMajemuk {
	r.refKecerdasanMajemukDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refKecerdasanMajemuk) As(alias string) *refKecerdasanMajemuk {
	r.refKecerdasanMajemukDo.DO = *(r.refKecerdasanMajemukDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refKecerdasanMajemuk) updateTableName(table string) *refKecerdasanMajemuk {
	r.ALL = field.NewAsterisk(table)
	r.No = field.NewString(table, "no")
	r.NamaKecerdasan = field.NewString(table, "nama_kecerdasan")
	r.NamaKecil = field.NewString(table, "nama_kecil")
	r.Icon = field.NewString(table, "icon")

	r.fillFieldMap()

	return r
}

func (r *refKecerdasanMajemuk) WithContext(ctx context.Context) *refKecerdasanMajemukDo {
	return r.refKecerdasanMajemukDo.WithContext(ctx)
}

func (r refKecerdasanMajemuk) TableName() string { return r.refKecerdasanMajemukDo.TableName() }

func (r refKecerdasanMajemuk) Alias() string { return r.refKecerdasanMajemukDo.Alias() }

func (r refKecerdasanMajemuk) Columns(cols ...field.Expr) gen.Columns {
	return r.refKecerdasanMajemukDo.Columns(cols...)
}

func (r *refKecerdasanMajemuk) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refKecerdasanMajemuk) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 4)
	r.fieldMap["no"] = r.No
	r.fieldMap["nama_kecerdasan"] = r.NamaKecerdasan
	r.fieldMap["nama_kecil"] = r.NamaKecil
	r.fieldMap["icon"] = r.Icon
}

func (r refKecerdasanMajemuk) clone(db *gorm.DB) refKecerdasanMajemuk {
	r.refKecerdasanMajemukDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refKecerdasanMajemuk) replaceDB(db *gorm.DB) refKecerdasanMajemuk {
	r.refKecerdasanMajemukDo.ReplaceDB(db)
	return r
}

type refKecerdasanMajemukDo struct{ gen.DO }

func (r refKecerdasanMajemukDo) Debug() *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Debug())
}

func (r refKecerdasanMajemukDo) WithContext(ctx context.Context) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refKecerdasanMajemukDo) ReadDB() *refKecerdasanMajemukDo {
	return r.Clauses(dbresolver.Read)
}

func (r refKecerdasanMajemukDo) WriteDB() *refKecerdasanMajemukDo {
	return r.Clauses(dbresolver.Write)
}

func (r refKecerdasanMajemukDo) Session(config *gorm.Session) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Session(config))
}

func (r refKecerdasanMajemukDo) Clauses(conds ...clause.Expression) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refKecerdasanMajemukDo) Returning(value interface{}, columns ...string) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refKecerdasanMajemukDo) Not(conds ...gen.Condition) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refKecerdasanMajemukDo) Or(conds ...gen.Condition) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refKecerdasanMajemukDo) Select(conds ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refKecerdasanMajemukDo) Where(conds ...gen.Condition) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refKecerdasanMajemukDo) Order(conds ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refKecerdasanMajemukDo) Distinct(cols ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refKecerdasanMajemukDo) Omit(cols ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refKecerdasanMajemukDo) Join(table schema.Tabler, on ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refKecerdasanMajemukDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refKecerdasanMajemukDo) RightJoin(table schema.Tabler, on ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refKecerdasanMajemukDo) Group(cols ...field.Expr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refKecerdasanMajemukDo) Having(conds ...gen.Condition) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refKecerdasanMajemukDo) Limit(limit int) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refKecerdasanMajemukDo) Offset(offset int) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refKecerdasanMajemukDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refKecerdasanMajemukDo) Unscoped() *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refKecerdasanMajemukDo) Create(values ...*entity.RefKecerdasanMajemuk) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refKecerdasanMajemukDo) CreateInBatches(values []*entity.RefKecerdasanMajemuk, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refKecerdasanMajemukDo) Save(values ...*entity.RefKecerdasanMajemuk) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refKecerdasanMajemukDo) First() (*entity.RefKecerdasanMajemuk, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKecerdasanMajemuk), nil
	}
}

func (r refKecerdasanMajemukDo) Take() (*entity.RefKecerdasanMajemuk, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKecerdasanMajemuk), nil
	}
}

func (r refKecerdasanMajemukDo) Last() (*entity.RefKecerdasanMajemuk, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKecerdasanMajemuk), nil
	}
}

func (r refKecerdasanMajemukDo) Find() ([]*entity.RefKecerdasanMajemuk, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefKecerdasanMajemuk), err
}

func (r refKecerdasanMajemukDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefKecerdasanMajemuk, err error) {
	buf := make([]*entity.RefKecerdasanMajemuk, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refKecerdasanMajemukDo) FindInBatches(result *[]*entity.RefKecerdasanMajemuk, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refKecerdasanMajemukDo) Attrs(attrs ...field.AssignExpr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refKecerdasanMajemukDo) Assign(attrs ...field.AssignExpr) *refKecerdasanMajemukDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refKecerdasanMajemukDo) Joins(fields ...field.RelationField) *refKecerdasanMajemukDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refKecerdasanMajemukDo) Preload(fields ...field.RelationField) *refKecerdasanMajemukDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refKecerdasanMajemukDo) FirstOrInit() (*entity.RefKecerdasanMajemuk, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKecerdasanMajemuk), nil
	}
}

func (r refKecerdasanMajemukDo) FirstOrCreate() (*entity.RefKecerdasanMajemuk, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKecerdasanMajemuk), nil
	}
}

func (r refKecerdasanMajemukDo) FindByPage(offset int, limit int) (result []*entity.RefKecerdasanMajemuk, count int64, err error) {
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

func (r refKecerdasanMajemukDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refKecerdasanMajemukDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refKecerdasanMajemukDo) Delete(models ...*entity.RefKecerdasanMajemuk) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refKecerdasanMajemukDo) withDO(do gen.Dao) *refKecerdasanMajemukDo {
	r.DO = *do.(*gen.DO)
	return r
}