package main

import (
	"net/http"

	"darkwilly08.com/mockapi/src/config/database"
	"darkwilly08.com/mockapi/src/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := database.InitDB("url")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/apps", controllers.NewAppController(db).Routes())

	http.ListenAndServe(":3333", r)
}
