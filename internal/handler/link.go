package handler

import (
	"encoding/json"
	"fmt"
	"lobster/internal/models"
	"lobster/internal/service"
	"lobster/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type LinkHandler struct {
	Service *service.LinkService
}

func (s *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req models.CreateLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Bad JSON")
		return
	}
	defer r.Body.Close()

	if req.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "password is required")
		return
	}

	fmt.Printf("Создается ссылка для пароля (длина: %d)\n", len(req.Password))

	resp, err := s.Service.CreateLinkService(req.Password)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, "CreateLinkService in handler")
		return
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func (s *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	resp, ok := s.Service.ExtractLinkService(id)
	if !ok {
		utils.WriteError(w, http.StatusNotFound, fmt.Sprintf("not found link with id: %s", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
