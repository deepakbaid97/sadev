package testbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	config "github.com/testBot/config"
)

var BallCount = 0

func (t *FieldBot) BowlingHandler(w http.ResponseWriter, r *http.Request) {

	varBall := t.Fd["BF"]
	smallBall := BallCount % 10
	ballZone := varBall[smallBall].Zone
	ballBowlingType := varBall[smallBall].BowingType
	ballBowlerType := varBall[smallBall].BowlerType
	ballSpeed := varBall[smallBall].Speed

	baller := []string{"Raghavendra Hosur Venkata Krishna", "Shruthi Shah", "Rohit Tiwari", "Amrita Dayal", "Gautam Bhatt", "Kiran Adekhandi Krishnamurthy", "Srinivas Kempanna", "Priyadarshini Angolkar", "Mayank Sekhar", "Rahul Sah", "Amit Jain"}
	ballingData := config.BallModel{Zone: ballZone, BowingType: ballBowlingType, BowlerType: ballBowlerType, Speed: ballSpeed, BowlerName: baller[BallCount/6]}
	json.NewEncoder(w).Encode(ballingData)
	fmt.Println("Endpint Hit: GET Bowling Handler")
	BallCount++
}
