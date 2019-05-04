package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  . "github.com/raidnav/movies/config"
  . "github.com/raidnav/movies/dao"
  . "github.com/raidnav/movies/models"
  "gopkg.in/mgo.v2/bson"
  "log"
  "net/http"
)

var config = Config{}
var dao = MoviesDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
  respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
  movies, err := dao.FindAll()
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, movies)
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "not implemented yet !")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  var movie Movie
  if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
    respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    return
  }
  movie.ID = bson.NewObjectId()
  dao.Insert(movie)
  respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "not implemented yet !")
}

func init() {
  config.Read()

  dao.Server = config.Server
  dao.Database = config.Database
  dao.Connect()
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
  r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
  r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
  r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
  r.HandleFunc("/movies /{id}", FindMovieEndpoint).Methods("GET")
  if err := http.ListenAndServe(":3000", r); err != nil {
    log.Fatal(err)
  }
}
