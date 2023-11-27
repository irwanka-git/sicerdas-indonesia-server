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

func newSoalModeKerjaEng(db *gorm.DB, opts ...gen.DOOption) soalModeKerjaEng {
	_soalModeKerjaEng := soalModeKerjaEng{}

	_soalModeKerjaEng.soalModeKerjaEngDo.UseDB(db, opts...)
	_soalModeKerjaEng.soalModeKerjaEngDo.UseModel(&entity.SoalModeKerjaEng{})

	tableName := _soalModeKerjaEng.soalModeKerjaEngDo.TableName()
	_soalModeKerjaEng.ALL = field.NewAsterisk(tableName)
	_soalModeKerjaEng.Urutan = field.NewInt16(tableName, "urutan")
	_soalModeKerjaEng.Soal = field.NewString(tableName, "soal")
	_soalModeKerjaEng.Deskripsi = field.NewString(tableName, "deskripsi")
	_soalModeKerjaEng.UUID = field.NewString(tableName, "uuid")
	_soalModeKerjaEng.PilihanA = field.NewString(tableName, "pilihan_a")
	_soalModeKerjaEng.PilihanB = field.NewString(tableName, "pilihan_b")
	_soalModeKerjaEng.PilihanC = field.NewString(tableName, "pilihan_c")
	_soalModeKerjaEng.PilihanD = field.NewString(tableName, "pilihan_d")
	_soalModeKerjaEng.PilihanE = field.NewString(tableName, "pilihan_e")
	_soalModeKerjaEng.ModeKerja = field.NewString(tableName, "mode_kerja")
	_soalModeKerjaEng.Kelompok = field.NewString(tableName, "kelompok")

	_soalModeKerjaEng.fillFieldMap()

	return _soalModeKerjaEng
}

type soalModeKerjaEng struct {
	soalModeKerjaEngDo soalModeKerjaEngDo

	ALL       field.Asterisk
	Urutan    field.Int16
	Soal      field.String
	Deskripsi field.String
	UUID      field.String
	PilihanA  field.String
	PilihanB  field.String
	PilihanC  field.String
	PilihanD  field.String
	PilihanE  field.String
	ModeKerja field.String
	Kelompok  field.String

	fieldMap map[string]field.Expr
}

func (s soalModeKerjaEng) Table(newTableName string) *soalModeKerjaEng {
	s.soalModeKerjaEngDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s soalModeKerjaEng) As(alias string) *soalModeKerjaEng {
	s.soalModeKerjaEngDo.DO = *(s.soalModeKerjaEngDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *soalModeKerjaEng) updateTableName(table string) *soalModeKerjaEng {
	s.ALL = field.NewAsterisk(table)
	s.Urutan = field.NewInt16(table, "urutan")
	s.Soal = field.NewString(table, "soal")
	s.Deskripsi = field.NewString(table, "deskripsi")
	s.UUID = field.NewString(table, "uuid")
	s.PilihanA = field.NewString(table, "pilihan_a")
	s.PilihanB = field.NewString(table, "pilihan_b")
	s.PilihanC = field.NewString(table, "pilihan_c")
	s.PilihanD = field.NewString(table, "pilihan_d")
	s.PilihanE = field.NewString(table, "pilihan_e")
	s.ModeKerja = field.NewString(table, "mode_kerja")
	s.Kelompok = field.NewString(table, "kelompok")

	s.fillFieldMap()

	return s
}

func (s *soalModeKerjaEng) WithContext(ctx context.Context) *soalModeKerjaEngDo {
	return s.soalModeKerjaEngDo.WithContext(ctx)
}

func (s soalModeKerjaEng) TableName() string { return s.soalModeKerjaEngDo.TableName() }

func (s soalModeKerjaEng) Alias() string { return s.soalModeKerjaEngDo.Alias() }

func (s soalModeKerjaEng) Columns(cols ...field.Expr) gen.Columns {
	return s.soalModeKerjaEngDo.Columns(cols...)
}

func (s *soalModeKerjaEng) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *soalModeKerjaEng) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["urutan"] = s.Urutan
	s.fieldMap["soal"] = s.Soal
	s.fieldMap["deskripsi"] = s.Deskripsi
	s.fieldMap["uuid"] = s.UUID
	s.fieldMap["pilihan_a"] = s.PilihanA
	s.fieldMap["pilihan_b"] = s.PilihanB
	s.fieldMap["pilihan_c"] = s.PilihanC
	s.fieldMap["pilihan_d"] = s.PilihanD
	s.fieldMap["pilihan_e"] = s.PilihanE
	s.fieldMap["mode_kerja"] = s.ModeKerja
	s.fieldMap["kelompok"] = s.Kelompok
}

