package days

import (
	"fmt"
	"log"
	"strings"

	utils "github.com/echosonusharma/aoc-2015/internal/common"
)

func Day1() {
	input, err := utils.LoadFile("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputArr := strings.Split(input, "")
	currentFloor := 0
	var basementPosition int

	fmt.Println(basementPosition)

	for i := 0; i < len(inputArr); i++ {
		if currentFloor == -1 && basementPosition == 0 {
			basementPosition = i
		}

		if inputArr[i] == "(" {
			currentFloor += 1
		} else if inputArr[i] == ")" {
			currentFloor -= 1
		}
	}

	fmt.Printf("Answer for day 1, part 1 is : %d\n", currentFloor)
	fmt.Printf("Answer for day 1, part 2 is : %d\n", basementPosition)
}
