package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/timeline_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type TimelineHandler struct {
	CreateTimelineUseCase     *timeline_usecase.CreateTimelineUseCase
	GetAllTimelinesUseCase    *timeline_usecase.GetAllTimelinesUseCase
	GetTimelineByIDUseCase    *timeline_usecase.GetTimelineByIDUseCase
	UpdateTimelineByIDUseCase *timeline_usecase.UpdateTimelineByIDUseCase
	DeleteTimelineUseCase     *timeline_usecase.DeleteTimelineUseCase
}

func NewTimelineHandler(
	createTimelineUseCase *timeline_usecase.CreateTimelineUseCase,
	getAllTimelinesUseCase *timeline_usecase.GetAllTimelinesUseCase,
	getTimelineByIDUseCase *timeline_usecase.GetTimelineByIDUseCase,
	updateTimelineByIDUseCase *timeline_usecase.UpdateTimelineByIDUseCase,
	deleteTimelineUseCase *timeline_usecase.DeleteTimelineUseCase,
) *TimelineHandler {
	return &TimelineHandler{
		CreateTimelineUseCase:     createTimelineUseCase,
		GetAllTimelinesUseCase:    getAllTimelinesUseCase,
		GetTimelineByIDUseCase:    getTimelineByIDUseCase,
		UpdateTimelineByIDUseCase: updateTimelineByIDUseCase,
		DeleteTimelineUseCase:     deleteTimelineUseCase,
	}
}

func (h *TimelineHandler) CreateTimeline(w http.ResponseWriter, r *http.Request) {
	var input timeline_usecase.CreateTimelineInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	output, err := h.CreateTimelineUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Timeline criada com sucesso!", output)
}

func (h *TimelineHandler) GetAllTimelines(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllTimelinesUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Timelines buscadas com sucesso!", output)
}

func (h *TimelineHandler) GetTimelineByID(w http.ResponseWriter, r *http.Request) {
	timelineID := chi.URLParam(r, "id")

	input := timeline_usecase.GetTimelineByIDInputDTO{ID: timelineID}

	output, err := h.GetTimelineByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Timeline encontrada com sucesso!", output)
}

func (h *TimelineHandler) UpdateTimeline(w http.ResponseWriter, r *http.Request) {
	timelineID := chi.URLParam(r, "id")
	var input timeline_usecase.UpdateTimelineByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	input.ID = timelineID

	output, err := h.UpdateTimelineByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Timeline atualizada com sucesso!", output)
}

func (h *TimelineHandler) DeleteTimeline(w http.ResponseWriter, r *http.Request) {
	timelineID := chi.URLParam(r, "id")

	input := timeline_usecase.DeleteTimelineInputDTO{ID: timelineID}

	if err := h.DeleteTimelineUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Timeline deletada com sucesso!", nil)
}
