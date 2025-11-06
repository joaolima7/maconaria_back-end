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

	// Converte []byte para sql.NullString
	fileDataNull := sql.NullString{Valid: false}
	if len(library.FileData) > 0 {
		fileDataNull = sql.NullString{
			String: string(library.FileData),
			Valid:  true,
		}
	}

	coverDataNull := sql.NullString{Valid: false}
	if len(library.CoverData) > 0 {
		coverDataNull = sql.NullString{
			String: string(library.CoverData),
			Valid:  true,
		}
	}

	params := db.UpdateLibraryParams{
		Title:            library.Title,
		SmallDescription: library.SmallDescription,
		Degree:           db.LibrariesDegree(library.Degree),
		FileData:         fileDataNull,
		CoverData:        coverDataNull,
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
