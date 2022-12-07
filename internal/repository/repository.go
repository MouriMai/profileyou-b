package repository

import (
	"database/sql"
	"profileyou/api/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllKeywords() ([]*models.Keyword, error)
	// GetUserByEmail(email string) (*models.User, error)
	// GetUserByID(id int) (*models.User, error)

	// OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	GetKeyword(id int) (*models.Keyword, error)
	// AllGenres() ([]*models.Genre, error)
	// InsertMovie(movie models.Movie) (int, error)
	// UpdateMovieGenres(id int, genreIDs []int) error
	// UpdateMovie(movie models.Movie) error
	// DeleteMovie(id int) error
}
