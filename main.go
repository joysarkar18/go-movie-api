package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Duration float64   `json:"duration"`
	Actors   []string  `json:"Actors"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies = []Movie{

	{
		Id:       "M3",
		Title:    "Movie 3",
		Duration: 1.78,
		Actors:   []string{"Actor1_3", "Actor2_3", "Actor3_3", "Actor4_3"},
		Director: &Director{FirstName: "Director3", LastName: "Smith"},
	},
	{
		Id:       "M4",
		Title:    "Movie 4",
		Duration: 2.11,
		Actors:   []string{"Actor1_4", "Actor2_4", "Actor3_4", "Actor4_4", "Actor5_4"},
		Director: &Director{FirstName: "Director4", LastName: "Smith"},
	},
	{
		Id:       "M5",
		Title:    "Movie 5",
		Duration: 0.67,
		Actors:   []string{"Actor1_5"},
		Director: &Director{FirstName: "Director5", LastName: "Smith"},
	},
	{
		Id:       "M6",
		Title:    "Movie 6",
		Duration: 1.23,
		Actors:   []string{"Actor1_6", "Actor2_6", "Actor3_6"},
		Director: &Director{FirstName: "Director6", LastName: "Smith"},
	},
	{
		Id:       "M7",
		Title:    "Movie 7",
		Duration: 2.95,
		Actors:   []string{"Actor1_7", "Actor2_7", "Actor3_7", "Actor4_7", "Actor5_7"},
		Director: &Director{FirstName: "Director7", LastName: "Smith"},
	},
	{
		Id:       "M8",
		Title:    "Movie 8",
		Duration: 0.45,
		Actors:   []string{"Actor1_8", "Actor2_8"},
		Director: &Director{FirstName: "Director8", LastName: "Smith"},
	},
	{
		Id:       "M9",
		Title:    "Movie 9",
		Duration: 2.10,
		Actors:   []string{"Actor1_9", "Actor2_9", "Actor3_9", "Actor4_9"},
		Director: &Director{FirstName: "Director9", LastName: "Smith"},
	},
	{
		Id:       "M10",
		Title:    "Movie 10",
		Duration: 2.67,
		Actors:   []string{"Actor1_10", "Actor2_10", "Actor3_10", "Actor4_10"},
		Director: &Director{FirstName: "Director10", LastName: "Smith"},
	},
	{
		Id:       "M11",
		Title:    "Movie 11",
		Duration: 1.34,
		Actors:   []string{"Actor1_11", "Actor2_11", "Actor3_11"},
		Director: &Director{FirstName: "Director11", LastName: "Smith"},
	},
	{
		Id:       "M12",
		Title:    "Movie 12",
		Duration: 0.92,
		Actors:   []string{"Actor1_12", "Actor2_12", "Actor3_12"},
		Director: &Director{FirstName: "Director12", LastName: "Smith"},
	},
	{
		Id:       "M13",
		Title:    "Movie 13",
		Duration: 2.21,
		Actors:   []string{"Actor1_13", "Actor2_13", "Actor3_13"},
		Director: &Director{FirstName: "Director13", LastName: "Smith"},
	},
	{
		Id:       "M14",
		Title:    "Movie 14",
		Duration: 1.56,
		Actors:   []string{"Actor1_14", "Actor2_14", "Actor3_14", "Actor4_14"},
		Director: &Director{FirstName: "Director14", LastName: "Smith"},
	},
	{
		Id:       "M15",
		Title:    "Movie 15",
		Duration: 0.78,
		Actors:   []string{"Actor1_15", "Actor2_15", "Actor3_15"},
		Director: &Director{FirstName: "Director15", LastName: "Smith"},
	},
	{
		Id:       "M16",
		Title:    "Movie 16",
		Duration: 2.43,
		Actors:   []string{"Actor1_16", "Actor2_16", "Actor3_16", "Actor4_16", "Actor5_16"},
		Director: &Director{FirstName: "Director16", LastName: "Smith"},
	},
	{
		Id:       "M17",
		Title:    "Movie 17",
		Duration: 1.89,
		Actors:   []string{"Actor1_17", "Actor2_17", "Actor3_17", "Actor4_17"},
		Director: &Director{FirstName: "Director17", LastName: "Smith"},
	},
	{
		Id:       "M18",
		Title:    "Movie 18",
		Duration: 0.57,
		Actors:   []string{"Actor1_18", "Actor2_18"},
		Director: &Director{FirstName: "Director18", LastName: "Smith"},
	},
	{
		Id:       "M19",
		Title:    "Movie 19",
		Duration: 1.76,
		Actors:   []string{"Actor1_19", "Actor2_19", "Actor3_19", "Actor4_19"},
		Director: &Director{FirstName: "Director19", LastName: "Smith"},
	},
	{
		Id:       "M20",
		Title:    "Movie 20",
		Duration: 2.88,
		Actors:   []string{"Actor1_20", "Actor2_20", "Actor3_20", "Actor4_20"},
		Director: &Director{FirstName: "Director20", LastName: "Smith"},
	},
}

func handleGetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func handelGetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, ele := range movies {
		if vars["id"] == ele.Id {

			json.NewEncoder(w).Encode(ele)
			return

		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, ele := range movies {
		if vars["id"] == ele.Id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(ele)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	_ = json.NewDecoder(r.Body).Decode(&m)
	movies = append(movies, m)
	json.NewEncoder(w).Encode(m)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	_ = json.NewDecoder(r.Body).Decode(&m)
	vars := mux.Vars(r)
	for index, ele := range movies {
		if vars["id"] == ele.Id {
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, m)
			json.NewEncoder(w).Encode(m)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", handleGetAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", handelGetMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/update/{id}", updateMovie).Methods("POST")

	http.ListenAndServe(":8080", r)

}
