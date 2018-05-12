package mmc

// MMC Calculates and returns the MMC
func MMC(numbers ...int) int {
	lenght := len(numbers)
	divider := 2
	dividerCount := 0
	var dividers []int

	finished := false

	for !finished {
		for i := 0; i < lenght; i++ {
			index := lenght - i - 1
			number := numbers[index]
			if number%divider == 0 {
				dividerCount++
				numbers[index] = number / divider
			}
		}

		for _, number := range numbers {
			finished = number == 1
			if !finished {
				break
			}
		}

		if dividerCount == 0 {
			divider++
		} else {
			dividerCount = 0
			dividers = append(dividers, divider)
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
