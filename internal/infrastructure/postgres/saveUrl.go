package postgres

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) (string, error) {
	const op = "postgres.SaveURL"
	var oldKey string

	err := s.db.QueryRowContext(ctx,
		"INSERT INTO Urls(alias, url_value) VALUES ($1, $2) ON CONFLICT (url_value) DO UPDATE SET url_value = EXCLUDED.url_value RETURNING alias",
		key, urlToSave.String(),
	).Scan(&oldKey)
	if err != nil {
		return oldKey, fmt.Errorf("%s: %w", op, err)
	}

	return oldKey, nil
}
