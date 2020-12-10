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
	}
	return days
}
