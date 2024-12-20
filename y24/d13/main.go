package d13

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	sum := 0
	xyRe := `X[+=]([0-9]*), Y[+=]([0-9]*)`
	re := regexp.MustCompile(xyRe)

	machines := []Machine{}

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 4 {
		btnA := re.FindStringSubmatch(lines[i])
		btnB := re.FindStringSubmatch(lines[i+1])
		price := re.FindStringSubmatch(lines[i+2])

		btnAx, _ := strconv.Atoi(btnA[1])
		btnAy, _ := strconv.Atoi(btnA[2])
		btnBx, _ := strconv.Atoi(btnB[1])
		btnBy, _ := strconv.Atoi(btnB[2])
		priceX, _ := strconv.Atoi(price[1])
		priceY, _ := strconv.Atoi(price[2])

		machines = append(machines, Machine{
			Button{
				btnAx,
				btnAy,
			},
			Button{
				btnBx,
				btnBy,
			},
			Price{
				priceX,
				priceY,
			},
		})
	}

	for _, machine := range machines {
		token := solve(machine)
		sum += token
	}

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	sum := 0
	xyRe := `X[+=]([0-9]*), Y[+=]([0-9]*)`
	re := regexp.MustCompile(xyRe)

	machines := []Machine{}

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 4 {
		btnA := re.FindStringSubmatch(lines[i])
		btnB := re.FindStringSubmatch(lines[i+1])
		price := re.FindStringSubmatch(lines[i+2])

		btnAx, _ := strconv.Atoi(btnA[1])
		btnAy, _ := strconv.Atoi(btnA[2])
		btnBx, _ := strconv.Atoi(btnB[1])
		btnBy, _ := strconv.Atoi(btnB[2])
		priceX, _ := strconv.Atoi(price[1])
		priceY, _ := strconv.Atoi(price[2])

		machines = append(machines, Machine{
			Button{
				btnAx,
				btnAy,
			},
			Button{
				btnBx,
				btnBy,
			},
			Price{
				priceX + 10000000000000,
				priceY + 10000000000000,
			},
		})
	}

	for _, machine := range machines {
		token := solve(machine)
		sum += token
	}

	return strconv.Itoa(sum)
}

type Button struct {
	x, y int
}

type Price struct {
	x, y int
}

type Machine struct {
	a, b Button
	p    Price
}

func solve(m Machine) int {
	x := m.p.x
	y := m.p.y
	ax := m.a.x
	ay := m.a.y
	bx := m.b.x
	by := m.b.y

	b := (-ax*y + ay*x) / (ay*bx - ax*by)
	a := (x - b*bx) / ax
	if a*ax+b*bx == x && a*ay+b*by == y {
		return 3*a + b
	}
	return 0
}
