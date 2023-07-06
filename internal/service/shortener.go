package service

import (
	"context"
	"time"

	"github.com/kiryu-dev/shorty/internal/libshorty/valuegen"
	"github.com/kiryu-dev/shorty/internal/model"
)

type shortenerStorage interface {
	Save(context.Context, *model.ShortURL) error
	GetURL(context.Context, string) (string, error)
	Delete(context.Context, string) (*model.ShortURL, error)
}

type Shortener struct {
	storage shortenerStorage
}

func NewShortener(storage shortenerStorage) *Shortener {
	return &Shortener{storage}
}

func (s *Shortener) MakeShort(ctx context.Context, url string) (string, error) {
	shortURL := &model.ShortURL{
		URL:       url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	var err error
	shortURL.Alias, err = valuegen.GenerateValue()
	if err != nil {
		return "", err
	}
	if err := s.storage.Save(ctx, shortURL); err != nil {
		return "", err
	}
	return shortURL.Alias, nil
}
