package config

type FieldingModel struct {
	Fp      FieldPosition `json:"fp"`
	Prvshot Shots         `json:"prvshot"`
}

type FieldPosition int

const (
	Z1 FieldPosition = iota
	Z2
	Z3
	Z4
)
