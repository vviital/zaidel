package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaveLengthBasedFastFinder(t *testing.T) {
	t.Run("FindElements by wave length", func(t *testing.T) {
		table := generateTestTable()
		finder := newWaveLengthBasedFastFinder(table)

		// check finder initialization
		typedFinder, ok := finder.(*waveLengthBasedFastFinder)
		assert.True(t, ok)
		assert.Len(t, typedFinder.table, 3)
		// // element A
		assert.Equal(t, float64(10), typedFinder.table["A"][0].WaveLength)
		assert.Equal(t, float64(20), typedFinder.table["A"][1].WaveLength)
		assert.Equal(t, float64(30), typedFinder.table["A"][2].WaveLength)
		assert.Equal(t, float64(40), typedFinder.table["A"][3].WaveLength)
		// // element B
		assert.Equal(t, float64(5), typedFinder.table["B"][0].WaveLength)
		assert.Equal(t, float64(10), typedFinder.table["B"][1].WaveLength)
		assert.Equal(t, float64(20), typedFinder.table["B"][2].WaveLength)
		// element C
		assert.Equal(t, float64(1), typedFinder.table["C"][0].WaveLength)
		assert.Equal(t, float64(2), typedFinder.table["C"][1].WaveLength)
		assert.Equal(t, float64(5), typedFinder.table["C"][2].WaveLength)

		// check found elements
		var elements Elements
		for e := range finder.FindElements(SearchCriteria{
			MinWaveLength: 10,
			MaxWaveLength: 30,
		}) {
			elements = append(elements, e)
		}

		assert.Len(t, elements, 5)
	})

	t.Run("FindElements by wave length and elements name", func(t *testing.T) {
		table := generateTestTable()
		finder := newWaveLengthBasedFastFinder(table)

		// check finder initialization
		typedFinder, ok := finder.(*waveLengthBasedFastFinder)
		assert.True(t, ok)
		assert.Len(t, typedFinder.table, 3)
		// // element A
		assert.Equal(t, float64(10), typedFinder.table["A"][0].WaveLength)
		assert.Equal(t, float64(20), typedFinder.table["A"][1].WaveLength)
		assert.Equal(t, float64(30), typedFinder.table["A"][2].WaveLength)
		assert.Equal(t, float64(40), typedFinder.table["A"][3].WaveLength)
		// // element B
		assert.Equal(t, float64(5), typedFinder.table["B"][0].WaveLength)
		assert.Equal(t, float64(10), typedFinder.table["B"][1].WaveLength)
		assert.Equal(t, float64(20), typedFinder.table["B"][2].WaveLength)
		// element C
		assert.Equal(t, float64(1), typedFinder.table["C"][0].WaveLength)
		assert.Equal(t, float64(2), typedFinder.table["C"][1].WaveLength)
		assert.Equal(t, float64(5), typedFinder.table["C"][2].WaveLength)

		// check found elements
		var elements Elements
		for e := range finder.FindElements(SearchCriteria{
			MinWaveLength: 10,
			MaxWaveLength: 30,
			Elements: map[string]struct{}{
				"B": struct{}{},
				"C": struct{}{},
			},
		}) {
			elements = append(elements, e)
		}

		assert.Len(t, elements, 2)
	})
}

// Query [10, 30]
func generateTestTable() (table TablePerElement) {
	table = make(TablePerElement, 3)

	table["A"] = Elements{
		Element{
			WaveLength: 30,
			Name:       "A",
		},
		Element{
			WaveLength: 40,
			Name:       "A",
		},
		Element{
			WaveLength: 10,
			Name:       "A",
		},
		Element{
			WaveLength: 20,
			Name:       "A",
		},
	}
	table["B"] = Elements{
		Element{
			WaveLength: 5,
			Name:       "B",
		},
		Element{
			WaveLength: 10,
			Name:       "B",
		},
		Element{
			WaveLength: 20,
			Name:       "B",
		},
	}
	table["C"] = Elements{
		Element{
			WaveLength: 5,
			Name:       "C",
		},
		Element{
			WaveLength: 1,
			Name:       "C",
		},
		Element{
			WaveLength: 2,
			Name:       "C",
		},
	}
	return
}
