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
	Amount             string
	Price              string
	TypeImpression     string
	VariableData       bool
	VariableDataFields int
	VariableDataPrice  string
	VariableDataStart  string
}

func formatMoney(amount string) string {
	var result strings.Builder
	n := len(amount)
	for i, r := range amount {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(r)
	}

	return result.String()
}

func roundToNearest(t *TarjetaPVC, cfg *ini.File, toNearest int) string {
	// save keys and values on keyMap
	keyMap := make(map[string]string)

	switch toNearest {
	case 1:
		if t.TypeImpression == "SIMPLE FAZ" {
			for _, key := range cfg.Section("SIMPLEFAZ").Keys() {
				keyMap[key.Name()] = key.Value()
			}
		} else {
			for _, key := range cfg.Section("DOBLEFAZ").Keys() {
				keyMap[key.Name()] = key.Value()
			}
		}
	case 2:
		if t.VariableData {
			for _, key := range cfg.Section("VARIABLEDATA").Keys() {
				keyMap[key.Name()] = key.Value()
			}
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
	result := keyMap[closestKeyStr]

	return result
}

func setTypeImpression(t *TarjetaPVC) {
	// get type of impression, simple or doble faz
	fmt.Print("Doble faz? (y/n): ")
	var dobleFaz string
	fmt.Scanln(&dobleFaz)

	if strings.ToLower(dobleFaz) == "y" {
		t.TypeImpression = "DOBLE FAZ"
	} else {
		t.TypeImpression = "SIMPLE FAZ"
	}
}

func setAmount(t *TarjetaPVC) {
	// get amount
	fmt.Print("Enter amount: ")
	fmt.Scanln(&t.Amount)
	t.Amount = strings.TrimSpace(t.Amount)
}

func loadAttributes(t *TarjetaPVC, cfg *ini.File) {
	setAmount(t)
	setTypeImpression(t)
	setVariableData(t, cfg)
	calculatePrice(t, cfg)
}

func setVariableData(t *TarjetaPVC, cfg *ini.File) {
	// get variable data
	fmt.Print("Contiene campos variables? (y/n): ")
	var variableData string
	fmt.Scanln(&variableData)

	if strings.ToLower(variableData) == "y" {
		t.VariableData = true

		// get quantity of variable fields
		fmt.Print("Cuantos campos variables?: ")
		fmt.Scanln(&t.VariableDataFields)
	} else {
		t.VariableData = false
	}

	if t.VariableData {
		// get variable price
		t.VariableDataPrice = roundToNearest(t, cfg, 2)
		fmt.Printf("Precio por campo variable: %s\n", t.VariableDataPrice)
		t.VariableDataStart = cfg.Section("VARIABLEDATA").Key("costoInicio").String()
		fmt.Printf("Costo de inicio: %s\n", t.VariableDataStart)
	}
}

func calculatePrice(t *TarjetaPVC, cfg *ini.File) {
	// temp vars
	var price int
	var amount int

	if t.TypeImpression == "DOBLE FAZ" {
		// round to nearest
		t.Price = roundToNearest(t, cfg, 1)

		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)

		price *= amount
		t.Price = strconv.Itoa(price)
	} else {
		// round to nearest
		t.Price = roundToNearest(t, cfg, 1)

		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)

		// calculate & set price
		price *= amount
		t.Price = strconv.Itoa(price)
	}

	if t.VariableData {
		// temp vars
		costoInicio, _ := strconv.Atoi(t.VariableDataStart)

		// calculate & set price
		variableDataPrice, _ := strconv.Atoi(t.VariableDataPrice)
		variableDataPrice = (variableDataPrice * t.VariableDataFields * amount) + costoInicio
		t.Price = strconv.Itoa(price + variableDataPrice)
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

	loadAttributes(&t, cfg)

	// format
	t.Price = formatMoney(t.Price)
	t.Amount = formatMoney(t.Amount)

	fmt.Printf("%su. Tarjetas PVC %s te salen $%s", t.Amount, t.TypeImpression, t.Price)
}
