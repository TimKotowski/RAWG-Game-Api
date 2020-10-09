package games

import (
	"boiler/cmd/cmd/models"
	"boiler/cmd/error"
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
)



var db *sql.DB

func New(db *sql.DB, r *chi.Mux) {
	r.Get("/api/v1/games", HandleGetAllGames(db))
}


func HandleGetAllGames(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		ctx := context.Background()
		games, err := models.Games().All(ctx, db)
		error.Error(err)

		jsonData, err := json.Marshal(games)
		error.Error(err)

		w.Write(jsonData)
	}
}
