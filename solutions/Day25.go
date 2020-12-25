package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Day25 struct {
	cardKey int
	doorKey int
}

func (d *Day25) init(s string) error {
	nums := strings.Split(s, "\n")
	cardKey, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil {
		return err
	}
	doorKey, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil {
		return err
	}
	d.cardKey = cardKey
	d.doorKey = doorKey
	return nil
}

func (d *Day25) executeA() int {
	cSize := 1
	res := 1
	for {
		res = (res * 7) % 20201227
		if res == d.cardKey {
			break
		}
		cSize++
	}
	value := 1
	for i := 0; i < cSize; i++ {
		value *= d.doorKey
		value %= 20201227
	}
	return value
}

func (d *Day25) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	return results, nil
}
