package service

import (
	"avito-segment/models"
	"avito-segment/pkg/repository"
)

type AvitoSegment interface {
	Create(segment models.Segment) (string, error)
	Delete(slug string) error
	GetUserSegments(userID int) ([]models.Segment, error)
}

type AvitoUser interface {
	UpdateUserSegments(userID int, addSegments, removeSegments []string) error
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
