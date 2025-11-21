package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/joaolima7/maconaria_back-end/config"
)

type FTPStorageService struct {
	host     string
	port     string
	user     string
	password string
	basePath string
	baseURL  string
	connPool sync.Pool
}

func NewFTPStorageService(cfg *config.Config) *FTPStorageService {
	service := &FTPStorageService{
		host:     cfg.FTPHost,
		port:     cfg.FTPPort,
		user:     cfg.FTPUser,
		password: cfg.FTPPassword,
		basePath: cfg.FTPBasePath,
		baseURL:  cfg.FTPBaseURL,
	}

	service.connPool.New = func() interface{} {
		conn, err := service.createConnection()
		if err != nil {
			log.Printf("Erro ao criar conexão FTP no pool: %v", err)
			return nil
		}
		return conn
	}

	return service
}

func (s *FTPStorageService) createConnection() (*ftp.ServerConn, error) {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao FTP: %w", err)
	}

	if err := conn.Login(s.user, s.password); err != nil {
		conn.Quit()
		return nil, fmt.Errorf("erro ao fazer login no FTP: %w", err)
	}

	return conn, nil
}

func (s *FTPStorageService) getConnection() (*ftp.ServerConn, error) {
	if conn := s.connPool.Get(); conn != nil {
		if ftpConn, ok := conn.(*ftp.ServerConn); ok {

			if err := ftpConn.NoOp(); err == nil {
				return ftpConn, nil
			}
			ftpConn.Quit()
		}
	}
	return s.createConnection()
}

func (s *FTPStorageService) releaseConnection(conn *ftp.ServerConn) {
	if conn != nil {
		s.connPool.Put(conn)
	}
}

func (s *FTPStorageService) ensureDirectoryExists(conn *ftp.ServerConn, dirPath string) error {
	dirPath = strings.TrimSuffix(dirPath, "/")

	err := conn.MakeDir(dirPath)
	if err != nil {
		if !strings.Contains(err.Error(), "550") && !strings.Contains(err.Error(), "exists") {
			log.Printf("Aviso ao criar diretório %s: %v", dirPath, err)
		}
	}
	return nil
}

func (s *FTPStorageService) UploadImage(imageData []byte, filename string, folder string) (string, error) {
	return s.uploadFile(context.Background(), imageData, filename, folder)
}

func (s *FTPStorageService) UploadImageWithContext(ctx context.Context, imageData []byte, filename string, folder string) (string, error) {
	return s.uploadFile(ctx, imageData, filename, folder)
}

func (s *FTPStorageService) UploadPDF(pdfData []byte, filename string, folder string) (string, error) {
	return s.uploadFile(context.Background(), pdfData, filename, folder)
}

func (s *FTPStorageService) uploadFile(ctx context.Context, fileData []byte, filename string, folder string) (string, error) {

	resultChan := make(chan struct {
		url string
		err error
	}, 1)

	go func() {
		conn, err := s.getConnection()
		if err != nil {
			resultChan <- struct {
				url string
				err error
			}{"", err}
			return
		}
		defer s.releaseConnection(conn)

		basePath := strings.TrimSuffix(s.basePath, "/")
		folder = strings.Trim(folder, "/")
		folderPath := basePath + "/" + folder

		if err := s.ensureDirectoryExists(conn, folderPath); err != nil {
			log.Printf("Aviso: não foi possível verificar diretório: %v", err)
		}

		safeName := filepath.Base(filename)
		remotePath := folderPath + "/" + safeName

		reader := bytes.NewReader(fileData)
		if err := conn.Stor(remotePath, reader); err != nil {
			resultChan <- struct {
				url string
				err error
			}{"", fmt.Errorf("erro ao fazer upload do arquivo: %w", err)}
			return
		}

		baseURL := strings.TrimSuffix(s.baseURL, "/")
		fileURL := baseURL + "/" + folder + "/" + safeName

		log.Printf("Arquivo salvo com sucesso: %s", fileURL)

		resultChan <- struct {
			url string
			err error
		}{fileURL, nil}
	}()

	select {
	case result := <-resultChan:
		return result.url, result.err
	case <-ctx.Done():
		return "", fmt.Errorf("upload cancelado: %w", ctx.Err())
	}
}

func (s *FTPStorageService) DeleteImage(imageURL string, folder string) error {
	return s.deleteFile(imageURL, folder)
}

func (s *FTPStorageService) DeletePDF(pdfURL string, folder string) error {
	return s.deleteFile(pdfURL, folder)
}

func (s *FTPStorageService) deleteFile(fileURL string, folder string) error {
	if fileURL == "" {
		return nil
	}

	conn, err := s.getConnection()
	if err != nil {
		return err
	}
	defer s.releaseConnection(conn)

	filename, err := s.extractFilenameFromURL(fileURL)
	if err != nil {
		return err
	}

	basePath := strings.TrimSuffix(s.basePath, "/")
	folder = strings.Trim(folder, "/")
	folderPath := basePath + "/" + folder
	remotePath := folderPath + "/" + filename

	if err := conn.Delete(remotePath); err != nil {
		if !strings.Contains(err.Error(), "550") && !strings.Contains(err.Error(), "No such file") {
			return fmt.Errorf("erro ao deletar arquivo do FTP: %w", err)
		}
		log.Printf("Aviso: arquivo %s não encontrado para deletar", remotePath)
	} else {
		log.Printf("Arquivo deletado com sucesso: %s", remotePath)
	}

	return nil
}

func (s *FTPStorageService) extractFilenameFromURL(fileURL string) (string, error) {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", fmt.Errorf("URL inválida: %w", err)
	}

	parts := strings.Split(parsedURL.Path, "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("URL não contém nome de arquivo")
	}

	filename := parts[len(parts)-1]
	if filename == "" {
		return "", fmt.Errorf("nome de arquivo vazio na URL")
	}

	return filename, nil
}
