package testbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	config "github.com/testBot/config"
)

type TestBot struct {
	Sd config.ShotDictionary
}

var ZoneKey string = ""
var ShotChosen int = int(config.Defensiveshot)
var BatsmenCount = 0

func (t *TestBot) BatsmanHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	var ballData config.BallModel

	err = json.Unmarshal(reqBody, &ballData)

	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Printf("%s", ballData.BowlerName)

	ZoneKey = strconv.Itoa(int(ballData.Zone)) + strconv.Itoa(int(ballData.BowingType))

	val, ok := t.Sd[ZoneKey]

	if ok {
		copyOfVal := val
		for _, item := range CurrentFieldingData {
			for k := range copyOfVal {
				if item.Prvshot == k {
					delete(copyOfVal, k)
				}
			}
		}

		fmt.Println(copyOfVal, "---- COPY OF VAL")

		if len(copyOfVal) == 0 {
			highestRank := 0
			ShotChosen = int(config.Uppercut)
			for k, v := range val {
				for _, rank := range v {
					if rank >= highestRank {
						ShotChosen = int(k)
					}
				}
			}
		} else {
			for k := range val {
				ShotChosen = int(k)
			}
		}
	}

	fmt.Println(val, "--- VLAUE")

	newBatSpeed := 120

	if ballData.Speed <= 120 {
		newBatSpeed = 120 - (ballData.Speed-120)/2
	}

	batsmen := []string{"Amit Jain", "Rahul Sah", "Mayank Sekhar", "Priyadarshini Angolkar", "Srinivas Kempanna", "Kiran Adekhandi Krishnamurthy", "Gautam Bhatt", "Amrita Dayal", "Rohit Tiwari", "Shruthi Shah", "Raghavendra Hosur Venkata Krishna"}

	batData := config.BatsmanModel{Shots: config.Shots(ShotChosen), BatSpeed: newBatSpeed, Batsman: batsmen[BatsmenCount]}
	json.NewEncoder(w).Encode(batData)
	fmt.Println("Endpint Hit: POST BatsmanHandler")
}
