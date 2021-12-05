package main

import (
	"fmt"
	"strings"
)

func Pentagram() string {
	fmt.Println("         *  ")
	fmt.Println("        * *  ")
	fmt.Println("       * * * ")
	fmt.Println(" * * * * * * * * * *  ")
	fmt.Println("    * * * * * * * ")
	fmt.Println("     * * * * * ")
	fmt.Println("    * *     * * ")
	fmt.Println("    *         * ")
	return ""
}
func AlgorithmPentagram() string {
	var y int = 10
	//i := "*"
	var line = strings.Repeat("-  ", y)
	for j := 0; j <= y; j++ {
		if j == y/2 {
			line := strings.ReplaceAll(line, "-  ", "*  ")
			fmt.Println(line)
			continue
		}
		if j < y/2 {
			stringlist := strings.Fields(line)
			//fmt.Printf("%q", stringlist)
			i := 0
			for _, num := range stringlist {
				if i == y/2+j && y/2+j < 10 {
					print("* ")
				} else if i == y/2-j && y/2-j >= 0 {
					print("* ")
				}
				//else if i== {

				//}
				fmt.Printf("%s ", num)
				i++

			}
			fmt.Println("")
			continue

		}
		fmt.Println(line)
	}
	return ""
}
