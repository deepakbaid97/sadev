package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	testbot "github.com/testBot/testbot"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("sdf")
	s := router.PathPrefix("/api/CricketBot").Subrouter()
	s.HandleFunc("/GetNextBall", testbot.BowlingHandler)
	s.HandleFunc("/PostBalldata", testbot.BatsmanHandler).Methods("POST")
	s.HandleFunc("/Postfieldsetting", testbot.FieldingHandler).Methods("POST")
	s.HandleFunc("/PostLastballStatus", testbot.LastballStatus).Methods("POST")
	s.HandleFunc("/Getfieldsetting", testbot.BatsmanHandler)
	s.HandleFunc("/Toss", testbot.TossHandler)
	log.Fatal(http.ListenAndServe(":2000", router))
}
