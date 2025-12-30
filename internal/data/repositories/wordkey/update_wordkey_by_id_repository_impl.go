package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateWordKeyByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateWordKeyByIDRepositoryImpl(database *sql.DB) *UpdateWordKeyByIDRepositoryImpl {
	return &UpdateWordKeyByIDRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *UpdateWordKeyByIDRepositoryImpl) UpdateWordKeyByID(wordkey *entity.WordKey) (*entity.WordKey, error) {
	ctx := context.Background()

	_, err := r.queries.UpdateWordKey(ctx, db.UpdateWordKeyParams{
		Wordkey: wordkey.WordKey,
		Active:  wordkey.Active,
		ID:      wordkey.ID,
	})
	if err != nil {
		return nil, apperrors.NewInternalError("Erro ao atualizar palavra chave!", err)
	}

	return wordkey, nil
}
