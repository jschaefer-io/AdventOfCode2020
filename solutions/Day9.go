package solutions

import (
	"fmt"
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

	// Create result list
	d.sums = make(map[int][][2]int)
	for i := 0; i < d.count; i++ {
		for u := i + 1; u < d.count; u++ {
			sum := d.nums[i] + d.nums[u]
			list, ok := d.sums[sum]
			if !ok {
				list = make([][2]int, 0)
			}
			d.sums[sum] = append(list, [2]int{
				i,
				u,
			})
		}
	}
	return nil
}

func (d *Day9) checkPosition(position int, inRange int) bool {
	pairs, ok := d.sums[d.nums[position]]
	if !ok {
		return false
	}

	minIndex := position - inRange
	for _, pair := range pairs {
		if pair[0] >= minIndex && pair[1] >= minIndex && pair[0] < position && pair[1] < position {
			return true
		}
	}

	return false
}

func (d *Day9) executeA() int {
	pre := 25
	check := 25
	for i := pre; i < d.count; i++ {
		if !d.checkPosition(i, check) {
			return d.nums[i]
		}
	}

	return -1
}

func (d *Day9) executeB(target int) int {
	for i := 0; i < d.count; i++ {
		min := d.nums[i]
		max := d.nums[i]
		sum := d.nums[i]
		for u := i + 1; u < d.count; u++ {
			sum += d.nums[u]

			// Update max values
			if d.nums[u] < min {
				min = d.nums[u]
			}
			if d.nums[u] > max {
				max = d.nums[u]
			}

			// Continue if over target
			if sum > target {
				continue
			}

			// return result if target reached
			if sum == target {
				return min + max
			}
		}
	}
	return 0
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
