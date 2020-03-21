package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vviital/zaidel/spectrumlines"

	"github.com/vviital/zaidel/geometry"
	"github.com/vviital/zaidel/peaks"
)

// V1RequestSpectrumPoint represents spectrum point data
type V1RequestSpectrumPoint struct {
	WaveLength float64 `json:"x"`
	Intensity  float64 `json:"y"`
}

// ToPoint returns point
func (p V1RequestSpectrumPoint) ToPoint() geometry.Coordinate {
	return geometry.Coordinate{
		X: p.WaveLength,
		Y: p.Intensity,
	}
}

// V1RequestSpectrumPoints represents slice of points
type V1RequestSpectrumPoints []V1RequestSpectrumPoint

// ToPoints returns points
func (pts V1RequestSpectrumPoints) ToPoints() geometry.Coordinates {
	res := make(geometry.Coordinates, len(pts))

	for i, p := range pts {
		res[i] = p.ToPoint()
	}

	return res
}

// V1CalculatePeaksRequest object structure
type V1CalculatePeaksRequest struct {
	FileID   string                   `json:"fileID"`
	Settings peaks.PeakSearchSettings `json:"settings"`
	OwnerID  string                   `json:"ownerId"`
}

// V1CalculatePeaksRequestFromRequestBody returns V1CalculatePeaksRequest object from request body
func V1CalculatePeaksRequestFromRequestBody(r *http.Request) (body V1CalculatePeaksRequest) {
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&body)
	return
}

// V1UpdatePeaksRequest object structure
type V1UpdatePeaksRequest struct {
	Peaks   []peaks.Peak `json:"peaks"`
	OwnerID string       `json:"ownerId"`
}

// V1UpdatePeaksRequestFromRequestBody returns V1CalculatePeaksRequest object from request body
func V1UpdatePeaksRequestFromRequestBody(r *http.Request) (body V1UpdatePeaksRequest) {
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&body)
	return
}

// V1CalculateSpectrumLinesRequest object structure
type V1CalculateSpectrumLinesRequest struct {
	Peaks    []peaks.Peak           `json:"peaks"`
	Settings spectrumlines.Settings `json:"settings"`
}

// V1CalculateSpectrumLinesRequestFromRequestBody returns V1CalculateSpectrumLinesRequest object from request body
func V1CalculateSpectrumLinesRequestFromRequestBody(r *http.Request) (body V1CalculateSpectrumLinesRequest) {
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&body)
	return
}

func extractInternalForwardableHeaders(r *http.Request) http.Header {
	var headers = make(http.Header)

	headers.Add("Authorization", r.Header.Get("Authorization"))

	return headers
}
