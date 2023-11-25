package days

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	utils "github.com/echosonusharma/aoc-2015/internal"
)

func Day2() {
	input, err := utils.LoadFile("day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputArr := strings.Split(input, "\n")

	var totalSurfaceArea int
	var ribbonLength int

	for _, value := range inputArr {
		// dimension - l * b * h
		dimension := strings.Split(value, "x")
		l, _ := strconv.Atoi(dimension[0])
		w, _ := strconv.Atoi(dimension[1])
		h, _ := strconv.Atoi(dimension[2])

		s1, s2, s3 := l*w, w*h, h*l

		surfaceArea := 2*s1 + 2*s2 + 2*s3

		// slack - lowest surface area
		slack := s1
		if s2 < slack {
			slack = s2
		}
		if s3 < slack {
			slack = s3
		}

		arr := []int{l, w, h}
		sort.Ints(arr)

		ribbonLength += (2*arr[0] + 2*arr[1]) + l*w*h

		totalSurfaceArea += surfaceArea + slack
	}

	fmt.Printf("Answer for day 2, part 1 is : %d\n", totalSurfaceArea)
	fmt.Printf("Answer for day 2, part 2 is : %d\n", ribbonLength)
}
