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

func newSoalKecerdasanMajemuk(db *gorm.DB, opts ...gen.DOOption) soalKecerdasanMajemuk {
	_soalKecerdasanMajemuk := soalKecerdasanMajemuk{}

	_soalKecerdasanMajemuk.soalKecerdasanMajemukDo.UseDB(db, opts...)
	_soalKecerdasanMajemuk.soalKecerdasanMajemukDo.UseModel(&entity.SoalKecerdasanMajemuk{})

	tableName := _soalKecerdasanMajemuk.soalKecerdasanMajemukDo.TableName()
	_soalKecerdasanMajemuk.ALL = field.NewAsterisk(tableName)
	_soalKecerdasanMajemuk.Nomor = field.NewString(tableName, "nomor")
	_soalKecerdasanMajemuk.Deskripsi = field.NewString(tableName, "deskripsi")
	_soalKecerdasanMajemuk.PernyataanA = field.NewString(tableName, "pernyataan_a")
	_soalKecerdasanMajemuk.PernyataanB = field.NewString(tableName, "pernyataan_b")
	_soalKecerdasanMajemuk.PernyataanC = field.NewString(tableName, "pernyataan_c")
	_soalKecerdasanMajemuk.PernyataanD = field.NewString(tableName, "pernyataan_d")
	_soalKecerdasanMajemuk.PernyataanE = field.NewString(tableName, "pernyataan_e")
	_soalKecerdasanMajemuk.PernyataanF = field.NewString(tableName, "pernyataan_f")
	_soalKecerdasanMajemuk.PernyataanG = field.NewString(tableName, "pernyataan_g")
	_soalKecerdasanMajemuk.PernyataanH = field.NewString(tableName, "pernyataan_h")
	_soalKecerdasanMajemuk.PernyataanI = field.NewString(tableName, "pernyataan_i")
	_soalKecerdasanMajemuk.PernyataanJ = field.NewString(tableName, "pernyataan_j")
	_soalKecerdasanMajemuk.PernyataanK = field.NewString(tableName, "pernyataan_k")
	_soalKecerdasanMajemuk.PernyataanL = field.NewString(tableName, "pernyataan_l")
	_soalKecerdasanMajemuk.UUID = field.NewString(tableName, "uuid")

	_soalKecerdasanMajemuk.fillFieldMap()

	return _soalKecerdasanMajemuk
}

type soalKecerdasanMajemuk struct {
	soalKecerdasanMajemukDo soalKecerdasanMajemukDo

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

func (s soalKecerdasanMajemuk) Table(newTableName string) *soalKecerdasanMajemuk {
	s.soalKecerdasanMajemukDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s soalKecerdasanMajemuk) As(alias string) *soalKecerdasanMajemuk {
	s.soalKecerdasanMajemukDo.DO = *(s.soalKecerdasanMajemukDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *soalKecerdasanMajemuk) updateTableName(table string) *soalKecerdasanMajemuk {
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

func (s *soalKecerdasanMajemuk) WithContext(ctx context.Context) *soalKecerdasanMajemukDo {
	return s.soalKecerdasanMajemukDo.WithContext(ctx)
}

func (s soalKecerdasanMajemuk) TableName() string { return s.soalKecerdasanMajemukDo.TableName() }

func (s soalKecerdasanMajemuk) Alias() string { return s.soalKecerdasanMajemukDo.Alias() }

func (s soalKecerdasanMajemuk) Columns(cols ...field.Expr) gen.Columns {
	return s.soalKecerdasanMajemukDo.Columns(cols...)
}

func (s *soalKecerdasanMajemuk) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *soalKecerdasanMajemuk) fillFieldMap() {
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

func (s soalKecerdasanMajemuk) clone(db *gorm.DB) soalKecerdasanMajemuk {
	s.soalKecerdasanMajemukDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s soalKecerdasanMajemuk) replaceDB(db *gorm.DB) soalKecerdasanMajemuk {
	s.soalKecerdasanMajemukDo.ReplaceDB(db)
	return s
}

type soalKecerdasanMajemukDo struct{ gen.DO }

func (s soalKecerdasanMajemukDo) Debug() *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Debug())
}

func (s soalKecerdasanMajemukDo) WithContext(ctx context.Context) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s soalKecerdasanMajemukDo) ReadDB() *soalKecerdasanMajemukDo {
	return s.Clauses(dbresolver.Read)
}

