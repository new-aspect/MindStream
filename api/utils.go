package api

import "net/http"

func GetUserIdInCookie(r *http.Request) (string, error) {
	userIdCookie, err := r.Cookie("user_id")
	if userIdCookie != nil {
		return userIdCookie.Value, err
	}
	return "", err
}
