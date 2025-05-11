package main

import (
	"context"
	"net/http"
	"encoding/json"
)

func (app *application)GetAllLogs(w http.ResponseWriter, r *http.Request){
	logs, err := app.Logs.GetAll(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)
}