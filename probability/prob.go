package main

import (
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

// case example : 8 ball : 5 white 3 blue, chance to get 3 balls 1 white 2 blue from the basket?
// format string must be : 3A,5B,4C and so on. on the totalTypesBalls
// format string must be 3A on takeRandomBalls
// NOTE : ... string argument we separate it with numberBalls arguments
func getProbDiffBalls(takeRandomBalls []string, numberBalls int, totalTypeBalls ...string) float32 {
	// getDataRandomBalls 1
	totalRandomBallsTaken1, _ := getDataBalls(takeRandomBalls[0])
	// getDataRandomBalls 2
	totalRandomBallsTaken2, _ := getDataBalls(takeRandomBalls[1])

	totalRandomBallsTaken := totalRandomBallsTaken1 + totalRandomBallsTaken2

	// get DataBalls

	// calculate whole combination
	factorialNumberBalls := getFactorial(numberBalls)
	divider := getFactorial(totalRandomBallsTaken) * getFactorial(numberBalls-totalRandomBallsTaken)

	wholeCombination := factorialNumberBalls / divider

	// calculate combination with same balls
	// get totalBalls in basket 1
	totalRandomBalls1, _ := getDataBalls(totalTypeBalls[0])
	totalRandomBalls2, _ := getDataBalls(totalTypeBalls[1])

	combinationBalls1 := getCombination(totalRandomBalls1, totalRandomBallsTaken1)

	combinationBalls2 := getCombination(totalRandomBalls2, totalRandomBallsTaken2)

	combinationCombine := combinationBalls1 * combinationBalls2

	return float32(combinationCombine / wholeCombination)
}

// calculate combination
func getCombination(totalNumber int, getNumber int) int {
	factorialTotalNumber := getFactorial(totalNumber)
	divider := getFactorial(getNumber) * getFactorial(totalNumber-getNumber)

	return factorialTotalNumber / divider

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
