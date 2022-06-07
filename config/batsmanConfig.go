package config

type BatsmanModel struct {
	Shots    Shots  `json:"shots"`
	BatSpeed int    `json:"batSpeed"`
	Batsman  string `json:"batsman"`
}

type Shots int

const (
	leave Shots = iota
	Defensiveshot
	Cut
	squarecut
	latecut
	Squaredrive
	pull
	hook
	Coverdrive
	Offdrive
	Straightdrive
	Ondrive
	Sweep
	Uppercut
)
