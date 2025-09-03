package user_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user_repository"

type UpdateUserPasswordInputDTO struct {
	ID          string `json:"id"`
	NewPassword string `json:"new_password"`
}

type UpdateUserPasswordUseCase struct {
	Repository user_repository.UpdateUserPasswordRepository
}

func NewUpdateUserPasswordUseCase(repository user_repository.UpdateUserPasswordRepository) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		Repository: repository,
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) error {
	return uc.Repository.UpdateUserPassword(input.ID, input.NewPassword)
}
