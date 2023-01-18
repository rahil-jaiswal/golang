package repository

import (
	"context"
	"database/sql"
	"fmt"

	domain "article-service/domain"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) domain.IAuthorRepository {
	return &AuthorRepository{
		db,
	}
}

func (a *AuthorRepository) GetById(context context.Context, id int64) (author domain.Author) {

	searchQuery := fmt.Sprintf("Select from author where id=?")
	result := a.db.QueryRowContext(context, searchQuery, id)
	result.Scan(&author.ID, &author.Name, &author.Created_At, &author.Updated_At)

	return
}
