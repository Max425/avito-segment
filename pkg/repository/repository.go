package repository

import (
	avito_segment "avito-segment"
	"github.com/jmoiron/sqlx"
)

type AvitoSegment interface {
	Create(segment avito_segment.Segment) (int, error)
	Delete(slug string) error
	GetUserSegments(userID int) ([]avito_segment.Segment, error)
}

type AvitoUser interface {
	UpdateUserSegments(userID int, addSegments, removeSegments []string) error
}

type Repository struct {
	AvitoSegment
	AvitoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
