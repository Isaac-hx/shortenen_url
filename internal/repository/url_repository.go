package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UrlRepository interface {
	CreateShortUrl(ctx context.Context, url string, code string) (string, error)
}

type urlRepository struct {
	conn *pgx.Conn
}

func NewUrlRepository(conn *pgx.Conn) UrlRepository {
	return &urlRepository{conn: conn}
}

func (u *urlRepository) CreateShortUrl(ctx context.Context, url string, code string) (string, error) {
	query := `INSERT INTO urls (url,code) VALUES($1,$2)`
	_, err := u.conn.Exec(ctx, query, url, code)
	if err != nil {
		return "", err
	}
	return code, nil
}
