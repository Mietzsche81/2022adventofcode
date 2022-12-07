package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//
	// Read input
	//

	headerSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var rawPacket string = strings.TrimSpace(os.Args[2])

	//
	// Find message start
	//

	var isUnique bool
	var messageStart int
	for i := range rawPacket {
		isUnique = CheckUnique(rawPacket[i : i+headerSize])
		if isUnique {
			messageStart = i + headerSize
			break
		}
	}

	//
	// Report message start
	//
	if isUnique {
		fmt.Printf("Message begins on character index: %d", messageStart)
	} else {
		fmt.Println("No packet header found")
	}
}

func CheckUnique(s string) bool {
	for i, forward := range s {
		for _, backward := range StringReverse((s[i+1:])) {
			if forward == backward {
				return false
				break
			}
		}
	}
	return true
}

func StringReverse(s string) (reverse string) {
	for _, c := range s {
		reverse = string(c) + reverse
	}
	return
}
