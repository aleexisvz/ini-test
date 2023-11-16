package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

type TarjetaPVC struct {
	Amount         string
	Price          string
	TypeImpression string
}

func roundToNearest(t *TarjetaPVC, cfg *ini.File) {
	// save keys and value on keyMap
	keyMap := make(map[string]string)

	if t.TypeImpression == "SIMPLE FAZ" {
		for _, key := range cfg.Section("SIMPLEFAZ").Keys() {
			keyMap[key.Name()] = key.Value()
		}
	} else {
		for _, key := range cfg.Section("DOBLEFAZ").Keys() {
			keyMap[key.Name()] = key.Value()
		}
	}

	// get closest key
	var closestKey int
	amount, _ := strconv.Atoi(t.Amount)
	minDiff := math.MaxInt
	for keyStr := range keyMap {
		keyInt, _ := strconv.Atoi(keyStr)
		diff := int(math.Abs(float64(keyInt - amount)))
		if diff < minDiff {
			minDiff = diff
			closestKey = keyInt
		}
	}

	closestKeyStr := strconv.Itoa(closestKey)
	result, _ := strconv.Atoi(keyMap[closestKeyStr])

	t.Price = strconv.Itoa(result)
}

func calculatePrice(t *TarjetaPVC, cfg *ini.File) {
	// vars
	var price int
	var amount int

	// set amount
	fmt.Print("Enter amount: ")
	fmt.Scanln(&t.Amount)
	t.Amount = strings.TrimSpace(t.Amount)

	// get type of impression, simple or doble faz
	fmt.Print("Doble faz? (y/n): ")
	var dobleFaz string
	fmt.Scanln(&dobleFaz)

	if strings.ToLower(dobleFaz) == "y" {
		t.TypeImpression = "DOBLE FAZ"

		// round to nearest
		roundToNearest(t, cfg)

		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)

		price = price * amount
		t.Price = strconv.Itoa(price)
	} else {
		t.TypeImpression = "SIMPLE FAZ"

		// round to nearest
		roundToNearest(t, cfg)

		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)

		// calculate & set price
		price = price * amount
		t.Price = strconv.Itoa(price)
	}
}

func main() {
	var t TarjetaPVC

	// load config file
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	calculatePrice(&t, cfg)

	fmt.Printf("%su. Tarjetas PVC %s te salen $%s", t.Amount, t.TypeImpression, t.Price)
}
