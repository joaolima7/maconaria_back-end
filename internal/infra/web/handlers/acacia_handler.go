package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/acacia_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type AcaciaHandler struct {
	CreateAcaciaUseCase     *acacia_usecase.CreateAcaciaUseCase
	GetAllAcaciasUseCase    *acacia_usecase.GetAllAcaciasUseCase
	GetAcaciaByIDUseCase    *acacia_usecase.GetAcaciaByIDUseCase
	UpdateAcaciaByIDUseCase *acacia_usecase.UpdateAcaciaByIDUseCase
	DeleteAcaciaUseCase     *acacia_usecase.DeleteAcaciaUseCase
}

func NewAcaciaHandler(
	createAcaciaUseCase *acacia_usecase.CreateAcaciaUseCase,
	getAllAcaciasUseCase *acacia_usecase.GetAllAcaciasUseCase,
	getAcaciaByIDUseCase *acacia_usecase.GetAcaciaByIDUseCase,
	updateAcaciaByIDUseCase *acacia_usecase.UpdateAcaciaByIDUseCase,
	deleteAcaciaUseCase *acacia_usecase.DeleteAcaciaUseCase,
) *AcaciaHandler {
	return &AcaciaHandler{
		CreateAcaciaUseCase:     createAcaciaUseCase,
		GetAllAcaciasUseCase:    getAllAcaciasUseCase,
		GetAcaciaByIDUseCase:    getAcaciaByIDUseCase,
		UpdateAcaciaByIDUseCase: updateAcaciaByIDUseCase,
		DeleteAcaciaUseCase:     deleteAcaciaUseCase,
	}
}

func (h *AcaciaHandler) CreateAcacia(w http.ResponseWriter, r *http.Request) {
	var input acacia_usecase.CreateAcaciaInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	output, err := h.CreateAcaciaUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Acácia criada com sucesso!", output)
}

func (h *AcaciaHandler) GetAllAcacias(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllAcaciasUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Acácias buscadas com sucesso!", output)
}

func (h *AcaciaHandler) GetAcaciaByID(w http.ResponseWriter, r *http.Request) {
	acaciaID := chi.URLParam(r, "id")

	input := acacia_usecase.GetAcaciaByIDInputDTO{ID: acaciaID}

	output, err := h.GetAcaciaByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Acácia encontrada com sucesso!", output)
}

func (h *AcaciaHandler) UpdateAcacia(w http.ResponseWriter, r *http.Request) {
	acaciaID := chi.URLParam(r, "id")
	var input acacia_usecase.UpdateAcaciaByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	input.ID = acaciaID

	output, err := h.UpdateAcaciaByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Acácia atualizada com sucesso!", output)
}

func (h *AcaciaHandler) DeleteAcacia(w http.ResponseWriter, r *http.Request) {
	acaciaID := chi.URLParam(r, "id")

	input := acacia_usecase.DeleteAcaciaInputDTO{ID: acaciaID}

	if err := h.DeleteAcaciaUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Acácia deletada com sucesso!", nil)
}
