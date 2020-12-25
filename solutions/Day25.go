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

func (d *Day25) transform(sNum, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= sNum
		value %= 20201227
	}
	return value
}

func (d *Day25) findLoopSize(target int) int {
	loopSize := 1
	res := 1
	for {
		res = (res * 7) % 20201227
		if res == target {
			return loopSize
		}
		loopSize++
	}
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
	cSize := d.findLoopSize(d.cardKey)
	return d.transform(d.doorKey, cSize)
}

func (d *Day25) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	//results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
