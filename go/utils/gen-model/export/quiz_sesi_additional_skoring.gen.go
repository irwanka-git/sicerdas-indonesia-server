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

func newQuizSesiAdditionalSkoring(db *gorm.DB, opts ...gen.DOOption) quizSesiAdditionalSkoring {
	_quizSesiAdditionalSkoring := quizSesiAdditionalSkoring{}

	_quizSesiAdditionalSkoring.quizSesiAdditionalSkoringDo.UseDB(db, opts...)
	_quizSesiAdditionalSkoring.quizSesiAdditionalSkoringDo.UseModel(&entity.QuizSesiAdditionalSkoring{})

	tableName := _quizSesiAdditionalSkoring.quizSesiAdditionalSkoringDo.TableName()
	_quizSesiAdditionalSkoring.ALL = field.NewAsterisk(tableName)
	_quizSesiAdditionalSkoring.ID = field.NewInt32(tableName, "id")
	_quizSesiAdditionalSkoring.NamaSkoring = field.NewString(tableName, "nama_skoring")
	_quizSesiAdditionalSkoring.DependTabel = field.NewString(tableName, "depend_tabel")

	_quizSesiAdditionalSkoring.fillFieldMap()

	return _quizSesiAdditionalSkoring
}

type quizSesiAdditionalSkoring struct {
	quizSesiAdditionalSkoringDo quizSesiAdditionalSkoringDo

	ALL         field.Asterisk
	ID          field.Int32
	NamaSkoring field.String
	DependTabel field.String

	fieldMap map[string]field.Expr
}

func (q quizSesiAdditionalSkoring) Table(newTableName string) *quizSesiAdditionalSkoring {
	q.quizSesiAdditionalSkoringDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q quizSesiAdditionalSkoring) As(alias string) *quizSesiAdditionalSkoring {
	q.quizSesiAdditionalSkoringDo.DO = *(q.quizSesiAdditionalSkoringDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *quizSesiAdditionalSkoring) updateTableName(table string) *quizSesiAdditionalSkoring {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt32(table, "id")
	q.NamaSkoring = field.NewString(table, "nama_skoring")
	q.DependTabel = field.NewString(table, "depend_tabel")

	q.fillFieldMap()

	return q
}

func (q *quizSesiAdditionalSkoring) WithContext(ctx context.Context) *quizSesiAdditionalSkoringDo {
	return q.quizSesiAdditionalSkoringDo.WithContext(ctx)
}

func (q quizSesiAdditionalSkoring) TableName() string {
	return q.quizSesiAdditionalSkoringDo.TableName()
}

func (q quizSesiAdditionalSkoring) Alias() string { return q.quizSesiAdditionalSkoringDo.Alias() }

func (q quizSesiAdditionalSkoring) Columns(cols ...field.Expr) gen.Columns {
	return q.quizSesiAdditionalSkoringDo.Columns(cols...)
}

func (q *quizSesiAdditionalSkoring) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *quizSesiAdditionalSkoring) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 3)
	q.fieldMap["id"] = q.ID
	q.fieldMap["nama_skoring"] = q.NamaSkoring
	q.fieldMap["depend_tabel"] = q.DependTabel
}

func (q quizSesiAdditionalSkoring) clone(db *gorm.DB) quizSesiAdditionalSkoring {
	q.quizSesiAdditionalSkoringDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q quizSesiAdditionalSkoring) replaceDB(db *gorm.DB) quizSesiAdditionalSkoring {
	q.quizSesiAdditionalSkoringDo.ReplaceDB(db)
	return q
}

type quizSesiAdditionalSkoringDo struct{ gen.DO }

func (q quizSesiAdditionalSkoringDo) Debug() *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Debug())
}

func (q quizSesiAdditionalSkoringDo) WithContext(ctx context.Context) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q quizSesiAdditionalSkoringDo) ReadDB() *quizSesiAdditionalSkoringDo {
	return q.Clauses(dbresolver.Read)
}

func (q quizSesiAdditionalSkoringDo) WriteDB() *quizSesiAdditionalSkoringDo {
	return q.Clauses(dbresolver.Write)
}

func (q quizSesiAdditionalSkoringDo) Session(config *gorm.Session) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Session(config))
}

