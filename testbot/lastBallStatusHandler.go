package testbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/testBot/config"
)

func (t *TestBot) LastballStatus(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	var statusData config.LastBallDataModel

	err = json.Unmarshal(reqBody, &statusData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%#v", statusData)

	if statusData.Iswicketlost {
		delete(t.Sd[ZoneKey], config.Shots(ShotChosen))
		BatsmenCount++
	}

	// val, ok:= t.Sd[ZoneKey]
	// fmt.Println(ZoneKey, "--- ZoneKey", config.Shots(ShotChosen), "--- SHOTCHOSEN", statusData.Runonlastball, "---statusData.Runonlastball")
	t.Sd[ZoneKey][config.Shots(ShotChosen)]["010"] = statusData.Runonlastball

	w.WriteHeader(100)
	fmt.Println("Endpint Hit: POST LastBallStatus")
}
