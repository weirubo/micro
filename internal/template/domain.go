package template

var (
	DomainSRV = `package domain

import (
	"context"
	"time"
)

// {{title .Alias}} is representing the Article data struct
type {{title .Alias}} struct {
	Id        int64 
	Title     string 
	UpdatedAt time.Time 
	CreatedAt time.Time 
}

// {{title .Alias}}Usecase represent the article's usecases
type {{title .Alias}}Usecase interface {
	Get(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	GetList(ctx context.Context, {{.Alias}} *{{title .Alias}}) ([]*{{title .Alias}}, error)
	Create(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	Update(ctx context.Context, {{.Alias}} *{{title .Alias}}, condition *{{title .Alias}}) (int64, error)
}

// {{title .Alias}}Repository represent the article's repository contract
type {{title .Alias}}Repository interface {
	Get(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	GetList(ctx context.Context, {{.Alias}} *{{title .Alias}}) ([]*{{title .Alias}}, error)
	Create(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	Update(ctx context.Context, {{.Alias}} *{{title .Alias}}, condition *{{title .Alias}}) (int64, error)
}`
)
