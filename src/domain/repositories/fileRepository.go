package repositories

import (
	"github.com/blacknikka/go-auth/domain/models/files"
)

// FileRepository Fileのリポジトリのインターフェース
type FileRepository interface {
	Create(files.File) (*files.File, error)
	FindByID(int) (*files.File, error)
}
