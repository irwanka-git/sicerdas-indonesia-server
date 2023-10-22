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

func newPetunjukSoal(db *gorm.DB, opts ...gen.DOOption) petunjukSoal {
	_petunjukSoal := petunjukSoal{}

	_petunjukSoal.petunjukSoalDo.UseDB(db, opts...)
	_petunjukSoal.petunjukSoalDo.UseModel(&entity.PetunjukSoal{})

	tableName := _petunjukSoal.petunjukSoalDo.TableName()
	_petunjukSoal.ALL = field.NewAsterisk(tableName)
	_petunjukSoal.IDPetunjuk = field.NewInt32(tableName, "id_petunjuk")
	_petunjukSoal.IsiPetunjuk = field.NewString(tableName, "isi_petunjuk")
	_petunjukSoal.UUID = field.NewString(tableName, "uuid")
	_petunjukSoal.Keterangan = field.NewString(tableName, "keterangan")

	_petunjukSoal.fillFieldMap()

	return _petunjukSoal
}

type petunjukSoal struct {
	petunjukSoalDo petunjukSoalDo

	ALL         field.Asterisk
	IDPetunjuk  field.Int32
	IsiPetunjuk field.String
	UUID        field.String
	Keterangan  field.String

	fieldMap map[string]field.Expr
}

func (p petunjukSoal) Table(newTableName string) *petunjukSoal {
	p.petunjukSoalDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p petunjukSoal) As(alias string) *petunjukSoal {
	p.petunjukSoalDo.DO = *(p.petunjukSoalDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *petunjukSoal) updateTableName(table string) *petunjukSoal {
	p.ALL = field.NewAsterisk(table)
	p.IDPetunjuk = field.NewInt32(table, "id_petunjuk")
	p.IsiPetunjuk = field.NewString(table, "isi_petunjuk")
	p.UUID = field.NewString(table, "uuid")
	p.Keterangan = field.NewString(table, "keterangan")

	p.fillFieldMap()

	return p
}

func (p *petunjukSoal) WithContext(ctx context.Context) *petunjukSoalDo {
	return p.petunjukSoalDo.WithContext(ctx)
}

func (p petunjukSoal) TableName() string { return p.petunjukSoalDo.TableName() }

func (p petunjukSoal) Alias() string { return p.petunjukSoalDo.Alias() }

func (p petunjukSoal) Columns(cols ...field.Expr) gen.Columns {
	return p.petunjukSoalDo.Columns(cols...)
}

func (p *petunjukSoal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *petunjukSoal) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id_petunjuk"] = p.IDPetunjuk
	p.fieldMap["isi_petunjuk"] = p.IsiPetunjuk
	p.fieldMap["uuid"] = p.UUID
	p.fieldMap["keterangan"] = p.Keterangan
}

func (p petunjukSoal) clone(db *gorm.DB) petunjukSoal {
	p.petunjukSoalDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p petunjukSoal) replaceDB(db *gorm.DB) petunjukSoal {
	p.petunjukSoalDo.ReplaceDB(db)
	return p
}

type petunjukSoalDo struct{ gen.DO }

func (p petunjukSoalDo) Debug() *petunjukSoalDo {
	return p.withDO(p.DO.Debug())
}

func (p petunjukSoalDo) WithContext(ctx context.Context) *petunjukSoalDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p petunjukSoalDo) ReadDB() *petunjukSoalDo {
	return p.Clauses(dbresolver.Read)
}

func (p petunjukSoalDo) WriteDB() *petunjukSoalDo {
	return p.Clauses(dbresolver.Write)
}

func (p petunjukSoalDo) Session(config *gorm.Session) *petunjukSoalDo {
	return p.withDO(p.DO.Session(config))
}

func (p petunjukSoalDo) Clauses(conds ...clause.Expression) *petunjukSoalDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p petunjukSoalDo) Returning(value interface{}, columns ...string) *petunjukSoalDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p petunjukSoalDo) Not(conds ...gen.Condition) *petunjukSoalDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p petunjukSoalDo) Or(conds ...gen.Condition) *petunjukSoalDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p petunjukSoalDo) Select(conds ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p petunjukSoalDo) Where(conds ...gen.Condition) *petunjukSoalDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p petunjukSoalDo) Order(conds ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p petunjukSoalDo) Distinct(cols ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p petunjukSoalDo) Omit(cols ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p petunjukSoalDo) Join(table schema.Tabler, on ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p petunjukSoalDo) LeftJoin(table schema.Tabler, on ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p petunjukSoalDo) RightJoin(table schema.Tabler, on ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p petunjukSoalDo) Group(cols ...field.Expr) *petunjukSoalDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p petunjukSoalDo) Having(conds ...gen.Condition) *petunjukSoalDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p petunjukSoalDo) Limit(limit int) *petunjukSoalDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p petunjukSoalDo) Offset(offset int) *petunjukSoalDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p petunjukSoalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *petunjukSoalDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p petunjukSoalDo) Unscoped() *petunjukSoalDo {
	return p.withDO(p.DO.Unscoped())
}

func (p petunjukSoalDo) Create(values ...*entity.PetunjukSoal) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p petunjukSoalDo) CreateInBatches(values []*entity.PetunjukSoal, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p petunjukSoalDo) Save(values ...*entity.PetunjukSoal) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p petunjukSoalDo) First() (*entity.PetunjukSoal, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PetunjukSoal), nil
	}
}

func (p petunjukSoalDo) Take() (*entity.PetunjukSoal, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PetunjukSoal), nil
	}
}

func (p petunjukSoalDo) Last() (*entity.PetunjukSoal, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PetunjukSoal), nil
	}
}

func (p petunjukSoalDo) Find() ([]*entity.PetunjukSoal, error) {
	result, err := p.DO.Find()
	return result.([]*entity.PetunjukSoal), err
}

func (p petunjukSoalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.PetunjukSoal, err error) {
	buf := make([]*entity.PetunjukSoal, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p petunjukSoalDo) FindInBatches(result *[]*entity.PetunjukSoal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p petunjukSoalDo) Attrs(attrs ...field.AssignExpr) *petunjukSoalDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p petunjukSoalDo) Assign(attrs ...field.AssignExpr) *petunjukSoalDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p petunjukSoalDo) Joins(fields ...field.RelationField) *petunjukSoalDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p petunjukSoalDo) Preload(fields ...field.RelationField) *petunjukSoalDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p petunjukSoalDo) FirstOrInit() (*entity.PetunjukSoal, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PetunjukSoal), nil
	}
}

func (p petunjukSoalDo) FirstOrCreate() (*entity.PetunjukSoal, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PetunjukSoal), nil
	}
}

func (p petunjukSoalDo) FindByPage(offset int, limit int) (result []*entity.PetunjukSoal, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p petunjukSoalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p petunjukSoalDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p petunjukSoalDo) Delete(models ...*entity.PetunjukSoal) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *petunjukSoalDo) withDO(do gen.Dao) *petunjukSoalDo {
	p.DO = *do.(*gen.DO)
	return p
}
