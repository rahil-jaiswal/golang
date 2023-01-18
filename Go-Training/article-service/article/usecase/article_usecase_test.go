package usecase

import (
	"article-service/domain"
	mocks "article-service/mock/domain"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// tests using testsuite
type ArticleUseCaseTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	mockArticleRepo *mocks.MockIArticleRepository
	mockAuthorRepo  *mocks.MockIAuthorRepository
}

func TestArticleUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleUseCaseTestSuite))
}

func (t *ArticleUseCaseTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.mockArticleRepo = mocks.NewMockIArticleRepository(t.mockCtrl)
	t.mockAuthorRepo = mocks.NewMockIAuthorRepository(t.mockCtrl)
}

func (t *ArticleUseCaseTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
}

func (t *ArticleUseCaseTestSuite) TestGetByID() {
	mockArticle := domain.Article{
		Title:   "Hello",
		Content: "Content",
	}
	mockAuthor := domain.Author{
		ID:   1,
		Name: "John Doe",
	}

	t.mockArticleRepo.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(mockArticle, nil).Times(1)
	t.mockAuthorRepo.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(mockAuthor, nil).Times(1)

	timeoutDuration := time.Duration(100) * time.Second
	articleUsecase := NewArticleUsecase(t.mockArticleRepo, t.mockAuthorRepo, timeoutDuration)

	article, err := articleUsecase.GetByID(context.TODO(), 1)

	assert.Nil(t.T(), err)
	assert.NotNil(t.T(), article)

	assert.Equal(t.T(), "John Doe", article.Author.Name)
}
