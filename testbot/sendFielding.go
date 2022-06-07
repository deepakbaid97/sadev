package testbot

import (
	"encoding/json"
	"net/http"

	"github.com/testBot/config"
)

func SendFieldingSettings(w http.ResponseWriter, r *http.Request) {

	Fielders := []config.FieldingModel{
		{Fp: 3, Prvshot: 11},
		{Fp: 3, Prvshot: 6},
		{Fp: 3, Prvshot: 11},
		{Fp: 3, Prvshot: 12},
		{Fp: 2, Prvshot: 11},
		{Fp: 2, Prvshot: 6},
		{Fp: 2, Prvshot: 10},
		{Fp: 2, Prvshot: 12},
		{Fp: 1, Prvshot: 1},
	}

	json.NewEncoder(w).Encode(Fielders)

}
