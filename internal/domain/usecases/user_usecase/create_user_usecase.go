package user_usecase

import (
	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
)

type CreateUserInputDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	IsAdmin  bool   `json:"is_admin"`
	IsActive bool   `json:"is_active"`
}

type CreateUserOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	IsActive bool   `json:"is_active"`
}

type CreateUserUseCase struct {
	Repository user.CreateUserRepository
}

func NewCreateUserUseCase(repository user.CreateUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		Repository: repository,
	}
}

func (uc *CreateUserUseCase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user, err := entity.NewUser(
		uuid.New().String(),
		input.Name,
		input.Email,
		input.Password,
		input.IsActive,
		input.IsAdmin,
	)
	if err != nil {
		return nil, err
	}

	userCreated, err := uc.Repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		ID:       userCreated.ID,
		Name:     userCreated.Name,
		Email:    userCreated.Email,
		IsAdmin:  userCreated.IsAdmin,
		IsActive: userCreated.IsActive,
	}, nil
}
