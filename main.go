package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	days := getDays()
	start := 1
	max := 8
	accTime := .0
	for i := start; i <= max; i++ {
		fName := fmt.Sprintf("./inputs/%d.txt", i)
		data, err := ioutil.ReadFile(fName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not read file %s", fName))
			continue
		}
		day, ok := days[i]
		if !ok {
			fmt.Println(fmt.Sprintf("Day%d Struct not defined yet", i))
			continue
		}
		startDay := time.Now()
		results, err := day.Handle(string(data))
		endDay := time.Since(startDay)
		accTime += endDay.Seconds()
		fmt.Println(fmt.Sprintf("# Day %d results (%s):", i, endDay))
		if err != nil {
			fmt.Println(err)
		} else {
			for index, result := range results {
				fmt.Println(fmt.Sprintf("%d: %s", index+1, result))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(fmt.Sprintf("---\nAOC Execution Time %fs", accTime))
}
