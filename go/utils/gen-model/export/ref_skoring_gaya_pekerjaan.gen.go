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

func newRefSkoringGayaPekerjaan(db *gorm.DB, opts ...gen.DOOption) refSkoringGayaPekerjaan {
	_refSkoringGayaPekerjaan := refSkoringGayaPekerjaan{}

	_refSkoringGayaPekerjaan.refSkoringGayaPekerjaanDo.UseDB(db, opts...)
	_refSkoringGayaPekerjaan.refSkoringGayaPekerjaanDo.UseModel(&entity.RefSkoringGayaPekerjaan{})

	tableName := _refSkoringGayaPekerjaan.refSkoringGayaPekerjaanDo.TableName()
	_refSkoringGayaPekerjaan.ALL = field.NewAsterisk(tableName)
	_refSkoringGayaPekerjaan.ID = field.NewInt64(tableName, "id")
	_refSkoringGayaPekerjaan.IDUser = field.NewInt32(tableName, "id_user")
	_refSkoringGayaPekerjaan.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_refSkoringGayaPekerjaan.Kode = field.NewString(tableName, "kode")
	_refSkoringGayaPekerjaan.Skor = field.NewInt16(tableName, "skor")
	_refSkoringGayaPekerjaan.Rangking = field.NewInt16(tableName, "rangking")
	_refSkoringGayaPekerjaan.Klasifikasi = field.NewString(tableName, "klasifikasi")

	_refSkoringGayaPekerjaan.fillFieldMap()

	return _refSkoringGayaPekerjaan
}

type refSkoringGayaPekerjaan struct {
	refSkoringGayaPekerjaanDo refSkoringGayaPekerjaanDo

	ALL         field.Asterisk
	ID          field.Int64
	IDUser      field.Int32
	IDQuiz      field.Int32
	Kode        field.String
	Skor        field.Int16
	Rangking    field.Int16
	Klasifikasi field.String

	fieldMap map[string]field.Expr
}

func (r refSkoringGayaPekerjaan) Table(newTableName string) *refSkoringGayaPekerjaan {
	r.refSkoringGayaPekerjaanDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refSkoringGayaPekerjaan) As(alias string) *refSkoringGayaPekerjaan {
	r.refSkoringGayaPekerjaanDo.DO = *(r.refSkoringGayaPekerjaanDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refSkoringGayaPekerjaan) updateTableName(table string) *refSkoringGayaPekerjaan {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.IDUser = field.NewInt32(table, "id_user")
	r.IDQuiz = field.NewInt32(table, "id_quiz")
	r.Kode = field.NewString(table, "kode")
	r.Skor = field.NewInt16(table, "skor")
	r.Rangking = field.NewInt16(table, "rangking")
	r.Klasifikasi = field.NewString(table, "klasifikasi")

	r.fillFieldMap()

	return r
}

func (r *refSkoringGayaPekerjaan) WithContext(ctx context.Context) *refSkoringGayaPekerjaanDo {
	return r.refSkoringGayaPekerjaanDo.WithContext(ctx)
}

func (r refSkoringGayaPekerjaan) TableName() string { return r.refSkoringGayaPekerjaanDo.TableName() }

func (r refSkoringGayaPekerjaan) Alias() string { return r.refSkoringGayaPekerjaanDo.Alias() }

func (r refSkoringGayaPekerjaan) Columns(cols ...field.Expr) gen.Columns {
	return r.refSkoringGayaPekerjaanDo.Columns(cols...)
}

func (r *refSkoringGayaPekerjaan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refSkoringGayaPekerjaan) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["id_user"] = r.IDUser
	r.fieldMap["id_quiz"] = r.IDQuiz
	r.fieldMap["kode"] = r.Kode
	r.fieldMap["skor"] = r.Skor
	r.fieldMap["rangking"] = r.Rangking
	r.fieldMap["klasifikasi"] = r.Klasifikasi
}

func (r refSkoringGayaPekerjaan) clone(db *gorm.DB) refSkoringGayaPekerjaan {
	r.refSkoringGayaPekerjaanDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refSkoringGayaPekerjaan) replaceDB(db *gorm.DB) refSkoringGayaPekerjaan {
	r.refSkoringGayaPekerjaanDo.ReplaceDB(db)
	return r
}

type refSkoringGayaPekerjaanDo struct{ gen.DO }

func (r refSkoringGayaPekerjaanDo) Debug() *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Debug())
}

