package repository

import (
	"avito-segment/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type AvitoSegment interface {
	Create(segment models.Segment) (string, error)
	Delete(slug string) error
	GetUserSegments(userID int) ([]models.Segment, error)
}

type AvitoUser interface {
	UpdateUserSegments(userID int, addSegments, removeSegments []string) error
	AddUserToSegmentWithTTL(userID int, segmentSlug string, ttl time.Duration) error
	GenerateUserSegmentHistoryReport(userID int, year int, month time.Month) ([]byte, error)
}

type Repository struct {
	AvitoSegment
	AvitoUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AvitoSegment: NewAvitoSegmentPostgres(db),
		AvitoUser:    NewAvitoUserPostgres(db),
	}
}
