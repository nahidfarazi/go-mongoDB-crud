package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nahidfarazi/go-mongo2/handlers"
)

func main() {

	r := chi.NewRouter()
	r.Get("/users", handlers.GetAllUsers)
	r.Get("/users/{id}", handlers.GetUserByID)
	r.Post("/users", handlers.CreateUser)
	r.Put("/users/{id}", handlers.UpdateUser)
	r.Delete("/users/{id}", handlers.DeleteUser)
	fmt.Println("server started at port : 9000")
	http.ListenAndServe(":9000", r)
}
