package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day10 struct {
	adapters []int
}

func (d *Day10) init(s string) error {
	d.adapters = make([]int, 0)
	for _, a := range strings.Split(s, "\n") {
		joltage, err := strconv.Atoi(strings.TrimSpace(a))
		if err != nil {
			return err
		}
		d.adapters = append(d.adapters, joltage)
	}
	sort.Ints(d.adapters)
	return nil
}

func (d *Day10) executeA() int {
	diffs := make(map[int]int)
	prev := 0

	for _, adapter := range d.adapters {
		diff := adapter - prev
		if diff > 3 {
			continue
		}
		diffs[diff]++
		prev = adapter
	}
	// built-in
	diffs[3]++

	return diffs[1] * diffs[3]
}

func (d *Day10) executeB() int {

	// Build sorted list of all powers numbers
	aCount := len(d.adapters)
	fullList := make([]int, 0)
	fullList = append(fullList, 0)
	fullList = append(fullList, d.adapters...)
	fullList = append(fullList, d.adapters[aCount-1]+3)
	fCount := len(fullList)

	// calculate the number of possibilities
	// for one adapter to to the next higher one
	counts := make(map[int][]int)
	for i := fCount - 2; i >= 0; i-- {
		pos := make([]int, 0)
		for u := i + 1; u < fCount; u++ {
			if fullList[i]+3 < fullList[u] {
				break
			}
			pos = append(pos, u)
		}
		counts[i] = pos
	}

	dP := make(map[int][2]int)
	var countPos2 func(index int) (int, int)
	countPos2 = func(index int) (int, int) {
		next := counts[index]
		count := len(next)

		// Ending condition
		if count == 0 {
			return 0, 1
		}

		// check dP
		v, ok := dP[index]
		if ok {
			return v[0], v[1]
		}

		// Resolve recursively
		res := 0
		for _, v := range next {
			_, r := countPos2(v)
			res += r
		}
		dP[index] = [2]int{
			count,
			res,
		}
		return count, res
	}

	_, res := countPos2(0)
	return res
}

func (d *Day10) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))

	return results, nil

}
