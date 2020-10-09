package main

import (
	"boiler/cmd/api/games"
	"boiler/cmd/api/users"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)


func main() {
	// Get the configuration environment variables.
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPass = os.Getenv("DB_PASS")


	r := chi.NewRouter()
	db, err := sql.Open("postgres", "postgres", cfg.DBUser+":"+cfg.DBPass+":"+cfg.DBHost+":"+cfg.DBPort+")/"+cfg.DBName+"sslmode=disable")
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

