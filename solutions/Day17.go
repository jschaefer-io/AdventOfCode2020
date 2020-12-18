package solutions

import (
	"fmt"
	"strings"
)

type Day17 struct {
	minX    int
	minY    int
	minZ    int
	minW    int
	maxX    int
	maxY    int
	maxZ    int
	maxW    int
	actives map[int]map[int]map[int]map[int]bool
	changes []changes
}

type changes struct {
	x     int
	y     int
	z     int
	w     int
	value bool
}

func (d *Day17) applyChanges() {
	for _, change := range d.changes {
		d.actives[change.z][change.y][change.x][change.w] = change.value
		if change.x <= d.minX {
			d.minX = change.x - 1
		}
		if change.y <= d.minY {
			d.minY = change.y - 1
		}
		if change.z <= d.minZ {
			d.minZ = change.z - 1
		}
		if change.w <= d.minW {
			d.minW = change.w - 1
		}
		if change.x >= d.maxX {
			d.maxX = change.x + 1
		}
		if change.y >= d.maxY {
			d.maxY = change.y + 1
		}
		if change.z >= d.maxZ {
			d.maxZ = change.z + 1
		}
		if change.w >= d.maxW {
			d.maxW = change.w + 1
		}
	}
	d.changes = make([]changes, 0)
}

func (d *Day17) setField(set bool, x, y, z, w int) {
	var ok bool
	_, ok = d.actives[z]
	if !ok {
		d.actives[z] = make(map[int]map[int]map[int]bool)
	}
	_, ok = d.actives[z][y]
	if !ok {
		d.actives[z][y] = make(map[int]map[int]bool)
	}

	_, ok = d.actives[z][y][x]
	if !ok {
		d.actives[z][y][x] = make(map[int]bool)
	}

	d.changes = append(d.changes, changes{
		x:     x,
		y:     y,
		z:     z,
		w:     w,
		value: set,
	})
}

func (d *Day17) getField(x, y, z, w int) bool {
	var ok bool
	_, ok = d.actives[z]
	if !ok {
		return false
	}
	_, ok = d.actives[z][y]
	if !ok {
		return false
	}
	return d.actives[z][y][x][w]
}

func (d *Day17) init(s string) {
	d.actives = make(map[int]map[int]map[int]map[int]bool)
	d.changes = make([]changes, 0)
	for y, row := range strings.Split(s, "\n") {
		for x, col := range strings.Split(strings.TrimSpace(row), "") {
			if col == "#" {
				d.setField(true, x, y, 0, 0)
			}
		}
	}
	d.applyChanges()
}

func (d *Day17) checkNeighborCount(x, y, z, w int) int {
	count := 0
	for xO := -1; xO <= 1; xO++ {
		for yO := -1; yO <= 1; yO++ {
			for zO := -1; zO <= 1; zO++ {
				for wO := -1; wO <= 1; wO++ {
					if xO == 0 && yO == 0 && zO == 0 && wO == 0 {
						continue
					}
					if d.getField(x+xO, y+yO, z+zO, w+wO) {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d *Day17) executeA() int {
	for i := 0; i < 6; i++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			for y := d.minY; y <= d.maxY; y++ {
				for x := d.minX; x <= d.maxX; x++ {
					count := d.checkNeighborCount(x, y, z, 0)
					if d.getField(x, y, z, 0) {
						if count != 2 && count != 3 {
							d.setField(false, x, y, z, 0)
						}
					} else {
						if count == 3 {
							d.setField(true, x, y, z, 0)
						}
					}
				}
			}
		}
		d.applyChanges()
	}
	return d.countActive()
}

func (d *Day17) executeB() int {
	for i := 0; i < 6; i++ {
		for z := d.minZ; z <= d.maxZ; z++ {
			for y := d.minY; y <= d.maxY; y++ {
				for x := d.minX; x <= d.maxX; x++ {
					for w := d.minW; w <= d.maxW; w++ {
						count := d.checkNeighborCount(x, y, z, w)
						if d.getField(x, y, z, w) {
							if count != 2 && count != 3 {
								d.setField(false, x, y, z, w)
							}
						} else {
							if count == 3 {
								d.setField(true, x, y, z, w)
							}
						}
					}
				}
			}
		}
		d.applyChanges()
	}

	return d.countActive()
}

func (d *Day17) countActive() int {
	count := 0
	for _, z := range d.actives {
		for _, x := range z {
			for _, y := range x {
				for _, w := range y {
					if w {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d *Day17) Handle(s string) ([]string, error) {
	results := make([]string, 0)
	d.init(s)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	d.init(s)
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
