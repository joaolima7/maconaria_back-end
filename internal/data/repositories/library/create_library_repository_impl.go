package library

import (
	"context"
	"database/sql"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

type CreateLibraryRepositoryImpl struct {
	queries *db.Queries
}

func NewCreateLibraryRepositoryImpl(queries *db.Queries) *CreateLibraryRepositoryImpl {
	return &CreateLibraryRepositoryImpl{queries: queries}
}

func (r *CreateLibraryRepositoryImpl) CreateLibrary(library *entity.Library) (*entity.Library, error) {
	ctx := context.Background()

	existingByTitle, err := r.queries.GetLibraryByTitle(ctx, library.Title)
	if err == nil && existingByTitle.Title == library.Title {
		return nil, apperrors.NewDuplicateError("título", library.Title)
	}

	params := db.CreateLibraryParams{
		ID:               library.ID,
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
	}

	_, err = r.queries.CreateLibrary(ctx, params)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "criar biblioteca")
	}

	libraryDB, err := r.queries.GetLibraryByID(ctx, library.ID)
	if err != nil {
		return nil, apperrors.WrapDatabaseError(err, "buscar biblioteca criada")
	}

	return dbLibraryToEntity(libraryDB), nil
}

func dbLibraryToEntity(libraryDB db.Library) *entity.Library {
	fileURL := ""
	if libraryDB.FileUrl.Valid {
		fileURL = libraryDB.FileUrl.String
	}

	coverURL := ""
	if libraryDB.CoverUrl.Valid {
		coverURL = libraryDB.CoverUrl.String
	}

	link := ""
	if libraryDB.Link.Valid {
		link = libraryDB.Link.String
	}

	return &entity.Library{
		ID:               libraryDB.ID,
		Title:            libraryDB.Title,
		SmallDescription: libraryDB.SmallDescription,
		Degree:           entity.UserDegree(libraryDB.Degree),
		FileURL:          fileURL,
		CoverURL:         coverURL,
		Link:             link,
		CreatedAt:        libraryDB.CreatedAt.Time,
		UpdatedAt:        libraryDB.UpdatedAt.Time,
	}
}
