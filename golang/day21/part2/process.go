package part2

// Boldly assume the variation is sufficiently linear at scale
func GradientDescent(fileName string) int {

	// Humn inputs
	x1, x2, x3 := 0, 0, 0
	// Difference outputs
	y1, y2, y3 := 0, 0, 0
	// Victory Flag
	success := false

	// Initialize guesses
	if success, y1 = Process(fileName, x1); success {
		// Zoinks, that's a lucky break.
		return x1
	}
	x2 = y1 / 2
	if success, y2 = Process(fileName, x2); success {
		// Let's get out of here Scoob!
		return x2
	}

	// Secant Search Method
	for !success {
		x3 = x2 - y2*(x2-x1)/(y2-y1)
		success, y3 = Process(fileName, x3)
		x1, y1 = x2, y2
		x2, y2 = x3, y3
	}

	// Check for int roundup, choose smallest
	if success, _ = Process(fileName, x3-1); success {
		return x3 - 1
	} else {
		return x3
	}
}

func Process(fileName string, humn int) (success bool, diff int) {

	monkeys := ParseInput(fileName)
	monkeys["humn"].value = &humn
	for monkeys["root"].value == nil {
		for _, monkey := range monkeys {
			monkey.operate()
		}
	}
	success = *(monkeys["root"].value) == 1
	diff = *(monkeys["root"].operands[0].value) - *(monkeys["root"].operands[1].value)

	return success, diff
}
