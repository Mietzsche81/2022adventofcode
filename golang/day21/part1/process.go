package part1

func Process(monkeys map[string]*Monkey) int {

	for monkeys["root"].value == nil {
		for _, monkey := range monkeys {
			monkey.operate()
		}
	}

	return *(monkeys["root"].value)
}
