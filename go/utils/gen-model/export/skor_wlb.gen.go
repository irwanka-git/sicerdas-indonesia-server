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

func newSkorWlb(db *gorm.DB, opts ...gen.DOOption) skorWlb {
	_skorWlb := skorWlb{}

	_skorWlb.skorWlbDo.UseDB(db, opts...)
	_skorWlb.skorWlbDo.UseModel(&entity.SkorWlb{})

	tableName := _skorWlb.skorWlbDo.TableName()
	_skorWlb.ALL = field.NewAsterisk(tableName)
	_skorWlb.IDUser = field.NewInt32(tableName, "id_user")
	_skorWlb.IDQuiz = field.NewInt32(tableName, "id_quiz")
	_skorWlb.KedamaianHati = field.NewInt32(tableName, "kedamaian_hati")
	_skorWlb.PengembanganDiri = field.NewInt32(tableName, "pengembangan_diri")
	_skorWlb.Ibadah = field.NewInt32(tableName, "ibadah")
	_skorWlb.Pendapatan = field.NewInt32(tableName, "pendapatan")
	_skorWlb.HubunganSosial = field.NewInt32(tableName, "hubungan_sosial")
	_skorWlb.Kesehatan = field.NewInt32(tableName, "kesehatan")
	_skorWlb.RekanKerja = field.NewInt32(tableName, "rekan_kerja")
	_skorWlb.NilaiKedamaianHati = field.NewInt32(tableName, "nilai_kedamaian_hati")
	_skorWlb.NilaiPengembanganDiri = field.NewInt32(tableName, "nilai_pengembangan_diri")
	_skorWlb.NilaiIbadah = field.NewInt32(tableName, "nilai_ibadah")
	_skorWlb.NilaiPendapatan = field.NewInt32(tableName, "nilai_pendapatan")
	_skorWlb.NilaiHubunganSosial = field.NewInt32(tableName, "nilai_hubungan_sosial")
	_skorWlb.NilaiKesehatan = field.NewInt32(tableName, "nilai_kesehatan")
	_skorWlb.NilaiRekanKerja = field.NewInt32(tableName, "nilai_rekan_kerja")

	_skorWlb.fillFieldMap()

	return _skorWlb
}

type skorWlb struct {
	skorWlbDo skorWlbDo

	ALL                   field.Asterisk
	IDUser                field.Int32
	IDQuiz                field.Int32
	KedamaianHati         field.Int32
	PengembanganDiri      field.Int32
	Ibadah                field.Int32
	Pendapatan            field.Int32
	HubunganSosial        field.Int32
	Kesehatan             field.Int32
	RekanKerja            field.Int32
	NilaiKedamaianHati    field.Int32
	NilaiPengembanganDiri field.Int32
	NilaiIbadah           field.Int32
	NilaiPendapatan       field.Int32
	NilaiHubunganSosial   field.Int32
	NilaiKesehatan        field.Int32
	NilaiRekanKerja       field.Int32

	fieldMap map[string]field.Expr
}

func (s skorWlb) Table(newTableName string) *skorWlb {
	s.skorWlbDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s skorWlb) As(alias string) *skorWlb {
	s.skorWlbDo.DO = *(s.skorWlbDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *skorWlb) updateTableName(table string) *skorWlb {
	s.ALL = field.NewAsterisk(table)
	s.IDUser = field.NewInt32(table, "id_user")
	s.IDQuiz = field.NewInt32(table, "id_quiz")
	s.KedamaianHati = field.NewInt32(table, "kedamaian_hati")
	s.PengembanganDiri = field.NewInt32(table, "pengembangan_diri")
	s.Ibadah = field.NewInt32(table, "ibadah")
	s.Pendapatan = field.NewInt32(table, "pendapatan")
	s.HubunganSosial = field.NewInt32(table, "hubungan_sosial")
	s.Kesehatan = field.NewInt32(table, "kesehatan")
	s.RekanKerja = field.NewInt32(table, "rekan_kerja")
	s.NilaiKedamaianHati = field.NewInt32(table, "nilai_kedamaian_hati")
	s.NilaiPengembanganDiri = field.NewInt32(table, "nilai_pengembangan_diri")
	s.NilaiIbadah = field.NewInt32(table, "nilai_ibadah")
	s.NilaiPendapatan = field.NewInt32(table, "nilai_pendapatan")
	s.NilaiHubunganSosial = field.NewInt32(table, "nilai_hubungan_sosial")
	s.NilaiKesehatan = field.NewInt32(table, "nilai_kesehatan")
	s.NilaiRekanKerja = field.NewInt32(table, "nilai_rekan_kerja")

	s.fillFieldMap()

	return s
}

func (s *skorWlb) WithContext(ctx context.Context) *skorWlbDo { return s.skorWlbDo.WithContext(ctx) }

func (s skorWlb) TableName() string { return s.skorWlbDo.TableName() }

func (s skorWlb) Alias() string { return s.skorWlbDo.Alias() }

func (s skorWlb) Columns(cols ...field.Expr) gen.Columns { return s.skorWlbDo.Columns(cols...) }

func (s *skorWlb) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *skorWlb) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["id_user"] = s.IDUser
	s.fieldMap["id_quiz"] = s.IDQuiz
	s.fieldMap["kedamaian_hati"] = s.KedamaianHati
	s.fieldMap["pengembangan_diri"] = s.PengembanganDiri
	s.fieldMap["ibadah"] = s.Ibadah
	s.fieldMap["pendapatan"] = s.Pendapatan
	s.fieldMap["hubungan_sosial"] = s.HubunganSosial
	s.fieldMap["kesehatan"] = s.Kesehatan
	s.fieldMap["rekan_kerja"] = s.RekanKerja
	s.fieldMap["nilai_kedamaian_hati"] = s.NilaiKedamaianHati
	s.fieldMap["nilai_pengembangan_diri"] = s.NilaiPengembanganDiri
	s.fieldMap["nilai_ibadah"] = s.NilaiIbadah
	s.fieldMap["nilai_pendapatan"] = s.NilaiPendapatan
	s.fieldMap["nilai_hubungan_sosial"] = s.NilaiHubunganSosial
	s.fieldMap["nilai_kesehatan"] = s.NilaiKesehatan
	s.fieldMap["nilai_rekan_kerja"] = s.NilaiRekanKerja
}

