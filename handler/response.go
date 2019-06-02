package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vviital/zaidel/spectrumlines"

	"github.com/vviital/zaidel/peaks"
)

// V1ResponseSpectrumPeaks represents peaks response to the client
type V1ResponseSpectrumPeaks struct {
	Settings               peaks.PeakSearchSettings `json:"settings"`
	PeaksCount             int                      `json:"peaksCount"`
	OriginalPeaksPositions []float64                `json:"originalPeaksPositions,omitempty"`
	Peaks                  []peaks.Peak             `json:"peaks,omitempty"`
	SpectrumArea           float64                  `json:"spectrumArea"`
}

// V1ResponseSpectrumPeaksFromFinderResult remap inner data structure to the client's expected result
func V1ResponseSpectrumPeaksFromFinderResult(result peaks.FinderResult, settings peaks.PeakSearchSettings) (resp V1ResponseSpectrumPeaks) {
	resp.OriginalPeaksPositions = result.OriginalPeaksPositions
	resp.Peaks = result.Peaks
	resp.PeaksCount = result.PeaksCount
	resp.Settings = settings
	resp.SpectrumArea = result.SpectrumArea
	return
}

// V1ResponseSpectrumLines represents spectrum lines response to the client
type V1ResponseSpectrumLines struct {
	Settings         spectrumlines.Settings         `json:"settings"`
	PeaksCount       int                            `json:"peaksCount"`
	PeakWithElements []spectrumlines.DeterminedPeak `json:"peakWithElements"`
}

// V1ResponseSpectrumLinesFromFinderResult remap inner data structure to the client's expected result
func V1ResponseSpectrumLinesFromFinderResult(result []spectrumlines.DeterminedPeak, settings spectrumlines.Settings) (resp V1ResponseSpectrumLines) {
	resp.PeakWithElements = result
	resp.PeaksCount = len(result)
	resp.Settings = settings
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
