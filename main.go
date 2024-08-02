package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"new-aspect/practice-memos/api"
	"new-aspect/practice-memos/store"
)

func init() {
	conf := zap.NewProductionConfig()
	conf.DisableStacktrace = true
	conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logg, _ := conf.Build()

	zap.ReplaceGlobals(logg)
	defer logg.Sync()
}
func main() {
	store.InitDBConn()

	r := mux.NewRouter().StrictSlash(true)

	api.RegisterUserRoutes(r)
	api.RegisterAuthRoutes(r)

	http.ListenAndServe("localhost:8080", r)
}
