package main

import (
	"github.com/iamkaroko/cinego/internal/adapters/redis"
	"github.com/iamkaroko/cinego/internal/booking"
	"github.com/iamkaroko/cinego/internal/utils"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)
	bookingHandler := booking.NewHandler(svc)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	mux.HandleFunc("GET /movies", listMovies)
	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}

var movies = []movieResponse{
	{"the-hobbit", "The Hobbit", 4, 8},
	{"the-matrix", "The Matrix", 4, 8},
	{"the-matrix-reloaded", "The Matrix Reloaded", 4, 8},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, movies)
}

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seatsPerRow"`
}
