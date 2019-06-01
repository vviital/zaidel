package database

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// table with element's info
var table TablePerElement

// WaveLengthBasedFinder is an interface for the Zaidel tables
var WaveLengthBasedFinder Zaidel

func parseRecord(slice []string) Element {
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

	return Element{
		Intensity:       intensity,
		IonizationStage: ionizationStage,
		Name:            slice[1],
		WaveLength:      waveLength,
	}
}

func parse() {
	// file, err := os.Open("./static/zaidel.txt")
	file, err := os.Open("/Users/vviital/go/src/github.com/vviital/zaidel/static/zaidel.txt")
	defer file.Close()

	if err != nil {
		log.Panic(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	count := 0

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

		element := parseRecord(record)
		table[element.Name] = append(table[element.Name], element)
		count++
	}

	log.Printf("%d elements loaded from CSV file. Table contains %d records.", len(table), count)
}

func init() {
	table = make(map[string]Elements)
	parse()
	WaveLengthBasedFinder = newWaveLengthBasedFastFinder(table)
}
