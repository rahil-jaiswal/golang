package repository

import (
	domain "article-service/domain"
	"context"
	"database/sql"
	"fmt"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) domain.IArticleRepository {
	return &ArticleRepository{
		db,
	}
}

func (m *ArticleRepository) GetById(ctx context.Context, id int64,
) (article domain.Article, err error) {

	query := `SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE id = ?`

	articles, err := m.fetch(ctx, query, id)
	if err != nil {
		return
	}

	if len(articles) > 0 {
		article = articles[0]
	}
	return
}

func (m *ArticleRepository) fetch(ctx context.Context, query string, args ...interface{},
) (result []domain.Article, err error) {

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		rowErr := rows.Close()
		if rowErr != nil {
			fmt.Println(rowErr)
		}
	}()

	result = make([]domain.Article, 0)
	for rows.Next() {
		t := domain.Article{}
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.Updated_At,
			&t.Created_At,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		t.Author = domain.Author{
			ID: authorID,
		}
		result = append(result, t)
	}

	return
}
