package service

import (
	"avito-segment/models"
	"avito-segment/pkg/repository"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

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

type Service struct {
	AvitoSegment
	AvitoUser
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AvitoSegment: NewAvitoSegmentService(repo.AvitoSegment),
		AvitoUser:    NewAvitoUserService(repo.AvitoUser),
	}
}
