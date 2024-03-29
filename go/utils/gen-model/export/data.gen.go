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

func newDatum(db *gorm.DB, opts ...gen.DOOption) datum {
	_datum := datum{}

	_datum.datumDo.UseDB(db, opts...)
	_datum.datumDo.UseModel(&entity.Datum{})

	tableName := _datum.datumDo.TableName()
	_datum.ALL = field.NewAsterisk(tableName)
	_datum.ID = field.NewInt32(tableName, "id")
	_datum.Nama = field.NewString(tableName, "nama")
	_datum.Keterangan = field.NewString(tableName, "keterangan")

	_datum.fillFieldMap()

	return _datum
}

type datum struct {
	datumDo datumDo

	ALL        field.Asterisk
	ID         field.Int32
	Nama       field.String
	Keterangan field.String

	fieldMap map[string]field.Expr
}

func (d datum) Table(newTableName string) *datum {
	d.datumDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d datum) As(alias string) *datum {
	d.datumDo.DO = *(d.datumDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *datum) updateTableName(table string) *datum {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.Nama = field.NewString(table, "nama")
	d.Keterangan = field.NewString(table, "keterangan")

	d.fillFieldMap()

	return d
}

func (d *datum) WithContext(ctx context.Context) *datumDo { return d.datumDo.WithContext(ctx) }

func (d datum) TableName() string { return d.datumDo.TableName() }

func (d datum) Alias() string { return d.datumDo.Alias() }

func (d datum) Columns(cols ...field.Expr) gen.Columns { return d.datumDo.Columns(cols...) }

func (d *datum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *datum) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["id"] = d.ID
	d.fieldMap["nama"] = d.Nama
	d.fieldMap["keterangan"] = d.Keterangan
}

func (d datum) clone(db *gorm.DB) datum {
	d.datumDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d datum) replaceDB(db *gorm.DB) datum {
	d.datumDo.ReplaceDB(db)
	return d
}

type datumDo struct{ gen.DO }

func (d datumDo) Debug() *datumDo {
	return d.withDO(d.DO.Debug())
}

func (d datumDo) WithContext(ctx context.Context) *datumDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d datumDo) ReadDB() *datumDo {
	return d.Clauses(dbresolver.Read)
}

func (d datumDo) WriteDB() *datumDo {
	return d.Clauses(dbresolver.Write)
}

func (d datumDo) Session(config *gorm.Session) *datumDo {
	return d.withDO(d.DO.Session(config))
}

func (d datumDo) Clauses(conds ...clause.Expression) *datumDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d datumDo) Returning(value interface{}, columns ...string) *datumDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d datumDo) Not(conds ...gen.Condition) *datumDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d datumDo) Or(conds ...gen.Condition) *datumDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d datumDo) Select(conds ...field.Expr) *datumDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d datumDo) Where(conds ...gen.Condition) *datumDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d datumDo) Order(conds ...field.Expr) *datumDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d datumDo) Distinct(cols ...field.Expr) *datumDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d datumDo) Omit(cols ...field.Expr) *datumDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d datumDo) Join(table schema.Tabler, on ...field.Expr) *datumDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d datumDo) LeftJoin(table schema.Tabler, on ...field.Expr) *datumDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d datumDo) RightJoin(table schema.Tabler, on ...field.Expr) *datumDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d datumDo) Group(cols ...field.Expr) *datumDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d datumDo) Having(conds ...gen.Condition) *datumDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d datumDo) Limit(limit int) *datumDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d datumDo) Offset(offset int) *datumDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d datumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *datumDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d datumDo) Unscoped() *datumDo {
	return d.withDO(d.DO.Unscoped())
}

func (d datumDo) Create(values ...*entity.Datum) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d datumDo) CreateInBatches(values []*entity.Datum, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d datumDo) Save(values ...*entity.Datum) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d datumDo) First() (*entity.Datum, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Datum), nil
	}
}

func (d datumDo) Take() (*entity.Datum, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Datum), nil
	}
}

func (d datumDo) Last() (*entity.Datum, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Datum), nil
	}
}

func (d datumDo) Find() ([]*entity.Datum, error) {
	result, err := d.DO.Find()
	return result.([]*entity.Datum), err
}

func (d datumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Datum, err error) {
	buf := make([]*entity.Datum, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d datumDo) FindInBatches(result *[]*entity.Datum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d datumDo) Attrs(attrs ...field.AssignExpr) *datumDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d datumDo) Assign(attrs ...field.AssignExpr) *datumDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d datumDo) Joins(fields ...field.RelationField) *datumDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d datumDo) Preload(fields ...field.RelationField) *datumDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d datumDo) FirstOrInit() (*entity.Datum, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Datum), nil
	}
}

func (d datumDo) FirstOrCreate() (*entity.Datum, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Datum), nil
	}
}

func (d datumDo) FindByPage(offset int, limit int) (result []*entity.Datum, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d datumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d datumDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d datumDo) Delete(models ...*entity.Datum) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *datumDo) withDO(do gen.Dao) *datumDo {
	d.DO = *do.(*gen.DO)
	return d
}
