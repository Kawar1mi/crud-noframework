package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Kawar1mi/crud-noframework/internal/models"
)

type BooksPostgresStorage struct {
	db *sql.DB
}

func NewBooksPostgresStorage(db *sql.DB) *BooksPostgresStorage {
	storage := new(BooksPostgresStorage)
	storage.db = db
	return storage
}

func (b *BooksPostgresStorage) Create(ctx context.Context, book models.Book) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO books (name) VALUES ($1)", book.Name)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (b *BooksPostgresStorage) Update(ctx context.Context, book models.Book) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE books SET name = $1 WHERE id = $2", book.Name, book.Id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (b *BooksPostgresStorage) Delete(ctx context.Context, bookId int) error {

	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "DELETE FROM books WHERE id = $1", bookId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (b *BooksPostgresStorage) GetAll(ctx context.Context) ([]models.Book, error) {

	var books []models.Book

	rows, err := b.db.QueryContext(ctx, "SELECT id, name FROM books")
	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.Id, &book.Name)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BooksPostgresStorage) GetById(ctx context.Context, bookId int) (models.Book, error) {

	var book models.Book

	rows, err := b.db.QueryContext(ctx, "SELECT id, name FROM books WHERE id = $1", bookId)
	if err != nil {
		return book, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&book.Id, &book.Name)
		if err != nil {
			return book, err
		}
	} else {
		err = fmt.Errorf("book id %v not found", bookId)
		return book, err
	}

	return book, nil
}
