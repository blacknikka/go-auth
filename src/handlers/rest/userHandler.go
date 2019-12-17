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

// NewUserHandler Userのハンドラを返す
func NewUserHandler(uu usecases.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// Index UserのIndexの処理
func (uh userHandler) Index(w http.ResponseWriter, r *http.Request) {
	type emailField struct {
		Email string `json:"email"`
	}

	type userField struct {
		Name  string     `json:"name"`
		Email emailField `json:"email"`
	}

	type response struct {
		Users []userField `json`
	}

	ctx := r.Context()
	users, err := uh.userUseCase.GetAll(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(response)
	for _, user := range users {
		uf := &userField{
			Name: user.Name,
			Email: emailField{
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
