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

func newQuizSesiMaster(db *gorm.DB, opts ...gen.DOOption) quizSesiMaster {
	_quizSesiMaster := quizSesiMaster{}

	_quizSesiMaster.quizSesiMasterDo.UseDB(db, opts...)
	_quizSesiMaster.quizSesiMasterDo.UseModel(&entity.QuizSesiMaster{})

	tableName := _quizSesiMaster.quizSesiMasterDo.TableName()
	_quizSesiMaster.ALL = field.NewAsterisk(tableName)
	_quizSesiMaster.IDSesiMaster = field.NewInt32(tableName, "id_sesi_master")
	_quizSesiMaster.Kategori = field.NewString(tableName, "kategori")
	_quizSesiMaster.NamaSesiUjian = field.NewString(tableName, "nama_sesi_ujian")
	_quizSesiMaster.Soal = field.NewString(tableName, "soal")
	_quizSesiMaster.Mode = field.NewString(tableName, "mode")
	_quizSesiMaster.Jawaban = field.NewInt32(tableName, "jawaban")
	_quizSesiMaster.PetunjukSesi = field.NewString(tableName, "petunjuk_sesi")
	_quizSesiMaster.MetodeSkoring = field.NewString(tableName, "metode_skoring")
	_quizSesiMaster.UUID = field.NewString(tableName, "uuid")

	_quizSesiMaster.fillFieldMap()

	return _quizSesiMaster
}

type quizSesiMaster struct {
	quizSesiMasterDo quizSesiMasterDo

	ALL           field.Asterisk
	IDSesiMaster  field.Int32
	Kategori      field.String
	NamaSesiUjian field.String
	Soal          field.String
	Mode          field.String
	Jawaban       field.Int32 // Jumlah Jawaban Ynag harus diisi / dikoreksi
	PetunjukSesi  field.String
	MetodeSkoring field.String
	UUID          field.String

	fieldMap map[string]field.Expr
}

func (q quizSesiMaster) Table(newTableName string) *quizSesiMaster {
	q.quizSesiMasterDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q quizSesiMaster) As(alias string) *quizSesiMaster {
	q.quizSesiMasterDo.DO = *(q.quizSesiMasterDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *quizSesiMaster) updateTableName(table string) *quizSesiMaster {
	q.ALL = field.NewAsterisk(table)
	q.IDSesiMaster = field.NewInt32(table, "id_sesi_master")
	q.Kategori = field.NewString(table, "kategori")
	q.NamaSesiUjian = field.NewString(table, "nama_sesi_ujian")
	q.Soal = field.NewString(table, "soal")
	q.Mode = field.NewString(table, "mode")
	q.Jawaban = field.NewInt32(table, "jawaban")
	q.PetunjukSesi = field.NewString(table, "petunjuk_sesi")
	q.MetodeSkoring = field.NewString(table, "metode_skoring")
	q.UUID = field.NewString(table, "uuid")

	q.fillFieldMap()

	return q
}

func (q *quizSesiMaster) WithContext(ctx context.Context) *quizSesiMasterDo {
	return q.quizSesiMasterDo.WithContext(ctx)
}

func (q quizSesiMaster) TableName() string { return q.quizSesiMasterDo.TableName() }

func (q quizSesiMaster) Alias() string { return q.quizSesiMasterDo.Alias() }

func (q quizSesiMaster) Columns(cols ...field.Expr) gen.Columns {
	return q.quizSesiMasterDo.Columns(cols...)
}

func (q *quizSesiMaster) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *quizSesiMaster) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 9)
	q.fieldMap["id_sesi_master"] = q.IDSesiMaster
	q.fieldMap["kategori"] = q.Kategori
	q.fieldMap["nama_sesi_ujian"] = q.NamaSesiUjian
	q.fieldMap["soal"] = q.Soal
	q.fieldMap["mode"] = q.Mode
	q.fieldMap["jawaban"] = q.Jawaban
	q.fieldMap["petunjuk_sesi"] = q.PetunjukSesi
	q.fieldMap["metode_skoring"] = q.MetodeSkoring
	q.fieldMap["uuid"] = q.UUID
}

func (q quizSesiMaster) clone(db *gorm.DB) quizSesiMaster {
	q.quizSesiMasterDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q quizSesiMaster) replaceDB(db *gorm.DB) quizSesiMaster {
	q.quizSesiMasterDo.ReplaceDB(db)
	return q
}

