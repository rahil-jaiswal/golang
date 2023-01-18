package usecase

import (
	"article-service/domain"
	"context"
	"time"
)

type articleUsecase struct {
	articleRepo    domain.IArticleRepository
	authorRepo     domain.IAuthorRepository
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(a domain.IArticleRepository, ar domain.IAuthorRepository, timeout time.Duration) domain.IArticleUseCase {
	return &articleUsecase{
		articleRepo:    a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (article domain.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	article, err = a.articleRepo.GetById(ctx, id)
	if err != nil {
		return
	}

	author := a.authorRepo.GetById(ctx, article.Author.ID)
	article.Author = author
	return
}
