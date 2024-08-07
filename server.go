package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
  Id string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Lead *Lead `json:"lead"`
}

type Lead struct{
  Fname string `json:"fname"`
  Lname string `json:"lname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json") 
  json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  
  for index, item := range movies{
    
    if item.Id == params["id"]{
      movies = append(movies[:index], movies[index+1:]...)
      break
    }
  }
}

func getMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for _, item := range movies{
    
    if item.Id == params["id"]{
      json.NewEncoder(w).Encode(item) 
      return
    }
  }
}

func createMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  var movie Movie
  _ = json.NewDecoder(r.Body).Decode(&movie)
  movie.Id = strconv.Itoa(rand.Intn(1000000000))
  movies = append(movies, movie)
  json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range movies{
    if item.Id == params["id"]{
      movies = append(movies[:index], movies[index+1:]...)
      var movie Movie
      _ = json.NewDecoder(r.Body).Decode(&movie)
      movie.Id = params["id"]
      movies = append(movies, movie)
      json.NewEncoder(w).Encode(movie)
      return
    }
  }
}

func main()  {
  r := mux.NewRouter()

  movies = append(movies, Movie{Id:"1", Isbn:"123", Title:"T1", Lead: &Lead{Fname:"L0F", Lname:"L0L"}})
  movies = append(movies, Movie{Id:"2", Isbn:"133", Title:"T4", Lead: &Lead{Fname:"L1F", Lname:"L3L"}})
  movies = append(movies, Movie{Id:"3", Isbn:"103", Title:"T2", Lead: &Lead{Fname:"L5F", Lname:"L1L"}})
  movies = append(movies, Movie{Id:"4", Isbn:"163", Title:"T9", Lead: &Lead{Fname:"L3F", Lname:"L5L"}})

  r.HandleFunc("/movies", getMovies).Methods("GET")
  r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  r.HandleFunc("/movie", createMovie).Methods("POST")
  r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
  r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

  fmt.Println("Starting server at port 8075...")
  log.Fatal(http.ListenAndServe(":8075",r))
}
