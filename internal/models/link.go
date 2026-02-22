package models

import "github.com/google/uuid"

type LinkStore struct {
	ID       string
	Password string
}

type CreateLinkRequest struct {
	Password string `json:"password"`
}

type CreateLinkResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetLinkRequest struct {
	ID string `json:"id"`
}

type GetLinkResponse struct {
	Password string `json:"password"`
}
