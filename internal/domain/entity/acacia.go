package entity

import (
	"encoding/json"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Acacia struct {
	ID          string
	Name        string
	Terms       []string
	IsPresident bool
	Deceased    bool
	IsActive    bool
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAcacia(
	id string,
	name string,
	terms []string,
	isPresident bool,
	deceased bool,
	imageURL string,
	isActive bool,
) (*Acacia, error) {
	acacia := &Acacia{
		ID:          id,
		Name:        name,
		Terms:       terms,
		IsPresident: isPresident,
		Deceased:    deceased,
		IsActive:    isActive,
		ImageURL:    imageURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := acacia.Validate(); err != nil {
		return nil, err
	}

	return acacia, nil
}

func (a *Acacia) Validate() error {
	if err := a.ValidateName(); err != nil {
		return err
	}
	if err := a.ValidateImage(); err != nil {
		return err
	}
	if err := a.ValidateTerms(); err != nil {
		return err
	}
	return nil
}

func (a *Acacia) ValidateName() error {
	if len(a.Name) == 0 {
		return apperrors.NewValidationError("nome", "O nome não pode ser vazio!")
	}
	if len(a.Name) < 3 {
		return apperrors.NewValidationError("nome", "O nome deve ter no mínimo 3 caracteres!")
	}
	return nil
}

func (a *Acacia) ValidateImage() error {
	if a.ImageURL == "" {
		return apperrors.NewValidationError("imagem", "A URL da imagem é obrigatória!")
	}
	return nil
}

func (a *Acacia) ValidateTerms() error {
	if a.IsPresident && len(a.Terms) == 0 {
		return apperrors.NewValidationError("mandatos", "Os períodos são obrigatórios para presidentes!")
	}
	return nil
}

func (a *Acacia) UpdateImage(imageURL string) {
	a.ImageURL = imageURL
	a.UpdatedAt = time.Now()
}

func (a *Acacia) MarkAsDeceased() {
	a.Deceased = true
	a.UpdatedAt = time.Now()
}

func (a *Acacia) MarkAsActive() {
	a.Deceased = false
	a.UpdatedAt = time.Now()
}

func (a *Acacia) TermsToJSON() (string, error) {
	if len(a.Terms) == 0 {
		return "[]", nil
	}
	bytes, err := json.Marshal(a.Terms)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func TermsFromJSON(jsonStr string) ([]string, error) {
	if jsonStr == "" || jsonStr == "null" {
		return []string{}, nil
	}
	var terms []string
	err := json.Unmarshal([]byte(jsonStr), &terms)
	if err != nil {
		return []string{}, err
	}
	return terms, nil
}
