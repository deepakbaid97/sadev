package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/testBot/config"
	testbot "github.com/testBot/testbot"
)

var Sd config.ShotDictionary
var Fd config.BallingAndFieldingPosition

func main() {
	router := mux.NewRouter()
	//os.Setenv("PORT", "2000")
	port := os.Getenv("PORT")
	preFillData()
	preFillFieldingData()
	tb := &testbot.TestBot{
		Sd: Sd,
	}
	fb := &testbot.FieldBot{
		Fd: Fd,
	}
	fmt.Println("Sdf")
	s := router.PathPrefix("/api/CricketBot").Subrouter()
	s.HandleFunc("/GetNextBall", fb.BowlingHandler)
	s.HandleFunc("/PostBalldata", tb.BatsmanHandler).Methods("POST")
	s.HandleFunc("/Postfieldsetting", testbot.FieldingHandler).Methods("POST")
	s.HandleFunc("/PostLastballStatus", tb.LastballStatus).Methods("POST")
	s.HandleFunc("/Getfieldsetting", fb.SendFieldingSettings)
	s.HandleFunc("/Toss", testbot.TossHandler)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(port), router))
}

func preFillData() {
	jsonFile, err := os.Open("BallToShot.json")

	if err != nil {
		log.Fatal(err)
	}
	byteData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteData, &Sd)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%#v", Sd)

}

func preFillFieldingData() {
	jsonFile, err := os.Open("FieldingPosition.json")

	if err != nil {
		log.Fatal(err)
	}
	byteData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteData, &Fd)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", Fd)
}
