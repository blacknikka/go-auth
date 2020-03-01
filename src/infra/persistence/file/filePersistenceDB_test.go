package file

import (
	"os"
	"reflect"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/files"
	"github.com/blacknikka/go-auth/infra/persistence"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func setup() {
	// DBに接続
	conn := persistence.NewConnectToDB(persistence.NewDBConnectionFactory())
	var err error
	db, err = conn.Connect()
	if err != nil {
		panic("Connect to DB failed.")
	}

	db.Delete(&files.File{})
}

func teardown() {
	defer db.Close()
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()

	os.Exit(ret)
}

func TestCreate(t *testing.T) {
	t.Run("Create正常系", func(t *testing.T) {
		setup()

		// DBを空にする
		db.Delete(&files.File{})

		// DBの登録数を取得
		var count = 0
		db.Model(&files.File{}).Count(&count)

		if count != 0 {
			t.Errorf("record count invalid: got %v want %v",
				count, 0)
		}

		// Create
		fileDB := &FilePersistanceDB{db: db}
		file := files.File{
			Name: "file-name1.jpg",
			Data: []byte{0x00, 0x01},
		}
		createdFile, err := fileDB.Create(file)
		if err != nil {
			t.Error("insert error")
		}

		db.Model(&files.File{}).Count(&count)

		if count != 1 {
			t.Errorf("insert error: got %v want %v", count, 1)
		}

		if createdFile.ID <= 0 {
			t.Errorf("created ID should be a positive value: %v", createdFile.ID)
		}

		teardown()
	})
}

func TestFindByID(t *testing.T) {
	t.Run("FindByID正常系", func(t *testing.T) {
		setup()

		// DBを空にする
		db.Delete(&files.File{})

		// DBの登録数を取得
		// Create
		fileDB := &FilePersistanceDB{db: db}
		file := files.File{
			Name: "file-name1.jpg",
			Data: []byte{0x00, 0x01},
		}
		createdFile, _ := fileDB.Create(file)

		// FindByID (test stub)
		found, err := fileDB.FindByID(int(createdFile.ID))

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if found.ID != createdFile.ID {
			t.Errorf("Found ID should be same as created ID: %v, %v", found.ID, createdFile.ID)
		}

		if found.Name != createdFile.Name {
			t.Errorf("Found Name should be same as created Name: %v, %v", found.Name, createdFile.Name)
		}

		if reflect.DeepEqual(found.Data, createdFile.Data) != true {
			t.Errorf("Found Data should be same as created Data: %v, %v", found.Data, createdFile.Data)
		}

		teardown()
	})
}
