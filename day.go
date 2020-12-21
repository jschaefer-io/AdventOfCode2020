package main

import "./solutions"

type Day interface {
	Handle(string) ([]string, error)
}

func getDays() []Day {
	days := []Day{
		&solutions.Day1{},
		&solutions.Day2{},
		&solutions.Day3{},
		&solutions.Day4{},
		&solutions.Day5{},
		&solutions.Day6{},
		&solutions.Day7{},
		&solutions.Day8{},
		&solutions.Day9{},
		&solutions.Day10{},
		&solutions.Day11{},
		&solutions.Day12{},
		&solutions.Day13{},
		&solutions.Day14{},
		&solutions.Day15{},
		&solutions.Day16{},
		&solutions.Day17{},
		&solutions.Day18{},
		&solutions.Day19{},
		&solutions.Day20{},
		&solutions.Day21{},
	}
	return days
}
