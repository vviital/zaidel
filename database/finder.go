package database

// SearchCriteria represents general query to the zaidel database
type SearchCriteria struct {
	Elements      map[string]struct{}
	MaxIntensity  float64
	MaxWaveLength float64
	MinIntensity  float64
	MinWaveLength float64
}

// Zaidel tables interface
type Zaidel interface {
	FindElements(criteria SearchCriteria) chan Element
}
