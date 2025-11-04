package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type UserHandler struct {
	CreateUserUseCase     *user_usecase.CreateUserUseCase
	GetAllUsersUseCase    *user_usecase.GetAllUsersUseCase
	UpdateUserByIdUseCase *user_usecase.UpdateUserByIdUseCase
	UpdatePasswordUseCase *user_usecase.UpdateUserPasswordUseCase
	GetUserByIdUseCase    *user_usecase.GetUserByIdUseCase
}

func NewUserHandler(
	createUserUseCase *user_usecase.CreateUserUseCase,
	getAllUsersUseCase *user_usecase.GetAllUsersUseCase,
	updateUserByIdUseCase *user_usecase.UpdateUserByIdUseCase,
	updatePasswordUseCase *user_usecase.UpdateUserPasswordUseCase,
	getUserByIdUseCase *user_usecase.GetUserByIdUseCase,
) *UserHandler {
	return &UserHandler{
		CreateUserUseCase:     createUserUseCase,
		GetAllUsersUseCase:    getAllUsersUseCase,
		UpdateUserByIdUseCase: updateUserByIdUseCase,
		UpdatePasswordUseCase: updatePasswordUseCase,
		GetUserByIdUseCase:    getUserByIdUseCase,
	}
}

// CreateUser godoc
// @Summary Criar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags users
// @Accept json
// @Produce json
// @Param user body user_usecase.CreateUserInputDTO true "Dados do usuário"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.CreateUserInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido"))
		return
	}

	output, err := h.CreateUserUseCase.Execute(input)
	if err != nil {
		response.Error(w, err) // AppError vem direto do usecase
		return
	}

	response.Created(w, "Usuário criado com sucesso!", output)
}

// GetAllUsers godoc
// @Summary Listar todos os usuários
// @Description Retorna todos os usuários cadastrados
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	output, err := h.GetAllUsersUseCase.Execute()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Todos os usuários buscados com sucesso!", output)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	output, err := h.GetUserByIdUseCase.Execute(user_usecase.GetUserByIdInputDTO{ID: userID})
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Usuário obtido com sucesso!", output)
}

// UpdateUser godoc
// @Summary Atualizar usuário
// @Description Atualiza os dados de um usuário
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Param user body user_usecase.UpdateUserByIdInputDTO true "Dados para atualização"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	var input user_usecase.UpdateUserByIdInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido"))
		return
	}

	input.ID = userID

	output, err := h.UpdateUserByIdUseCase.Execute(input)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Usuário atualizado com sucesso!", output)
}

// UpdatePassword godoc
// @Summary Atualizar senha
// @Description Atualiza a senha de um usuário
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Param password body user_usecase.UpdateUserPasswordInputDTO true "Nova senha"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id}/password [patch]
func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	var input user_usecase.UpdateUserPasswordInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, apperrors.NewValidationError("request", "Corpo da requisição inválido"))
		return
	}

	input.ID = userID

	if err := h.UpdatePasswordUseCase.Execute(input); err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, "Senha do usuário atualizada com sucesso!", nil)
}
