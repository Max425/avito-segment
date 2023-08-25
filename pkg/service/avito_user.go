package service

import (
	"avito-segment/pkg/repository"
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
