package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/wordkey_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type WordKeyHandler struct {
	CreateWordKeyUseCase      *wordkey_usecase.CreateWordKeyUseCase
	GetAllWordKeysUseCase     *wordkey_usecase.GetAllWordKeysUseCase
	GetWordKeyByIDUseCase     *wordkey_usecase.GetWordKeyByIDUseCase
	GetWordKeyByActiveUseCase *wordkey_usecase.GetWordKeyByActiveUseCase
	UpdateWordKeyByIDUseCase  *wordkey_usecase.UpdateWordKeyByIDUseCase
	DeleteWordKeyUseCase      *wordkey_usecase.DeleteWordKeyUseCase
}

func NewWordKeyHandler(
	createWordKeyUseCase *wordkey_usecase.CreateWordKeyUseCase,
	getAllWordKeysUseCase *wordkey_usecase.GetAllWordKeysUseCase,
	getWordKeyByIDUseCase *wordkey_usecase.GetWordKeyByIDUseCase,
	getWordKeyByActiveUseCase *wordkey_usecase.GetWordKeyByActiveUseCase,
	updateWordKeyByIDUseCase *wordkey_usecase.UpdateWordKeyByIDUseCase,
	deleteWordKeyUseCase *wordkey_usecase.DeleteWordKeyUseCase,
) *WordKeyHandler {
	return &WordKeyHandler{
		CreateWordKeyUseCase:      createWordKeyUseCase,
		GetAllWordKeysUseCase:     getAllWordKeysUseCase,
		GetWordKeyByIDUseCase:     getWordKeyByIDUseCase,
		GetWordKeyByActiveUseCase: getWordKeyByActiveUseCase,
		UpdateWordKeyByIDUseCase:  updateWordKeyByIDUseCase,
		DeleteWordKeyUseCase:      deleteWordKeyUseCase,
	}
}

func (h *WordKeyHandler) CreateWordKey(w http.ResponseWriter, r *http.Request) {
	var input wordkey_usecase.CreateWordKeyInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	output, err := h.CreateWordKeyUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Created(w, "Palavra chave criada com sucesso!", output)
}

func (h *WordKeyHandler) GetAllWordKeys(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllWordKeysUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Palavras chave buscadas com sucesso!", output)
}

func (h *WordKeyHandler) GetWordKeyByID(w http.ResponseWriter, r *http.Request) {
	wordkeyID := chi.URLParam(r, "id")

	input := wordkey_usecase.GetWordKeyByIDInputDTO{ID: wordkeyID}

	output, err := h.GetWordKeyByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Palavra chave encontrada com sucesso!", output)
}

func (h *WordKeyHandler) GetWordKeyByActive(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetWordKeyByActiveUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Palavra chave ativa encontrada com sucesso!", output)
}

func (h *WordKeyHandler) UpdateWordKey(w http.ResponseWriter, r *http.Request) {
	wordkeyID := chi.URLParam(r, "id")
	var input wordkey_usecase.UpdateWordKeyByIDInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido!"))
		return
	}

	input.ID = wordkeyID

	output, err := h.UpdateWordKeyByIDUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Palavra chave atualizada com sucesso!", output)
}

func (h *WordKeyHandler) DeleteWordKey(w http.ResponseWriter, r *http.Request) {
	wordkeyID := chi.URLParam(r, "id")

	input := wordkey_usecase.DeleteWordKeyInputDTO{ID: wordkeyID}

	if err := h.DeleteWordKeyUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Palavra chave deletada com sucesso!", nil)
}
