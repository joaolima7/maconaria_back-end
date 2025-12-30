package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/data/mappers"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetWordKeyByActiveRepositoryImpl struct {
	queries *db.Queries
}

func NewGetWordKeyByActiveRepositoryImpl(database *sql.DB) *GetWordKeyByActiveRepositoryImpl {
	return &GetWordKeyByActiveRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *GetWordKeyByActiveRepositoryImpl) GetWordKeyByActive() (*entity.WordKey, error) {
	ctx := context.Background()

	wordkeyDB, err := r.queries.GetWordKeyByActive(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.NewNotFoundError("Nenhuma palavra chave ativa encontrada!")
		}
		return nil, apperrors.NewInternalError("Erro ao buscar palavra chave ativa!", err)
	}

	return mappers.DbWordKeyToEntity(wordkeyDB), nil
}
