package config

type RankOfShot map[string]int

type ShotListForAPerticularBall map[Shots]RankOfShot

type ShotDictionary map[string]ShotListForAPerticularBall

type BF struct {
	Zone       BallPitchZone `json:"zone"`
	BowingType BowlingType   `json:"bowingType"`
	BowlerType BowlerTypes   `json:"bowlerType"`
	Speed      int           `json:"speed"`
	Fielders   []int         `json:"fielders"`
}

type BallingAndFieldingPosition map[string][]BF
