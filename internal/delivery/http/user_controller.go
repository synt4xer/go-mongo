package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/synt4xer/go-mongo/internal/domain"
	"github.com/synt4xer/go-mongo/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(usecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.usecase.Save(r.Context(), &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
