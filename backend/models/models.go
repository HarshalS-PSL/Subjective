package models

type Settings struct {
	NoOfWells       int      `json:"NoOfWells"`
	NoOfWavelengths int      `json:"NoOfWavelengths"`
	Lm              []string `json:"Lm"`
}
