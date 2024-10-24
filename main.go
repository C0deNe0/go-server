package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string`json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID ==params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}

}
	func getMovie(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","application/json")
		params:= mux.Vars(r)
		for _, item := range movies {
			if item.ID == params["id"]{
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
	return

}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID =params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

var movies []Movie

func main() {
 r := mux.NewRouter()

 movies = append(movies, Movie{ID: "1",Isbn: "1234",Title: "one", Director: &Director{firstName:"neo",lastName:"jo"}})
 movies = append(movies, Movie{ID: "2",Isbn: "12345",Title: "two", Director: &Director{firstName:"neon",lastName:"yo"}})
 
 r.HandleFunc("/movies",getMovies).Methods("GET")
 r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
 r.HandleFunc("/movies",createMovie).Methods("POST")
 r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
 r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

 fmt.Printf("starting  server at port 8000\n")
 log.Fatal(http.ListenAndServe(":8000",r))
}
