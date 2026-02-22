package repository

import (
	"lobster/internal/models"
	"sync"

	"github.com/google/uuid"
)

type LinkStore struct {
	mu sync.RWMutex
	// uuid = password
	links map[uuid.UUID]string
}

func NewLinkStore() *LinkStore {
	return &LinkStore{
		links: make(map[uuid.UUID]string),
	}
}

func (s *LinkStore) Save(id uuid.UUID, password string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.links[id] = password
}

// сразу получает и удаляет ссылку
func (s *LinkStore) ExtractLink(id uuid.UUID) (models.GetLinkResponse, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	password, ok := s.links[id]
	if !ok {
		return models.GetLinkResponse{}, false
	}

	delete(s.links, id)

	resp := models.GetLinkResponse{
		Password: password,
	}

	return resp, true
}
