package app

import (
	"music/app/controller"
	"music/db"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Init() {
	a.DB = db.Init()
	a.Router = mux.NewRouter()
	a.setRoutes()
}

func (a *App) setRoutes() {
	a.get("/api/v1/tracks", a.getTracks)
	a.get("/api/v1/queue", a.getQueue)
	a.post("/api/v1/setQueue", a.setQueue)
	a.put("/api/v1/updateQueue", a.updateQueue)
	a.delete("/api/v1/deleteTrack", a.deleteTrack)
	a.post("/api/v1/upload", a.upload)
}

func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) getTracks(w http.ResponseWriter, r *http.Request) {
	controller.GetTracks(a.DB, w, r)
}

func (a *App) getQueue(w http.ResponseWriter, r *http.Request) {
	controller.GetQueue(a.DB, w, r)
}

func (a *App) setQueue(w http.ResponseWriter, r *http.Request) {
	controller.SetQueue(a.DB, w, r)
}

func (a *App) updateQueue(w http.ResponseWriter, r *http.Request) {
	controller.UpdateQueue(a.DB, w, r)
}

func (a *App) deleteTrack(w http.ResponseWriter, r *http.Request) {
	controller.DeleteTrack(a.DB, w, r)
}

func (a *App) upload(w http.ResponseWriter, r *http.Request) {
	controller.Upload(a.DB, w, r)
}
