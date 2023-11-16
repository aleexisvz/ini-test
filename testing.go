package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

func testing() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// save keys and value on keyMap
	keyMap := make(map[string]string)
	for _, key := range cfg.Section("SIMPLEFAZ").Keys() {
		keyMap[key.Name()] = key.Value()
	}

	var value int
	value = 1565

	var closestKey int
	minDiff := math.MaxInt
	for keyStr := range keyMap {
		keyInt, _ := strconv.Atoi(keyStr)
		diff := int(math.Abs(float64(keyInt - value)))
		if diff < minDiff {
			minDiff = diff
			closestKey = keyInt
		}
	}

	closestKeyStr := strconv.Itoa(closestKey)
	result, _ := strconv.Atoi(keyMap[closestKeyStr])

	fmt.Printf("%du. Tarjetas PVC SIMPLE FAZ te salen $%d", value, result)
}

func main() {
	testing()
}
