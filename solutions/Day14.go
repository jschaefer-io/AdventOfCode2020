package solutions

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day14 struct {
	orMap    int
	andMap   int
	floatMap []int
	memoryA  map[int]int
	memoryB  map[int]int
}

func (d *Day14) init(s string) error {
	d.memoryA = make(map[int]int)
	d.memoryB = make(map[int]int)
	r := regexp.MustCompile("\\d+")
	for _, line := range strings.Split(s, "\n") {
		eq := strings.Split(strings.TrimSpace(line), "=")
		if strings.TrimSpace(eq[0]) == "mask" {
			d.orMap = 0
			d.andMap = 0
			d.floatMap = make([]int, 0)
			mapStr := strings.TrimSpace(eq[1])
			mapLength := len(mapStr)
			for i, c := range mapStr {
				switch c {
				case '1':
					d.orMap += 1 << (mapLength - i - 1)
					break
				case '0':
					d.andMap += 1 << (mapLength-i-1)
					break
				default:
					d.floatMap = append(d.floatMap, mapLength-i-1)
					sort.Ints(d.floatMap)
				}
			}
		} else {
			index, err := strconv.Atoi(r.FindString(eq[0]))
			if err != nil {
				return err
			}
			value, err := strconv.Atoi(strings.TrimSpace(eq[1]))
			if err != nil {
				return err
			}

			// Solution A
			d.memoryA[index] = value&(^d.andMap) | d.orMap

			// Solution B
			index = index | d.orMap
			d.memoryB[index] = value
			lenMap := len(d.floatMap)
			for i := 0; i < 1 << lenMap; i++ {
				target := index
				for u, n := range d.floatMap {
					num := i & (1 << u) << (n - u)
					if num != 0 {
						target |= num
					} else {
						target &= ^(1 << n)
					}
				}
				d.memoryB[target] = value
			}
		}
	}
	return nil
}

func (d *Day14) executeA() int {
	sum := 0
	for _, v := range d.memoryA {
		sum += v
	}
	return sum
}

func (d *Day14) executeB() int {
	sum := 0
	for _, v := range d.memoryB {
		sum += v
	}
	return sum
}

func (d *Day14) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
