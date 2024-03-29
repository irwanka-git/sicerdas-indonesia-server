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

func newRefSkorKarakterPribadi(db *gorm.DB, opts ...gen.DOOption) refSkorKarakterPribadi {
	_refSkorKarakterPribadi := refSkorKarakterPribadi{}

	_refSkorKarakterPribadi.refSkorKarakterPribadiDo.UseDB(db, opts...)
	_refSkorKarakterPribadi.refSkorKarakterPribadiDo.UseModel(&entity.RefSkorKarakterPribadi{})

	tableName := _refSkorKarakterPribadi.refSkorKarakterPribadiDo.TableName()
	_refSkorKarakterPribadi.ALL = field.NewAsterisk(tableName)
	_refSkorKarakterPribadi.ID = field.NewInt32(tableName, "id")
	_refSkorKarakterPribadi.Respon = field.NewString(tableName, "respon")
	_refSkorKarakterPribadi.Skor = field.NewInt32(tableName, "skor")
	_refSkorKarakterPribadi.Jawaban = field.NewString(tableName, "jawaban")

	_refSkorKarakterPribadi.fillFieldMap()

	return _refSkorKarakterPribadi
}

type refSkorKarakterPribadi struct {
	refSkorKarakterPribadiDo refSkorKarakterPribadiDo

	ALL     field.Asterisk
	ID      field.Int32
	Respon  field.String
	Skor    field.Int32
	Jawaban field.String

	fieldMap map[string]field.Expr
}

func (r refSkorKarakterPribadi) Table(newTableName string) *refSkorKarakterPribadi {
	r.refSkorKarakterPribadiDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkorKarakterPribadi) As(alias string) *refSkorKarakterPribadi {
	r.refSkorKarakterPribadiDo.DO = *(r.refSkorKarakterPribadiDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkorKarakterPribadi) updateTableName(table string) *refSkorKarakterPribadi {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Respon = field.NewString(table, "respon")
	r.Skor = field.NewInt32(table, "skor")
	r.Jawaban = field.NewString(table, "jawaban")

	r.fillFieldMap()

	return r
}

func (r *refSkorKarakterPribadi) WithContext(ctx context.Context) *refSkorKarakterPribadiDo {
	return r.refSkorKarakterPribadiDo.WithContext(ctx)
}

func (r refSkorKarakterPribadi) TableName() string { return r.refSkorKarakterPribadiDo.TableName() }

func (r refSkorKarakterPribadi) Alias() string { return r.refSkorKarakterPribadiDo.Alias() }

func (r refSkorKarakterPribadi) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkorKarakterPribadiDo.Columns(cols...)
}

func (r *refSkorKarakterPribadi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkorKarakterPribadi) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 4)
	r.fieldMap["id"] = r.ID
	r.fieldMap["respon"] = r.Respon
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["jawaban"] = r.Jawaban
}

func (r refSkorKarakterPribadi) clone(db *gorm.DB) refSkorKarakterPribadi {
	r.refSkorKarakterPribadiDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkorKarakterPribadi) replaceDB(db *gorm.DB) refSkorKarakterPribadi {
	r.refSkorKarakterPribadiDo.ReplaceDB(db)
	return r
}

type refSkorKarakterPribadiDo struct{ gen.DO }

func (r refSkorKarakterPribadiDo) Debug() *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkorKarakterPribadiDo) WithContext(ctx context.Context) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkorKarakterPribadiDo) ReadDB() *refSkorKarakterPribadiDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkorKarakterPribadiDo) WriteDB() *refSkorKarakterPribadiDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkorKarakterPribadiDo) Session(config *gorm.Session) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkorKarakterPribadiDo) Clauses(conds ...clause.Expression) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkorKarakterPribadiDo) Returning(value interface{}, columns ...string) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkorKarakterPribadiDo) Not(conds ...gen.Condition) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkorKarakterPribadiDo) Or(conds ...gen.Condition) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkorKarakterPribadiDo) Select(conds ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkorKarakterPribadiDo) Where(conds ...gen.Condition) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkorKarakterPribadiDo) Order(conds ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkorKarakterPribadiDo) Distinct(cols ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkorKarakterPribadiDo) Omit(cols ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkorKarakterPribadiDo) Join(table schema.Tabler, on ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkorKarakterPribadiDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkorKarakterPribadiDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkorKarakterPribadiDo) Group(cols ...field.Expr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkorKarakterPribadiDo) Having(conds ...gen.Condition) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkorKarakterPribadiDo) Limit(limit int) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkorKarakterPribadiDo) Offset(offset int) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkorKarakterPribadiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkorKarakterPribadiDo) Unscoped() *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkorKarakterPribadiDo) Create(values ...*entity.RefSkorKarakterPribadi) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkorKarakterPribadiDo) CreateInBatches(values []*entity.RefSkorKarakterPribadi, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkorKarakterPribadiDo) Save(values ...*entity.RefSkorKarakterPribadi) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkorKarakterPribadiDo) First() (*entity.RefSkorKarakterPribadi, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorKarakterPribadi), nil
	}
}

func (r refSkorKarakterPribadiDo) Take() (*entity.RefSkorKarakterPribadi, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorKarakterPribadi), nil
	}
}

func (r refSkorKarakterPribadiDo) Last() (*entity.RefSkorKarakterPribadi, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorKarakterPribadi), nil
	}
}

func (r refSkorKarakterPribadiDo) Find() ([]*entity.RefSkorKarakterPribadi, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkorKarakterPribadi), err
}

func (r refSkorKarakterPribadiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkorKarakterPribadi, err error) {
	buf := make([]*entity.RefSkorKarakterPribadi, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkorKarakterPribadiDo) FindInBatches(result *[]*entity.RefSkorKarakterPribadi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkorKarakterPribadiDo) Attrs(attrs ...field.AssignExpr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkorKarakterPribadiDo) Assign(attrs ...field.AssignExpr) *refSkorKarakterPribadiDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkorKarakterPribadiDo) Joins(fields ...field.RelationField) *refSkorKarakterPribadiDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkorKarakterPribadiDo) Preload(fields ...field.RelationField) *refSkorKarakterPribadiDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkorKarakterPribadiDo) FirstOrInit() (*entity.RefSkorKarakterPribadi, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorKarakterPribadi), nil
	}
}

func (r refSkorKarakterPribadiDo) FirstOrCreate() (*entity.RefSkorKarakterPribadi, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkorKarakterPribadi), nil
	}
}

func (r refSkorKarakterPribadiDo) FindByPage(offset int, limit int) (result []*entity.RefSkorKarakterPribadi, count int64, err error) {
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

func (r refSkorKarakterPribadiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkorKarakterPribadiDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkorKarakterPribadiDo) Delete(models ...*entity.RefSkorKarakterPribadi) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkorKarakterPribadiDo) withDO(do gen.Dao) *refSkorKarakterPribadiDo {
	r.DO = *do.(*gen.DO)
	return r
}
