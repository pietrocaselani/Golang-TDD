package mmc

// MMC Calculates and returns the MMC
func MMC(numbers ...int) int {
	currentDivider := 2
	dividerCount := 0
	var dividers []int

	finished := false

	for !finished {
		for index, number := range numbers {
			if number%currentDivider == 0 {
				dividerCount++
				numbers[index] = number / currentDivider
			}
		}

		for _, number := range numbers {
			finished = number == 1
			if !finished {
				break
			}
		}

		if dividerCount == 0 {
			currentDivider++
		} else {
			dividerCount = 0
			dividers = append(dividers, currentDivider)
		}
	}

	result := 1

	for _, number := range dividers {
		result *= number
	}

	if result == 1 {
		return 0
	}

	return result
}
