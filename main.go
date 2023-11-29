package main

import (
	"net/http"
	"os"

	"github.com/Kawar1mi/crud-noframework/internal/api"
	"github.com/Kawar1mi/crud-noframework/internal/db"
	"github.com/Kawar1mi/crud-noframework/internal/handlers"
	loghelper "github.com/Kawar1mi/crud-noframework/internal/log_helper"
	"github.com/Kawar1mi/crud-noframework/internal/processors"
	"github.com/Kawar1mi/crud-noframework/internal/storage"
)

func main() {

	db := db.DatabaseConnection()

	booksStorage := storage.NewBooksPostgresStorage(db)

	booksProcessor := processors.NewBooksProcessor(booksStorage)

	booksHandler := handlers.NewBooksHandler(booksProcessor)

	router := api.CreateRouter(booksHandler)

	addr := os.Getenv("ADDR")
	server := http.Server{Addr: addr, Handler: router}

	loghelper.InfoMsg("starting server on: " + addr)

	err := server.ListenAndServe()
	loghelper.FatalIfError(err)
}
