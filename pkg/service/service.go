package service

import (
	avito_segment "avito-segment"
	"avito-segment/pkg/repository"
)

type AvitoSegment interface {
	Create(segment avito_segment.Segment) (int, error)
	Delete(slug string) error
	GetUserSegments(userID int) ([]avito_segment.Segment, error)
}

type AvitoUser interface {
	UpdateUserSegments(userID int, addSegments, removeSegments []string) error
}

type Service struct {
	AvitoSegment
	AvitoUser
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
