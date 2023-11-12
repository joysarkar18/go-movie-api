package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
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
		Title:    "Inception",
		Duration: 148.5,
		Actors:   []string{"Leonardo DiCaprio", "Elliot Page", "Tom Hardy"},
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	},
	{
		Title:    "The Shawshank Redemption",
		Duration: 142.0,
		Actors:   []string{"Tim Robbins", "Morgan Freeman"},
		Director: &Director{FirstName: "Frank", LastName: "Darabont"},
	},
	{
		Title:    "Pulp Fiction",
		Duration: 154.5,
		Actors:   []string{"John Travolta", "Samuel L. Jackson", "Uma Thurman"},
		Director: &Director{FirstName: "Quentin", LastName: "Tarantino"},
	},
	{
		Title:    "The Dark Knight",
		Duration: 152.0,
		Actors:   []string{"Christian Bale", "Heath Ledger", "Aaron Eckhart"},
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	},
	{
		Title:    "Forrest Gump",
		Duration: 142.5,
		Actors:   []string{"Tom Hanks", "Robin Wright", "Gary Sinise"},
		Director: &Director{FirstName: "Robert", LastName: "Zemeckis"},
	},
	{
		Title:    "The Matrix",
		Duration: 136.0,
		Actors:   []string{"Keanu Reeves", "Laurence Fishburne", "Carrie-Anne Moss"},
		Director: &Director{FirstName: "Lana", LastName: "Wachowski"},
	},
	{
		Title:    "The Godfather",
		Duration: 175.0,
		Actors:   []string{"Marlon Brando", "Al Pacino", "James Caan"},
		Director: &Director{FirstName: "Francis Ford", LastName: "Coppola"},
	},
	{
		Title:    "Titanic",
		Duration: 195.5,
		Actors:   []string{"Leonardo DiCaprio", "Kate Winslet", "Billy Zane"},
		Director: &Director{FirstName: "James", LastName: "Cameron"},
	},
	{
		Title:    "Inglourious Basterds",
		Duration: 153.5,
		Actors:   []string{"Brad Pitt", "Christoph Waltz", "Diane Kruger"},
		Director: &Director{FirstName: "Quentin", LastName: "Tarantino"},
	},
	{
		Title:    "The Silence of the Lambs",
		Duration: 118.0,
		Actors:   []string{"Jodie Foster", "Anthony Hopkins", "Scott Glenn"},
		Director: &Director{FirstName: "Jonathan", LastName: "Demme"},
	},
}

func handleGetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", handleGetAllMovies).Methods("GET")

	http.ListenAndServe(":8080", r)

}
