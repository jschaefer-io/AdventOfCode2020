package solutions

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type Day8 struct {
	acc          int
	ip           int
	instructions []func(day *Day8)
	executed     map[int]struct{}
}

func (d *Day8) nextTick() (int, error, bool) {
	currentIp := d.ip
	if _, ok := d.executed[currentIp]; ok {
		return currentIp, errors.New("loop detected"), false
	}
	if currentIp == len(d.instructions) {
		return currentIp, nil, true
	}
	op := d.instructions[currentIp]
	op(d)
	d.executed[currentIp] = struct{}{}
	d.ip++
	return currentIp, nil, false
}

func (d *Day8) reset(s string) {
	r := regexp.MustCompile("(\\w+) (-|\\+)(\\d+)")
	d.executed = make(map[int]struct{})
	d.instructions = make([]func(day *Day8), 0)
	d.acc = 0
	d.ip = 0
	for _, line := range strings.Split(s, "\n") {
		g := r.FindStringSubmatch(line)
		op, _ := strconv.Atoi(g[3])

		if g[2] == "-" {
			op *= -1
		}

		switch g[1] {
		case "acc":
			d.instructions = append(d.instructions, func(day *Day8) {
				day.acc += op
			})
			break
		case "jmp":
			d.instructions = append(d.instructions, func(day *Day8) {
				day.ip += op - 1
			})
			break
		default:
			d.instructions = append(d.instructions, func(day *Day8) {
				// noop
			})
			break
		}
	}
}

func (d *Day8) executeA(s string) (int, bool, bool) {
	d.reset(s)
	loop := false
	success := false
	for true {
		_, err, terminated := d.nextTick()
		if err != nil {
			loop = true
			break
		}
		if terminated {
			success = true
			break
		}
	}
	return d.acc, loop, success
}

func (d *Day8) executeB(s string) (int, error) {
	lines := strings.SplitAfter(s, "\n")
	for i, _ := range lines {
		var strBuilder strings.Builder
		for ln, line := range lines {
			if ln == i {
				if strings.Contains(line, "jmp") {
					line = strings.ReplaceAll(line, "jmp", "nop")
				} else if strings.Contains(line, "nop") {
					line = strings.ReplaceAll(line, "nop", "jmp")
				}
			}
			strBuilder.WriteString(line)
		}
		acc, _, success := d.executeA(strBuilder.String())
		if success {
			return acc, nil
		}
	}
	return 0, errors.New("no terminating set found")
}

func (d *Day8) Handle(s string) ([]string, error) {

	results := make([]string, 0)

	a, _, _ := d.executeA(s)
	results = append(results, fmt.Sprintf("%d", a))

	b, err := d.executeB(s)
	results = append(results, fmt.Sprintf("%d", b))

	return results, err
}
