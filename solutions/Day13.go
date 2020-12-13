package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day13 struct {
	minTime   int
	busList   []int
	offsetMap map[int]int
}

func (d *Day13) init(s string) error {
	d.busList = make([]int, 0)
	d.offsetMap = make(map[int]int)

	lists := strings.Split(s, "\n")
	minTime, err := strconv.Atoi(strings.TrimSpace(lists[0]))
	if err != nil {
		return err
	}
	d.minTime = minTime
	r := regexp.MustCompile("(\\d+|\\w)")

	for i, v := range r.FindAllString(lists[1], -1) {
		value, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		d.offsetMap[value] = i
		//d.offsetMap[value] = (value - i) % value
		d.busList = append(d.busList, value)
	}
	return nil
}

func (d *Day13) executeA() int {
	min := -1
	id := 0
	for _, v := range d.busList {
		value := v - (d.minTime % v)
		if min < 0 || value < min {
			min = value
			id = v
		}
	}
	return min * id
}

func (d *Day13) executeB() int {
	time := 0
	acc := 1
	for _, busTime := range d.busList {
		i := d.offsetMap[busTime]
		// double the time until phase matches
		// doubling the time conserves the previous phase state
		for (time+i)%busTime > 0 {
			time += acc
		}
		// equal phase time
		acc *= busTime
	}
	return time
}

func (d *Day13) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
