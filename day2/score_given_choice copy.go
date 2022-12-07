package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type RoundRecord struct {
	OpponentChoice string
	MyChoice       string
	ChoicePoints   int
	WinPoints      int
	RecordPoints   int
}

func CreateAllRecords(data [][]string) []RoundRecord {
	var results []RoundRecord
	for _, line := range data {
		results = append(results, CreateRecord(line))
	}
	return results
}

func CreateRecord(line []string) RoundRecord {
	ChoicePoints := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	OutcomePoints := map[string]int{
		"AX": 3,
		"AY": 6,
		"AZ": 0,
		"BX": 0,
		"BY": 3,
		"BZ": 6,
		"CX": 6,
		"CY": 0,
		"CZ": 3,
	}
	round := RoundRecord{
		OpponentChoice: line[0],
		MyChoice:       line[1],
	}
	round.ChoicePoints = ChoicePoints[round.MyChoice]
	round.WinPoints = OutcomePoints[round.OpponentChoice+round.MyChoice]
	round.RecordPoints = round.ChoicePoints + round.WinPoints
	return round
}

func main() {

	//
	// Read input
	//

	var fileName string = strings.TrimSpace(os.Args[1])
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	csvReader := csv.NewReader(fin)
	csvReader.Comma = ' '
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Create records
	//

	records := CreateAllRecords(data)
	myTotalScore := 0
	for _, record := range records {
		myTotalScore += record.RecordPoints
	}

	//
	// Report
	//

	fmt.Printf("Total Points earned: %d\n", myTotalScore)
}
