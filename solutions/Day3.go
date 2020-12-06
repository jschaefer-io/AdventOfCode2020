package solutions

import (
	"fmt"
	"strings"
)

type toboggan struct {
	x  int
	y  int
	vX int
	vY int
}

func (t toboggan) countTrees(forest *[][]bool, width int, height int) int {
	count := 0
	ref := *forest
	for t.y+t.vY < height {
		t.x += t.vX
		t.x %= width
		t.y += t.vY
		if ref[t.y][t.x] {
			count++
		}
	}
	return count
}

func newToboggan(vX int, vY int) toboggan {
	return toboggan{
		x:  0,
		y:  0,
		vX: vX,
		vY: vY,
	}
}

type Day3 struct {
	width  int
	height int
	forest [][]bool
}

func (d *Day3) init(s string) {
	for _, row := range strings.Split(s, "\n") {
		row = strings.TrimSpace(row)
		rowSlice := make([]bool, len(row))
		for i, r := range row {
			if r == '#' {
				rowSlice[i] = true
			}
		}
		d.forest = append(d.forest, rowSlice)
	}
	d.height = len(d.forest)
	d.width = len(d.forest[0])
}

func (d *Day3) executeA() int {
	t := newToboggan(3, 1)
	return t.countTrees(&d.forest, d.width, d.height)
}

func (d *Day3) executeB() int {
	slopes := make([][]int, 0)
	slopes = append(slopes, []int{1, 1})
	slopes = append(slopes, []int{3, 1})
	slopes = append(slopes, []int{5, 1})
	slopes = append(slopes, []int{7, 1})
	slopes = append(slopes, []int{1, 2})

	calc := 1
	for _, slope := range slopes {
		t := newToboggan(slope[0], slope[1])
		calc *= t.countTrees(&d.forest, d.width, d.height)
	}
	return calc
}

func (d *Day3) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))

	return results, nil
}
