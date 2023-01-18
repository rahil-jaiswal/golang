package router

import (
	"article-service/article/handlers"
	"article-service/article/middleware"
	repository2 "article-service/article/repository"
	"article-service/article/usecase"
	"article-service/author/repository"
	"article-service/config"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter(dbCon *sql.DB, config *config.Configuration) *gin.Engine {
	router := gin.Default()

	authorRepo := repository.NewAuthorRepository(dbCon)
	articleRepo := repository2.NewArticleRepository(dbCon)

	timeoutContext := time.Duration(config.Context.Timeout) * time.Second
	articleUseCase := usecase.NewArticleUsecase(articleRepo, authorRepo, timeoutContext)
	articleHandler := handler.NewArticleHandler(articleUseCase)

	router.RedirectTrailingSlash = true

	router.Use(middleware.CORS())

	articleGroup := router.Group("/api/v1/article")
	{
		articleGroup.GET("/:id", articleHandler.GetByID)

		//articleGroup.POST("", articleHandler.Store)
	}

	return router

}