func (s skorWlb) clone(db *gorm.DB) skorWlb {
	s.skorWlbDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s skorWlb) replaceDB(db *gorm.DB) skorWlb {
	s.skorWlbDo.ReplaceDB(db)
	return s
}

type skorWlbDo struct{ gen.DO }

func (s skorWlbDo) Debug() *skorWlbDo {
	return s.withDO(s.DO.Debug())
}

func (s skorWlbDo) WithContext(ctx context.Context) *skorWlbDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s skorWlbDo) ReadDB() *skorWlbDo {
	return s.Clauses(dbresolver.Read)
}

func (s skorWlbDo) WriteDB() *skorWlbDo {
	return s.Clauses(dbresolver.Write)
}

func (s skorWlbDo) Session(config *gorm.Session) *skorWlbDo {
	return s.withDO(s.DO.Session(config))
}

func (s skorWlbDo) Clauses(conds ...clause.Expression) *skorWlbDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s skorWlbDo) Returning(value interface{}, columns ...string) *skorWlbDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s skorWlbDo) Not(conds ...gen.Condition) *skorWlbDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s skorWlbDo) Or(conds ...gen.Condition) *skorWlbDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s skorWlbDo) Select(conds ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s skorWlbDo) Where(conds ...gen.Condition) *skorWlbDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s skorWlbDo) Order(conds ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s skorWlbDo) Distinct(cols ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s skorWlbDo) Omit(cols ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s skorWlbDo) Join(table schema.Tabler, on ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s skorWlbDo) LeftJoin(table schema.Tabler, on ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s skorWlbDo) RightJoin(table schema.Tabler, on ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s skorWlbDo) Group(cols ...field.Expr) *skorWlbDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s skorWlbDo) Having(conds ...gen.Condition) *skorWlbDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s skorWlbDo) Limit(limit int) *skorWlbDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s skorWlbDo) Offset(offset int) *skorWlbDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s skorWlbDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *skorWlbDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s skorWlbDo) Unscoped() *skorWlbDo {
	return s.withDO(s.DO.Unscoped())
}

func (s skorWlbDo) Create(values ...*entity.SkorWlb) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s skorWlbDo) CreateInBatches(values []*entity.SkorWlb, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s skorWlbDo) Save(values ...*entity.SkorWlb) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s skorWlbDo) First() (*entity.SkorWlb, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorWlb), nil
	}
}

func (s skorWlbDo) Take() (*entity.SkorWlb, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorWlb), nil
	}
}

func (s skorWlbDo) Last() (*entity.SkorWlb, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorWlb), nil
	}
}

func (s skorWlbDo) Find() ([]*entity.SkorWlb, error) {
	result, err := s.DO.Find()
	return result.([]*entity.SkorWlb), err
}

func (s skorWlbDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.SkorWlb, err error) {
	buf := make([]*entity.SkorWlb, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s skorWlbDo) FindInBatches(result *[]*entity.SkorWlb, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s skorWlbDo) Attrs(attrs ...field.AssignExpr) *skorWlbDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s skorWlbDo) Assign(attrs ...field.AssignExpr) *skorWlbDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s skorWlbDo) Joins(fields ...field.RelationField) *skorWlbDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s skorWlbDo) Preload(fields ...field.RelationField) *skorWlbDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s skorWlbDo) FirstOrInit() (*entity.SkorWlb, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorWlb), nil
	}
}

func (s skorWlbDo) FirstOrCreate() (*entity.SkorWlb, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.SkorWlb), nil
	}
}

func (s skorWlbDo) FindByPage(offset int, limit int) (result []*entity.SkorWlb, count int64, err error) {
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

func (s skorWlbDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s skorWlbDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s skorWlbDo) Delete(models ...*entity.SkorWlb) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *skorWlbDo) withDO(do gen.Dao) *skorWlbDo {
	s.DO = *do.(*gen.DO)
	return s
}
