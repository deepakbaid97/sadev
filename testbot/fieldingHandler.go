package testbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/testBot/config"
)

func FieldingHandler(w http.ResponseWriter, r *http.Request) {
	// reqBody, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Fprintf(w, err.Error())
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	var fielderData []config.FieldingModel

	// fmt.Printf("___ %s", string(reqBody))

	// err = json.Unmarshal(reqBody, &fielderData)

	err := json.NewDecoder(r.Body).Decode(&fielderData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	// for _, item := range fielderData {

	fmt.Printf(" %#v", fielderData)

	w.WriteHeader(200)
	fmt.Println("Endpint Hit: FieldingHangler")
}
