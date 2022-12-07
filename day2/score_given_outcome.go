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
	OutcomePoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	ChoicePoints := map[string]int{
		"AX": 3,
		"AY": 1,
		"AZ": 2,
		"BX": 1,
		"BY": 2,
		"BZ": 3,
		"CX": 2,
		"CY": 3,
		"CZ": 1,
	}
	round := RoundRecord{
		OpponentChoice: line[0],
		MyChoice:       line[1],
	}
	round.WinPoints = OutcomePoints[round.MyChoice]
	round.ChoicePoints = ChoicePoints[round.OpponentChoice+round.MyChoice]
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
