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

func newRefKomponenKarakteristikPribadiEng(db *gorm.DB, opts ...gen.DOOption) refKomponenKarakteristikPribadiEng {
	_refKomponenKarakteristikPribadiEng := refKomponenKarakteristikPribadiEng{}

	_refKomponenKarakteristikPribadiEng.refKomponenKarakteristikPribadiEngDo.UseDB(db, opts...)
	_refKomponenKarakteristikPribadiEng.refKomponenKarakteristikPribadiEngDo.UseModel(&entity.RefKomponenKarakteristikPribadiEng{})

	tableName := _refKomponenKarakteristikPribadiEng.refKomponenKarakteristikPribadiEngDo.TableName()
	_refKomponenKarakteristikPribadiEng.ALL = field.NewAsterisk(tableName)
	_refKomponenKarakteristikPribadiEng.IDKomponen = field.NewInt32(tableName, "id_komponen")
	_refKomponenKarakteristikPribadiEng.NamaKomponen = field.NewString(tableName, "nama_komponen")
	_refKomponenKarakteristikPribadiEng.Keterangan = field.NewString(tableName, "keterangan")
	_refKomponenKarakteristikPribadiEng.FieldSkoring = field.NewString(tableName, "field_skoring")
	_refKomponenKarakteristikPribadiEng.UUID = field.NewString(tableName, "uuid")
	_refKomponenKarakteristikPribadiEng.Icon = field.NewString(tableName, "icon")

	_refKomponenKarakteristikPribadiEng.fillFieldMap()

	return _refKomponenKarakteristikPribadiEng
}

type refKomponenKarakteristikPribadiEng struct {
	refKomponenKarakteristikPribadiEngDo refKomponenKarakteristikPribadiEngDo

	ALL          field.Asterisk
	IDKomponen   field.Int32
	NamaKomponen field.String
	Keterangan   field.String
	FieldSkoring field.String
	UUID         field.String
	Icon         field.String

	fieldMap map[string]field.Expr
}

func (r refKomponenKarakteristikPribadiEng) Table(newTableName string) *refKomponenKarakteristikPribadiEng {
	r.refKomponenKarakteristikPribadiEngDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refKomponenKarakteristikPribadiEng) As(alias string) *refKomponenKarakteristikPribadiEng {
	r.refKomponenKarakteristikPribadiEngDo.DO = *(r.refKomponenKarakteristikPribadiEngDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refKomponenKarakteristikPribadiEng) updateTableName(table string) *refKomponenKarakteristikPribadiEng {
	r.ALL = field.NewAsterisk(table)
	r.IDKomponen = field.NewInt32(table, "id_komponen")
	r.NamaKomponen = field.NewString(table, "nama_komponen")
	r.Keterangan = field.NewString(table, "keterangan")
	r.FieldSkoring = field.NewString(table, "field_skoring")
	r.UUID = field.NewString(table, "uuid")
	r.Icon = field.NewString(table, "icon")

	r.fillFieldMap()

	return r
}

func (r *refKomponenKarakteristikPribadiEng) WithContext(ctx context.Context) *refKomponenKarakteristikPribadiEngDo {
	return r.refKomponenKarakteristikPribadiEngDo.WithContext(ctx)
}

func (r refKomponenKarakteristikPribadiEng) TableName() string {
	return r.refKomponenKarakteristikPribadiEngDo.TableName()
}

func (r refKomponenKarakteristikPribadiEng) Alias() string {
	return r.refKomponenKarakteristikPribadiEngDo.Alias()
}

func (r refKomponenKarakteristikPribadiEng) Columns(cols ...field.Expr) gen.Columns {
	return r.refKomponenKarakteristikPribadiEngDo.Columns(cols...)
}

func (r *refKomponenKarakteristikPribadiEng) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refKomponenKarakteristikPribadiEng) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id_komponen"] = r.IDKomponen
	r.fieldMap["nama_komponen"] = r.NamaKomponen
	r.fieldMap["keterangan"] = r.Keterangan
	r.fieldMap["field_skoring"] = r.FieldSkoring
	r.fieldMap["uuid"] = r.UUID
	r.fieldMap["icon"] = r.Icon
}

func (r refKomponenKarakteristikPribadiEng) clone(db *gorm.DB) refKomponenKarakteristikPribadiEng {
	r.refKomponenKarakteristikPribadiEngDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refKomponenKarakteristikPribadiEng) replaceDB(db *gorm.DB) refKomponenKarakteristikPribadiEng {
	r.refKomponenKarakteristikPribadiEngDo.ReplaceDB(db)
	return r
}

type refKomponenKarakteristikPribadiEngDo struct{ gen.DO }

func (r refKomponenKarakteristikPribadiEngDo) Debug() *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Debug())
}

