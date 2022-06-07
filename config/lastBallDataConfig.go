package config

type LastBallDataModel struct {
	IsFieldingValid      bool   `json:"isfieldingvalid"`
	Isballvalid          bool   `json:"isballvalid"`
	Isshotvalid          bool   `json:"isshotvalid"`
	Iswicketlost         bool   `json:"iswicketlost"`
	IsMissed             bool   `json:"isMissed"`
	Runonlastball        int    `json:"runonlastball"`
	Totalruns            int    `json:"totalruns"`
	Totalwicketremaining int    `json:"totalwicketremaining"`
	Totalballremaining   int    `json:"totalballremaining"`
	OutReason            string `json:"outReason"`
}
