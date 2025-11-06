package user_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/auth"
	"golang.org/x/crypto/bcrypt"
)

type LoginInputDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginOutputDTO struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CIM      string `json:"cim"`
	Degree   string `json:"degree"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginUseCase struct {
	Repository user.GetUserByEmailRepository
	JWTService *auth.JWTService
}

func NewLoginUseCase(
	repository user.GetUserByEmailRepository,
	jwtService *auth.JWTService,
) *LoginUseCase {
	return &LoginUseCase{
		Repository: repository,
		JWTService: jwtService,
	}
}

func (uc *LoginUseCase) Execute(input LoginInputDTO) (*LoginOutputDTO, error) {
	user, err := uc.Repository.GetUserByEmail(input.Email)
	if err != nil {
		return nil, apperrors.NewUnauthorizedError("E-mail ou senha incorretos")
	}

	if !user.IsActive {
		return nil, apperrors.NewForbiddenError("Usuário inativo. Entre em contato com o administrador!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, apperrors.NewUnauthorizedError("E-mail ou senha incorretos!")
	}

	token, err := uc.JWTService.GenerateToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao gerar token de autenticação!", err)
	}

	return &LoginOutputDTO{
		User: UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			CIM:      user.CIM,
			Degree:   string(user.Degree),
			IsActive: user.IsActive,
			IsAdmin:  user.IsAdmin,
		},
		Token: token,
	}, nil
}
