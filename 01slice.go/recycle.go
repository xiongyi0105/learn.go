package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitalRegexp = regexp.MustCompile("[0-9]+")

func main() {
	mostdigital := FindDigitals("C:\\Users\\10194\\LightWAN\\cpeagent.log")
	fmt.Println(mostdigital)
}

func FindDigitals(filename string) []byte {
	stringarry, _ := ioutil.ReadFile(filename)
	want_slice := digitalRegexp.FindAll(stringarry, len(stringarry))
	release_for_arry_use := make([]byte, 0)
	for _, bytes := range want_slice {
		release_for_arry_use = append(release_for_arry_use, bytes...)
	}
	return release_for_arry_use

}
