package user_usecase

import "github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"

type GetAllUsersOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type GetAllUsersUseCase struct {
	Repository user.GetAllUsersRepository
}

func NewGetAllUsersUseCase(repository user.GetAllUsersRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{
		Repository: repository,
	}
}

func (uc *GetAllUsersUseCase) Execute() ([]*GetAllUsersOutputDTO, error) {
	users, err := uc.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var output []*GetAllUsersOutputDTO

	for _, out := range users {
		output = append(output, &GetAllUsersOutputDTO{
			ID:       out.ID,
			Name:     out.Name,
			Email:    out.Email,
			IsActive: out.IsActive,
			IsAdmin:  out.IsAdmin,
		})
	}

	return output, nil
}
