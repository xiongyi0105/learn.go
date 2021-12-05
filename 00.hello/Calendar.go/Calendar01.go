package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		calendar = [12][]int{}
	)
	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Year:")
	year, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("Year is %s\n", year)
	}
	// calendar[1] = []int{}
	for month, _ := range calendar {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			for i := 1; i <= 31; i++ {
				calendar[month] = append(calendar[month], i)
			}
		case 2:
			year, err := strconv.Atoi(year)
			if err == nil {
				fmt.Println("year convert to int false:", year, err)
			}
			if year%4 == 0 {
				for i := 1; i <= 29; i++ {
					calendar[month] = append(calendar[month], i)
				}
			} else {
				for i := 1; i <= 28; i++ {
					calendar[month] = append(calendar[month], i)
				}
			}
		default:
			for i := 1; i <= 30; i++ {
				calendar[month] = append(calendar[month], i)
			}

		}

	}
	for month, days := range calendar {
		fmt.Printf("%s月\n", strconv.Itoa(month+1))
		for index, day := range days {
			if len(strconv.Itoa(day)) == 2 {
				fmt.Printf("%d ", day)
			} else {
				fmt.Printf("%d  ", day)
			}

			if day%7 == 0 {
				fmt.Println("")
			} else if index == len(days)-1 {
				fmt.Println("")

			}

		}
	}
}
