package domain

import "context"

type Article struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Updated_At string `json:"updated_at"`
	Created_At string `json:"created_at"`
	Author     Author `json:"author"`
}

type IArticleRepository interface {
	GetById(context context.Context, id int64) (Article, error)
}

type IArticleUseCase interface {
	GetByID(ctx context.Context, id int64) (Article, error)
}
