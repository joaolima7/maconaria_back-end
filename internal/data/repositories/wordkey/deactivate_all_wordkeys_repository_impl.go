package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeactivateAllWordKeysRepositoryImpl struct {
	queries *db.Queries
}

func NewDeactivateAllWordKeysRepositoryImpl(database *sql.DB) *DeactivateAllWordKeysRepositoryImpl {
	return &DeactivateAllWordKeysRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *DeactivateAllWordKeysRepositoryImpl) DeactivateAllWordKeys() error {
	ctx := context.Background()

	err := r.queries.DeactivateAllWordKeys(ctx)
	if err != nil {
		return apperrors.NewInternalError("Erro ao desativar palavras chave!", err)
	}

	return nil
}
