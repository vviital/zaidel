package researches

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Similarities struct {
	ExperimentID   string  `json:"eID"`
	ExperimentName string  `json:"eName"`
	Percentage     float64 `json:"p"`
	ResearchID     string  `json:"rID"`
	ResearchName   string  `json:"rName"`
}

// Comparison struct
type Comparison struct {
	BaseResearchID string         `json:"researchID"`
	CreatedAt      time.Time      `json:"createdAt"`
	ExperimentID   string         `json:"experimentID"`
	Finished       bool           `json:"finished"`
	ID             string         `json:"id"`
	LockedAt       time.Time      `json:"lockedAt"`
	OwnerID        string         `json:"ownerID"`
	Processed      int64          `json:"processed"`
	ResearchID     string         `json:"baseResearchID"`
	Similarities   []Similarities `json:"similarities"`
	Total          int64          `json:"total"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

var ComparisonIsNotAvailableError = errors.New("Comparison is not available")
var ResearchIDsAreNotAvailableError = errors.New("Research IDs are not available")
var ExperimentResultsAreNotAvailableError = errors.New("Experiment results are not available")
var ActualizationFailureError = errors.New("Actualization failure")
var FinalizationFailureError = errors.New("Finalization failure")

// GetComparisonByID Returns comparison with parsed content by ID
func GetComparisonByID(id string, headers http.Header) (comparison Comparison, err error) {
	req, err := http.NewRequest("GET", url+"/comparisons/"+id, nil)

	if err != nil {
		log.Print(err)
		return comparison, ComparisonIsNotAvailableError
	}

	for k, v := range headers {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return comparison, ComparisonIsNotAvailableError
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return comparison, ComparisonIsNotAvailableError
	}

	err = json.Unmarshal(content, &comparison)

	if err != nil {
		log.Print(err)
		return Comparison{}, ComparisonIsNotAvailableError
	}

	return comparison, nil
}

// ExperimentDs struct
type ExperimentDs struct {
	IDs []string `json:"identifiers"`
}

// GetExperimentIDs returns list of available research ids
func GetExperimentIDs(headers http.Header) (ids []string, err error) {
	req, err := http.NewRequest("GET", url+"/experiments/ids", nil)

	if err != nil {
		log.Print(err)
		return ids, ResearchIDsAreNotAvailableError
	}

	for k, v := range headers {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return ids, ResearchIDsAreNotAvailableError
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return ids, ResearchIDsAreNotAvailableError
	}

	r := ExperimentDs{}

	err = json.Unmarshal(content, &r)

	if err != nil {
		log.Print(err)
		return ids, ResearchIDsAreNotAvailableError
	}

	return r.IDs, nil
}

type ExperimentResults struct {
	Element           string  `json:"element"`
	ElementIntensity  float64 `json:"elementIntensity"`
	ElementWaveLength float64 `json:"elementWaveLength"`
	ExperimentName    string  `json:"experimentName"`
	PeakIntensity     float64 `json:"peakIntensity"`
	PeakWaveLength    float64 `json:"peakWaveLength"`
	ResearchID        string  `json:"researchID"`
	ResearchName      string  `json:"researchName"`
}

// GetExperimentResults returns experiment results
func GetExperimentResults(id string, headers http.Header) (results []ExperimentResults, err error) {
	req, err := http.NewRequest("GET", url+"/experiments/"+id+"/results", nil)

	if err != nil {
		log.Print(err)
		return results, ExperimentResultsAreNotAvailableError
	}

	for k, v := range headers {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return results, ExperimentResultsAreNotAvailableError
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return results, ExperimentResultsAreNotAvailableError
	}

	var slice []ExperimentResults
	err = json.Unmarshal(content, &slice)

	if err != nil {
		log.Print(err)
		log.Print("Raw results: ", string(content))
		return []ExperimentResults{}, ExperimentResultsAreNotAvailableError
	}

	return slice, nil
}

// ActualizeComparisonResults sends number of processed items to the researches service
func ActualizeComparisonResults(id string, total, processed int64, headers http.Header) (err error) {
	body, _ := json.Marshal(struct {
		Total     int64 `json:"total"`
		Processed int64 `json:"processed"`
	}{
		Total:     total,
		Processed: processed,
	})

	req, err := http.NewRequest("PATCH", url+"/comparisons/"+id+"/actualization", bytes.NewBufferString(string(body)))

	if err != nil {
		log.Print(err)
		return ActualizationFailureError
	}

	for k, v := range headers {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return ActualizationFailureError
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return ActualizationFailureError
	}

	return err
}

// FinalizeComparisonResults sends final results to the researches service
func FinalizeComparisonResults(id string, total, processed int64, similarities []Similarities, headers http.Header) (err error) {
	body, _ := json.Marshal(struct {
		Total        int64          `json:"total"`
		Processed    int64          `json:"processed"`
		Similarities []Similarities `json:"similarities"`
	}{
		Total:        total,
		Processed:    processed,
		Similarities: similarities,
	})

	req, err := http.NewRequest("PATCH", url+"/comparisons/"+id+"/finalization", bytes.NewBufferString(string(body)))

	if err != nil {
		log.Print(err)
		return FinalizationFailureError
	}

	for k, v := range headers {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return FinalizationFailureError
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return FinalizationFailureError
	}

	return err
}
