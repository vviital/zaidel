package database

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// ElementDefinition represents element's info
type ElementDefinition struct {
	Intensity       float64
	IonizationStage int
	Name            string
	WaveLength      float64
}

// TablePerElement represent table per element
type TablePerElement map[string][]ElementDefinition

// Table with element's info
var Table TablePerElement

func parseRecord(slice []string) (ElementDefinition, string) {
	waveLength, err := strconv.ParseFloat(slice[0], 64)
	if err != nil {
		log.Panic(err)
	}
	ionizationStage, err := strconv.Atoi(slice[2])
	if err != nil {
		log.Panic(err)
	}
	intensity, err := strconv.ParseFloat(slice[2], 64)
	if err != nil {
		log.Panic(err)
	}

	name := slice[1]
	return ElementDefinition{
		Intensity:       intensity,
		IonizationStage: ionizationStage,
		Name:            name,
		WaveLength:      waveLength,
	}, name
}

func parse() {
	// file, err := os.Open("./static/zaidel.txt")
	file, err := os.Open("/Users/vviital/go/src/github.com/vviital/zaidel/static/zaidel.txt")
	defer file.Close()

	if err != nil {
		log.Panic(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '	'

	for i := 0; ; i++ {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}

		if i == 0 {
			continue
		}

		element, name := parseRecord(record)
		Table[name] = append(Table[name], element)
	}

	log.Printf("%d records loaded from CSV file", len(Table))
}

// Predicate type
type Predicate func(ElementDefinition) bool

// GetRecordsByPredicate returns channels of elements which match predicate
func GetRecordsByPredicate(fn Predicate) chan ElementDefinition {
	ch := make(chan ElementDefinition, 10)

	go func() {
		defer close(ch)

		for name, slice := range Table {
			log.Print("processing ", name, " ", len(slice))
			for _, el := range slice {
				if fn(el) {
					ch <- el
				}
			}
		}
	}()

	return ch
}

func init() {
	Table = make(map[string][]ElementDefinition)
	parse()
}