func (s soalModeKerjaEng) clone(db *gorm.DB) soalModeKerjaEng {
	s.soalModeKerjaEngDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s soalModeKerjaEng) replaceDB(db *gorm.DB) soalModeKerjaEng {
	s.soalModeKerjaEngDo.ReplaceDB(db)
	return s
}

type soalModeKerjaEngDo struct{ gen.DO }

func (s soalModeKerjaEngDo) Debug() *soalModeKerjaEngDo {
	return s.withDO(s.DO.Debug())
}

func (s soalModeKerjaEngDo) WithContext(ctx context.Context) *soalModeKerjaEngDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s soalModeKerjaEngDo) ReadDB() *soalModeKerjaEngDo {
	return s.Clauses(dbresolver.Read)
}

func (s soalModeKerjaEngDo) WriteDB() *soalModeKerjaEngDo {
	return s.Clauses(dbresolver.Write)
}

func (s soalModeKerjaEngDo) Session(config *gorm.Session) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Session(config))
}

func (s soalModeKerjaEngDo) Clauses(conds ...clause.Expression) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s soalModeKerjaEngDo) Returning(value interface{}, columns ...string) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s soalModeKerjaEngDo) Not(conds ...gen.Condition) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s soalModeKerjaEngDo) Or(conds ...gen.Condition) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s soalModeKerjaEngDo) Select(conds ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s soalModeKerjaEngDo) Where(conds ...gen.Condition) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s soalModeKerjaEngDo) Order(conds ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s soalModeKerjaEngDo) Distinct(cols ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s soalModeKerjaEngDo) Omit(cols ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s soalModeKerjaEngDo) Join(table schema.Tabler, on ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s soalModeKerjaEngDo) LeftJoin(table schema.Tabler, on ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s soalModeKerjaEngDo) RightJoin(table schema.Tabler, on ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s soalModeKerjaEngDo) Group(cols ...field.Expr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s soalModeKerjaEngDo) Having(conds ...gen.Condition) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s soalModeKerjaEngDo) Limit(limit int) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s soalModeKerjaEngDo) Offset(offset int) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s soalModeKerjaEngDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s soalModeKerjaEngDo) Unscoped() *soalModeKerjaEngDo {
	return s.withDO(s.DO.Unscoped())
}

func (s soalModeKerjaEngDo) Create(values ...*entity.SoalModeKerjaEng) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s soalModeKerjaEngDo) CreateInBatches(values []*entity.SoalModeKerjaEng, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s soalModeKerjaEngDo) Save(values ...*entity.SoalModeKerjaEng) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s soalModeKerjaEngDo) First() (*entity.SoalModeKerjaEng, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalModeKerjaEng), nil
	}
}

func (s soalModeKerjaEngDo) Take() (*entity.SoalModeKerjaEng, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalModeKerjaEng), nil
	}
}

func (s soalModeKerjaEngDo) Last() (*entity.SoalModeKerjaEng, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalModeKerjaEng), nil
	}
}

func (s soalModeKerjaEngDo) Find() ([]*entity.SoalModeKerjaEng, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SoalModeKerjaEng), err
}

func (s soalModeKerjaEngDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SoalModeKerjaEng, err error) {
	buf := make([]*entity.SoalModeKerjaEng, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s soalModeKerjaEngDo) FindInBatches(result *[]*entity.SoalModeKerjaEng, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s soalModeKerjaEngDo) Attrs(attrs ...field.AssignExpr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s soalModeKerjaEngDo) Assign(attrs ...field.AssignExpr) *soalModeKerjaEngDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s soalModeKerjaEngDo) Joins(fields ...field.RelationField) *soalModeKerjaEngDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s soalModeKerjaEngDo) Preload(fields ...field.RelationField) *soalModeKerjaEngDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s soalModeKerjaEngDo) FirstOrInit() (*entity.SoalModeKerjaEng, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalModeKerjaEng), nil
	}
}

func (s soalModeKerjaEngDo) FirstOrCreate() (*entity.SoalModeKerjaEng, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SoalModeKerjaEng), nil
	}
}

func (s soalModeKerjaEngDo) FindByPage(offset int, limit int) (result []*entity.SoalModeKerjaEng, count int64, err error) {
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

func (s soalModeKerjaEngDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s soalModeKerjaEngDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s soalModeKerjaEngDo) Delete(models ...*entity.SoalModeKerjaEng) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *soalModeKerjaEngDo) withDO(do gen.Dao) *soalModeKerjaEngDo {
	s.DO = *do.(*gen.DO)
	return s
}