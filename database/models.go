package database

import "math"

const epsilon = 0.000000001

// Element represents element's info
type Element struct {
	Intensity       float64
	IonizationStage int
	Name            string
	WaveLength      float64
}

// Elements represent slice of elements
type Elements []Element

// Len is the number of elements in the collection.
func (e Elements) Len() int {
	return len(e)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e Elements) Less(i, j int) bool {
	if e[i].WaveLength < e[j].WaveLength {
		return true
	}
	if math.Abs(e[i].WaveLength-e[j].WaveLength) < epsilon {
		return e[i].Intensity < e[j].Intensity
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (e Elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// TablePerElement represent table per element
type TablePerElement map[string]Elements
