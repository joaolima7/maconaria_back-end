package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joaolima7/maconaria_back-end/internal/domain/usecases/user_usecase"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type AuthHandler struct {
	LoginUseCase *user_usecase.LoginUseCase
}

func NewAuthHandler(loginUseCase *user_usecase.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		LoginUseCase: loginUseCase,
	}
}

// Login godoc
// @Summary Login de usuário
// @Description Autentica usuário e retorna token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body user_usecase.LoginInputDTO true "Credenciais de login"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input user_usecase.LoginInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.BadRequest(w, "Dados inválidos!", err)
		return
	}

	output, err := h.LoginUseCase.Execute(input)
	if err != nil {
		response.Unauthorized(w, err.Error(), nil)
		return
	}

	response.OK(w, "Login realizado com sucesso!", output)
}
