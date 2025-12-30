package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/data/mappers"
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type GetAllWordKeysRepositoryImpl struct {
	queries *db.Queries
}

func NewGetAllWordKeysRepositoryImpl(database *sql.DB) *GetAllWordKeysRepositoryImpl {
	return &GetAllWordKeysRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *GetAllWordKeysRepositoryImpl) GetAllWordKeys() ([]*entity.WordKey, error) {
	ctx := context.Background()

	wordkeysDB, err := r.queries.GetAllWordKeys(ctx)
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao buscar palavras chave!", err)
	}

	wordkeys := make([]*entity.WordKey, len(wordkeysDB))
	for i, wk := range wordkeysDB {
		wordkeys[i] = mappers.DbWordKeyToEntity(wk)
	}

	return wordkeys, nil
}
