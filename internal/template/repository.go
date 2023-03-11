package template

var (
	RepositoryMysqlSRV = `package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"{{.Dir}}/{{.Alias}}/repository"
	"{{.Dir}}/domain"
)

type mysql{{title .Alias}}Repository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysql{{title .Alias}}Repository(conn *sql.DB) domain.{{title .Alias}}Repository {
	return &mysql{{title .Alias}}Repository{conn}
}

func (m *mysql{{title .Alias}}Repository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.{{title .Alias}}, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.{{title .Alias}}, 0)
	for rows.Next() {
		t := domain.{{title .Alias}}{}
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = domain.Author{
			ID: authorID,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysql{{title .Alias}}Repository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.{{title .Alias}}, nextCursor string, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
FROM {{.Alias}} WHERE created_at ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return
}
func (m *mysql{{title .Alias}}Repository) GetByID(ctx context.Context, id int64) (res domain.{{title .Alias}}, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
FROM {{.Alias}} WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.{{title .Alias}}{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysql{{title .Alias}}Repository) GetByTitle(ctx context.Context, title string) (res domain.{{title .Alias}}, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
FROM {{.Alias}} WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysql{{title .Alias}}Repository) Store(ctx context.Context, a *domain.{{title .Alias}}) (err error) {
	query := `INSERT  {{.Alias}} SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}

func (m *mysql{{title .Alias}}Repository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM {{.Alias}} WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysql{{title .Alias}}Repository) Update(ctx context.Context, ar *domain.{{title .Alias}}) (err error) {
	query := `UPDATE {{.Alias}} set title=?, content=?, author_id=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.Content, ar.Author.ID, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}`
)
