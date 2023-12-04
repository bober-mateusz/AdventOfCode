package main

import (
	"fmt"
	"strconv"

	"github.com/bober-mateusz/AdventOfCode/Utils"
)

func main() {
	mySlice := Utils.FileReader("./input.txt")
	mySum := 0

	for _, line := range mySlice {
		mySum += getCalibrationValues(line)
	}

	fmt.Println("The calibration value is : ", mySum)
}

func getCalibrationValues(line string) int {

	first := getFirstDigitChar(line)
	last := getLastDigitChar(line)
	lineCalibration := first + last
	number, error := strconv.Atoi(lineCalibration)
	if error != nil {
		fmt.Println(error)
	}
	return number
}

func getFirstDigitChar(line string) string {
	for i := 0; i < len(line); i++ {
		if byteIsNumber(line[i]) {
			return string(line[i])
		}
	}
	return ""
}

func getLastDigitChar(line string) string {
	for i := len(line) - 1; i < len(line); i-- {
		if byteIsNumber(line[i]) {
			return string(line[i])
		}
	}
	return ""
}

func byteIsNumber(myByte byte) bool {
	if myByte > 47 && myByte < 72 {
		return true
	}
	return false
}
