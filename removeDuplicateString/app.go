package main

func removeDulicateString(input string) string {
	// convert it to array
	splitString := []rune(input)

	//create temp slice and init value
	var tempSplitString []rune
	tempSplitString = append(tempSplitString, splitString[0])

	for index := range splitString {
		for index2 := index + 1; index2 < len(splitString); index2++ {
			if index2 < len(splitString) {
				if splitString[index] == splitString[index2] {
					// remove specifi index array
					splitString = removeSpecifiArray(splitString, index2)

				}
			}
		}
	}

	return string(splitString)
}

func removeSpecifiArray(data []rune, index int) []rune {
	data = append(data[:index], data[index+1:]...)
	return data
}
