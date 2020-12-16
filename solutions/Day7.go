package solutions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day7 struct {
	bags map[string][]string
}

func (d *Day7) init(s string) {
	r := regexp.MustCompile("(\\d+) (\\w+ \\w+)")
	bagList := make(map[string][]string)
	for _, bag := range strings.Split(s, "\n") {
		data := strings.Split(bag, "bags contain")
		key := strings.TrimSpace(data[0])
		val := make([]string, 0)
		for _, typeDef := range strings.Split(data[1], ",") {
			typeData := r.FindStringSubmatch(typeDef)
			if len(typeData) == 0 {
				continue
			}
			size, _ := strconv.Atoi(typeData[1])
			for i := 0; i < size; i++ {
				val = append(val, strings.TrimSpace(typeData[2]))
			}
		}
		bagList[key] = val
	}
	d.bags = bagList
}

func (d *Day7) getBagVariations(bag string) []string {
	res := make([]string, 0)
	for name, list := range d.bags {
		if sliceContains(list, bag) {
			res = append(res, name)
			res = append(res, d.getBagVariations(name)...)
		}
	}
	return res
}

func (d *Day7) countBagsInside(bag string) (int, error) {
	list, ok := d.bags[bag]
	if !ok {
		return 0, errors.New("bag not found")
	}
	count := len(list)
	for _, subBag := range list {
		sCount, err := d.countBagsInside(subBag)
		if err != nil {
			return 0, err
		}
		count += sCount
	}
	return count, nil
}

func removeDuplications(s []string) []string {
	set := make(map[string]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	res := make([]string, 0, len(set))
	for v := range set {
		res = append(res, v)
	}
	return res
}

func sliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func (d *Day7) executeA() int {
	res := d.getBagVariations("shiny gold")
	res = removeDuplications(res)
	return len(res)
}

func (d *Day7) executeB() int {
	res, err := d.countBagsInside("shiny gold")
	if err != nil {
		return 0
	}
	return res
}

func (d *Day7) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	//
	return results, nil
}
