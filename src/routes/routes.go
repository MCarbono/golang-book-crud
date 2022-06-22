package routes

import (
	"go-bookstore/src/controllers"
	"go-bookstore/src/middleware"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.StrictSlash(true)
	router.Use(middleware.JSON)
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