func (s soalKecerdasanMajemukDo) WriteDB() *soalKecerdasanMajemukDo {
	return s.Clauses(dbresolver.Write)
}

func (s soalKecerdasanMajemukDo) Session(config *gorm.Session) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Session(config))
}

func (s soalKecerdasanMajemukDo) Clauses(conds ...clause.Expression) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s soalKecerdasanMajemukDo) Returning(value interface{}, columns ...string) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s soalKecerdasanMajemukDo) Not(conds ...gen.Condition) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s soalKecerdasanMajemukDo) Or(conds ...gen.Condition) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s soalKecerdasanMajemukDo) Select(conds ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s soalKecerdasanMajemukDo) Where(conds ...gen.Condition) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s soalKecerdasanMajemukDo) Order(conds ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s soalKecerdasanMajemukDo) Distinct(cols ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s soalKecerdasanMajemukDo) Omit(cols ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s soalKecerdasanMajemukDo) Join(table schema.Tabler, on ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s soalKecerdasanMajemukDo) LeftJoin(table schema.Tabler, on ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s soalKecerdasanMajemukDo) RightJoin(table schema.Tabler, on ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s soalKecerdasanMajemukDo) Group(cols ...field.Expr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s soalKecerdasanMajemukDo) Having(conds ...gen.Condition) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s soalKecerdasanMajemukDo) Limit(limit int) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s soalKecerdasanMajemukDo) Offset(offset int) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s soalKecerdasanMajemukDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s soalKecerdasanMajemukDo) Unscoped() *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Unscoped())
}

func (s soalKecerdasanMajemukDo) Create(values ...*entity.SoalKecerdasanMajemuk) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s soalKecerdasanMajemukDo) CreateInBatches(values []*entity.SoalKecerdasanMajemuk, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s soalKecerdasanMajemukDo) Save(values ...*entity.SoalKecerdasanMajemuk) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s soalKecerdasanMajemukDo) First() (*entity.SoalKecerdasanMajemuk, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKecerdasanMajemuk), nil
	}
}

func (s soalKecerdasanMajemukDo) Take() (*entity.SoalKecerdasanMajemuk, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKecerdasanMajemuk), nil
	}
}

func (s soalKecerdasanMajemukDo) Last() (*entity.SoalKecerdasanMajemuk, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKecerdasanMajemuk), nil
	}
}

func (s soalKecerdasanMajemukDo) Find() ([]*entity.SoalKecerdasanMajemuk, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SoalKecerdasanMajemuk), err
}

func (s soalKecerdasanMajemukDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SoalKecerdasanMajemuk, err error) {
	buf := make([]*entity.SoalKecerdasanMajemuk, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s soalKecerdasanMajemukDo) FindInBatches(result *[]*entity.SoalKecerdasanMajemuk, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s soalKecerdasanMajemukDo) Attrs(attrs ...field.AssignExpr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s soalKecerdasanMajemukDo) Assign(attrs ...field.AssignExpr) *soalKecerdasanMajemukDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s soalKecerdasanMajemukDo) Joins(fields ...field.RelationField) *soalKecerdasanMajemukDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s soalKecerdasanMajemukDo) Preload(fields ...field.RelationField) *soalKecerdasanMajemukDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s soalKecerdasanMajemukDo) FirstOrInit() (*entity.SoalKecerdasanMajemuk, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKecerdasanMajemuk), nil
	}
}

func (s soalKecerdasanMajemukDo) FirstOrCreate() (*entity.SoalKecerdasanMajemuk, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalKecerdasanMajemuk), nil
	}
}

func (s soalKecerdasanMajemukDo) FindByPage(offset int, limit int) (result []*entity.SoalKecerdasanMajemuk, count int64, err error) {
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

func (s soalKecerdasanMajemukDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s soalKecerdasanMajemukDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s soalKecerdasanMajemukDo) Delete(models ...*entity.SoalKecerdasanMajemuk) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *soalKecerdasanMajemukDo) withDO(do gen.Dao) *soalKecerdasanMajemukDo {
	s.DO = *do.(*gen.DO)
	return s
}
