package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/worker_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type WorkerHandler struct {
	CreateWorkerUseCase     *worker_usecase.CreateWorkerUseCase
	GetAllWorkersUseCase    *worker_usecase.GetAllWorkersUseCase
	GetWorkerByIDUseCase    *worker_usecase.GetWorkerByIDUseCase
	UpdateWorkerByIDUseCase *worker_usecase.UpdateWorkerByIDUseCase
	DeleteWorkerUseCase     *worker_usecase.DeleteWorkerUseCase
}

func NewWorkerHandler(
	createWorkerUseCase *worker_usecase.CreateWorkerUseCase,
	getAllWorkersUseCase *worker_usecase.GetAllWorkersUseCase,
	getWorkerByIDUseCase *worker_usecase.GetWorkerByIDUseCase,
	updateWorkerByIDUseCase *worker_usecase.UpdateWorkerByIDUseCase,
	deleteWorkerUseCase *worker_usecase.DeleteWorkerUseCase,
) *WorkerHandler {
	return &WorkerHandler{
		CreateWorkerUseCase:     createWorkerUseCase,
		GetAllWorkersUseCase:    getAllWorkersUseCase,
		GetWorkerByIDUseCase:    getWorkerByIDUseCase,
		UpdateWorkerByIDUseCase: updateWorkerByIDUseCase,
		DeleteWorkerUseCase:     deleteWorkerUseCase,
	}
}

func (h *WorkerHandler) CreateWorker(w http.ResponseWriter, r *http.Request) {
	var input worker_usecase.CreateWorkerInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	output, err := h.CreateWorkerUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Obreiro criado com sucesso!", output)
}

func (h *WorkerHandler) GetAllWorkers(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllWorkersUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Obreiros buscados com sucesso!", output)
}

func (h *WorkerHandler) GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	workerID := chi.URLParam(r, "id")

	input := worker_usecase.GetWorkerByIDInputDTO{ID: workerID}

	output, err := h.GetWorkerByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Obreiro encontrado com sucesso!", output)
}

func (h *WorkerHandler) UpdateWorker(w http.ResponseWriter, r *http.Request) {
	workerID := chi.URLParam(r, "id")
	var input worker_usecase.UpdateWorkerByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	input.ID = workerID

	output, err := h.UpdateWorkerByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Obreiro atualizado com sucesso!", output)
}

func (h *WorkerHandler) DeleteWorker(w http.ResponseWriter, r *http.Request) {
	workerID := chi.URLParam(r, "id")

	input := worker_usecase.DeleteWorkerInputDTO{ID: workerID}

	if err := h.DeleteWorkerUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Obreiro deletado com sucesso!", nil)
}
