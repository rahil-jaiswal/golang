package handler

import (
	"article-service/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	ArticleUseCase domain.IArticleUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(articleUseCase domain.IArticleUseCase) IArticleHandler {
	return &ArticleHandler{
		ArticleUseCase: articleUseCase,
	}
}

type IArticleHandler interface {
	GetByID(ctx *gin.Context)
	Store(ctx *gin.Context)
}

// GetByID will get article by given id
func (a *ArticleHandler) GetByID(ctx *gin.Context) {
	idP, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	id := int64(idP)
	article, err := a.ArticleUseCase.GetByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &ResponseError{domain.ErrNotFound.Error()})
		return
	}
	ctx.JSON(http.StatusOK, article)
	return
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *ArticleHandler) Store(ctx *gin.Context) {
	var article domain.Article

	err := ctx.Bind(&article)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	if ok, err := isRequestValid(&article); !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	//err = a.ArticleUseCase.Store(ctx, &article)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(getStatusCode(err), ResponseError{Message: err.Error()})
	//	return
	//}

	ctx.JSON(http.StatusCreated, article)
	return
}
