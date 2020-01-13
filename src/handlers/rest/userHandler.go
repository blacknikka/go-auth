package rest

import (
	"encoding/json"
	"net/http"

	"github.com/blacknikka/go-auth/usecases"
)

// UserHandler Userのハンドラ
type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecases.UserUseCase
}

// ------------------

// EmailField email
type EmailField struct {
	Email string `json:"email"`
}

// UserField ユーザー情報
type UserField struct {
	Name  string     `json:"name"`
	Email EmailField `json:"email"`
}

// Response ユーザー情報配列
type Response struct {
	Users []UserField `json`
}

// NewUserHandler Userのハンドラを返す
func NewUserHandler(uu usecases.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// Index UserのIndexの処理
func (uh userHandler) Index(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := uh.userUseCase.GetAll(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(Response)
	for _, user := range users {
		uf := &UserField{
			Name: user.Name,
			Email: EmailField{
				Email: user.Email,
			},
		}
		res.Users = append(res.Users, *uf)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
