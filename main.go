package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"new-aspect/MindStream/api"
	"new-aspect/MindStream/store"
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
	api.RegisterQueryRoutes(r)
	api.RegisterMemoRoutes(r)

	spa := api.SPAHandler{
		StaticPath: "./web/dist",
		IndexPath:  "index.html",
	}
	r.PathPrefix("/").Handler(spa)

	http.ListenAndServe("localhost:8080", r)
}
