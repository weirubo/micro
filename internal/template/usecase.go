package template

var (
	UsecaseSRV = `package usecase

import (
	"context"
	"{{.Dir}}/domain"
)

type {{.Alias}}Usecase struct {
	{{.Alias}}Repo domain.{{title .Alias}}Repository
}

// New{{title .Alias}}Usecase will create new an {{.Alias}}Usecase object representation of domain.{{title .Alias}}Usecase interface
func New{{title .Alias}}Usecase({{.Alias}}Repo domain.{{title .Alias}}Repository) domain.{{title .Alias}}Usecase {
	return &{{.Alias}}Usecase{
		{{.Alias}}Repo: {{.Alias}}Repo,
	}
}

func (u *{{.Alias}}Usecase) Get(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (err error) {
	err = u.{{.Alias}}Repo.Get(ctx, {{.Alias}})
	return
}

func (u *{{.Alias}}Usecase) GetList(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (list []*domain.{{title .Alias}}, err error) {
	list, err = u.{{.Alias}}Repo.GetList(ctx, {{.Alias}})
	return
}

func (u *{{.Alias}}Usecase) Create(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}) (err error) {
	err = u.{{.Alias}}Repo.Create(ctx, {{.Alias}})
	return
}

func (u *{{.Alias}}Usecase) Update(ctx context.Context, {{.Alias}} *domain.{{title .Alias}}, condition *domain.{{title .Alias}}) (rows int64, err error) {
	rows, err = u.{{.Alias}}Repo.Update(ctx, {{.Alias}}, condition)
	return
}`
)
