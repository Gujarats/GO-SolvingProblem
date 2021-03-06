package main

import "fmt"

func main() {

}

func solutions(A []int, B []int, M int, X int, Y int) int {
	var totalWeight int
	var stopNumber int
	var totalPeople int
	indexLift := 0

	// this variable is used for checking the previous lift if we have visit them
	var previousLift map[int]bool

	// check the len of Persons and the len of LiftValue they must be the same
	if len(A) != len(B) {
		return -1
	}

	for indexPeople, weight := range A {
		totalWeight += weight // counting total weight
		totalPeople++         // counting total people

		// check if we got the righ total persons
		if totalPeople == X || indexPeople == len(A)-1 {
			totalPeople = 0 // reset total people

			if totalWeight <= Y {
				totalWeight = 0 // reset total weight

				// proceed to lift person to specifc lift
				limitLift := indexLift + X - 1
				for indexLift <= limitLift {

					// avoid index of bound
					if indexLift <= len(B)-1 {
						if B[indexLift] > M {
							return -1
						}

						if B[indexLift] == 0 {
							break
						}

						if previousLift == nil {
							previousLift = make(map[int]bool)
							previousLift[B[indexLift]] = true
							stopNumber++
						} else if previousLift != nil {
							if !previousLift[B[indexLift]] {
								stopNumber++
							}
						}

						indexLift++

					} else {
						break
					}
				}

				previousLift = nil
				// go to the ground if the last index is not ground
				stopNumber++

				fmt.Println("stopNumber = ", stopNumber)

			}
		}
	}
	return stopNumber
}
