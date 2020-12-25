package solutions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day23 struct {
	cups     []*linkedCup
	circle   *linkedCup
	maxValue int
	minValue int
	cupMap   []*linkedCup
}

type linkedCup struct {
	prev  *linkedCup
	next  *linkedCup
	value int
}

func (d *Day23) init(s string) {
	cups := strings.Split(strings.TrimSpace(s), "")
	count := len(cups)
	d.cups = make([]*linkedCup, count)

	for i, cupStr := range cups {
		value, err := strconv.Atoi(cupStr)
		if err != nil {
			panic(err)
		}
		cup := linkedCup{
			value: value,
		}
		d.cups[i] = &cup

		if i == 0 || cup.value < d.minValue {
			d.minValue = cup.value
		}
		if i == 0 || cup.value > d.maxValue {
			d.maxValue = cup.value
		}
	}
	for i, cup := range d.cups {
		cup.next = d.cups[(i+1+count)%count]
		cup.prev = d.cups[(i-1+count)%count]
	}
	d.circle = d.cups[0]
}

func (d *Day23) printCircle() {
	done := make(map[int]struct{})
	cup := d.circle
	fmt.Print("cups: ")
	for {
		fmt.Printf(" %d", cup.value)
		done[cup.value] = struct{}{}
		cup = cup.next

		if _, ok := done[cup.value]; ok {
			break
		}
	}
	fmt.Print("\n")
}

func (d *Day23) findCup(value int, excluded []*linkedCup) (*linkedCup, error) {
	cup := d.cupMap[value-1]
	for _, c := range excluded {
		if c.value == value {
			return nil, errors.New("cup not found")
		}
	}
	return cup, nil
}

func (d *Day23) execute(times int) {
	pickup := make([]*linkedCup, 3)
	for move := 0; move < times; move++ {
		currentCup := d.circle
		nextCup := currentCup.next
		for i := 0; i < 3; i++ {
			pickup[i] = nextCup
			nextCup = nextCup.next
		}
		currentCup.next = nextCup
		nextCup.prev = currentCup

		var destinationCup *linkedCup
		destination := currentCup.value
		for {
			destination = destination - 1
			if destination < d.minValue {
				destination = d.maxValue
			}
			cup, err := d.findCup(destination, pickup)
			if err != nil {
				continue
			}
			destinationCup = cup
			break
		}
		pickup[0].prev = destinationCup
		pickup[2].next = destinationCup.next
		destinationCup.next = pickup[0]
		destinationCup.prev = pickup[2]
		d.circle = currentCup.next
	}
}

func (d *Day23) executeA(s string) int {
	d.init(s)
	d.buildMap(len(d.cups))
	d.execute(100)

	res := 0
	cup, _ := d.findCup(1, make([]*linkedCup, 0))
	for {
		cup = cup.next
		if cup.value == 1 {
			break
		}
		res *= 10
		res += cup.value
	}
	return res
}

func (d *Day23) buildMap(count int) {
	cup := d.circle
	con := cup.value
	cupMap := make([]*linkedCup, count)
	for {
		cupMap[cup.value-1] = cup
		cup = cup.next
		if cup.value == con {
			break
		}
	}
	d.cupMap = cupMap
}

func (d *Day23) executeB(s string) int {
	d.init(s)

	// Fill circle to 1000000
	limit := 1000000
	prev := d.cups[len(d.cups)-1]
	for i := d.maxValue + 1; i <= limit; i++ {
		cup := &linkedCup{
			prev:  prev,
			value: i,
		}
		prev.next = cup
		prev = cup
	}
	prev.next = d.cups[0]
	d.cups[0].prev = prev
	d.maxValue = limit

	d.buildMap(limit)
	d.execute(10000000)

	cup, _ := d.findCup(1, make([]*linkedCup, 0))
	return cup.next.value * cup.next.next.value
}

func (d *Day23) Handle(s string) ([]string, error) {
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA(s)))
	results = append(results, fmt.Sprintf("%d", d.executeB(s)))
	return results, nil
}
