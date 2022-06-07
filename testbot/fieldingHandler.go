package testbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/testBot/config"
)

func FieldingHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	var fielderData []config.FieldPosition

	err = json.Unmarshal(reqBody, &fielderData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%#v", fielderData)

	w.WriteHeader(200)
	fmt.Println("Endpint Hit: FieldingHangler")
}
