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
	return 1
	//aCount := len(d.adapters)
	//fullList := make([]int, 0)
	//fullList = append(fullList, 0)
	//fullList = append(fullList, d.adapters...)
	//fullList = append(fullList, d.adapters[aCount-1]+3)
	//fCount := len(fullList)
	//
	//res := 1
	//for i := 0; i < fCount - 1; i++{
	//	pos := 0
	//	for u := i+1; u < fCount; u++{
	//		if fullList[i] + 3 < fullList[u] {
	//			break
	//		}
	//		pos++
	//	}
	//	fmt.Printf("Possible Steprs from %d : %d\n", fullList[i], pos)
	//	res *= pos
	//}
	//return res
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
