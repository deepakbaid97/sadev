package config

type BallModel struct {
	Zone       BallPitchZone `json:"zone"`
	BowingType BowlingType   `json:"bowingType"`
	BowlerType BowlerTypes   `json:"bowlerType"`
	Speed      int           `json:"speed"`
	BowlerName string        `json:"bowlerName"`
}

type BallPitchZone int
type BowlingType int
type BowlerTypes int

const (
	Zone1 BallPitchZone = iota
	Zone2
)

const (
	Bouncer BowlingType = iota
	Outswinger
	Inswingers
	LegBreak
	OffBreak
	Googly
)

const (
	RAF BowlerTypes = iota
	RAFM
	RAS
	OB
	LB
)
