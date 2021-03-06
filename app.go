package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/raidnav/movies/config"
	. "github.com/raidnav/movies/dao"
	. "github.com/raidnav/movies/models"
	. "github.com/raidnav/movies/util"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var config = Config{}
var dao = MoviesDAO{}

func AllMovies(w http.ResponseWriter, _ *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		HttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	HttpResponse(w, http.StatusOK, movies)
}

func FindMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		HttpResponse(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	HttpResponse(w, http.StatusOK, movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		HttpResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	dao.Insert(movie)
	HttpResponse(w, http.StatusCreated, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		HttpResponse(w, http.StatusBadRequest, "Invalid payload request")
		return
	}
	if err := dao.Update(movie); err != nil {
		HttpResponse(w, http.StatusInternalServerError, err.Error())
	}
	HttpResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		HttpResponse(w, http.StatusBadRequest, "Invalid payload request")
		return
	}
	if err := dao.Delete(movie); err != nil {
		HttpResponse(w, http.StatusInternalServerError, err.Error())
	}
	HttpResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies /{id}", FindMovie).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
