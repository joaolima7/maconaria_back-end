package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
)

type Worker struct {
	ID                string
	Number            int32
	Name              string
	Registration      string
	BirthDate         time.Time
	InitiationDate    time.Time
	ElevationDate     time.Time
	ExaltationDate    time.Time
	AffiliationDate   time.Time
	InstallationDate  time.Time
	EmeritusMasonDate *time.Time
	ProvectMasonDate  *time.Time
	ImageData         []byte
	Deceased          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewWorker(
	id string,
	number int32,
	name string,
	registration string,
	birthDate time.Time,
	initiationDate time.Time,
	elevationDate time.Time,
	exaltationDate time.Time,
	affiliationDate time.Time,
	installationDate time.Time,
	emeritusMasonDate *time.Time,
	provectMasonDate *time.Time,
	imageData []byte,
	deceased bool,
) (*Worker, error) {
	worker := &Worker{
		ID:                id,
		Number:            number,
		Name:              name,
		Registration:      registration,
		BirthDate:         birthDate,
		InitiationDate:    initiationDate,
		ElevationDate:     elevationDate,
		ExaltationDate:    exaltationDate,
		AffiliationDate:   affiliationDate,
		InstallationDate:  installationDate,
		EmeritusMasonDate: emeritusMasonDate,
		ProvectMasonDate:  provectMasonDate,
		ImageData:         imageData,
		Deceased:          deceased,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := worker.Validate(); err != nil {
		return nil, err
	}

	return worker, nil
}

func (w *Worker) Validate() error {
	if err := w.ValidateNumber(); err != nil {
		return err
	}
	if err := w.ValidateName(); err != nil {
		return err
	}
	if err := w.ValidateRegistration(); err != nil {
		return err
	}
	if err := w.ValidateImage(); err != nil {
		return err
	}
	if err := w.ValidateDates(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) ValidateNumber() error {
	if w.Number <= 0 {
		return apperrors.NewValidationError("número", "O número deve ser maior que zero!")
	}
	return nil
}

func (w *Worker) ValidateName() error {
	if len(w.Name) == 0 {
		return apperrors.NewValidationError("nome", "O nome não pode ser vazio!")
	}
	if len(w.Name) < 3 {
		return apperrors.NewValidationError("nome", "O nome deve ter no mínimo 3 caracteres!")
	}
	return nil
}

func (w *Worker) ValidateRegistration() error {
	if len(w.Registration) == 0 {
		return apperrors.NewValidationError("registro", "O cadastro não pode ser vazio!")
	}
	return nil
}

func (w *Worker) ValidateImage() error {
	if len(w.ImageData) == 0 {
		return apperrors.NewValidationError("imagem", "A imagem é obrigatória!")
	}
	return nil
}

func (w *Worker) ValidateDates() error {
	now := time.Now()

	if w.BirthDate.After(now) {
		return apperrors.NewValidationError("data de nascimento", "A data de nascimento não pode ser futura!")
	}

	if w.InitiationDate.Before(w.BirthDate) {
		return apperrors.NewValidationError("data de iniciação", "A data de iniciação não pode ser anterior à data de nascimento!")
	}

	if w.ElevationDate.Before(w.InitiationDate) {
		return apperrors.NewValidationError("data de elevação", "A data de elevação não pode ser anterior à iniciação!")
	}

	if w.ExaltationDate.Before(w.ElevationDate) {
		return apperrors.NewValidationError("data de exaltação", "A data de exaltação não pode ser anterior à elevação!")
	}

	if w.AffiliationDate.Before(w.BirthDate) {
		return apperrors.NewValidationError("data de afiliação", "A data de filiação não pode ser anterior à data de nascimento!")
	}

	if w.InstallationDate.Before(w.BirthDate) {
		return apperrors.NewValidationError("data de instalação", "A data de instalação não pode ser anterior à data de nascimento!")
	}

	return nil
}

func (w *Worker) UpdateImage(imageData []byte) {
	w.ImageData = imageData
	w.UpdatedAt = time.Now()
}

func (w *Worker) MarkAsDeceased() {
	w.Deceased = true
	w.UpdatedAt = time.Now()
}

func (w *Worker) MarkAsActive() {
	w.Deceased = false
	w.UpdatedAt = time.Now()
}
