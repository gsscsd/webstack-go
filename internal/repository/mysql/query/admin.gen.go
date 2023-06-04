// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
)

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&model.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewInt64(tableName, "id")
	_admin.Username = field.NewString(tableName, "username")
	_admin.Password = field.NewString(tableName, "password")
	_admin.Nickname = field.NewString(tableName, "nickname")
	_admin.Mobile = field.NewString(tableName, "mobile")
	_admin.IsUsed = field.NewInt64(tableName, "is_used")
	_admin.IsDeleted = field.NewInt64(tableName, "is_deleted")
	_admin.CreatedAt = field.NewTime(tableName, "created_at")
	_admin.CreatedUser = field.NewString(tableName, "created_user")
	_admin.UpdatedAt = field.NewTime(tableName, "updated_at")
	_admin.UpdatedUser = field.NewString(tableName, "updated_user")

	_admin.fillFieldMap()

	return _admin
}

type admin struct {
	adminDo

	ALL         field.Asterisk
	ID          field.Int64  // 主键
	Username    field.String // 用户名
	Password    field.String // 密码
	Nickname    field.String // 昵称
	Mobile      field.String // 手机号
	IsUsed      field.Int64  // 是否启用 1:是  -1:否
	IsDeleted   field.Int64  // 是否删除 1:是  -1:否
	CreatedAt   field.Time   // 创建时间
	CreatedUser field.String // 创建人
	UpdatedAt   field.Time   // 更新时间
	UpdatedUser field.String // 更新人

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.Nickname = field.NewString(table, "nickname")
	a.Mobile = field.NewString(table, "mobile")
	a.IsUsed = field.NewInt64(table, "is_used")
	a.IsDeleted = field.NewInt64(table, "is_deleted")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.CreatedUser = field.NewString(table, "created_user")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.UpdatedUser = field.NewString(table, "updated_user")

	a.fillFieldMap()

	return a
}

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["nickname"] = a.Nickname
	a.fieldMap["mobile"] = a.Mobile
	a.fieldMap["is_used"] = a.IsUsed
	a.fieldMap["is_deleted"] = a.IsDeleted
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["created_user"] = a.CreatedUser
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["updated_user"] = a.UpdatedUser
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

type IAdminDo interface {
	gen.SubQuery
	Debug() IAdminDo
	WithContext(ctx context.Context) IAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminDo
	WriteDB() IAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminDo
	Not(conds ...gen.Condition) IAdminDo
	Or(conds ...gen.Condition) IAdminDo
	Select(conds ...field.Expr) IAdminDo
	Where(conds ...gen.Condition) IAdminDo
	Order(conds ...field.Expr) IAdminDo
	Distinct(cols ...field.Expr) IAdminDo
	Omit(cols ...field.Expr) IAdminDo
	Join(table schema.Tabler, on ...field.Expr) IAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	Group(cols ...field.Expr) IAdminDo
	Having(conds ...gen.Condition) IAdminDo
	Limit(limit int) IAdminDo
	Offset(offset int) IAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo
	Unscoped() IAdminDo
	Create(values ...*model.Admin) error
	CreateInBatches(values []*model.Admin, batchSize int) error
	Save(values ...*model.Admin) error
	First() (*model.Admin, error)
	Take() (*model.Admin, error)
	Last() (*model.Admin, error)
	Find() ([]*model.Admin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error)
	FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Admin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminDo
	Assign(attrs ...field.AssignExpr) IAdminDo
	Joins(fields ...field.RelationField) IAdminDo
	Preload(fields ...field.RelationField) IAdminDo
	FirstOrInit() (*model.Admin, error)
	FirstOrCreate() (*model.Admin, error)
	FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminDo) Debug() IAdminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) IAdminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() IAdminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() IAdminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) IAdminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) IAdminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) IAdminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAdminDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a adminDo) Order(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) IAdminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) IAdminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() IAdminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*model.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*model.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Take() (*model.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Last() (*model.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Find() ([]*model.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*model.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error) {
	buf := make([]*model.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*model.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*model.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*model.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}
