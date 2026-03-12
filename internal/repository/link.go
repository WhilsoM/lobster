package repository

import (
	"fmt"
	"lobster/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type LinkStore struct {
	Mu    sync.Mutex
	Links map[uuid.UUID]string
}

func NewLinkStore() *LinkStore {
	return &LinkStore{
		Links: make(map[uuid.UUID]string),
	}
}

func (s *LinkStore) StartCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			s.Mu.Lock()
			for k := range s.Links {
				delete(s.Links, k)
			}
			s.Mu.Unlock()
			fmt.Println("Repository cleanup performed")
		}
	}()
}

func (s *LinkStore) Save(id uuid.UUID, password string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Links[id] = password
}

func (s *LinkStore) ExtractLink(id uuid.UUID) (models.GetLinkResponse, bool) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	password, ok := s.Links[id]
	if !ok {
		return models.GetLinkResponse{}, false
	}

	delete(s.Links, id)

	resp := models.GetLinkResponse{
		Password: password,
	}

	return resp, true
}
