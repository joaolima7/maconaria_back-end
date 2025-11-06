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

	// Verifica duplicação por título
	existingByTitle, err := r.queries.GetLibraryByTitle(ctx, library.Title)
	if err == nil && existingByTitle.Title == library.Title {
		return nil, apperrors.NewDuplicateError("título", library.Title)
	}

	// Converte []byte para sql.NullString (base64)
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

	params := db.CreateLibraryParams{
		ID:               library.ID,
		Title:            library.Title,
		SmallDescription: library.SmallDescription,
		Degree:           db.LibrariesDegree(library.Degree),
		FileData:         fileDataNull,
		CoverData:        coverDataNull,
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
	link := ""
	if libraryDB.Link.Valid {
		link = libraryDB.Link.String
	}

	// Converte sql.NullString de volta para []byte
	var fileData []byte
	if libraryDB.FileData.Valid {
		fileData = []byte(libraryDB.FileData.String)
	}

	var coverData []byte
	if libraryDB.CoverData.Valid {
		coverData = []byte(libraryDB.CoverData.String)
	}

	return &entity.Library{
		ID:               libraryDB.ID,
		Title:            libraryDB.Title,
		SmallDescription: libraryDB.SmallDescription,
		Degree:           entity.UserDegree(libraryDB.Degree),
		FileData:         fileData,
		CoverData:        coverData,
		Link:             link,
		CreatedAt:        libraryDB.CreatedAt.Time,
		UpdatedAt:        libraryDB.UpdatedAt.Time,
	}
}
