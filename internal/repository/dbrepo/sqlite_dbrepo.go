package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"profileyou/api/models"
	"time"
)

type SQliteDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// Connection returns underlying connection pool.
func (m *SQliteDBRepo) Connection() *sql.DB {
	return m.DB
}

// AllMovies returns a slice of movies, sorted by name. If the optional parameter genre
// is supplied, then only all movies for a particular genre is returned.
func (m *SQliteDBRepo) AllKeywords(genre ...int) ([]*models.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		// where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genre[0])
		where = fmt.Sprintf("keywords")
	}

	query := fmt.Sprintf(`
		select
			id, word, description, coalesce(image_url, ''),
			created_at, updated_at
		from
			keywords %s
		order by
			created_at
	`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keywords []*models.Keyword

	for rows.Next() {
		var keyword models.Keyword
		err := rows.Scan(
			&keyword.ID,
			&keyword.Word,
			&keyword.Description,
			&keyword.ImageUrl,
			&keyword.CreatedAt,
			&keyword.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		keywords = append(keywords, &keyword)
	}

	return keywords, nil
}

func (m *SQliteDBRepo) OneMovie(id int) (*models.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, word, description, image_url, created_at, 
		updated_at, coalesce(image, ''), deleted_at
		from keywords where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var keyword models.Keyword

	err := row.Scan(
		&keyword.ID,
		&keyword.Word,
		&keyword.Description,
		&keyword.ImageUrl,
		&keyword.CreatedAt,
		&keyword.UpdatedAt,
		&keyword.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &keyword, err
}
