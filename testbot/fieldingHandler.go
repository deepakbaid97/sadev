package testbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/testBot/config"
)

var CurrentFieldingData []config.FieldingModel

func FieldingHandler(w http.ResponseWriter, r *http.Request) {

	var fielderData []config.FieldingModel

	err := json.NewDecoder(r.Body).Decode(&fielderData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	CurrentFieldingData = fielderData

	fmt.Printf(" %#v", fielderData)

	w.WriteHeader(200)
	fmt.Println("Endpint Hit: POST FieldingHangler")
}
