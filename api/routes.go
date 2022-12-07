package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func (app *application) routes() http.Handler {
func (app *application) routes() {
	r := gin.Default()

	r.GET("/keywords", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	// create a router mux
	// mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	// mux.Use(app.enableCORS)

	// mux.Get("/", app.Home)

	// mux.Post("/authenticate", app.authenticate)
	// mux.Get("/refresh", app.refreshToken)
	// mux.Get("/logout", app.logout)

	// mux.Get("/keywords", app.GetAllKeywords)
	// mux.Get("/keywords/{id}", app.GetKeyword)

	// mux.Get("/genres", app.AllGenres)
	// mux.Get("/movies/genres/{id}", app.AllMoviesByGenre)

	// mux.Post("/graph", app.moviesGraphQL)

	// mux.Route("/admin", func(mux chi.Router) {
	// 	mux.Use(app.authRequired)

	// mux.Get("/movies", app.MovieCatalog)
	// mux.Get("/movies/{id}", app.MovieForEdit)
	// mux.Put("/movies/0", app.InsertMovie)
	// mux.Patch("/movies/{id}", app.UpdateMovie)
	// mux.Delete("/movies/{id}", app.DeleteMovie)
	// })

	// return mux
}
