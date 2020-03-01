package file

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-auth/domain/models/files"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// FilePersistanceDB DBでの永続化を行う
type FilePersistanceDB struct {
	db *gorm.DB
}

// NewFilePersistence 永続化Objectを返す
func NewFilePersistence(db *gorm.DB) repositories.FileRepository {
	return &FilePersistanceDB{
		db: db,
	}
}

// Create 登録
func (fileDB FilePersistanceDB) Create(file files.File) (*files.File, error) {
	if result := fileDB.db.NewRecord(file); result == false {
		return nil, errors.New("creating file failed")
	}
	fileDB.db.Create(&file)

	return &file, nil
}

// FindByID Idから取得する
func (fileDB FilePersistanceDB) FindByID(id int) (*files.File, error) {
	file := files.File{}
	fileDB.db.First(&file, id)

	if file.ID <= 0 {
		return nil, errors.New("File not found")
	}

	return &file, nil
}