func (q quizSesiAdditionalSkoringDo) Clauses(conds ...clause.Expression) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q quizSesiAdditionalSkoringDo) Returning(value interface{}, columns ...string) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q quizSesiAdditionalSkoringDo) Not(conds ...gen.Condition) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q quizSesiAdditionalSkoringDo) Or(conds ...gen.Condition) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q quizSesiAdditionalSkoringDo) Select(conds ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q quizSesiAdditionalSkoringDo) Where(conds ...gen.Condition) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q quizSesiAdditionalSkoringDo) Order(conds ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q quizSesiAdditionalSkoringDo) Distinct(cols ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q quizSesiAdditionalSkoringDo) Omit(cols ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q quizSesiAdditionalSkoringDo) Join(table schema.Tabler, on ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q quizSesiAdditionalSkoringDo) LeftJoin(table schema.Tabler, on ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q quizSesiAdditionalSkoringDo) RightJoin(table schema.Tabler, on ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q quizSesiAdditionalSkoringDo) Group(cols ...field.Expr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q quizSesiAdditionalSkoringDo) Having(conds ...gen.Condition) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q quizSesiAdditionalSkoringDo) Limit(limit int) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q quizSesiAdditionalSkoringDo) Offset(offset int) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q quizSesiAdditionalSkoringDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q quizSesiAdditionalSkoringDo) Unscoped() *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Unscoped())
}

func (q quizSesiAdditionalSkoringDo) Create(values ...*entity.QuizSesiAdditionalSkoring) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q quizSesiAdditionalSkoringDo) CreateInBatches(values []*entity.QuizSesiAdditionalSkoring, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q quizSesiAdditionalSkoringDo) Save(values ...*entity.QuizSesiAdditionalSkoring) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q quizSesiAdditionalSkoringDo) First() (*entity.QuizSesiAdditionalSkoring, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiAdditionalSkoring), nil
	}
}

func (q quizSesiAdditionalSkoringDo) Take() (*entity.QuizSesiAdditionalSkoring, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiAdditionalSkoring), nil
	}
}

func (q quizSesiAdditionalSkoringDo) Last() (*entity.QuizSesiAdditionalSkoring, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiAdditionalSkoring), nil
	}
}

func (q quizSesiAdditionalSkoringDo) Find() ([]*entity.QuizSesiAdditionalSkoring, error) {
	result, err := q.DO.Find()
	return result.([]*entity.QuizSesiAdditionalSkoring), err
}

func (q quizSesiAdditionalSkoringDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.QuizSesiAdditionalSkoring, err error) {
	buf := make([]*entity.QuizSesiAdditionalSkoring, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q quizSesiAdditionalSkoringDo) FindInBatches(result *[]*entity.QuizSesiAdditionalSkoring, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q quizSesiAdditionalSkoringDo) Attrs(attrs ...field.AssignExpr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q quizSesiAdditionalSkoringDo) Assign(attrs ...field.AssignExpr) *quizSesiAdditionalSkoringDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q quizSesiAdditionalSkoringDo) Joins(fields ...field.RelationField) *quizSesiAdditionalSkoringDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q quizSesiAdditionalSkoringDo) Preload(fields ...field.RelationField) *quizSesiAdditionalSkoringDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q quizSesiAdditionalSkoringDo) FirstOrInit() (*entity.QuizSesiAdditionalSkoring, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiAdditionalSkoring), nil
	}
}

func (q quizSesiAdditionalSkoringDo) FirstOrCreate() (*entity.QuizSesiAdditionalSkoring, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.QuizSesiAdditionalSkoring), nil
	}
}

func (q quizSesiAdditionalSkoringDo) FindByPage(offset int, limit int) (result []*entity.QuizSesiAdditionalSkoring, count int64, err error) {
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

func (q quizSesiAdditionalSkoringDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q quizSesiAdditionalSkoringDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q quizSesiAdditionalSkoringDo) Delete(models ...*entity.QuizSesiAdditionalSkoring) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *quizSesiAdditionalSkoringDo) withDO(do gen.Dao) *quizSesiAdditionalSkoringDo {
	q.DO = *do.(*gen.DO)
	return q
}
