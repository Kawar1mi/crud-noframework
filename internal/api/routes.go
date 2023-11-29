package api

import (
	"fmt"
	"net/http"

	"github.com/Kawar1mi/crud-noframework/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

func CreateRouter(booksHandler *handlers.BooksHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Welcome home!")
	})

	router.POST("/api/books", booksHandler.Create)
	router.PATCH("/api/books", booksHandler.Update)
	router.DELETE("/api/books/:bookId", booksHandler.Delete)
	router.GET("/api/books", booksHandler.GetAll)
	router.GET("/api/books/:bookId", booksHandler.GetById)

	return router
}
