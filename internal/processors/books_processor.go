package processors

import (
	"context"
	"errors"

	loghelper "github.com/Kawar1mi/crud-noframework/internal/log_helper"
	"github.com/Kawar1mi/crud-noframework/internal/models"
	"github.com/Kawar1mi/crud-noframework/internal/storage"
)

type BooksProcessor struct {
	storage *storage.BooksPostgresStorage
}

func NewBooksProcessor(storage *storage.BooksPostgresStorage) *BooksProcessor {
	processor := new(BooksProcessor)
	processor.storage = storage
	return processor
}

func (p *BooksProcessor) Create(ctx context.Context, book models.Book) error {
	if book.Name == "" {
		err := errors.New("name should not be empty")
		loghelper.AddError(err)
		return err
	}

	return p.storage.Create(ctx, book)
}

func (p *BooksProcessor) Update(ctx context.Context, book models.Book) error {
	if book.Name == "" {
		err := errors.New("name should not be empty")
		loghelper.AddError(err)
		return err
	}

	_, err := p.storage.GetById(ctx, book.Id)
	if err != nil {
		return err
	}

	return p.storage.Update(ctx, book)
}

func (p *BooksProcessor) Delete(ctx context.Context, bookId int) error {

	_, err := p.storage.GetById(ctx, bookId)
	if err != nil {
		return err
	}

	return p.storage.Delete(ctx, bookId)
}

func (p *BooksProcessor) GetAll(ctx context.Context) ([]models.Book, error) {
	return p.storage.GetAll(ctx)
}

func (p *BooksProcessor) GetById(ctx context.Context, bookId int) (models.Book, error) {
	return p.storage.GetById(ctx, bookId)
}
