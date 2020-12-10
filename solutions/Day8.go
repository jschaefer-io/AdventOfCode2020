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
	instructions []instruction
	executed     map[int]struct{}
	flip         int
}

type instruction struct {
	cmd string
	op  int
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
	d.doInstruction(op)
	d.executed[currentIp] = struct{}{}
	d.ip++
	return currentIp, nil, false
}

func (d *Day8) reset(flip int) {
	d.executed = make(map[int]struct{})
	d.acc = 0
	d.ip = 0
	d.flip = flip
}

func (d *Day8) init(s string) {
	r := regexp.MustCompile("(\\w+) (-|\\+)(\\d+)")
	d.instructions = make([]instruction, 0)
	for _, line := range strings.Split(s, "\n") {
		g := r.FindStringSubmatch(line)
		op, _ := strconv.Atoi(g[3])

		if g[2] == "-" {
			op *= -1
		}

		d.instructions = append(d.instructions, instruction{
			cmd: g[1],
			op:  op,
		})
	}
}

func (d *Day8) doInstruction(i instruction) {
	if d.flip == d.ip {
		if i.cmd == "jmp" {
			i.cmd = "nop"
		} else if i.cmd == "nop" {
			i.cmd = "jmp"
		}
	}
	switch i.cmd {
	case "acc":
		d.acc += i.op
		break
	case "jmp":
		d.ip += i.op - 1
		break
	default:
		break
	}
}

func (d *Day8) executeA(flip int) (int, bool, bool) {
	d.reset(flip)
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
		acc, _, success := d.executeA(i)
		if success {
			return acc, nil
		}
	}
	return 0, errors.New("no terminating set found")
}

func (d *Day8) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)

	a, _, _ := d.executeA(-1)
	results = append(results, fmt.Sprintf("%d", a))

	b, err := d.executeB(s)
	results = append(results, fmt.Sprintf("%d", b))

	return results, err
}
