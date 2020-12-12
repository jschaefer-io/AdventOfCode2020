package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day12 struct {
	instructions []navInstruction
	x            int
	y            int
	wX           int
	wY           int
	dir          int // mod 4 -> north = 0
}

type navInstruction struct {
	op    rune
	value int
}

func (d *Day12) init(s string) error {
	r := regexp.MustCompile("(\\w)(\\d+)")
	d.instructions = make([]navInstruction, 0)
	for _, i := range strings.Split(s, "\n") {
		res := r.FindStringSubmatch(i)
		op := rune(res[1][0])
		value, err := strconv.Atoi(res[2])
		if err != nil {
			return err
		}
		d.instructions = append(d.instructions, navInstruction{
			op:    op,
			value: value,
		})
	}
	return nil
}

func (d *Day12) executeA() int {
	d.x = 0
	d.y = 0
	d.dir = 1
	for _, instruction := range d.instructions {
		switch instruction.op {
		case 'N':
			d.y += instruction.value
			break
		case 'S':
			d.y -= instruction.value
			break
		case 'E':
			d.x += instruction.value
			break
		case 'W':
			d.x -= instruction.value
			break
		case 'L',
			'R':
			rotation := instruction.value / 90
			if instruction.op == 'L' {
				rotation *= -1
			}
			d.dir += rotation
			d.dir = ((d.dir % 4) + 4) % 4
			break
		default:
			switch d.dir {
			case 0:
				d.y += instruction.value
				break
			case 1:
				d.x += instruction.value
				break
			case 2:
				d.y -= instruction.value
				break
			case 3:
				d.x -= instruction.value
				break
			}
		}
	}
	return int(math.Abs(float64(d.x)) + math.Abs(float64(d.y)))
}

func (d *Day12) executeB() int {
	d.x = 0
	d.y = 0
	d.wX = 10
	d.wY = 1
	for _, instruction := range d.instructions {
		switch instruction.op {
		case 'N':
			d.wY += instruction.value
			break
		case 'S':
			d.wY -= instruction.value
			break
		case 'E':
			d.wX += instruction.value
			break
		case 'W':
			d.wX -= instruction.value
			break
		case 'L',
			'R':
			rotation := instruction.value
			if instruction.op == 'R' {
				rotation *= -1
			}
			angle := float64(rotation) * math.Pi / 180

			fWx := float64(d.wX)
			fWy := float64(d.wY)

			nX := math.Round(fWx*math.Cos(angle) - fWy*math.Sin(angle))
			nY := math.Round(fWy*math.Cos(angle) + fWx*math.Sin(angle))

			d.wX = int(nX)
			d.wY = int(nY)
			break
		default:
			d.x += instruction.value * d.wX
			d.y += instruction.value * d.wY
		}
	}

	return int(math.Abs(float64(d.x)) + math.Abs(float64(d.y)))
}

func (d *Day12) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
