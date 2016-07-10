package main

import (
	"encoding/json"
	"net/http"
)

type (
	IndexController struct{}
)

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c IndexController) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Welcome"); err != nil {
		panic(err)
	}
}
