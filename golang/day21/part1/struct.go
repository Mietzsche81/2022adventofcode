package part1

import "fmt"

type Monkey struct {
	name      string
	value     *int
	operation func(int, int) int
	operands  [2]*Monkey
}

func (m *Monkey) done() bool {
	return m.value != nil
}

func (m *Monkey) operate() bool {
	// Shortcut if done for efficiency
	if m.done() {
		return true
	}

	// If operands both ready
	if m.operands[0].operate() && m.operands[1].operate() {
		// Execute
		value := m.operation(*m.operands[0].value, *m.operands[1].value)
		m.value = &value
		return true
	}

	// Fallthrough: not done
	return false
}

func (m *Monkey) Str() string {
	out := fmt.Sprintf("{ name: '%s' ", m.name)
	if m.value != nil {
		out += fmt.Sprintf("value: %3d ", *(m.value))
	} else {
		out += "value: --- "
	}
	out += "operands: [ "
	if m.operands[0] != nil {
		out += fmt.Sprintf("%s ", m.operands[0].name)
	} else {
		out += "---- "
	}
	if m.operands[1] != nil {
		out += fmt.Sprintf("%s ", m.operands[1].name)
	} else {
		out += "---- "
	}
	out += "] }"
	return out
}
