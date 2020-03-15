package files

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vviital/zaidel/geometry"
)

type File struct {
	ID      string `json:"id"`
	OwnerID string `json"ownerID"`
	Content []struct {
		WaveLength float64 `json:"waveLength"`
		Intensity  float64 `json:"Intensity"`
	}
}

// GetPoints returns points from the file
func (file File) GetPoints() geometry.Coordinates {
	res := make(geometry.Coordinates, len(file.Content))

	for i, p := range file.Content {
		res[i] = geometry.Coordinate{
			X: p.WaveLength,
			Y: p.Intensity,
		}
	}

	return res
}

var FileIsNotAvailableError = errors.New("File is not available")

// Returns file with parsed content by ID
func GetByID(id string, headers http.Header) (file File, err error) {
	req, err := http.NewRequest("GET", url+"/files/"+id+"?limit=999999", nil)

	if err != nil {
		log.Print(err)
		return file, FileIsNotAvailableError
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
		return file, FileIsNotAvailableError
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		return file, FileIsNotAvailableError
	}

	err = json.Unmarshal(content, &file)

	if err != nil {
		log.Print(err)
		return File{}, FileIsNotAvailableError
	}

	return file, nil
}