func (r refSkoringGayaPekerjaanDo) WithContext(ctx context.Context) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refSkoringGayaPekerjaanDo) ReadDB() *refSkoringGayaPekerjaanDo {
	return r.Clauses(dbresolver.Read)
}

func (r refSkoringGayaPekerjaanDo) WriteDB() *refSkoringGayaPekerjaanDo {
	return r.Clauses(dbresolver.Write)
}

func (r refSkoringGayaPekerjaanDo) Session(config *gorm.Session) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Session(config))
}

func (r refSkoringGayaPekerjaanDo) Clauses(conds ...clause.Expression) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refSkoringGayaPekerjaanDo) Returning(value interface{}, columns ...string) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refSkoringGayaPekerjaanDo) Not(conds ...gen.Condition) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refSkoringGayaPekerjaanDo) Or(conds ...gen.Condition) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refSkoringGayaPekerjaanDo) Select(conds ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refSkoringGayaPekerjaanDo) Where(conds ...gen.Condition) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refSkoringGayaPekerjaanDo) Order(conds ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refSkoringGayaPekerjaanDo) Distinct(cols ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refSkoringGayaPekerjaanDo) Omit(cols ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refSkoringGayaPekerjaanDo) Join(table schema.Tabler, on ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refSkoringGayaPekerjaanDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refSkoringGayaPekerjaanDo) RightJoin(table schema.Tabler, on ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refSkoringGayaPekerjaanDo) Group(cols ...field.Expr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refSkoringGayaPekerjaanDo) Having(conds ...gen.Condition) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refSkoringGayaPekerjaanDo) Limit(limit int) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refSkoringGayaPekerjaanDo) Offset(offset int) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refSkoringGayaPekerjaanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refSkoringGayaPekerjaanDo) Unscoped() *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refSkoringGayaPekerjaanDo) Create(values ...*entity.RefSkoringGayaPekerjaan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refSkoringGayaPekerjaanDo) CreateInBatches(values []*entity.RefSkoringGayaPekerjaan, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refSkoringGayaPekerjaanDo) Save(values ...*entity.RefSkoringGayaPekerjaan) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refSkoringGayaPekerjaanDo) First() (*entity.RefSkoringGayaPekerjaan, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringGayaPekerjaan), nil
	}
}

func (r refSkoringGayaPekerjaanDo) Take() (*entity.RefSkoringGayaPekerjaan, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringGayaPekerjaan), nil
	}
}

func (r refSkoringGayaPekerjaanDo) Last() (*entity.RefSkoringGayaPekerjaan, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringGayaPekerjaan), nil
	}
}

func (r refSkoringGayaPekerjaanDo) Find() ([]*entity.RefSkoringGayaPekerjaan, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefSkoringGayaPekerjaan), err
}

func (r refSkoringGayaPekerjaanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefSkoringGayaPekerjaan, err error) {
	buf := make([]*entity.RefSkoringGayaPekerjaan, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refSkoringGayaPekerjaanDo) FindInBatches(result *[]*entity.RefSkoringGayaPekerjaan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refSkoringGayaPekerjaanDo) Attrs(attrs ...field.AssignExpr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refSkoringGayaPekerjaanDo) Assign(attrs ...field.AssignExpr) *refSkoringGayaPekerjaanDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refSkoringGayaPekerjaanDo) Joins(fields ...field.RelationField) *refSkoringGayaPekerjaanDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refSkoringGayaPekerjaanDo) Preload(fields ...field.RelationField) *refSkoringGayaPekerjaanDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refSkoringGayaPekerjaanDo) FirstOrInit() (*entity.RefSkoringGayaPekerjaan, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringGayaPekerjaan), nil
	}
}

func (r refSkoringGayaPekerjaanDo) FirstOrCreate() (*entity.RefSkoringGayaPekerjaan, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefSkoringGayaPekerjaan), nil
	}
}

func (r refSkoringGayaPekerjaanDo) FindByPage(offset int, limit int) (result []*entity.RefSkoringGayaPekerjaan, count int64, err error) {
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

func (r refSkoringGayaPekerjaanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refSkoringGayaPekerjaanDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refSkoringGayaPekerjaanDo) Delete(models ...*entity.RefSkoringGayaPekerjaan) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refSkoringGayaPekerjaanDo) withDO(do gen.Dao) *refSkoringGayaPekerjaanDo {
	r.DO = *do.(*gen.DO)
	return r
}
