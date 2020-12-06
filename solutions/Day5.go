package solutions

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strings"
)

func binaryReduction(min int, max int, list []bool) int {
	mid := float64(max+min) / 2
	if min == max {
		return max
	}
	if list[0] {
		return binaryReduction(int(math.Ceil(mid)), max, list[1:len(list)])
	} else {
		return binaryReduction(min, int(math.Floor(mid)), list[1:len(list)])
	}
}

type boardingPass struct {
	num string
	row int
	col int
	id  int
}

func newBoardingPass(number string, numberRows int, numberCols int) boardingPass {
	rowList := number[0:7]
	colList := number[7:10]

	rows := make([]bool, len(rowList))
	cols := make([]bool, len(colList))

	for i, f := range rowList {
		if f == 'B' {
			rows[i] = true
		}
	}

	for i, f := range colList {
		if f == 'R' {
			cols[i] = true
		}
	}

	rowId := binaryReduction(0, numberRows-1, rows)
	colId := binaryReduction(0, numberCols-1, cols)

	return boardingPass{
		num: number,
		row: rowId,
		col: colId,
		id:  rowId*8 + colId,
	}
}

func (b *boardingPass) exist() bool {
	return b.id != 0 && b.num != ""
}

type Day5 struct {
	passes []boardingPass
	ids    map[int]boardingPass
	rows   int
	cols   int
	seats  [][]boardingPass
}

func (d *Day5) init(s string) {
	// setup plane data
	d.ids = make(map[int]boardingPass)
	d.rows = 128
	d.cols = 8

	// setup boarding passes
	for _, num := range strings.Split(s, "\n") {
		newPass := newBoardingPass(strings.TrimSpace(num), d.rows, d.cols)
		d.passes = append(d.passes, newPass)
		d.ids[newPass.id] = newPass
	}

	// fill seats
	d.seats = make([][]boardingPass, d.rows)
	for i, _ := range d.seats {
		d.seats[i] = make([]boardingPass, d.cols)
	}

	// assign passes to seats
	for _, pass := range d.passes {
		d.seats[pass.row][pass.col] = pass
	}
}

func (d *Day5) executeA() int {
	max := -1
	for _, pass := range d.passes {
		if pass.id > max {
			max = pass.id
		}
	}
	return max
}

func (d *Day5) executeB() (int, error) {
	// find first row
	firstRow := func() int {
		for i := 0; i < len(d.seats); i++ {
			for _, seat := range d.seats[i] {
				if seat.exist() {
					return i
				}
			}
		}
		return -1
	}()
	if firstRow == -1 {
		return 0, errors.New("no filled seats found")
	}

	// find first row
	lastRow := func() int {
		for i := len(d.seats) - 1; i >= 0; i-- {
			for _, seat := range d.seats[i] {
				if seat.exist() {
					return i
				}
			}
		}
		return -1
	}()
	if lastRow == -1 {
		return 0, errors.New("no filled seats found")
	}

	// check for empty seats
	ignoreEmpty := true
	for u, row := range d.seats[firstRow:lastRow] {
		for i, seat := range row {
			if seat.exist() {
				ignoreEmpty = false
			} else if !seat.exist() && !ignoreEmpty {
				return (firstRow+u)*8 + i, nil
			}
		}
	}
	return 1, errors.New("no empty seat found")
}

func (d *Day5) Handle(s string) ([]string, error) {
	d.init(s)
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	b, err := d.executeB()
	if err != nil {
		return nil, err
	}
	results = append(results, fmt.Sprintf("%d", b))
	return results, nil
}
