package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day9 struct {
	nums  []int
	sums  map[int][][2]int
	count int
}

func (d *Day9) init(s string) error {
	nums := strings.Split(s, "\n")
	d.count = len(nums)

	// parse input data
	d.nums = make([]int, d.count)
	for i, num := range nums {
		num, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return err
		}
		d.nums[i] = num
	}
	return nil
}

func (d *Day9) checkPosition(position int, inRange int, target int) bool {
	list := d.nums[position-inRange : position]
	for i := 0; i < inRange; i++ {
		for u := i + 1; u < inRange; u++ {
			if list[u]+list[i] == target {
				return true
			}
		}
	}
	return false
}

func (d *Day9) executeA() int {
	pre := 25
	check := 25
	for i := pre; i < d.count; i++ {
		if !d.checkPosition(i, check, d.nums[i]) {
			return d.nums[i]
		}
	}

	return -1
}

func (d *Day9) executeB(target int) int {
	a := 0
	b := 1
	sum := d.nums[a]
	for a < d.count && b < d.count {
		if sum < target {
			sum += d.nums[b]
			b++
			continue
		} else if sum > target {
			sum -= d.nums[a]
			a++
			continue
		} else {
			resList := make([]int, b-a)
			for i := 0; i < b-a; i++ {
				resList[i] = d.nums[i+a]
			}
			sort.Ints(resList)
			return resList[0] + resList[b-a-1]
		}
	}
	return -1
}

func (d *Day9) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	a := d.executeA()
	results = append(results, fmt.Sprintf("%d", a))
	results = append(results, fmt.Sprintf("%d", d.executeB(a)))

	return results, nil
}
