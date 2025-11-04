package user_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
)

type GetUserByIdInputDTO struct {
	ID string
}

type GetUserByIdOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByIdUseCase struct {
	Repository user.GetUserByIdRepository
}

func NewGetUserByIdUseCase(repository user.GetUserByIdRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		Repository: repository,
	}
}

func (uc *GetUserByIdUseCase) Execute(input GetUserByIdInputDTO) (*GetUserByIdOutputDTO, error) {
	user, err := uc.Repository.GetUserById(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetUserByIdOutputDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}
