package solutions

import (
	"fmt"
	"regexp"
	"strings"
)

type Day24 struct {
	tiles        map[int]map[int]bool
	descriptions [][]string
	minX         int
	maxX         int
	minY         int
	maxY         int
}

func (d *Day24) setTile(x, y int) {
	if _, ok := d.tiles[y]; !ok {
		d.tiles[y] = make(map[int]bool)
	}
	d.tiles[y][x] = !d.tiles[y][x]

	// Set Bounds
	if x-1 < d.minX {
		d.minX = x - 1
	}
	if x+1 > d.maxX {
		d.maxX = x + 1
	}
	if y-1 < d.minY {
		d.minY = y - 1
	}
	if y+1 > d.maxY {
		d.maxY = y + 1
	}
}

func (d *Day24) getTile(x, y int) bool {
	return d.tiles[y][x]
}

func (d *Day24) init(s string) {
	d.tiles = make(map[int]map[int]bool)
	d.descriptions = make([][]string, 0)
	dirs := regexp.MustCompile("(se|sw|nw|ne|w|e)")
	for _, desc := range strings.Split(s, "\n") {
		movements := dirs.FindAllStringSubmatch(strings.TrimSpace(desc), -1)
		dirs := make([]string, len(movements))
		for i, m := range movements {
			dirs[i] = m[0]
		}
		d.descriptions = append(d.descriptions, dirs)
	}
}

func (d *Day24) countAdjacentBlacks(x, y int) int {
	count := 0
	for yO := -1; yO <= 1; yO++ {
		for xO := -1; xO <= 1; xO++ {
			if (xO == 0 && yO == 0) || (xO == -1 && yO == -1) || (xO == 1 && yO == 1) {
				continue
			}
			if d.getTile(x+xO, y+yO) {
				count++
			}
		}
	}
	return count
}

func (d *Day24) countBlackTiles() int {
	count := 0
	for _, y := range d.tiles {
		for _, x := range y {
			if x {
				count++
			}
		}
	}
	return count
}

func (d *Day24) executeA() int {
	for _, path := range d.descriptions {
		x := 0
		y := 0
		for _, dir := range path {
			switch dir {
			case "e":
				x++
				break
			case "se":
				y++
				break
			case "sw":
				x--
				y++
				break
			case "w":
				x--
				break
			case "nw":
				y--
				break
			case "ne":
				x++
				y--
				break
			}
		}
		d.setTile(x, y)
	}
	return d.countBlackTiles()
}

func (d *Day24) executeB() int {
	for i := 0; i < 100; i++ {
		flip := make([][2]int, 0)
		for y := d.minY; y <= d.maxY; y++ {
			for x := d.minX; x <= d.maxX; x++ {
				c := d.countAdjacentBlacks(x, y)
				isset := d.getTile(x, y)
				if (isset && (c == 0 || c > 2)) || (!isset && c == 2) {
					flip = append(flip, [2]int{x, y})
				}
			}
		}
		for _, f := range flip {
			d.setTile(f[0], f[1])
		}
	}
	return d.countBlackTiles()
}

func (d *Day24) Handle(s string) ([]string, error) {
	d.init(s)
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
