package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)
import "gopkg.in/DATA-DOG/go-sqlmock.v1"

func TestAuthorRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating connection")
	}
	defer db.Close()
	mock.ExpectPrepare("select from author where id=?;").ExpectQuery().WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(1, "Lionel Messi", time.Now(), time.Now()))

	authRepo := NewAuthorRepository(db)
	actualAuthor := authRepo.GetById(context.TODO(), 1)
	assert.Equal(t, actualAuthor.Name, "Lionel Mssi")
}
