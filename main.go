package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	days := getDays()
	start := 6
	max := 6
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
		results, err := day.Handle(string(data))
		fmt.Println(fmt.Sprintf("Day %d results:", i))
		if err != nil {
			fmt.Println(err)
		} else {
			for index, result := range results {
				fmt.Println(fmt.Sprintf("%d: %s", index+1, result))
			}
		}
		fmt.Print("\n")
	}
}
