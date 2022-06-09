package testbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/testBot/config"
)

type FieldBot struct {
	Fd config.BallingAndFieldingPosition
}

func (t *FieldBot) SendFieldingSettings(w http.ResponseWriter, r *http.Request) {

	Fielders := []config.FieldingModel{}
	varFielders := t.Fd["BF"][BallCount%10].Fielders

	for i := 0; i < 9; i++ {
		fp := config.Z1
		if i%2 == 0 {
			fp = config.Z3
		} else {
			fp = config.Z4
		}
		Fielders = append(Fielders, config.FieldingModel{Fp: fp, Prvshot: config.Shots(varFielders[i])})
	}

	json.NewEncoder(w).Encode(Fielders)

	fmt.Println("Endpint Hit: GET FieldingHangler")

}
