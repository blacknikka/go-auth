package file

import (
	"errors"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/files"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// for mock
type fakeFileRepository struct {
	repositories.FileRepository
	FakeCreate   func(files.File) (*files.File, error)
	FakeFindByID func(int) (*files.File, error)
}

func (s *fakeFileRepository) Create(file files.File) (*files.File, error) {
	return s.FakeCreate(file)
}

func (s *fakeFileRepository) FindByID(id int) (*files.File, error) {
	return s.FakeFindByID(id)
}

func TestCreate(t *testing.T) {
	t.Run("Create正常系", func(t *testing.T) {
		fileForCreate := files.File{
			ID:   1,
			Name: "fileName.jpg",
			Data: []byte{0x01, 0x02},
		}
		spyFileRepository := &fakeFileRepository{
			FakeCreate: func(file files.File) (*files.File, error) {
				return &fileForCreate, nil
			},
		}

		usecase := NewFileUseCase(spyFileRepository)

		creationResult, err := usecase.Create(fileForCreate)

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if creationResult == nil {
			t.Errorf("Returned File shouldn't be nil: %v", creationResult)
		}
	})
}

func TestFindByID(t *testing.T) {
	t.Run("FindByID正常系", func(t *testing.T) {
		fileForFindByID := files.File{
			ID:   1,
			Name: "fileName.jpg",
			Data: []byte{0x01, 0x02},
		}
		spyFileRepository := &fakeFileRepository{
			FakeFindByID: func(id int) (*files.File, error) {
				return &fileForFindByID, nil
			},
		}

		usecase := NewFileUseCase(spyFileRepository)

		foundResult, err := usecase.FindByID(int(fileForFindByID.ID))

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if foundResult == nil {
			t.Errorf("foundResult shouldn't be nil: %v", foundResult)
		}

		if foundResult.ID != fileForFindByID.ID {
			t.Errorf("The ID should be same between two targets: %v, %v", foundResult.ID, fileForFindByID.ID)
		}
	})

	t.Run("FindByID異常系", func(t *testing.T) {
		fileForFindByID := files.File{
			ID:   1,
			Name: "fileName.jpg",
			Data: []byte{0x01, 0x02},
		}
		spyFileRepository := &fakeFileRepository{
			FakeFindByID: func(id int) (*files.File, error) {
				return nil, errors.New("error")
			},
		}

		usecase := NewFileUseCase(spyFileRepository)

		foundResult, err := usecase.FindByID(int(fileForFindByID.ID + 1))

		if err == nil {
			t.Errorf("err shouldn't be nil: %v", err)
		}

		if foundResult != nil {
			t.Errorf("foundResult should be nil: %v", foundResult)
		}
	})
}
