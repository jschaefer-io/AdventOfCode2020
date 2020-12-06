package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type group struct {
	answers []int
}

type Day6 struct {
	groups []group
}

func (d *Day6) init(s string) {
	splitG := regexp.MustCompile("((|\\r)\\n){2}")
	for _, grp := range splitG.Split(s, -1) {
		g := group{}
		for _, p := range strings.Split(grp, "\n") {
			answer := 0
			for _, q := range strings.TrimSpace(p) {
				answer |= int(math.Pow(2, float64(q-'a')))
			}
			g.answers = append(g.answers, answer)
		}
		d.groups = append(d.groups, g)
	}
}

func countBinaryOnes(number int) int {
	count := 0
	for number > 0 {
		count += number & 1
		number = number >> 1
	}
	return count
}

func (d *Day6) executeA() int {
	sum := 0
	for _, group := range d.groups {
		res := 0
		for _, answer := range group.answers {
			res |= answer
		}
		sum += countBinaryOnes(res)
	}
	return sum
}

func (d *Day6) executeB() int {
	sum := 0
	for _, group := range d.groups {
		res := ^0
		for _, answer := range group.answers {
			res &= answer
		}
		sum += countBinaryOnes(res)
	}
	return sum
}

func (d *Day6) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))

	return results, nil
}
