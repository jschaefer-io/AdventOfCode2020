package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	password string
	char     rune
	min      int
	max      int
}

func (p password) testOccurances() bool {
	count := 0
	for _, char := range p.password {
		if char == p.char {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func (p password) testPositions() bool {
	a := rune(p.password[p.min-1]) == p.char
	b := rune(p.password[p.max-1]) == p.char
	return a != b
}

type Day2 struct {
	passwords []password
}

func (d *Day2) init(s string) {
	list := strings.Split(s, "\n")
	r := regexp.MustCompile("\\d+|\\w+")
	for _, pass := range list {
		m := r.FindAllString(pass, -1)
		obj := password{}
		obj.password = m[3]
		obj.char = rune(m[2][0])
		obj.min, _ = strconv.Atoi(m[0])
		obj.max, _ = strconv.Atoi(m[1])
		d.passwords = append(d.passwords, obj)
	}
}

func (d *Day2) executeA() int {
	count := 0
	for _, pw := range d.passwords {
		if pw.testOccurances() {
			count++
		}
	}
	return count
}

func (d *Day2) executeB() int {
	count := 0
	for _, pw := range d.passwords {
		if pw.testPositions() {
			count++
		}
	}
	return count
}

func (d *Day2) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))

	return results, nil
}
