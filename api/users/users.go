package users

import (
	"boiler/cmd/cmd/models"
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
	"github.com/volatiletech/sqlboiler/queries"
)

var db *sql.DB

func New(db *sql.DB, r *chi.Mux) {
	r.Get("/api/v1/users", HandleGetAllUsers(db))
	r.Get("/api/v1/users/all", HandleGetAll(db))
	r.Get("/api/v1/games/single/{id}", HandleGetUser(db))
}

type Results struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Background_Image string `json:"background_image"`
}

type Games struct {
	Next    string    `json:"next"`
	Results []Results `json:"results"`
}

func HandleGetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var json = jsoniter.ConfigFastest
		url := "https://rawg-video-games-database.p.rapidapi.com/games?page=1&page_size=40&filter=true&comments=true"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "rawg-video-games-database.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "712a907377msha629f432939448ep111056jsn177dfc17e2f3")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("rerr %v", err)
		}
		defer res.Body.Close()
		var results Games
		if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
			log.Fatalf("errs %v", err)
			return
		}
		jsonData, err := json.Marshal(results)
		if err != nil {
			log.Fatalf("error %v", err)
			return
		}

		w.Write(jsonData)
	}
}

func HandleGetAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var json = jsoniter.ConfigFastest
		ctx := context.Background()

		users, err := models.Users().All(ctx, db)
		if err != nil {
			log.Fatalf("err %v", err)
			return
		}
		jsonData, err := json.Marshal(users)
		if err != nil {
			log.Fatalf("err %v", err)
			return
		}
		w.Write(jsonData)
	}
}

func HandleGetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var json = jsoniter.ConfigFastest

		ctx := context.Background()
		params := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(params)

		var single models.User
		err := queries.Raw(`select id, name from users where id = $1`, id).Bind(ctx, db, &single)
		if err != nil {
			log.Fatalf("err %v", err)
			return
		}

		jsonData, err := json.Marshal(single)
		if err != nil {
			log.Fatalf("err %v", err)
			return
		}

		w.Write(jsonData)
	}
}
