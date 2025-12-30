package wordkey

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type DeleteWordKeyRepositoryImpl struct {
	queries *db.Queries
}

func NewDeleteWordKeyRepositoryImpl(database *sql.DB) *DeleteWordKeyRepositoryImpl {
	return &DeleteWordKeyRepositoryImpl{
		queries: db.New(database),
	}
}

func (r *DeleteWordKeyRepositoryImpl) DeleteWordKey(id string) error {
	ctx := context.Background()

	err := r.queries.DeleteWordKey(ctx, id)
	if err != nil {
		return apperrors.NewInternalError("Erro ao deletar palavra chave!", err)
	}

	return nil
}
