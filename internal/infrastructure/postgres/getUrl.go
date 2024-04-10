package postgres

import (
	"context"
	"fmt"
	"net/url"
)

func (s *Storage) GetURL(ctx context.Context, key string) (url.URL, error) {
	const op = "postgres.GetURL"

	var data string
	err := s.db.QueryRowContext(ctx, "SELECT url_value FROM Urls WHERE alias = $1", key).Scan(&data)
	if err != nil {
		return url.URL{}, fmt.Errorf("%s : %w", op, err)
	}

	outputURL, _ := url.Parse(data)
	return *outputURL, nil
}
