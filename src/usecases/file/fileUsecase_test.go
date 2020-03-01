package file

import (
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
