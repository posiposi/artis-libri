package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dsn := "sqlmock_db_0"
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := &bookRepository{db: gormDB}

	t.Run("successfully delete book", func(t *testing.T) {
		bookId := "1"
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Errorf("Error was not expected while deleting")

		}
	})

	t.Run("book not found", func(t *testing.T) {
		bookId := "2"
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("record not found"))).
			WillReturnError(errors.New("record not found"))

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Error("An error should have occurred.")
		}
	})

	t.Run("database error", func(t *testing.T) {
		bookId := "3"
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `books` WHERE id = ?").
			WithArgs(bookId).
			WillReturnError(errors.New("db error"))
		mock.ExpectCommit()

		err := repo.DeleteBook(bookId)
		if err == nil {
			t.Errorf("Error was not expected while deleting")

		}
	})
}
