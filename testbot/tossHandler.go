package testbot

import (
	"fmt"
	"net/http"
)

func TossHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0")
	fmt.Println("Endpint Hit: Toss Done")
}
