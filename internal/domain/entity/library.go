package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Library struct {
	ID               string
	Title            string
	SmallDescription string
	Degree           UserDegree
	FileData         []byte
	CoverData        []byte
	Link             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewLibrary(
	id string,
	title string,
	smallDescription string,
	degree UserDegree,
	fileData []byte,
	coverData []byte,
	link string,
) (*Library, error) {
	library := &Library{
		ID:               id,
		Title:            title,
		SmallDescription: smallDescription,
		Degree:           degree,
		FileData:         fileData,
		CoverData:        coverData,
		Link:             link,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := library.Validate(); err != nil {
		return nil, err
	}

	return library, nil
}

func (l *Library) Validate() error {
	if err := l.ValidateTitle(); err != nil {
		return err
	}
	if err := l.ValidateSmallDescription(); err != nil {
		return err
	}
	if err := l.ValidateDegree(); err != nil {
		return err
	}
	return nil
}

func (l *Library) ValidateTitle() error {
	if len(l.Title) == 0 {
		return apperrors.NewValidationError("título", "O título não pode ser vazio!")
	}
	if len(l.Title) > 255 {
		return apperrors.NewValidationError("título", "O título deve ter no máximo 255 caracteres!")
	}
	return nil
}

func (l *Library) ValidateSmallDescription() error {
	if len(l.SmallDescription) == 0 {
		return apperrors.NewValidationError("descrição", "A descrição não pode ser vazia!")
	}
	return nil
}

func (l *Library) ValidateDegree() error {
	validDegrees := []UserDegree{DegreeApprentice, DegreeCompanion, DegreeMaster}
	for _, valid := range validDegrees {
		if l.Degree == valid {
			return nil
		}
	}
	return apperrors.NewValidationError("grau", "O grau deve ser 'apprentice', 'companion' ou 'master'!")
}

func (l *Library) UpdateFile(fileData []byte) {
	l.FileData = fileData
	l.UpdatedAt = time.Now()
}

func (l *Library) UpdateCover(coverData []byte) {
	l.CoverData = coverData
	l.UpdatedAt = time.Now()
}
