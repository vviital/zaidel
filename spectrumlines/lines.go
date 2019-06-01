package spectrumlines

import (
	"context"
	"math"
	"sync"

	"github.com/vviital/zaidel/database"
	"github.com/vviital/zaidel/peaks"
)

// Settings represents settings for the spectrum line search
type Settings struct {
	MaxIntensity              float64
	MaxIonizationLevel        int
	MinIntensity              float64
	SearchInMostSuitableGroup bool
	WaveLengthRange           float64
	waveLengthCoefficient     float64
}

// Element struct
type Element struct {
	database.ElementDefinition
	IsSearchCriteriaMatched bool
	Distance                float64
}

func (s *Settings) setWaveLengthCoefficient(peaks []peaks.Peak) {
	s.waveLengthCoefficient = 1

	for _, p := range peaks {
		if p.Point.X > 2000 {
			s.waveLengthCoefficient = 0.1
			break
		}
	}
}

// DefaultSettings returns default settings
func DefaultSettings() Settings {
	return Settings{
		WaveLengthRange:           0.05,
		MaxIntensity:              100,
		MinIntensity:              1,
		SearchInMostSuitableGroup: false,
		MaxIonizationLevel:        100,
	}
}

// DeterminedPeak represent final result struct
type DeterminedPeak struct {
	peaks.Peak
	Elements []Element
}

func matcheElementAndPeak(peak peaks.Peak, element Element, settings Settings) (float64, bool) {
	baseValue := element.WaveLength / settings.waveLengthCoefficient
	distance := math.Abs(peak.Point.X - baseValue)

	if distance > settings.WaveLengthRange {
		return 0, false
	}

	return 1 - distance/settings.WaveLengthRange, true
}

func getMatchedBySettingsPredicate(settings Settings) database.Predicate {
	return func(element database.ElementDefinition) bool {
		isMatched := (settings.MaxIntensity > 0 && element.Intensity < settings.MaxIntensity) ||
			(settings.MinIntensity > 0 && element.Intensity > settings.MinIntensity) ||
			(settings.MaxIonizationLevel > 0 && element.IonizationStage < settings.MaxIonizationLevel)
		return isMatched
	}
}

func getNotMatchedBySettingsPredicate(settings Settings) database.Predicate {
	isMatched := getMatchedBySettingsPredicate(settings)
	return func(element database.ElementDefinition) bool {
		return !isMatched(element)
	}
}

func combineElements(matched chan database.ElementDefinition, notMatched chan database.ElementDefinition) chan Element {
	elements := make(chan Element, 10)
	var wg sync.WaitGroup

	push := func(els chan database.ElementDefinition, isMatched bool) {
		for e := range els {
			elements <- Element{ElementDefinition: e, IsSearchCriteriaMatched: isMatched}
		}
		wg.Done()
	}

	go func() {
		defer close(elements)
		wg.Add(2)

		go push(matched, true)
		go push(notMatched, false)

		wg.Wait()
	}()

	return elements
}

func processElements(pts []peaks.Peak, settings Settings, elements chan Element) chan DeterminedPeak {
	res := make(chan DeterminedPeak, len(pts))
	detPeaks := make([]DeterminedPeak, len(pts))

	for i, pt := range pts {
		detPeaks[i].Peak = pt
	}

	go func() {
		defer close(res)

		for el := range elements {
			for i, pt := range pts {
				if d, ok := matcheElementAndPeak(pt, el, settings); ok {
					e := el
					e.Distance = d
					detPeaks[i].Elements = append(detPeaks[i].Elements, e)
				}
			}
		}

		for _, peak := range detPeaks {
			res <- peak
		}
	}()

	return res
}

// Find returns enhanced peaks with chemical elements
func Find(ctx context.Context, points []peaks.Peak, settings Settings) (res []DeterminedPeak) {
	settings.setWaveLengthCoefficient(points)

	elements := combineElements(
		database.GetRecordsByPredicate(getMatchedBySettingsPredicate(settings)),
		database.GetRecordsByPredicate(getNotMatchedBySettingsPredicate(settings)),
	)

	detPeaks := processElements(points, settings, elements)

	for p := range detPeaks {
		res = append(res, p)
	}

	return
}
