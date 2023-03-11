package template

var (
	RepositoryMysqlSRV = `package mysql

import (
	"context"
	"github.com/go-xorm/xorm"
	"{{.Dir}}/domain"
)

type mysql{{title .Alias}}Repository struct {
	dbEngine *xorm.Engine
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysql{{title .Alias}}Repository() domain.{{title .Alias}}Repository {
    e, _ := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	return &mysql{{title .Alias}}Repository{
		dbEngine: e
	}
}

func (m *mysql{{title .Alias}}Repository) Get(ctx context.Context, {{.Alias}} *{{title .Alias}}) (err error) {
	_, err = m.dbEngine.Get({{.Alias}})
	return
}

func (m *mysql{{title .Alias}}Repository) GetList(ctx context.Context, {{.Alias}} *{{title .Alias}}) (list []*{{title .Alias}}, err error) {
	err = m.dbEngine.Find(&list)
	return
}

func (m *mysql{{title .Alias}}Repository) Create(ctx context.Context, {{.Alias}} *{{title .Alias}}) (err error) {
	_, err = m.dbEngine.InsertOne({{.Alias}})
	return
}

func (m *mysql{{title .Alias}}Repository) Update(ctx context.Context, {{.Alias}} *{{title .Alias}}, condition *{{title .Alias}}) (int, error) {
	rows, err = m.dbEngine.Update({{.Alias}}, condition)
	return
}`
)
