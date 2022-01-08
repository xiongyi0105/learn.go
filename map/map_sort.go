package main

import (
	"fmt"
	"sort"
)

func main() {
	var deviceMap = make(map[string]string, 5)
	deviceMap["RF1_LW3001_lw3001"] = "10.5.5.10"
	deviceMap["RF1_LW2302_lw3001"] = "10.5.5.40"
	deviceMap["RF2_LW1305_lw3001"] = "10.5.5.70"
	deviceMap["RF3_LW2302_lw3001"] = "10.5.5.100"
	deviceMap["RF3_LW1301_lw3001"] = "10.5.5.130"
	var keySlice = make([]string, len(deviceMap))
	i := 0
	for key := range deviceMap {
		keySlice[i] = key
		i++
	}
	sort.Strings(keySlice)
	for _, key := range keySlice {
		fmt.Printf("The key is %s,The value is %v\n", key, deviceMap[key])
	}

}