type quizSesiMasterDo struct{ gen.DO }

func (q quizSesiMasterDo) Debug() *quizSesiMasterDo {
	return q.withDO(q.DO.Debug())
}

func (q quizSesiMasterDo) WithContext(ctx context.Context) *quizSesiMasterDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q quizSesiMasterDo) ReadDB() *quizSesiMasterDo {
	return q.Clauses(dbresolver.Read)
}

func (q quizSesiMasterDo) WriteDB() *quizSesiMasterDo {
	return q.Clauses(dbresolver.Write)
}

func (q quizSesiMasterDo) Session(config *gorm.Session) *quizSesiMasterDo {
	return q.withDO(q.DO.Session(config))
}

func (q quizSesiMasterDo) Clauses(conds ...clause.Expression) *quizSesiMasterDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q quizSesiMasterDo) Returning(value interface{}, columns ...string) *quizSesiMasterDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q quizSesiMasterDo) Not(conds ...gen.Condition) *quizSesiMasterDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q quizSesiMasterDo) Or(conds ...gen.Condition) *quizSesiMasterDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q quizSesiMasterDo) Select(conds ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q quizSesiMasterDo) Where(conds ...gen.Condition) *quizSesiMasterDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q quizSesiMasterDo) Order(conds ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q quizSesiMasterDo) Distinct(cols ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q quizSesiMasterDo) Omit(cols ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q quizSesiMasterDo) Join(table schema.Tabler, on ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q quizSesiMasterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q quizSesiMasterDo) RightJoin(table schema.Tabler, on ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q quizSesiMasterDo) Group(cols ...field.Expr) *quizSesiMasterDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q quizSesiMasterDo) Having(conds ...gen.Condition) *quizSesiMasterDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q quizSesiMasterDo) Limit(limit int) *quizSesiMasterDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q quizSesiMasterDo) Offset(offset int) *quizSesiMasterDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q quizSesiMasterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *quizSesiMasterDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q quizSesiMasterDo) Unscoped() *quizSesiMasterDo {
	return q.withDO(q.DO.Unscoped())
}

func (q quizSesiMasterDo) Create(values ...*entity.QuizSesiMaster) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q quizSesiMasterDo) CreateInBatches(values []*entity.QuizSesiMaster, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q quizSesiMasterDo) Save(values ...*entity.QuizSesiMaster) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q quizSesiMasterDo) First() (*entity.QuizSesiMaster, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiMaster), nil
	}
}

func (q quizSesiMasterDo) Take() (*entity.QuizSesiMaster, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiMaster), nil
	}
}

func (q quizSesiMasterDo) Last() (*entity.QuizSesiMaster, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiMaster), nil
	}
}

func (q quizSesiMasterDo) Find() ([]*entity.QuizSesiMaster, error) {
	result, err := q.DO.Find()
	return result.([]*entity.QuizSesiMaster), err
}

func (q quizSesiMasterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.QuizSesiMaster, err error) {
	buf := make([]*entity.QuizSesiMaster, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q quizSesiMasterDo) FindInBatches(result *[]*entity.QuizSesiMaster, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q quizSesiMasterDo) Attrs(attrs ...field.AssignExpr) *quizSesiMasterDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q quizSesiMasterDo) Assign(attrs ...field.AssignExpr) *quizSesiMasterDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q quizSesiMasterDo) Joins(fields ...field.RelationField) *quizSesiMasterDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q quizSesiMasterDo) Preload(fields ...field.RelationField) *quizSesiMasterDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q quizSesiMasterDo) FirstOrInit() (*entity.QuizSesiMaster, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiMaster), nil
	}
}

func (q quizSesiMasterDo) FirstOrCreate() (*entity.QuizSesiMaster, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiMaster), nil
	}
}

func (q quizSesiMasterDo) FindByPage(offset int, limit int) (result []*entity.QuizSesiMaster, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q quizSesiMasterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q quizSesiMasterDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q quizSesiMasterDo) Delete(models ...*entity.QuizSesiMaster) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *quizSesiMasterDo) withDO(do gen.Dao) *quizSesiMasterDo {
	q.DO = *do.(*gen.DO)
	return q
}