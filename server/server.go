package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Wendller/clientServerAPI/server/database"
	"github.com/Wendller/clientServerAPI/server/services"
)

var SERVER_PORT string = ":8080"

type Error struct {
	Message string `json:"message"`
}

func Init() {
	http.HandleFunc("/cotacao", GetCotationHandler)
	http.ListenAndServe(SERVER_PORT, nil)
}

func GetCotationHandler(w http.ResponseWriter, r *http.Request) {
	cotationsDB := database.NewCotationsDB().DB

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	cotation, err := services.GetUSDToBRL(ctx)
	if err != nil {
		error := Error{Message: err.Error()}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = database.InsertCotation(ctx, cotationsDB, cotation.Bid)
	if err != nil {
		error := Error{Message: err.Error()}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotation)
}
