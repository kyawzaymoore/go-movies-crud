package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	model "go-movies-crud/model"

	"github.com/gorilla/mux"
)

var movies []model.Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := model.ResponseModelList[model.Movie]{Code: 1, Message: "success", Data: movies}
	json.NewEncoder(w).Encode(res)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		res := model.ResponseModelList[model.Movie]{Code: 1, Message: "success", Data: movies}
		json.NewEncoder(w).Encode(res)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			res := model.ResponseModel[model.Movie]{Code: 1, Message: "success", Data: item}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil{ //Error not catch need to search more
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, model.Movie{ID: "1", Isbn: "93472", Title: "Movie One", Director: &model.Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, model.Movie{ID: "2", Isbn: "42321", Title: "Movie Two", Director: &model.Director{Firstname: "Fido", Lastname: "Dido"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8001\n")
	log.Fatal(http.ListenAndServe(":8001", r))
}
