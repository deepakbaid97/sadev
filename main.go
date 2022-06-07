package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	testbot "github.com/testBot/testbot"
)

func main() {
	router := mux.NewRouter()

	port := flag.Int("PORT", -1, "specify a port")
	flag.Parse()

	fmt.Println("sdf")
	s := router.PathPrefix("/api/CricketBot").Subrouter()
	s.HandleFunc("/GetNextBall", testbot.BowlingHandler)
	s.HandleFunc("/PostBalldata", testbot.BatsmanHandler).Methods("POST")
	s.HandleFunc("/Postfieldsetting", testbot.FieldingHandler).Methods("POST")
	s.HandleFunc("/PostLastballStatus", testbot.LastballStatus).Methods("POST")
	s.HandleFunc("/Getfieldsetting", testbot.BatsmanHandler)
	s.HandleFunc("/Toss", testbot.TossHandler)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(*port), router))
}
