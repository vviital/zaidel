package peaks

import "github.com/vviital/zaidel/geometry"

// SpectrumRecord struct
type SpectrumRecord struct {
	WaveLength float64 `json:"x"`
	Intensity  float64 `json:"y"`
}

// ToPoint returns point from record
func (r SpectrumRecord) ToPoint() geometry.Coordinate {
	return geometry.Coordinate{
		X: r.WaveLength,
		Y: r.Intensity,
	}
}

// Spectrum type
type Spectrum []SpectrumRecord

// GetPoints from the spectrum
func (s Spectrum) GetPoints() (coords geometry.Coordinates) {
	if s == nil {
		return
	}

	for _, r := range s {
		coords = append(coords, r.ToPoint())
	}

	return
}

// PeakSearchSettings contains all availabe settings for the search
type PeakSearchSettings struct {
	AverageWindow           int     `json:"averageWindow"`
	CalculateBackground     bool    `json:"calculateBackground"`
	DeconvolutionIterations int     `json:"deconvolutionIterations"`
	Sigma                   float64 `json:"sigma"`
	SmoothMarkov            bool    `json:"smoothMarkov"`
	Threshold               float64 `json:"threshold"`
}

type searchSettings struct {
	PeakSearchSettings
	Size             int
	NumberIterations int
	SizeExtended     int
}

func newSearchSettings(points geometry.Coordinates, settings PeakSearchSettings) searchSettings {
	size := len(points)
	numberIterations := int(7*settings.Sigma + 0.5)
	sizeExtended := size + 2*numberIterations

	return searchSettings{
		PeakSearchSettings: settings,
		Size:               size,
		NumberIterations:   numberIterations,
		SizeExtended:       sizeExtended,
	}
}

// NewDefaultSettings returns default settings for the search.
func NewDefaultSettings() PeakSearchSettings {
	return PeakSearchSettings{
		AverageWindow:           2,
		CalculateBackground:     true,
		DeconvolutionIterations: 2,
		Sigma:                   2,
		SmoothMarkov:            true,
		Threshold:               10,
	}
}

// Peak struct
type Peak struct {
	Point geometry.Coordinate `json:"peak"`
	Left  geometry.Coordinate `json:"left"`
	Right geometry.Coordinate `json:"right"`
	Area  float64             `json:"area"`
}

// FinderResult struct
type FinderResult struct {
	OriginalPeaksPositions []float64            `json:"originalPeaksPositions,omitempty"`
	PeaksCount             int                  `json:"peaksCount"`
	WaveletData            geometry.Coordinates `json:"waveletData,omitempty"`
	BackgroundData         geometry.Coordinates `json:"backgroundData,omitempty"`
	Peaks                  []Peak               `json:"peaks,omitempty"`
	SpectrumArea           float64              `json:"spectrumArea"`
	Destination            []float64            `json:"destination,omitempty"`
}
