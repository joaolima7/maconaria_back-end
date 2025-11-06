package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/library_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type LibraryHandler struct {
	CreateLibraryUseCase        *library_usecase.CreateLibraryUseCase
	GetAllLibrariesUseCase      *library_usecase.GetAllLibrariesUseCase
	GetLibraryByIDUseCase       *library_usecase.GetLibraryByIDUseCase
	GetLibrariesByDegreeUseCase *library_usecase.GetLibrariesByDegreeUseCase
	UpdateLibraryByIDUseCase    *library_usecase.UpdateLibraryByIDUseCase
	DeleteLibraryUseCase        *library_usecase.DeleteLibraryUseCase
}

func NewLibraryHandler(
	createLibraryUseCase *library_usecase.CreateLibraryUseCase,
	getAllLibrariesUseCase *library_usecase.GetAllLibrariesUseCase,
	getLibraryByIDUseCase *library_usecase.GetLibraryByIDUseCase,
	getLibrariesByDegreeUseCase *library_usecase.GetLibrariesByDegreeUseCase,
	updateLibraryByIDUseCase *library_usecase.UpdateLibraryByIDUseCase,
	deleteLibraryUseCase *library_usecase.DeleteLibraryUseCase,
) *LibraryHandler {
	return &LibraryHandler{
		CreateLibraryUseCase:        createLibraryUseCase,
		GetAllLibrariesUseCase:      getAllLibrariesUseCase,
		GetLibraryByIDUseCase:       getLibraryByIDUseCase,
		GetLibrariesByDegreeUseCase: getLibrariesByDegreeUseCase,
		UpdateLibraryByIDUseCase:    updateLibraryByIDUseCase,
		DeleteLibraryUseCase:        deleteLibraryUseCase,
	}
}

func (h *LibraryHandler) CreateLibrary(w http.ResponseWriter, r *http.Request) {
	var input library_usecase.CreateLibraryInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	output, err := h.CreateLibraryUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Biblioteca criada com sucesso!", output)
}

func (h *LibraryHandler) GetAllLibraries(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllLibrariesUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Bibliotecas buscadas com sucesso!", output)
}

func (h *LibraryHandler) GetLibraryByID(w http.ResponseWriter, r *http.Request) {
	libraryID := chi.URLParam(r, "id")

	input := library_usecase.GetLibraryByIDInputDTO{ID: libraryID}

	output, err := h.GetLibraryByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Biblioteca encontrada com sucesso!", output)
}

func (h *LibraryHandler) GetLibrariesByDegree(w http.ResponseWriter, r *http.Request) {
	degree := chi.URLParam(r, "degree")

	input := library_usecase.GetLibrariesByDegreeInputDTO{Degree: degree}

	output, err := h.GetLibrariesByDegreeUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Bibliotecas buscadas com sucesso!", output)
}

func (h *LibraryHandler) UpdateLibrary(w http.ResponseWriter, r *http.Request) {
	libraryID := chi.URLParam(r, "id")
	var input library_usecase.UpdateLibraryByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	input.ID = libraryID

	output, err := h.UpdateLibraryByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Biblioteca atualizada com sucesso!", output)
}

func (h *LibraryHandler) DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	libraryID := chi.URLParam(r, "id")

	input := library_usecase.DeleteLibraryInputDTO{ID: libraryID}

	if err := h.DeleteLibraryUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Biblioteca deletada com sucesso!", nil)
}
