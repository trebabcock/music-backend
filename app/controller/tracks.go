package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetTracks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func GetQueue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func SetQueue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func UpdateQueue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeleteTrack(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func Upload(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(104857600)
	file, handler, err := r.FormFile()
}
