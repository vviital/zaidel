package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vviital/zaidel/peaks"
	"github.com/vviital/zaidel/spectrumlines"
)

// V1ResponseSpectrumPeaks represents peaks response to the client
type V1ResponseSpectrumPeaks struct {
	OriginalPeaksPositions []float64                `json:"originalPeaksPositions,omitempty"`
	Peaks                  []peaks.Peak             `json:"peaks,omitempty"`
	PeaksCount             int                      `json:"peaksCount"`
	Settings               peaks.PeakSearchSettings `json:"settings"`
	SpectrumArea           float64                  `json:"spectrumArea"`
}

// V1ResponseSpectrumPeaksFromDatasource remap inner data structure to the client's expected result
func V1ResponseSpectrumPeaksResult(peaks peaks.FinderResult, settings peaks.PeakSearchSettings) (resp V1ResponseSpectrumPeaks) {
	resp.OriginalPeaksPositions = peaks.OriginalPeaksPositions
	resp.Peaks = peaks.Peaks
	resp.PeaksCount = peaks.PeaksCount
	resp.Settings = settings
	resp.SpectrumArea = peaks.SpectrumArea
	return
}

// V1ResponseSpectrumLines represents spectrum lines response to the client
type V1ResponseSpectrumLines struct {
	Settings         spectrumlines.Settings         `json:"settings"`
	PeaksCount       int                            `json:"peaksCount"`
	PeakWithElements []spectrumlines.DeterminedPeak `json:"peaksWithElements"`
	AutoSuggestions  []spectrumlines.AutoSuggestion `json:"autoSuggestions"`
}

// V1ResponseSpectrumLinesFromFinderResult remap inner data structure to the client's expected result
func V1ResponseSpectrumLinesFromFinderResult(result []spectrumlines.DeterminedPeak, settings spectrumlines.Settings, suggestions []spectrumlines.AutoSuggestion) (resp V1ResponseSpectrumLines) {
	resp.PeakWithElements = result
	resp.PeaksCount = len(result)
	resp.Settings = settings
	resp.AutoSuggestions = suggestions
	return
}

func sendJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{Message: err.Error()})
}
