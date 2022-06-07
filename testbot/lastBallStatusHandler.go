package testbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/testBot/config"
)

func LastballStatus(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(100)
	fmt.Println("Endpint Hit: HomePage")
}
