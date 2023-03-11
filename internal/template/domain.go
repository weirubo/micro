package template

var (
	DomainSRV = `package domain

import (
	"context"
	"time"
)

// Article is representing the Article data struct
type {{title .Alias}} struct {
	ID        int64 
	Title     string 
	Content   string 
	Author    Author 
	UpdatedAt time.Time 
	CreatedAt time.Time 
}

// ArticleUsecase represent the article's usecases
type {{.Alias}}Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]{{.Alias}}, string, error)
	GetByID(ctx context.Context, id int64) ({{.Alias}}, error)
	Update(ctx context.Context, ar *{{.Alias}}) error
	GetByTitle(ctx context.Context, title string) ({{.Alias}}, error)
	Store(context.Context, *{{.Alias}}) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type {{.Alias}}Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []{{.Alias}}, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) ({{.Alias}}, error)
	GetByTitle(ctx context.Context, title string) ({{.Alias}}, error)
	Update(ctx context.Context, ar *{{.Alias}}) error
	Store(ctx context.Context, a *{{.Alias}}) error
	Delete(ctx context.Context, id int64) error
}`
)
