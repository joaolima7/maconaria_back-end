package storage

import "context"

type StorageService interface {
	UploadImage(imageData []byte, filename string, folder string) (string, error)
	UploadImageWithContext(ctx context.Context, imageData []byte, filename string, folder string) (string, error)
	UploadPDF(pdfData []byte, filename string, folder string) (string, error)
	DeleteImage(imageURL string, folder string) error
	DeletePDF(pdfURL string, folder string) error
}
