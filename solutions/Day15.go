package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day15 struct {
	record  map[int][2]int
	turns   map[int]int
	current int
	first   bool
}

func (d *Day15) init(s string) error {
	d.record = make(map[int][2]int)
	d.turns = make(map[int]int)
	d.first = true
	d.current = 1
	for _, n := range strings.Split(s, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil
		}
		d.record[num] = [2]int{0, d.current}
		d.current++
	}
	return nil
}

func (d *Day15) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)

	for  {
		var say int
		var list [2]int
		if !d.first {
			list = d.record[d.turns[d.current-1]]
			say = list[1] - list[0]
		}
		list, ok := d.record[say]
		if !ok {
			list = [2]int{}
		}
		d.first = !ok
		list[0] = list[1]
		list[1] = d.current
		d.record[say] = list
		d.turns[d.current] = say
		if d.current == 2020 {
			results = append(results, fmt.Sprintf("%d", say))
		}
		if d.current == 3000000 {
			results = append(results, fmt.Sprintf("%d", say))
			break
		}
		d.current++
	}

	return results, nil
}
