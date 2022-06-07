package testbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	config "github.com/testBot/config"
)

func BatsmanHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	var ballData config.BallModel

	err = json.Unmarshal(reqBody, &ballData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%s", ballData.BowlerName)

	batData := config.BatsmanModel{Shots: 6, BatSpeed: 38, Batsman: "Deepak Baid"}
	json.NewEncoder(w).Encode(batData)
	fmt.Println("Endpint Hit: BatsmanHandler")
}
