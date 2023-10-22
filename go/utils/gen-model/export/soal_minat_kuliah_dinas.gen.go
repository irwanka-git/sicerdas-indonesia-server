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

func newSoalMinatKuliahDina(db *gorm.DB, opts ...gen.DOOption) soalMinatKuliahDina {
	_soalMinatKuliahDina := soalMinatKuliahDina{}

	_soalMinatKuliahDina.soalMinatKuliahDinaDo.UseDB(db, opts...)
	_soalMinatKuliahDina.soalMinatKuliahDinaDo.UseModel(&entity.SoalMinatKuliahDina{})

	tableName := _soalMinatKuliahDina.soalMinatKuliahDinaDo.TableName()
	_soalMinatKuliahDina.ALL = field.NewAsterisk(tableName)
	_soalMinatKuliahDina.Nomor = field.NewString(tableName, "nomor")
	_soalMinatKuliahDina.Deskripsi = field.NewString(tableName, "deskripsi")
	_soalMinatKuliahDina.PernyataanA = field.NewString(tableName, "pernyataan_a")
	_soalMinatKuliahDina.PernyataanB = field.NewString(tableName, "pernyataan_b")
	_soalMinatKuliahDina.PernyataanC = field.NewString(tableName, "pernyataan_c")
	_soalMinatKuliahDina.PernyataanD = field.NewString(tableName, "pernyataan_d")
	_soalMinatKuliahDina.PernyataanE = field.NewString(tableName, "pernyataan_e")
	_soalMinatKuliahDina.PernyataanF = field.NewString(tableName, "pernyataan_f")
	_soalMinatKuliahDina.PernyataanG = field.NewString(tableName, "pernyataan_g")
	_soalMinatKuliahDina.PernyataanH = field.NewString(tableName, "pernyataan_h")
	_soalMinatKuliahDina.PernyataanI = field.NewString(tableName, "pernyataan_i")
	_soalMinatKuliahDina.PernyataanJ = field.NewString(tableName, "pernyataan_j")
	_soalMinatKuliahDina.PernyataanK = field.NewString(tableName, "pernyataan_k")
	_soalMinatKuliahDina.PernyataanL = field.NewString(tableName, "pernyataan_l")
	_soalMinatKuliahDina.UUID = field.NewString(tableName, "uuid")

	_soalMinatKuliahDina.fillFieldMap()

	return _soalMinatKuliahDina
}

type soalMinatKuliahDina struct {
	soalMinatKuliahDinaDo soalMinatKuliahDinaDo

	ALL         field.Asterisk
	Nomor       field.String
	Deskripsi   field.String
	PernyataanA field.String
	PernyataanB field.String
	PernyataanC field.String
	PernyataanD field.String
	PernyataanE field.String
	PernyataanF field.String
	PernyataanG field.String
	PernyataanH field.String
	PernyataanI field.String
	PernyataanJ field.String
	PernyataanK field.String
	PernyataanL field.String
	UUID        field.String

	fieldMap map[string]field.Expr
}

func (s soalMinatKuliahDina) Table(newTableName string) *soalMinatKuliahDina {
	s.soalMinatKuliahDinaDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s soalMinatKuliahDina) As(alias string) *soalMinatKuliahDina {
	s.soalMinatKuliahDinaDo.DO = *(s.soalMinatKuliahDinaDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *soalMinatKuliahDina) updateTableName(table string) *soalMinatKuliahDina {
	s.ALL = field.NewAsterisk(table)
	s.Nomor = field.NewString(table, "nomor")
	s.Deskripsi = field.NewString(table, "deskripsi")
	s.PernyataanA = field.NewString(table, "pernyataan_a")
	s.PernyataanB = field.NewString(table, "pernyataan_b")
	s.PernyataanC = field.NewString(table, "pernyataan_c")
	s.PernyataanD = field.NewString(table, "pernyataan_d")
	s.PernyataanE = field.NewString(table, "pernyataan_e")
	s.PernyataanF = field.NewString(table, "pernyataan_f")
	s.PernyataanG = field.NewString(table, "pernyataan_g")
	s.PernyataanH = field.NewString(table, "pernyataan_h")
	s.PernyataanI = field.NewString(table, "pernyataan_i")
	s.PernyataanJ = field.NewString(table, "pernyataan_j")
	s.PernyataanK = field.NewString(table, "pernyataan_k")
	s.PernyataanL = field.NewString(table, "pernyataan_l")
	s.UUID = field.NewString(table, "uuid")

	s.fillFieldMap()

	return s
}

func (s *soalMinatKuliahDina) WithContext(ctx context.Context) *soalMinatKuliahDinaDo {
	return s.soalMinatKuliahDinaDo.WithContext(ctx)
}

func (s soalMinatKuliahDina) TableName() string { return s.soalMinatKuliahDinaDo.TableName() }

func (s soalMinatKuliahDina) Alias() string { return s.soalMinatKuliahDinaDo.Alias() }

func (s soalMinatKuliahDina) Columns(cols ...field.Expr) gen.Columns {
	return s.soalMinatKuliahDinaDo.Columns(cols...)
}

func (s *soalMinatKuliahDina) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *soalMinatKuliahDina) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["nomor"] = s.Nomor
	s.fieldMap["deskripsi"] = s.Deskripsi
	s.fieldMap["pernyataan_a"] = s.PernyataanA
	s.fieldMap["pernyataan_b"] = s.PernyataanB
	s.fieldMap["pernyataan_c"] = s.PernyataanC
	s.fieldMap["pernyataan_d"] = s.PernyataanD
	s.fieldMap["pernyataan_e"] = s.PernyataanE
	s.fieldMap["pernyataan_f"] = s.PernyataanF
	s.fieldMap["pernyataan_g"] = s.PernyataanG
	s.fieldMap["pernyataan_h"] = s.PernyataanH
	s.fieldMap["pernyataan_i"] = s.PernyataanI
	s.fieldMap["pernyataan_j"] = s.PernyataanJ
	s.fieldMap["pernyataan_k"] = s.PernyataanK
	s.fieldMap["pernyataan_l"] = s.PernyataanL
	s.fieldMap["uuid"] = s.UUID
}

