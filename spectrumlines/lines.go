package spectrumlines

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/vviital/zaidel/database"
	"github.com/vviital/zaidel/peaks"
)

const epsilon = 0.000000001

var (
	zaidel = database.WaveLengthBasedFinder
)

// NewDefaultSettings returns default settings
func NewDefaultSettings() Settings {
	return Settings{
		MaxElementsPerPeak:        20,
		MaxIntensity:              100,
		MaxIonizationLevel:        100,
		MinIntensity:              1,
		SearchInMostSuitableGroup: false,
		WaveLengthRange:           0.05,
	}
}

func getPeakWithElements(peak peaks.Peak, settings Settings, elements chan database.Element) chan DeterminedPeak {
	ch := make(chan DeterminedPeak)

	go func() {
		defer close(ch)
		detPeak := DeterminedPeak{Peak: peak}

		for el := range elements {
			detPeak.Elements = append(detPeak.Elements, Element{
				Element:                 el,
				Distance:                1 - math.Abs(peak.Point.X-el.WaveLength)/settings.WaveLengthRange,
				IsSearchCriteriaMatched: isMatched(el, settings),
			})
		}

		if detPeak.Elements != nil {
			// sort elements in desc order by similarity to the given peak
			sort.Slice(detPeak.Elements, func(i, j int) bool {
				if diff := detPeak.Elements[i].Distance - detPeak.Elements[j].Distance; diff >= epsilon {
					return true
				} else if math.Abs(diff) <= epsilon {
					return detPeak.Elements[i].Name < detPeak.Elements[j].Name
				}
				return false
			})

			nextSize := len(detPeak.Elements)
			detPeak.TotalElementsCount = nextSize
			if nextSize > settings.MaxElementsPerPeak {
				nextSize = settings.MaxElementsPerPeak
			}
			detPeak.Elements = detPeak.Elements[0:nextSize]
		}

		ch <- detPeak
	}()

	return ch
}

func isMatched(element database.Element, settings Settings) bool {
	return (settings.MaxIntensity > 0 && element.Intensity < settings.MaxIntensity) &&
		(settings.MinIntensity > 0 && element.Intensity > settings.MinIntensity) &&
		(settings.MaxIonizationLevel > 0 && element.IonizationStage < settings.MaxIonizationLevel)
}

// Find returns enhanced peaks with chemical elements
func Find(ctx context.Context, points []peaks.Peak, settings Settings) (res []DeterminedPeak) {
	var channels []chan DeterminedPeak
	res = make([]DeterminedPeak, len(points))

	for _, peak := range points {
		criteria := database.SearchCriteria{
			MinWaveLength: peak.Point.X - settings.WaveLengthRange,
			MaxWaveLength: peak.Point.X + settings.WaveLengthRange,
		}
		c := zaidel.FindElements(criteria)

		channels = append(channels, getPeakWithElements(peak, settings, c))
	}

	for i, ch := range channels {
		res[i] = <-ch
	}

	return
}

type elementWithDistance struct {
	distance float64
	element  Element
}

// AuthMatch matches peaks with spectrum lines
func AuthMatch(peaks []DeterminedPeak) (suggestions []AutoSuggestion) {
	suggestionsPerPeak := make(map[int][]elementWithDistance)

	for index, peak := range peaks {
		linesPerElements := make(map[string]*elementWithDistance)
		for _, element := range peak.Elements {
			if !element.IsSearchCriteriaMatched {
				continue
			}

			currentDistance := math.Abs(peak.Peak.Point.X - element.WaveLength)

			fmt.Println("Current distance:", currentDistance, element.Name)

			if linesPerElements[element.Name] == nil {
				fmt.Println("Appending first", currentDistance)
				value := elementWithDistance{
					distance: currentDistance,
					element:  element,
				}
				linesPerElements[element.Name] = &value
				continue
			}

			if linesPerElements[element.Name].distance > currentDistance {
				value := elementWithDistance{
					distance: currentDistance,
					element:  element,
				}
				linesPerElements[element.Name] = &value
			}
		}

		var candidates []elementWithDistance
		for _, candidate := range linesPerElements {
			candidates = append(candidates, *candidate)
		}

		suggestionsPerPeak[index] = candidates
	}

	alreadyMatchedElements := make(map[string]int)
	alreadyMatchedPeaks := make(map[int]struct{})

	calculateDistance := func(candidate elementWithDistance) float64 {
		if alreadyMatchedElements[candidate.element.Name] != 0 {
			return candidate.distance * (1 + 0.1*float64(alreadyMatchedElements[candidate.element.Name])) // Make match less likely
		}
		return candidate.distance
	}

	// Number of iterations
	for range peaks {
		match := -1
		elementWithDistance := elementWithDistance{}

		for index := range peaks {
			if _, exists := alreadyMatchedPeaks[index]; exists {
				continue
			}

			peakCandidates := suggestionsPerPeak[index]
			if len(peakCandidates) == 0 {
				continue
			}

			sort.SliceStable(peakCandidates, func(i, j int) bool {
				return calculateDistance(peakCandidates[i]) < calculateDistance(peakCandidates[j])
			})

			currentCandidate := peakCandidates[0]

			if match == -1 || calculateDistance(currentCandidate) < calculateDistance(elementWithDistance) {
				fmt.Println("Choosing the best", " Prev: ", calculateDistance(elementWithDistance), " Next: ", calculateDistance(currentCandidate))
				match = index
				elementWithDistance = currentCandidate
			}
		}

		if match != -1 {
			fmt.Println("--- match ---", match, peaks[match].Peak.Point.X, elementWithDistance.element.Name)

			alreadyMatchedPeaks[match] = struct{}{}
			alreadyMatchedElements[elementWithDistance.element.Name] = alreadyMatchedElements[elementWithDistance.element.Name] + 1

			suggestions = append(suggestions, AutoSuggestion{
				Peak:    peaks[match].Peak,
				Element: elementWithDistance.element,
			})
		}
	}

	return suggestions
}
