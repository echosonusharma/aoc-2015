package days

import (
	"fmt"
	"log"
	"slices"
	"strings"

	utils "github.com/echosonusharma/aoc-2015/internal"
)

var vowels []string = []string{"a", "e", "i", "o", "u"}
var strCk []string = []string{"ab", "cd", "pq", "xy"}

func checkForNiceString1(value string) bool {
	valueLen := len(value)
	vowelCheckIdx := []int{}
	doubleSameCharCheck := false
	containsStrCk := false

	for i := 0; i < valueLen; i++ {
		containsVowel := slices.Contains(vowels, string(value[i]))

		if containsVowel {
			vowelCheckIdx = append(vowelCheckIdx, i)
		}

		if i != valueLen-1 && value[i] == value[i+1] {
			doubleSameCharCheck = true
		}

		if i != valueLen-1 {
			checkStrCk := slices.Contains(strCk, value[i:i+2])

			if checkStrCk {
				containsStrCk = true
			}
		}
	}

	if len(vowelCheckIdx) >= 3 && doubleSameCharCheck && !containsStrCk {
		return true
	}

	return false
}

func checkForNiceString2(value string) bool {
	valueLen := len(value)
	pairCheck := false
	repeatCheck := false

	pairStrArr := make(map[string]int)

	for i := 0; i < valueLen-1; i++ {
		pair := value[i : i+2]
		if prevIndex, ok := pairStrArr[pair]; ok && prevIndex < i-1 {
			pairCheck = true
			break
		}

		pairStrArr[pair] = i
	}

	for i := 0; i < valueLen-2; i++ {
		if value[i] == value[i+2] {
			repeatCheck = true
			break
		}
	}

	return pairCheck && repeatCheck
}

func Day5() {
	input, err := utils.LoadFile("day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	wordList := strings.SplitAfter(input, "\n")

	niceWordCount1 := 0
	niceWordCount2 := 0

	for _, item := range wordList {
		wordCheck1 := checkForNiceString1(item)
		wordCheck2 := checkForNiceString2(item)

		if wordCheck1 {
			niceWordCount1++
		}
		if wordCheck2 {
			niceWordCount2++
		}
	}

	fmt.Printf("Answer for day 5, part 1 is : %d\n", niceWordCount1)
	fmt.Printf("Answer for day 5, part 2 is : %d\n", niceWordCount2)
}