func (s soalMinatKuliahDina) clone(db *gorm.DB) soalMinatKuliahDina {
	s.soalMinatKuliahDinaDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s soalMinatKuliahDina) replaceDB(db *gorm.DB) soalMinatKuliahDina {
	s.soalMinatKuliahDinaDo.ReplaceDB(db)
	return s
}

type soalMinatKuliahDinaDo struct{ gen.DO }

func (s soalMinatKuliahDinaDo) Debug() *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Debug())
}

func (s soalMinatKuliahDinaDo) WithContext(ctx context.Context) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s soalMinatKuliahDinaDo) ReadDB() *soalMinatKuliahDinaDo {
	return s.Clauses(dbresolver.Read)
}

func (s soalMinatKuliahDinaDo) WriteDB() *soalMinatKuliahDinaDo {
	return s.Clauses(dbresolver.Write)
}

func (s soalMinatKuliahDinaDo) Session(config *gorm.Session) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Session(config))
}

func (s soalMinatKuliahDinaDo) Clauses(conds ...clause.Expression) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s soalMinatKuliahDinaDo) Returning(value interface{}, columns ...string) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s soalMinatKuliahDinaDo) Not(conds ...gen.Condition) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s soalMinatKuliahDinaDo) Or(conds ...gen.Condition) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s soalMinatKuliahDinaDo) Select(conds ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s soalMinatKuliahDinaDo) Where(conds ...gen.Condition) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s soalMinatKuliahDinaDo) Order(conds ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s soalMinatKuliahDinaDo) Distinct(cols ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s soalMinatKuliahDinaDo) Omit(cols ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s soalMinatKuliahDinaDo) Join(table schema.Tabler, on ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s soalMinatKuliahDinaDo) LeftJoin(table schema.Tabler, on ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s soalMinatKuliahDinaDo) RightJoin(table schema.Tabler, on ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s soalMinatKuliahDinaDo) Group(cols ...field.Expr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s soalMinatKuliahDinaDo) Having(conds ...gen.Condition) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s soalMinatKuliahDinaDo) Limit(limit int) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s soalMinatKuliahDinaDo) Offset(offset int) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s soalMinatKuliahDinaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s soalMinatKuliahDinaDo) Unscoped() *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Unscoped())
}

func (s soalMinatKuliahDinaDo) Create(values ...*entity.SoalMinatKuliahDina) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s soalMinatKuliahDinaDo) CreateInBatches(values []*entity.SoalMinatKuliahDina, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s soalMinatKuliahDinaDo) Save(values ...*entity.SoalMinatKuliahDina) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s soalMinatKuliahDinaDo) First() (*entity.SoalMinatKuliahDina, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalMinatKuliahDina), nil
	}
}

func (s soalMinatKuliahDinaDo) Take() (*entity.SoalMinatKuliahDina, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalMinatKuliahDina), nil
	}
}

func (s soalMinatKuliahDinaDo) Last() (*entity.SoalMinatKuliahDina, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalMinatKuliahDina), nil
	}
}

func (s soalMinatKuliahDinaDo) Find() ([]*entity.SoalMinatKuliahDina, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SoalMinatKuliahDina), err
}

func (s soalMinatKuliahDinaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SoalMinatKuliahDina, err error) {
	buf := make([]*entity.SoalMinatKuliahDina, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s soalMinatKuliahDinaDo) FindInBatches(result *[]*entity.SoalMinatKuliahDina, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s soalMinatKuliahDinaDo) Attrs(attrs ...field.AssignExpr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s soalMinatKuliahDinaDo) Assign(attrs ...field.AssignExpr) *soalMinatKuliahDinaDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s soalMinatKuliahDinaDo) Joins(fields ...field.RelationField) *soalMinatKuliahDinaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s soalMinatKuliahDinaDo) Preload(fields ...field.RelationField) *soalMinatKuliahDinaDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s soalMinatKuliahDinaDo) FirstOrInit() (*entity.SoalMinatKuliahDina, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalMinatKuliahDina), nil
	}
}

func (s soalMinatKuliahDinaDo) FirstOrCreate() (*entity.SoalMinatKuliahDina, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalMinatKuliahDina), nil
	}
}

func (s soalMinatKuliahDinaDo) FindByPage(offset int, limit int) (result []*entity.SoalMinatKuliahDina, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s soalMinatKuliahDinaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s soalMinatKuliahDinaDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s soalMinatKuliahDinaDo) Delete(models ...*entity.SoalMinatKuliahDina) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *soalMinatKuliahDinaDo) withDO(do gen.Dao) *soalMinatKuliahDinaDo {
	s.DO = *do.(*gen.DO)
	return s
}
