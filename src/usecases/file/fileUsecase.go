package usecases

import (
	"errors"

	"github.com/blacknikka/go-auth/domain/models/files"
	"github.com/blacknikka/go-auth/domain/repositories"
)

var (
	// ErrConnectingToDB is an error message for connection failed.
	ErrConnectingToDB = errors.New("Connect to DB failed.")
)

// FileUseCase Userのユースケース
type FileUseCase interface {
	Create(files.File) (*files.File, error)
	FindByID(int) (*files.File, error)
}

type fileUseCase struct {
	fileRepository repositories.FileRepository
}

// NewFileUseCase Fileのユースケースを取得する
func NewFileUseCase(fileRepository repositories.FileRepository) FileUseCase {
	return &fileUseCase{
		fileRepository: fileRepository,
	}
}

// Create Fileを登録する
func (fu fileUseCase) Create(file files.File) (*files.File, error) {
	return fu.fileRepository.Create(file)
}

// FindByID Fileを探す
func (fu fileUseCase) FindByID(id int) (*files.File, error) {
	return fu.fileRepository.FindByID(id)
}
