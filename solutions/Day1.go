package solutions

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Day1 struct {
	max     int
	results []string
	intMap  map[int]bool
}

func (d *Day1) init(s string) error {
	d.max = 2020
	d.results = make([]string, 0)
	d.intMap = make(map[int]bool)
	for _, num := range strings.Split(s, "\n") {
		n, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return err
		}
		d.intMap[n] = true
	}
	return nil
}

func (d Day1) findCounterPart(search int) bool {
	_, ok := d.intMap[search]
	return ok
}

func (d *Day1) executeA() error {
	for num, _ := range d.intMap {
		search := d.max - num
		if d.findCounterPart(search) {
			d.results = append(d.results, fmt.Sprintf("%d", search*num))
			return nil
		}
	}
	return errors.New("no result for first part found")
}

func (d *Day1) executeB() error {
	for a, _ := range d.intMap {
		for b, _ := range d.intMap {
			search := d.max - a - b
			if d.findCounterPart(search) {
				d.results = append(d.results, fmt.Sprintf("%d", search*a*b))
				return nil
			}
		}
	}
	return errors.New("no result for second part found")
}

func (d *Day1) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}

	// Execute Day 1
	err = d.executeA()
	if err != nil {
		return nil, err
	}

	// Execute Day 2
	err = d.executeB()
	if err != nil {
		return nil, err
	}

	return d.results, nil
}
