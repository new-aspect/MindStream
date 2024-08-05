package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"new-aspect/MindStream/common/errors"
	"new-aspect/MindStream/store"
)

func handleGetMyUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserIdInCookie(r)

	user, err := store.GetUserById(userId)

	if err != nil {
		zap.S().Errorf("数据库报错 %v", err)
		errors.ErrorHandler(w, "DATABASE_ERROR")
		return
	}

	json.NewEncoder(w).Encode(user)
}

type UpdateUser struct {
	Username   string
	Password   string
	GithubName string
	WxOpenId   string
}

// type 我猜测这个方法还没有写完
func handleUpdateMyUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, _ := GetUserIdInCookie(r)

	user, err := store.GetUserById(userId)
	if err != nil {
		zap.S().Errorf("数据库报错 %v", err)
		errors.ErrorHandler(w, "DATABASE_ERROR")
		return
	}

	var updateUser UpdateUser
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		zap.S().Errorf("报错 %v", err)
		errors.ErrorHandler(w, "REQUEST_BODY_ERROR")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func RegisterUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/user").Subrouter()

	userRouter.Use(AuthCheckerMiddleWare)

	userRouter.HandleFunc("/me", handleGetMyUserInfo).Methods("GET")
	userRouter.HandleFunc("/me", handleUpdateMyUserInfo).Methods("PATCH")
}
