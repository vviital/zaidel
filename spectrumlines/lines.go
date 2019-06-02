package spectrumlines

import (
	"context"
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
	return (settings.MaxIntensity > 0 && element.Intensity < settings.MaxIntensity) ||
		(settings.MinIntensity > 0 && element.Intensity > settings.MinIntensity) ||
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
