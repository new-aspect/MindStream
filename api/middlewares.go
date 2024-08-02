package api

import (
	"net/http"
	"new-aspect/practice-memos/common/errors"
)

func AuthCheckerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := GetUserIdInCookie(r)

		if err != nil || userId == "" {
			errors.ErrorHandler(w, "NOT_AUTH")
			return
		}

		next.ServeHTTP(w, r)
	})
}
