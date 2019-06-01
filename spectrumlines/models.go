package spectrumlines

import (
	"encoding/json"

	"github.com/vviital/zaidel/database"
	"github.com/vviital/zaidel/peaks"
)

// Settings represents settings for the spectrum line search
type Settings struct {
	MaxIntensity              float64 `json:"maxIntensity"`
	MaxIonizationLevel        int     `json:"maxIonizationLevel"`
	MinIntensity              float64 `json:"minIntensity"`
	SearchInMostSuitableGroup bool    `json:"searchInMostSuitableGroup"`
	WaveLengthRange           float64 `json:"waveLengthRange"`
}

// Element struct
type Element struct {
	database.Element
	IsSearchCriteriaMatched bool    `json:"isSearchCriteriaMatched"`
	Distance                float64 `json:"similarity"`
}

// DeterminedPeak represent final result struct
type DeterminedPeak struct {
	peaks.Peak
	Elements []Element
}

// ToJSON returns json serialization of peak with matched elements
func (p *DeterminedPeak) ToJSON() string {
	str, _ := json.Marshal(p)
	return string(str)
}
