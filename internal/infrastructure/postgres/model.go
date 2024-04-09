package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq"
	"github.com/smakkking/url-shortener/internal/app"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(cfg app.Config) (*Storage, error) {
	time.Sleep(5 * time.Second) // для корректного подключения в докере

	database_url := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgDBName,
		cfg.PgUser,
		cfg.PgPassword,
		cfg.PgSSLMode,
	)
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		return nil, err
	}

	// проверка, что подключились
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) SaveURL(ctx context.Context, key string, urlToSave url.URL) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO Urls(alias, url_value) VALUES $1, $2",
		key, urlToSave.String(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetURL(ctx context.Context, key string) (url.URL, error) {
	var data string
	err := s.db.QueryRowContext(ctx, "SELECT url_value FROM Urls WHERE alias = $1", key).Scan(&data)
	if err != nil {
		return url.URL{}, err
	}

	outputURL, _ := url.Parse(data)
	return *outputURL, nil
}
