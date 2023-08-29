package service

import (
	"avito-segment/models"
	"avito-segment/pkg/repository"
)

type AvitoSegmentService struct {
	repo repository.AvitoSegment
}

func NewAvitoSegmentService(repo repository.AvitoSegment) *AvitoSegmentService {
	return &AvitoSegmentService{repo: repo}
}

func (s *AvitoSegmentService) Create(segment models.Segment) (string, error) {
	return s.repo.Create(segment)
}

func (s *AvitoSegmentService) Delete(slug string) error {
	return s.repo.Delete(slug)
}

func (s *AvitoSegmentService) GetUserSegments(userId int) ([]models.Segment, error) {
	return s.repo.GetUserSegments(userId)
}
