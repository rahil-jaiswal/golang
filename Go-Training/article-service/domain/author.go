package domain

import "context"

type Author struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type IAuthorRepository interface {
	GetById(context context.Context, id int64) Author
}
