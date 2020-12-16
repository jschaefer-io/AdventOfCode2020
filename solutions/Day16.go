package solutions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day16 struct {
	rules     map[string]map[int]struct{}
	validNums map[int]struct{}
	ticket    []int
	tickets   [][]int
}

func (d *Day16) init(s string) error {
	d.rules = make(map[string]map[int]struct{})
	d.validNums = make(map[int]struct{})
	d.ticket = make([]int, 0)
	d.tickets = make([][]int, 0)

	splitG := regexp.MustCompile("((|\\r)\\n){2}")
	findNums := regexp.MustCompile("(\\d+)-(\\d+)")
	groups := splitG.Split(s, -1)

	if len(groups) != 3 {
		return errors.New("not enough input groups")
	}

	// rules
	for _, rule := range strings.Split(groups[0], "\n") {
		rulePartial := strings.Split(strings.TrimSpace(rule), ": ")
		ruleBounds := findNums.FindAllStringSubmatch(rulePartial[1], -1)

		list := make(map[int]struct{})

		for _, v := range ruleBounds {
			a, err := strconv.Atoi(v[1])
			if err != nil {
				return err
			}
			b, err := strconv.Atoi(v[2])
			if err != nil {
				return err
			}
			for i := a; i <= b; i++ {
				list[i] = struct{}{}
				d.validNums[i] = struct{}{}
			}
		}

		d.rules[rulePartial[0]] = list
	}

	// ticket
	ticketInfo := strings.Split(groups[1], "\n")
	for _, num := range strings.Split(ticketInfo[1], ",") {
		ticketNum, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return err
		}
		d.ticket = append(d.ticket, ticketNum)
	}

	// other tickets
	ticketList := strings.Split(groups[2], "\n")
	for _, v := range ticketList[1:] {
		ticket := make([]int, 0)
		for _, num := range strings.Split(v, ",") {
			ticketNum, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				return err
			}
			ticket = append(ticket, ticketNum)
		}
		d.tickets = append(d.tickets, ticket)
	}
	return nil
}

func (d *Day16) executeA() (int, map[int]struct{}) {
	wrongIndex := make(map[int]struct{})
	errRate := 0
	for index, ticket := range d.tickets {
		for _, ticketNum := range ticket {
			if _, ok := d.validNums[ticketNum]; !ok {
				errRate += ticketNum
				wrongIndex[index] = struct{}{}
			}
		}
	}
	return errRate, wrongIndex
}

func (d *Day16) executeB(ignoreList map[int]struct{}) int {
	foundList := make(map[string][]int)

	for i, tValue := range d.ticket {
		for name, rule := range d.rules {
			var found bool
			for tId, ticket := range d.tickets {
				if _, ok := ignoreList[tId]; ok {
					continue
				}

				found = true
				if _, ok := rule[ticket[i]]; !ok {
					found = false
					break
				}
			}
			if found {
				foundList[name] = append(foundList[name], tValue)
			}

		}
	}

	// resolve duplicated assignments
	fixed := make(map[int]struct{})
	changed := true
	for changed {
		changed = false
		for fIndex, list := range foundList {
			if len(list) == 1 {
				if _, ok := fixed[list[0]]; ok {
					continue
				}
				fixed[list[0]] = struct{}{}
			} else {
				newList := make([]int, 0)
				for _, v := range list {
					if _, ok := fixed[v]; !ok {
						newList = append(newList, v)
					}
				}
				foundList[fIndex] = newList
			}
			changed = true
		}
	}

	res := 1
	for name, value := range foundList {
		if strings.Contains(name, "departure") {
			res *= value[0]
		}
	}
	return res
}

func (d *Day16) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)

	a, ignore := d.executeA()
	results = append(results, fmt.Sprintf("%d", a))
	results = append(results, fmt.Sprintf("%d", d.executeB(ignore)))
	return results, nil
}
