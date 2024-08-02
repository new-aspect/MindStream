package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"new-aspect/practice-memos/common/errors"
	"new-aspect/practice-memos/store"
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

func RegisterUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/user").Subrouter()

	userRouter.Use(AuthCheckerMiddleWare)

	userRouter.HandleFunc("/me", handleGetMyUserInfo).Methods("GET")
}
