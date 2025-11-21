package storage

type StorageService interface {
	UploadImage(imageData []byte, filename string, folder string) (string, error)
	UploadPDF(pdfData []byte, filename string, folder string) (string, error)
	DeleteImage(imageURL string, folder string) error
	DeletePDF(pdfURL string, folder string) error
}
