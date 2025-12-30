package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateWordKeyRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateWordKeyRepositoryImpl(database *sql.DB) *CreateWordKeyRepositoryImpl {
	return &CreateWordKeyRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *CreateWordKeyRepositoryImpl) CreateWordKey(wordkey *entity.WordKey) (*entity.WordKey, error) {
	ctx := context.Background()

	_, err := r.queries.CreateWordKey(ctx, db.CreateWordKeyParams{
		ID:      wordkey.ID,
		Wordkey: wordkey.WordKey,
		Active:  wordkey.Active,
	})
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao criar palavra chave!", err)
	}

	return wordkey, nil
}
