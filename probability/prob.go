package main

import (
	"log"
	"strconv"
	"unicode/utf8"
)

func main() {}

// case example : 5 students 3 chairs. how many combinations ?
func getCombinationSitOnChair(numberElement int, chair int) int {
	comparator := numberElement - chair

	//calculate factorialElement
	factorialElement := getFactorial(numberElement)

	//calculate factorial comparator
	factorialComparator := getFactorial(comparator)

	return factorialElement / factorialComparator

}

func getFactorial(number int) int {
	resultFactorial := 1
	for number > 0 {
		resultFactorial = resultFactorial * number
		number -= 1
	}

	return resultFactorial
}

// case example : 8 ball : 5 white 3 blue, chance to get 2 white balls from the basket?
// format string must be : 3A,5B,4C and so on. on the totalTypesBalls
// format string must be 3A on takeRandomBalls
func getProbSameBalls(numberBalls int, takeRandomBalls string, totalTypeBalls ...string) float32 {
	// getDataRandomBalls
	totalRandomBallsTaken, typeRandomBallsTaken := getDataBalls(takeRandomBalls)

	log.Println("totalRandomBallsTaken = ", totalRandomBallsTaken)

	// calculate whole combination
	factorNumberBalls := getFactorial(numberBalls)
	divider := getFactorial(totalRandomBallsTaken) * getFactorial(numberBalls-totalRandomBallsTaken)

	wholeCombination := factorNumberBalls / divider

	// find totalTypeRandomBalls
	var totalTypesRandomBalls int
	for index := range totalTypeBalls {
		dataTotalTypesBalls := []rune(totalTypeBalls[index])
		// check type
		if typeRandomBallsTaken == dataTotalTypesBalls[1] {
			totalTypesRandomBalls = getInt(dataTotalTypesBalls[0])
			break
		}
	}

	factorialTotalTypeRandomBalls := getFactorial(totalTypesRandomBalls)
	dividerTaken := getFactorial(totalRandomBallsTaken) * getFactorial(totalTypesRandomBalls-totalRandomBallsTaken)

	specificCombination := factorialTotalTypeRandomBalls / dividerTaken

	// chance to get the same balls from the basket
	return float32(specificCombination / wholeCombination)
}

// get the data from format string
func getDataBalls(dataBalls string) (int, rune) {
	data := []rune(dataBalls)
	totalBalls := getInt(data[0])
	typeBalls := data[1]

	return totalBalls, typeBalls
}

// convert rune to int
func getInt(data rune) int {
	buf := make([]byte, 1)

	_ = utf8.EncodeRune(buf, data)
	value, _ := strconv.Atoi(string(buf))
	return value
}

func getProbDiffBalls(totalBalls int, totalTypeBalls int, takeRandomBalls ...int) {

}
