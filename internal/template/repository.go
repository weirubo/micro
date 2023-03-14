package template

var (
	RepositoryMysqlSRV = `package mysql

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"{{.Dir}}/domain"
)

type mysql{{title .Alias}}Repository struct {
	dbEngine *xorm.Engine
}

// NewMysql{{title .Alias}}Repository will create an object that represent the article.Repository interface
func NewMysql{{title .Alias}}Repository(engine *xorm.Engine) domain.{{title .Alias}}Repository {
	return &mysql{{title .Alias}}Repository{
		dbEngine: engine,
	}
}

func (m *mysql{{title .Alias}}Repository) Get(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (err error) {
	_, err = m.dbEngine.Get({{.Alias}})
	return
}

func (m *mysql{{title .Alias}}Repository) GetList(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (list []*domain.{{title .Alias}}, err error) {
	err = m.dbEngine.Find(&list)
	return
}

func (m *mysql{{title .Alias}}Repository) Create(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (err error) {
	_, err = m.dbEngine.InsertOne({{.Alias}})
	return
}

func (m *mysql{{title .Alias}}Repository) Update(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}, condition *domain.{{title .Alias}}) (rows int64, err error) {
	rows, err = m.dbEngine.Update({{.Alias}}, condition)
	return
}`
)
