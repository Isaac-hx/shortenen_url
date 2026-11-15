package usecase

import (
	"context"
	"url_shortener_api/internal/repository"
)

type UrlUsecase interface {
	CreateUrl(ctx context.Context, longUrl string, customCode string) (string, error)
}
type urlUsecase struct {
	urlRepo repository.UrlRepository
}

func NewUrlUsecase(urlRepo repository.UrlRepository) UrlUsecase {
	return &urlUsecase{
		urlRepo: urlRepo,
	}
}

func (u *urlUsecase) CreateUrl(ctx context.Context, longUrl string, customCode string) (string, error) {
	return u.urlRepo.CreateShortUrl(ctx, longUrl, customCode)
}
