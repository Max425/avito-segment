package service

import (
	avito_segment "avito-segment"
	"avito-segment/pkg/repository"
)

type AvitoSegmentService struct {
	repo repository.AvitoSegment
}

func NewAvitoSegmentService(repo repository.AvitoSegment) *AvitoSegmentService {
	return &AvitoSegmentService{repo: repo}
}

func (s *AvitoSegmentService) Create(segment avito_segment.Segment) (string, error) {
	return s.repo.Create(segment)
}

func (s *AvitoSegmentService) Delete(slug string) error {
	return s.repo.Delete(slug)
}

func (s *AvitoSegmentService) GetUserSegments(userId int) ([]avito_segment.Segment, error) {
	return s.repo.GetUserSegments(userId)
}
