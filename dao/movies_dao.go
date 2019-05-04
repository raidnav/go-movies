package dao

import (
  . "github.com/raidnav/movies/models"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
)

type MoviesDAO struct {
  Server    string
  Database  string
}

var db *mgo.Database

const (
  COLLECTION = "movies"
)

func (m *MoviesDAO) Connect() {
  session, err := mgo.Dial(m.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]Movie, error) {
  var movies []Movie
  err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
  if err != nil {
    log.Fatal("There was an error when getting all of movies")
  }
  return movies, nil
}

func (m *MoviesDAO) FindById(id string) (Movie, error) {
  var movie Movie
  err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
  if err != nil {
    log.Fatal("There was an error when getting movie with id: ", id)
  }
  return movie, nil
}

func (m *MoviesDAO) Insert(movie Movie) {
  err := db.C(COLLECTION).Insert(&movie)
  if err != nil {
    log.Fatal("There was an error when inserting movie with id:", movie.ID)
  }
}

func (m *MoviesDAO) Delete(movie Movie) {
  err := db.C(COLLECTION).Remove(&movie)
  if err != nil {
    log.Fatal("There was an error when deleting movie with id:", movie.ID)
  }
}

func (m *MoviesDAO) Update(movie Movie) {
  err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
  if err != nil {
    log.Fatal("There was an error when updating movie with id:", movie.ID)
  }
}