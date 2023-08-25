package service

import (
	avito_segment "avito-segment"
	"avito-segment/pkg/repository"
)

type AvitoSegment interface {
	Create(segment avito_segment.Segment) (int, error)
	Delete(slug string) error
}

type AvitoUser interface {
}

type Service struct {
	AvitoSegment
	AvitoUser
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
