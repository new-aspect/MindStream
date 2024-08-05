package common

import (
	"github.com/google/uuid"
	"time"
)

func GetNowDateTimeStr() string {
	return time.Now().Format(time.RFC3339)
}

func GetUUID() string {
	return uuid.New().String()
}
