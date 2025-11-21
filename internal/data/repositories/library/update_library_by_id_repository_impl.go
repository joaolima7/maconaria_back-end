package library

import (
	"context"
	"database/sql"
	"errors"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type UpdateLibraryByIDRepositoryImpl struct {
	queries *db.Queries
}

func NewUpdateLibraryByIDRepositoryImpl(queries *db.Queries) *UpdateLibraryByIDRepositoryImpl {
	return &UpdateLibraryByIDRepositoryImpl{queries: queries}
}

func (r *UpdateLibraryByIDRepositoryImpl) UpdateLibraryByID(library *entity.Library) (*entity.Library, error) {
	ctx := context.Background()

	_, err := r.queries.GetLibraryByID(ctx, library.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewNotFoundError("Biblioteca")
		}
		return nil, apperrors.WrapDatabaseError(err, "buscar biblioteca")
	}

	params := db.UpdateLibraryParams{
		Title:            library.Title,
		SmallDescription: library.SmallDescription,
		Degree:           db.LibrariesDegree(library.Degree),
		FileUrl: sql.NullString{
			String: library.FileURL,
			Valid:  library.FileURL != "",
		},
		CoverUrl: sql.NullString{
			String: library.CoverURL,
			Valid:  library.CoverURL != "",
		},
		Link: sql.NullString{
			String: library.Link,
			Valid:  library.Link != "",
		},
		ID: library.ID,
	}

	_, err = r.queries.UpdateLibrary(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "atualizar biblioteca")
	}

	libraryDB, err := r.queries.GetLibraryByID(ctx, library.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar biblioteca atualizada")
	}

	return dbLibraryToEntity(libraryDB), nil
}
