package days

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"time"

	utils "github.com/echosonusharma/aoc-2015/internal"
)

func md5HashAString(srt *string) string {
	hashBytes := md5.Sum([]byte(*srt))
	hashString := fmt.Sprintf("%x", hashBytes)
	return hashString
}

func Day4() {
	input, err := utils.LoadFile("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	startTime := time.Now()

	hashValue := input
	checkValue := 0
	foundCheck := true
	var md5HashValue string

	var part1 string
	var part2 string

	part1FoundCheck := true

	for foundCheck {
		hashValue = input + strconv.Itoa(checkValue)
		md5HashValue = md5HashAString(&hashValue)
		checkValue++

		if checkValue > 1e10 {
			log.Fatalln("🦀")
		}

		if md5HashValue[0:5] == "00000" && part1FoundCheck {
			part1FoundCheck = false
			part1 = hashValue
		}

		if md5HashValue[0:6] == "000000" {
			foundCheck = false
			part2 = hashValue
		}
	}

	fmt.Printf("Answer for day 4, part 1 is : %s\n", part1)
	fmt.Printf("Answer for day 4, part 2 is : %s\n", part2)
	fmt.Println("Time Elapsed for day 4 :", time.Since(startTime))
}
