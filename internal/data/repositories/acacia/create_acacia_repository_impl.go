package acacia

import (
	"context"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateAcaciaRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateAcaciaRepositoryImpl(queries *db.Queries) *CreateAcaciaRepositoryImpl {
	return &CreateAcaciaRepositoryImpl{queries: queries}
}

func (r *CreateAcaciaRepositoryImpl) CreateAcacia(acacia *entity.Acacia) (*entity.Acacia, error) {
	ctx := context.Background()

	existingByName, err := r.queries.GetAcaciaByName(ctx, acacia.Name)
	if err == nil && existingByName.Name == acacia.Name {
		return nil, apperrors.NewDuplicateError("nome", acacia.Name)
	}

	termsJSON, err := acacia.TermsToJSON()
	if err != nil {
		return nil, apperrors.NewValidationError("mandatos", "Formato de mandatos inválido!")
	}

	params := db.CreateAcaciaParams{
		ID:          acacia.ID,
		Name:        acacia.Name,
		Terms:       []byte(termsJSON),
		IsPresident: acacia.IsPresident,
		Deceased:    acacia.Deceased,
		ImageData:   acacia.ImageData,
	}

	_, err = r.queries.CreateAcacia(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar acácia")
	}

	acaciaDB, err := r.queries.GetAcaciaByID(ctx, acacia.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar acácia criada")
	}

	return dbAcaciaToEntity(acaciaDB)
}

func dbAcaciaToEntity(acaciaDB db.Acacia) (*entity.Acacia, error) {
	terms, err := entity.TermsFromJSON(string(acaciaDB.Terms))
	if err != nil {
		return nil, apperrors.NewValidationError("mandatos", "Erro ao processar mandatos!")
	}

	return &entity.Acacia{
		ID:          acaciaDB.ID,
		Name:        acaciaDB.Name,
		Terms:       terms,
		IsPresident: acaciaDB.IsPresident,
		Deceased:    acaciaDB.Deceased,
		ImageData:   acaciaDB.ImageData,
		CreatedAt:   acaciaDB.CreatedAt.Time,
		UpdatedAt:   acaciaDB.UpdatedAt.Time,
	}, nil
}
