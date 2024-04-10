package postgres

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) error {
	const op = "postgres.SaveURL"

	_, err := s.db.ExecContext(ctx,
		"INSERT INTO Urls(alias, url_value) VALUES ($1, $2)",
		key, urlToSave.String(),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
