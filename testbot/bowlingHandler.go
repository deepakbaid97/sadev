package testbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	config "github.com/testBot/config"
)

func BowlingHandler(w http.ResponseWriter, r *http.Request) {
	ballingData := config.BallModel{Zone: 1, BowingType: 1, BowlerType: 2, Speed: 130, BowlerName: "Rohit"}
	json.NewEncoder(w).Encode(ballingData)
	fmt.Println("Endpint Hit: Bowling Handler")
}
