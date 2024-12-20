package d18

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	w, h := 71, 71
	bytes := 1024
	area := newArea(w, h)
	area.prepare(input, bytes)
	path := area.path(pos{0, 0}, pos{w - 1, h - 1})
	return strconv.Itoa(len(path) - 1)
}

func Part2(input string) string {
	w, h := 71, 71
	area := newArea(w, h)
	lines := strings.Split(input, "\n")
	area.prepare(input, len(lines)-1)

	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) == 0 {
			continue
		}
		area.remove(lines[i])
		path := area.path(pos{0, 0}, pos{w - 1, h - 1})
		if len(path) != 0 {
			return lines[i]
		}
	}
	return "???"
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

func newArea(width, height int) Area {
	area := Area{
		width:  width,
		height: height,
		field:  make([]rune, width*height),
	}
	for i := range area.field {
		area.field[i] = '.'
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
func (a *Area) prepare(input string, fallen int) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		if len(lines[i]) == 0 {
			continue
		}
		if i >= fallen {
			break
		}
		a.add(lines[i])
	}
}
func (a *Area) add(line string) {
	grp := strings.Split(line, ",")
	x, _ := strconv.Atoi(grp[0])
	y, _ := strconv.Atoi(grp[1])
	a.set(x, y, '#')
}
func (a *Area) remove(line string) {
	grp := strings.Split(line, ",")
	x, _ := strconv.Atoi(grp[0])
	y, _ := strconv.Atoi(grp[1])
	a.set(x, y, '.')
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
func (a *Area) path(src, dst pos) path {
	pathsToTarget := paths{}

	type Node struct {
		pos  pos
		path []pos
	}
	visited := map[pos]bool{}
	visited[src] = true
	queue := []Node{{pos: src, path: []pos{src}}}
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

			if a.get(next.x, next.y, '#') == '#' {
				continue
			}

			if _, ok := visited[next]; ok {
				continue
			}

			visited[next] = true

			newPath := make([]pos, len(node.path))
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
