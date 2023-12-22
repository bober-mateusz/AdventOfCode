package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bober-mateusz/AdventOfCode/Utils"
)

func main() {
	mySlice := Utils.FileReader("./input.txt")

	// Get Game number & rounds in 0 and 1 for line 0
	sum := 0
	//Rounds is in array 1
	for _, line := range mySlice {
		gameLine := strings.Split(line, ":")
		gameNum := gameLine[0]
		games := strings.Split(gameLine[1], ";")
		sum += checkGame(gameNum, games)
	}
	fmt.Println(sum)
}

func checkGame(gameNum string, games []string) int {
	firstNumSplit := strings.Split(gameNum, " ")
	lastNum, err := strconv.Atoi(firstNumSplit[1])
	if err != nil {
		fmt.Println("Number error")
	}
	flag := true
	for _, round := range games {
		if checkRound(round) == false {
			flag = false
		}
	}
	if flag == true {
		return lastNum
	}
	return 0
}

func checkRound(round string) bool {
	colours := strings.Split(round, ",")
	for _, colour := range colours {
		//14 Red Splits into 14 and Red so we can get the first number to work with
		firstNumSplit := strings.Split(colour, " ")
		firstNum, err := strconv.Atoi(firstNumSplit[1])

		if err != nil {
			fmt.Println("Number error")
		}
		if strings.Contains(colour, "bl") {
			if firstNum > 14 {
				return false
			}
		}
		if strings.Contains(colour, "gr") {
			if firstNum > 13 {
				return false
			}
		}
		if strings.Contains(colour, "red") {
			if firstNum > 12 {
				return false
			}
		}
	}
	return true
}
