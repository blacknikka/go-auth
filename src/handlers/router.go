package handlers

import (
	"net/http"

	fileHTML "github.com/blacknikka/go-auth/handlers/html/file"
	userHTML "github.com/blacknikka/go-auth/handlers/html/user"
	"github.com/blacknikka/go-auth/handlers/rest"
	"github.com/blacknikka/go-auth/infra/persistence"
	filePersistence "github.com/blacknikka/go-auth/infra/persistence/file"
	userPersistence "github.com/blacknikka/go-auth/infra/persistence/user"
	fileUsecase "github.com/blacknikka/go-auth/usecases/file"
	userUsecase "github.com/blacknikka/go-auth/usecases/user"
)

// InitializeRouting routing
func InitializeRouting() {
	conn := persistence.NewConnectToDB(persistence.NewDBConnectionFactory())
	db, err := conn.Connect()
	if err != nil {
		panic(err.Error())
	}

	userPersistence := userPersistence.NewUserPersistence(db)
	usecaseForUser := userUsecase.UserUseCase(userPersistence)
	userHandler := rest.NewUserHandler(usecaseForUser)
	userHTMLHandler := userHTML.NewUserHTMLHandler(usecaseForUser)

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/api/user", userHandler.Index)
	http.HandleFunc("/user/all", userHTMLHandler.ShowAllUsers)

	filePersistence := filePersistence.NewFilePersistence(db)
	usecaseForFile := fileUsecase.FileUseCase(filePersistence)
	fileHTMLHandler := fileHTML.NewFileHTMLHandler(usecaseForFile)

	http.HandleFunc("/file/find_form", fileHTMLHandler.FindFileForm)
	http.HandleFunc("/file/find_request", fileHTMLHandler.FindFileRequest)
	http.HandleFunc("/file/upload_form", fileHTMLHandler.UploadFileForm)
	http.HandleFunc("/file/upload_request", fileHTMLHandler.UploadFileRequest)
}
