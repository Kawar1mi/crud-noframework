package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Kawar1mi/crud-noframework/internal/models"
	"github.com/Kawar1mi/crud-noframework/internal/processors"
	"github.com/julienschmidt/httprouter"
)

type BooksHandler struct {
	processor *processors.BooksProcessor
}

func NewBooksHandler(processor *processors.BooksProcessor) *BooksHandler {
	handler := new(BooksHandler)
	handler.processor = processor
	return handler
}

func (h *BooksHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = h.processor.Create(r.Context(), newBook)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   "",
	}

	WrapOk(w, m)
}

func (h *BooksHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = h.processor.Update(r.Context(), book)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   "",
	}

	WrapOk(w, m)
}

func (h *BooksHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	bookId := p.ByName("bookId")
	if bookId == "" {
		err := errors.New("empty parameter bookId")
		WrapError(w, err)
		return
	}

	id, err := strconv.Atoi(bookId)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = h.processor.Delete(r.Context(), id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   "",
	}

	WrapOk(w, m)
}

func (h *BooksHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	books, err := h.processor.GetAll(r.Context())
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   books,
	}

	WrapOk(w, m)
}

func (h *BooksHandler) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	bookId := p.ByName("bookId")
	if bookId == "" {
		err := errors.New("empty parameter bookId")
		WrapError(w, err)
		return
	}

	id, err := strconv.Atoi(bookId)
	if err != nil {
		WrapError(w, err)
		return
	}

	book, err := h.processor.GetById(r.Context(), id)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   book,
	}

	WrapOk(w, m)
}
