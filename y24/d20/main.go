package d20

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	var (
		start, end pos
	)
	area := newArea(input)

	{ // find start and end
		for y := 0; y < area.height; y++ {
			for x := 0; x < area.width; x++ {
				switch area.get(x, y, '.') {
				case 'S':
					start.x = x
					start.y = y
				case 'E':
					end.x = x
					end.y = y
				}
			}
		}
	}

	path := area.path(start, end, pos{0, 0})
	shortcuts := shortcuts(path, 2)

	return strconv.Itoa(shortcuts)
}

func Part2(input string) string {
	var (
		start, end pos
	)
	area := newArea(input)

	{ // find start and end
		for y := 0; y < area.height; y++ {
			for x := 0; x < area.width; x++ {
				switch area.get(x, y, '.') {
				case 'S':
					start.x = x
					start.y = y
				case 'E':
					end.x = x
					end.y = y
				}
			}
		}
	}

	path := area.path(start, end, pos{0, 0})
	shortcuts := shortcuts(path, 20)

	return strconv.Itoa(shortcuts)
}

func shortcuts(path path, allowedDistance int) (shortcuts int) {
	// FIXME: Could this be optimized?
	for i := 0; i < len(path)-1; i++ {
		for j := i; j < len(path); j++ {
			a := path[i]
			b := path[j]
			delta := pos{
				x: int(math.Abs(float64(a.x - b.x))),
				y: int(math.Abs(float64(a.y - b.y))),
			}

			distance := delta.x + delta.y
			if distance > allowedDistance {
				continue
			}

			newPathLength := j - i - distance
			if newPathLength < 100 {
				continue
			}

			shortcuts++
		}
	}
	return
}

type pos struct {
	x, y int
}
type path []pos
type paths []path

type Area struct {
	width, height int
	field         []rune
}

func newArea(input string) Area {
	lines := strings.Split(input, "\n")
	height := len(lines) - 1
	width := len(lines[0])

	area := Area{
		width:  width,
		height: height,
		field:  make([]rune, width*height),
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			area.set(x, y, rune(lines[y][x]))
		}
	}
	return area
}
func (a *Area) print() string {
	builder := strings.Builder{}
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			builder.WriteRune(a.get(x, y, 'X'))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
func (a *Area) set(x, y int, r rune) {
	if x < 0 || x >= a.width {
		return
	}

	if y < 0 || y >= a.height {
		return
	}

	a.field[y*a.width+x] = r
}
func (a *Area) get(x, y int, dr rune) rune {
	if x < 0 || y >= a.width {
		return dr
	}

	if x < 0 || y >= a.height {
		return dr
	}

	return a.field[y*a.width+x]
}
func (a *Area) path(src, dst pos, ignore pos) path {
	pathsToTarget := paths{}

	type Node struct {
		pos  pos
		path []pos
	}
	visited := map[pos]bool{}
	visited[src] = true
	queue := make([]Node, 0, 1024)
	queue = append(queue, Node{pos: src, path: make([]pos, 0, 1024)})
	queue[0].path = append(queue[0].path, src)
	for len(queue) > 0 {
		var node Node
		node, queue = queue[0], queue[1:]

		if node.pos == dst {
			pathsToTarget = append(pathsToTarget, node.path)
			break
		}
		for _, dir := range []pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := pos{
				x: node.pos.x + dir.x,
				y: node.pos.y + dir.y,
			}

			if next.x < 0 || next.y < 0 || next.x >= a.width || next.y >= a.height {
				continue
			}

			if next != ignore {
				if a.get(next.x, next.y, '#') == '#' {
					continue
				}
			}

			if _, ok := visited[next]; ok {
				continue
			}

			visited[next] = true

			newPath := make([]pos, len(node.path), len(node.path)*2)
			copy(newPath, node.path)
			newPath = append(newPath, next)

			queue = append(queue, Node{next, newPath})
		}
	}
	slices.SortFunc(pathsToTarget, func(a, b path) int {
		return cmp.Compare(len(a), len(b))
	})
	if len(pathsToTarget) == 0 {
		return path{}
	}
	return pathsToTarget[0]
}
