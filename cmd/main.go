package main

import (
	"boiler/cmd/api/games"
	"boiler/cmd/api/users"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)


func main() {

	r := chi.NewRouter()
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=timkotowski password=butter333 dbname=gaming sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	// Create a new API using our router.
	users.New(db, r)
	games.New(db, r)

	// Create a new HTTP server.
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Running server...")

	// Start the HTTP server.
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

