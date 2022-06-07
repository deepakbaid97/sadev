package testbot

import (
	"fmt"
	"net/http"
)

func FieldingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("Endpint Hit: FieldingHangler")
}
