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

	var ballData config.LastBallDataModel

	err = json.Unmarshal(reqBody, &ballData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%#v", ballData)
	w.WriteHeader(100)
	fmt.Println("Endpint Hit: HomePage")
}
