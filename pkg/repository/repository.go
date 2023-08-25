package repository

import (
	avito_segment "avito-segment"
	"github.com/jmoiron/sqlx"
)

type AvitoSegment interface {
	Create(segment avito_segment.Segment) (int, error)
	Delete(slug string) error
}

type AvitoUser interface {
}

type Repository struct {
	AvitoSegment
	AvitoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
