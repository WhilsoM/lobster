package service

import (
	"lobster/internal/models"

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
	namespace := uuid.Must(uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	name := password
	id := uuid.NewSHA1(namespace, []byte(name))
	resp := models.CreateLinkResponse{ID: id}

	s.Storage.Save(id, password)

	return resp, nil
}

func (s *LinkService) ExtractLinkService(id uuid.UUID) (models.GetLinkResponse, bool) {
	return s.Storage.ExtractLink(id)
}
