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
type {{title .Alias}}Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]{{title .Alias}}, string, error)
	GetByID(ctx context.Context, id int64) ({{title .Alias}}, error)
	Update(ctx context.Context, ar *{{title .Alias}}) error
	GetByTitle(ctx context.Context, title string) ({{title .Alias}}, error)
	Store(context.Context, *{{title .Alias}}) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type {{title .Alias}}Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []{{title .Alias}}, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) ({{title .Alias}}, error)
	GetByTitle(ctx context.Context, title string) ({{title .Alias}}, error)
	Update(ctx context.Context, ar *{{title .Alias}}) error
	Store(ctx context.Context, a *{{title .Alias}}) error
	Delete(ctx context.Context, id int64) error
}`
)
