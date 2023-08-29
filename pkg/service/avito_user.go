package service

import (
	"avito-segment/pkg/repository"
	"time"
)

type AvitoUserService struct {
	repo repository.AvitoUser
}

func NewAvitoUserService(repo repository.AvitoUser) *AvitoUserService {
	return &AvitoUserService{repo: repo}
}

func (s *AvitoUserService) UpdateUserSegments(userID int, addSegments, removeSegments []string) error {
	return s.repo.UpdateUserSegments(userID, addSegments, removeSegments)
}

func (s *AvitoUserService) AddUserToSegmentWithTTL(userID int, segmentSlug string, ttl time.Duration) error {
	return s.repo.AddUserToSegmentWithTTL(userID, segmentSlug, ttl)
}
