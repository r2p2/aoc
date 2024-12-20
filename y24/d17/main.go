package d17

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	regAstr := strings.Fields(lines[0])[2]
	regA, _ := strconv.Atoi(regAstr)
	regBstr := strings.Fields(lines[1])[2]
	regB, _ := strconv.Atoi(regBstr)
	regCstr := strings.Fields(lines[2])[2]
	regC, _ := strconv.Atoi(regCstr)
	prgStr := strings.Fields(lines[3])[1]

	mem := []byte{}
	ops := strings.Split(prgStr, ",")
	for i := range ops {
		n, _ := strconv.Atoi(ops[i])
		mem = append(mem, byte(n))
	}
	machine := Machine{
		reg: [3]int{regA, regB, regC},
		pc:  0,
		mem: mem,
		out: []byte{},
	}

	for !machine.Step() {
	}

	return machine.print()
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	prgStr := strings.Fields(lines[3])[1]

	mem := []byte{}
	ops := strings.Split(prgStr, ",")
	for i := range ops {
		n, _ := strconv.Atoi(ops[i])
		mem = append(mem, byte(n))
	}
	machine := Machine{
		reg: [3]int{0, 0, 0},
		pc:  0,
		mem: mem,
		out: []byte{},
	}

	// FIXME: Should work if we start with 0?
	register := 1

	for {
		machine.reg[0] = register
		machine.reg[1] = 0
		machine.reg[2] = 0
		machine.pc = 0
		machine.out = []byte{}

		for !machine.Step() {
		}

		if slices.Compare(machine.mem, machine.out) == 0 {
			return strconv.Itoa(register)
		}

		if comp(machine.mem, machine.out) {
			register *= 8
		} else {
			register++
		}
	}
}

type Machine struct {
	reg [3]int
	pc  int
	mem []byte
	out []byte
}

func (m *Machine) Step() (halt bool) {
	if m.pc >= len(m.mem) {
		halt = true
		return
	}

	switch m.mem[m.pc] {
	case 0: // adv
		a := float64(m.reg[0])
		b := math.Pow(2.0, float64(m.operand(int(m.mem[m.pc+1]))))
		m.reg[0] = int(a / b)
		m.pc += 2
	case 1: // bxl
		m.reg[1] ^= int(m.mem[m.pc+1])
		m.pc += 2
	case 2: // bst
		m.reg[1] = int(m.operand(int(m.mem[m.pc+1])) % 8)
		m.pc += 2
	case 3: // jnz
		if m.reg[0] != 0 {
			m.pc = int(m.mem[m.pc+1])
		} else {
			m.pc += 2
		}
	case 4: // bxc
		m.reg[1] ^= m.reg[2]
		m.pc += 2
	case 5: // out
		m.out = append(m.out, byte(m.operand(int(m.mem[m.pc+1]))%8))
		m.pc += 2
	case 6: // bdv
		a := float64(m.reg[0])
		b := math.Pow(2.0, float64(m.operand(int(m.mem[m.pc+1]))))
		m.reg[1] = int(a / b)
		m.pc += 2
	case 7: // cdv
		a := float64(m.reg[0])
		b := math.Pow(2.0, float64(m.operand(int(m.mem[m.pc+1]))))
		m.reg[2] = int(a / b)
		m.pc += 2
	}

	return
}

func (m *Machine) operand(operand int) int {
	if operand <= 3 {
		return operand
	} else {
		return m.reg[operand-4]
	}
}

func (m *Machine) print() string {
	sx := []string{}
	for i := range m.out {
		s := strconv.Itoa(int(m.out[i]))
		sx = append(sx, s)
	}
	return strings.Join(sx, ",")
}

// TODO: I don't like this function, makes more sensor to compare the values instead of translating to string
func comp(left []byte, right []byte) bool {
	var leftS, rightS string

	lx := []string{}
	for i := range left {
		l := strconv.Itoa(int(left[i]))
		lx = append(lx, l)
	}
	rx := []string{}
	for i := range right {
		r := strconv.Itoa(int(right[i]))
		rx = append(rx, r)
	}
	leftS = strings.Join(lx, ",")
	rightS = strings.Join(rx, ",")

	return strings.HasSuffix(leftS, rightS)
}
