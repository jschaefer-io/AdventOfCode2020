package solutions

import (
	"fmt"
	"math"
	"strings"
)

type Day11 struct {
	seats     [][]int
	backSeats [][]int
	width     int
	height    int
}

func (d *Day11) init(s string) {
	d.seats = make([][]int, 0)
	d.backSeats = make([][]int, 0)
	rows := strings.Split(s, "\n")
	d.height = len(rows)
	for _, row := range rows {
		sRow := make([]int, 0)
		sBackRow := make([]int, 0)
		seats := strings.Split(strings.TrimSpace(row), "")
		d.width = len(seats)
		for _, seat := range seats {
			v := 0
			if seat == "L" {
				v = 1
			} else if seat == "#" {
				v = 2
			}
			sRow = append(sRow, v)
			sBackRow = append(sBackRow, v)
		}
		d.seats = append(d.seats, sRow)
		d.backSeats = append(d.backSeats, sBackRow)
	}
}

func (d *Day11) switchPlane(reset bool) int {
	count := 0
	for y, r := range d.backSeats {
		for x, s := range r {
			d.seats[y][x] = s
			if s == 2 {
				count++
				if reset {
					d.seats[y][x] = 1
					d.backSeats[y][x] = 1
				}
			}
		}
	}
	return count
}

func (d *Day11) executeWithMethod(limit int, countFunction func(int, int) int) int {
	change := true
	var count int
	for change {
		change = false
		for y := 0; y < d.height; y++ {
			for x := 0; x < d.width; x++ {
				if d.seats[y][x] == 1 && countFunction(x, y) == 0 {
					d.backSeats[y][x] = 2
					change = true
				} else if d.seats[y][x] == 2 && countFunction(x, y) >= limit {
					d.backSeats[y][x] = 1
					change = true
				}
			}
		}
		count = d.switchPlane(false)
	}
	return count
}

func (d *Day11) printSeats() {
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.height; x++ {
			v := '.'
			if d.seats[y][x] == 2 {
				v = '#'
			}
			if d.seats[y][x] == 1 {
				v = 'L'
			}
			fmt.Printf("%c", v)
		}
		fmt.Print("\n")
	}
	fmt.Println()
	fmt.Println()
}

func (d *Day11) executeA() int {
	d.switchPlane(true)
	return d.executeWithMethod(4, func(x int, y int) int {
		count := 0
		for oX := -1; oX <= 1; oX++ {
			for oY := -1; oY <= 1; oY++ {
				cX := x + oX
				cY := y + oY
				if (cX == x && cY == y) || cX < 0 || cY < 0 || cX >= d.width || cY >= d.height {
					continue
				}
				if d.seats[cY][cX] == 2 {
					count++
				}
			}
		}
		return count
	})
}

func (d *Day11) executeB() int {
	d.switchPlane(true)
	return d.executeWithMethod(5, func(x int, y int) int {
		count := 0
		for oX := -1; oX <= 1; oX++ {
			for oY := -1; oY <= 1; oY++ {
				for oC := 1; oC < int(math.Max(float64(d.width), float64(d.height))); oC++ {
					cX := oX*oC + x
					cY := oY*oC + y
					if (cX == x && cY == y) || cX < 0 || cY < 0 || cX >= d.width || cY >= d.height {
						break
					}
					if d.seats[cY][cX] > 0 {
						if d.seats[cY][cX] == 2 {
							count++
						}
						break
					}
				}
			}
		}
		return count
	})
}

func (d *Day11) Handle(s string) ([]string, error) {
	d.init(s)
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
