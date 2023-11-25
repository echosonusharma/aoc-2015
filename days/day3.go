package days

import (
	"fmt"
	"log"
	"strings"

	utils "github.com/echosonusharma/aoc-2015/internal"
)

type Coordinates struct {
	x int
	y int
}

func Day3() {
	input, err := utils.LoadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputArr := strings.Split(input, "")

	// (+x,+y)
	// (-x,+y)
	// (+x,-y)
	// (-x,-y)

	position := []Coordinates{}
	position = append(position, Coordinates{x: 0, y: 0}) // santa's initial position

	for _, value := range inputArr {
		lastPosition := &position[len(position)-1]
		var currentPosition Coordinates

		if value == ">" {
			// east
			currentPosition = Coordinates{x: lastPosition.x + 1, y: lastPosition.y}
		} else if value == "<" {
			// west
			currentPosition = Coordinates{x: lastPosition.x - 1, y: lastPosition.y}
		} else if value == "^" {
			// north
			currentPosition = Coordinates{x: lastPosition.x, y: lastPosition.y + 1}
		} else if value == "v" {
			// south
			currentPosition = Coordinates{x: lastPosition.x, y: lastPosition.y - 1}
		}

		position = append(position, currentPosition)
	}

	houses := make(map[Coordinates]int)

	for _, coordinate := range position {
		houses[coordinate] += 1
	}

	fmt.Printf("Answer for day 3, part 1 is : %d\n", len(houses))
	fmt.Printf("Answer for day 3, part 2 is : %d\n", part2(&inputArr))
}

func part2(inputArr *[]string) int {
	santaPosition := []Coordinates{}
	roboSantaPosition := []Coordinates{}

	santaPosition = append(santaPosition, Coordinates{x: 0, y: 0})
	roboSantaPosition = append(roboSantaPosition, Coordinates{x: 0, y: 0})

	turn := true

	for _, value := range *inputArr {
		var lastPosition *Coordinates
		var currentPosition Coordinates

		if turn {
			lastPosition = &santaPosition[len(santaPosition)-1]
		} else {
			lastPosition = &roboSantaPosition[len(roboSantaPosition)-1]
		}

		if value == ">" {
			currentPosition = Coordinates{x: lastPosition.x + 1, y: lastPosition.y}
		} else if value == "<" {
			currentPosition = Coordinates{x: lastPosition.x - 1, y: lastPosition.y}
		} else if value == "^" {
			currentPosition = Coordinates{x: lastPosition.x, y: lastPosition.y + 1}
		} else if value == "v" {
			currentPosition = Coordinates{x: lastPosition.x, y: lastPosition.y - 1}
		}

		if turn {
			santaPosition = append(santaPosition, currentPosition)
		} else {
			roboSantaPosition = append(roboSantaPosition, currentPosition)
		}

		turn = !turn
	}

	houses := make(map[Coordinates]int)
	concatenated := append(santaPosition, roboSantaPosition...)

	for _, coordinate := range concatenated {
		houses[coordinate] += 1
	}

	return len(houses)
}