func (r refKomponenKarakteristikPribadiEngDo) WithContext(ctx context.Context) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refKomponenKarakteristikPribadiEngDo) ReadDB() *refKomponenKarakteristikPribadiEngDo {
	return r.Clauses(dbresolver.Read)
}

func (r refKomponenKarakteristikPribadiEngDo) WriteDB() *refKomponenKarakteristikPribadiEngDo {
	return r.Clauses(dbresolver.Write)
}

func (r refKomponenKarakteristikPribadiEngDo) Session(config *gorm.Session) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Session(config))
}

func (r refKomponenKarakteristikPribadiEngDo) Clauses(conds ...clause.Expression) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Returning(value interface{}, columns ...string) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refKomponenKarakteristikPribadiEngDo) Not(conds ...gen.Condition) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Or(conds ...gen.Condition) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Select(conds ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Where(conds ...gen.Condition) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Order(conds ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Distinct(cols ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refKomponenKarakteristikPribadiEngDo) Omit(cols ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refKomponenKarakteristikPribadiEngDo) Join(table schema.Tabler, on ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refKomponenKarakteristikPribadiEngDo) LeftJoin(table schema.Tabler, on ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refKomponenKarakteristikPribadiEngDo) RightJoin(table schema.Tabler, on ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refKomponenKarakteristikPribadiEngDo) Group(cols ...field.Expr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refKomponenKarakteristikPribadiEngDo) Having(conds ...gen.Condition) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refKomponenKarakteristikPribadiEngDo) Limit(limit int) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refKomponenKarakteristikPribadiEngDo) Offset(offset int) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refKomponenKarakteristikPribadiEngDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refKomponenKarakteristikPribadiEngDo) Unscoped() *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refKomponenKarakteristikPribadiEngDo) Create(values ...*entity.RefKomponenKarakteristikPribadiEng) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refKomponenKarakteristikPribadiEngDo) CreateInBatches(values []*entity.RefKomponenKarakteristikPribadiEng, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refKomponenKarakteristikPribadiEngDo) Save(values ...*entity.RefKomponenKarakteristikPribadiEng) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refKomponenKarakteristikPribadiEngDo) First() (*entity.RefKomponenKarakteristikPribadiEng, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKomponenKarakteristikPribadiEng), nil
	}
}

func (r refKomponenKarakteristikPribadiEngDo) Take() (*entity.RefKomponenKarakteristikPribadiEng, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKomponenKarakteristikPribadiEng), nil
	}
}

func (r refKomponenKarakteristikPribadiEngDo) Last() (*entity.RefKomponenKarakteristikPribadiEng, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKomponenKarakteristikPribadiEng), nil
	}
}

func (r refKomponenKarakteristikPribadiEngDo) Find() ([]*entity.RefKomponenKarakteristikPribadiEng, error) {
	result, err := r.DO.Find()
	return result.([]*entity.RefKomponenKarakteristikPribadiEng), err
}

func (r refKomponenKarakteristikPribadiEngDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.RefKomponenKarakteristikPribadiEng, err error) {
	buf := make([]*entity.RefKomponenKarakteristikPribadiEng, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refKomponenKarakteristikPribadiEngDo) FindInBatches(result *[]*entity.RefKomponenKarakteristikPribadiEng, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refKomponenKarakteristikPribadiEngDo) Attrs(attrs ...field.AssignExpr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refKomponenKarakteristikPribadiEngDo) Assign(attrs ...field.AssignExpr) *refKomponenKarakteristikPribadiEngDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refKomponenKarakteristikPribadiEngDo) Joins(fields ...field.RelationField) *refKomponenKarakteristikPribadiEngDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refKomponenKarakteristikPribadiEngDo) Preload(fields ...field.RelationField) *refKomponenKarakteristikPribadiEngDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refKomponenKarakteristikPribadiEngDo) FirstOrInit() (*entity.RefKomponenKarakteristikPribadiEng, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKomponenKarakteristikPribadiEng), nil
	}
}

func (r refKomponenKarakteristikPribadiEngDo) FirstOrCreate() (*entity.RefKomponenKarakteristikPribadiEng, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.RefKomponenKarakteristikPribadiEng), nil
	}
}

func (r refKomponenKarakteristikPribadiEngDo) FindByPage(offset int, limit int) (result []*entity.RefKomponenKarakteristikPribadiEng, count int64, err error) {
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

func (r refKomponenKarakteristikPribadiEngDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refKomponenKarakteristikPribadiEngDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refKomponenKarakteristikPribadiEngDo) Delete(models ...*entity.RefKomponenKarakteristikPribadiEng) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refKomponenKarakteristikPribadiEngDo) withDO(do gen.Dao) *refKomponenKarakteristikPribadiEngDo {
	r.DO = *do.(*gen.DO)
	return r
}