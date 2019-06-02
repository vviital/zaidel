package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vviital/zaidel/peaks"
	"github.com/vviital/zaidel/spectrumlines"

	peaksdatasource "github.com/vviital/zaidel/datasource/peaks"
)

var peaksDatasource peaksdatasource.Peaks

// V1CalculatePeaks responds with calculated peaks to the client
func V1CalculatePeaks(w http.ResponseWriter, r *http.Request) {
	body := V1CalculatePeaksRequestFromRequestBody(r)
	foundPeaks, err := peaks.Find(body.Points.ToPoints(), body.Settings)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	savedPeaks, err := peaksDatasource.InsertOne(r.Context(), foundPeaks, body.Settings, body.OwnerID)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	sendJSON(w, V1ResponseSpectrumPeaksFromDatasource(savedPeaks))
}

// V1GetPeaksByID responds with peaks by ID
func V1GetPeaksByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		sendError(w, errors.New("ID must be provided"), 400)
		return
	}
	foundPeaks, err := peaksDatasource.FindByID(r.Context(), id)
	if err != nil {
		sendError(w, err, 400)
		return
	}
	sendJSON(w, V1ResponseSpectrumPeaksFromDatasource(foundPeaks))
}

// V1UpdatePeaksByID responds with peaks by ID
func V1UpdatePeaksByID(w http.ResponseWriter, r *http.Request) {

}

// V1GetDefaultSettingsForPeaks responds with default settings for the peaks finder to the client
func V1GetDefaultSettingsForPeaks(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, peaks.NewDefaultSettings())
}

// V1CalculateSpectrumLines responds with calculated spectrumlines to the client
func V1CalculateSpectrumLines(w http.ResponseWriter, r *http.Request) {
	body := V1CalculateSpectrumLinesRequestFromRequestBody(r)
	foundLines := spectrumlines.Find(r.Context(), body.Peaks, body.Settings)
	sendJSON(w, V1ResponseSpectrumLinesFromFinderResult(foundLines, body.Settings))
}

// V1GetDefaultSettingsForSpectrumLines responds with default settings for the spectum lines finder to the client
func V1GetDefaultSettingsForSpectrumLines(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, spectrumlines.NewDefaultSettings())
}

// V1GetSpectrumLinesByID responds with peaks by ID
func V1GetSpectrumLinesByID(w http.ResponseWriter, r *http.Request) {

}

// V1UpdateSpectrumLinesByID responds with peaks by ID
func V1UpdateSpectrumLinesByID(w http.ResponseWriter, r *http.Request) {

}

func init() {
	peaksDatasource = peaksdatasource.NewMongoPeaks()
}
