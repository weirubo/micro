package template

var (
	UsecaseSRV = `package usecase

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"{{.Dir}}/domain"
)

type {{.Alias}}Usecase struct {
	{{.Alias}}Repo    domain.{{title .Alias}}Repository
	authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func New{{title .Alias}}Usecase(a domain.{{title .Alias}}Repository, ar domain.AuthorRepository, timeout time.Duration) domain.{{title .Alias}}Usecase {
	return &{{.Alias}}Usecase{
		{{.Alias}}Repo:    a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (a *{{.Alias}}Usecase) fillAuthorDetails(c context.Context, data []domain.{{title .Alias}}) ([]domain.{{title .Alias}}, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the author's id
	mapAuthors := map[int64]domain.Author{}

	for _, {{.Alias}} := range data { //nolint
		mapAuthors[{{.Alias}}.Author.ID] = domain.Author{}
	}
	// Using goroutine to fetch the author's detail
	chanAuthor := make(chan domain.Author)
	for authorID := range mapAuthors {
		authorID := authorID
		g.Go(func() error {
			res, err := a.authorRepo.GetByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanAuthor <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanAuthor)
	}()

	for author := range chanAuthor {
		if author != (domain.Author{}) {
			mapAuthors[author.ID] = author
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the author's data
	for index, item := range data { //nolint
		if a, ok := mapAuthors[item.Author.ID]; ok {
			data[index].Author = a
		}
	}
	return data, nil
}

func (a *{{.Alias}}Usecase) Fetch(c context.Context, cursor string, num int64) (res []domain.{{title .Alias}}, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.{{.Alias}}Repo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	res, err = a.fillAuthorDetails(ctx, res)
	if err != nil {
		nextCursor = ""
	}
	return
}

func (a *{{.Alias}}Usecase) GetByID(c context.Context, id int64) (res domain.{{title .Alias}}, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.{{.Alias}}Repo.GetByID(ctx, id)
	if err != nil {
		return
	}

	resAuthor, err := a.authorRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return domain.{{title .Alias}}{}, err
	}
	res.Author = resAuthor
	return
}

func (a *{{.Alias}}Usecase) Update(c context.Context, ar *domain.{{title .Alias}}) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.{{.Alias}}Repo.Update(ctx, ar)
}

func (a *{{.Alias}}Usecase) GetByTitle(c context.Context, title string) (res domain.{{title .Alias}}, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.{{.Alias}}Repo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	resAuthor, err := a.authorRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return domain.{{title .Alias}}{}, err
	}

	res.Author = resAuthor
	return
}

func (a *{{.Alias}}Usecase) Store(c context.Context, m *domain.{{title .Alias}}) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existed{{title .Alias}}, _ := a.GetByTitle(ctx, m.Title)
	if existed{{title .Alias}} != (domain.{{title .Alias}}{}) {
		return domain.ErrConflict
	}

	err = a.{{.Alias}}Repo.Store(ctx, m)
	return
}

func (a *{{.Alias}}Usecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existed{{title .Alias}}, err := a.{{.Alias}}Repo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existed{{title .Alias}} == (domain.{{title .Alias}}{}) {
		return domain.ErrNotFound
	}
	return a.{{.Alias}}Repo.Delete(ctx, id)
}`
)
