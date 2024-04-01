package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Wendller/clientServerAPI/server/database"
	"github.com/Wendller/clientServerAPI/server/services"
)

type CotationsHandler struct {
	DB *sql.DB
}

type Error struct {
	Message string `json:"message"`
}

func main() {
	cotationsDB, err := database.NewCotationsDB()
	if err != nil {
		log.Fatal(err)
	}
	defer cotationsDB.Close()

	mux := http.NewServeMux()
	mux.Handle("/cotacao", &CotationsHandler{DB: cotationsDB})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func (h *CotationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cotation, err := services.GetUSDToBRLCotation()
	if err != nil {
		error := Error{Message: err.Error()}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}

	err = database.InsertCotation(h.DB, cotation.Bid)
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
