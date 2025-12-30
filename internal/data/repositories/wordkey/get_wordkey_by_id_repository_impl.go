package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/data/mappers"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetWordKeyByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewGetWordKeyByIDRepositoryImpl(database *sql.DB) *GetWordKeyByIDRepositoryImpl {
	return &GetWordKeyByIDRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *GetWordKeyByIDRepositoryImpl) GetWordKeyByID(id string) (*entity.WordKey, error) {
	ctx := context.Background()

	wordkeyDB, err := r.queries.GetWordKeyByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.NewNotFoundError("Palavra chave não encontrada!")
		}
		return nil, apperrors.NewInternalError("Erro ao buscar palavra chave!", err)
	}

	return mappers.DbWordKeyToEntity(wordkeyDB), nil
}
