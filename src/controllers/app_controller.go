package controllers

import (
	"net/http"

	"darkwilly08.com/mockapi/src/config/database"
	"darkwilly08.com/mockapi/src/dtos"
	"darkwilly08.com/mockapi/src/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AppController struct {
	appService *services.AppService
}

func NewAppController(db *database.DB) *AppController {
	return &AppController{
		appService: services.NewAppService(db),
	}
}

var arr []dtos.AppDto = []dtos.AppDto{{
	Name: "hola",
}}

// Routes creates a REST router for the todos resource
func (c *AppController) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", c.findAll)
	r.Post("/", c.create) // POST /users - create a new user and persist it
	// r.Put("/", rs.Delete)

	// r.Route("/{id}", func(r chi.Router) {
	// 	// r.Use(rs.TodoCtx) // lets have a users map, and lets actually load/manipulate
	// 	r.Get("/", rs.Get)       // GET /users/{id} - read a single user by :id
	// 	r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
	// 	r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	// })

	return r
}

func (c *AppController) findAll(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, arr)
}

func (c *AppController) create(w http.ResponseWriter, r *http.Request) {
	var bodyResponse dtos.AppDto
	err := render.DecodeJSON(r.Body, &bodyResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, err.Error())
		return
	}

	appDto := c.appService.CreateOrUpdate(&bodyResponse)

	render.JSON(w, r, appDto)
}
