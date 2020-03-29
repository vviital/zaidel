package comparisons

import (
	"log"
	"math"
	"net/http"
	"time"

	"github.com/vviital/zaidel/datasource/researches"
)

// Compare research from the comparion request
func Compare(comparisonID string, headers http.Header) {
	comparison, err := researches.GetComparisonByID(comparisonID, headers)
	if err != nil {
		log.Println(err)
		return
	}

	experimentIDs, err := researches.GetExperimentIDs(headers)
	if err != nil {
		log.Println(err)
		return
	}

	baseExperimentResults, err := researches.GetExperimentResults(comparison.ExperimentID, headers)

	if err != nil {
		log.Println(err)
		return
	}

	total := int64(len(experimentIDs))
	similarities := make([]researches.Similarities, 0)

	for index, id := range experimentIDs {
		results, err := researches.GetExperimentResults(id, headers)

		if err != nil {
			log.Println(err)
			continue
		}

		if len(results) == 0 {
			continue
		}

		similarities = append(similarities, researches.Similarities{
			ExperimentID:   id,
			ExperimentName: results[0].ExperimentName,
			Percentage:     calculateSimilarity(baseExperimentResults, results),
			ResearchID:     results[0].ResearchID,
			ResearchName:   results[0].ResearchName,
		})

		if index%10 == 0 {
			err = researches.ActualizeComparisonResults(comparisonID, total, int64(index+1), headers)
			if err != nil {
				log.Println(err)
			}
		}

		<-time.After(5 * time.Millisecond)
	}

	normalizeResults(similarities)

	err = researches.FinalizeComparisonResults(comparisonID, total, total, similarities, headers)
	if err != nil {
		log.Println(err)
	}

	log.Print("Finished to compare results for the research: ", comparisonID)
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
}

func calculateDistanceBetweenPeaks(r1, r2 researches.ExperimentResults) float64 {
	return calculateDistance(r1.PeakWaveLength, r1.PeakIntensity, r2.PeakWaveLength, r2.PeakIntensity)
}

func calculateSimilarity(source, target []researches.ExperimentResults) (score float64) {
	for _, resultItem := range target {
		bestMatchIndex := -1

		for i, sourceResultItem := range source {
			if sourceResultItem.Element != resultItem.Element {
				continue
			}

			if bestMatchIndex == -1 && calculateDistanceBetweenPeaks(resultItem, sourceResultItem) < calculateDistanceBetweenPeaks(resultItem, source[bestMatchIndex]) {
				bestMatchIndex = i
			}
		}

		if bestMatchIndex == -1 {
			continue
		}

		score += calculateDistanceBetweenPeaks(resultItem, source[bestMatchIndex])
	}

	return
}

func normalizeResults(similarities []researches.Similarities) {
	maxScore := -1.0

	for _, similarity := range similarities {
		if similarity.Percentage > maxScore {
			maxScore = similarity.Percentage
		}
	}

	for _, similarity := range similarities {
		similarity.Percentage = similarity.Percentage / maxScore
	}
}
