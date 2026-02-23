package service

import (
	"lobster/internal/models"
	"errors"
	"github.com/google/uuid"
)

type LinkStorage interface {
	Save(id uuid.UUID, password string)
	ExtractLink(id uuid.UUID) (models.GetLinkResponse, bool)
}

type LinkService struct {
	Storage LinkStorage
}

func (s *LinkService) CreateLinkService(password string) (models.CreateLinkResponse, error) {
	if password == "" { 
		return models.CreateLinkResponse{}, errors.New("password cannot be empty")
	}

	id := uuid.New()
	resp := models.CreateLinkResponse{ID: id}

	s.Storage.Save(id, password)

	return resp, nil
}

func (s *LinkService) ExtractLinkService(id uuid.UUID) (models.GetLinkResponse, bool) {
	return s.Storage.ExtractLink(id)
}
