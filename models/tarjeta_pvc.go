package models

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

type TarjetaPVC struct {
	Amount              int
	Price               float64
	PriceTotal          float64
	TypeImpression      string
	VariableData        bool
	VariableDataFields  int
	VariableDataPrice   float64
	VariableDataStart   float64
	VariablePhoto       bool
	VariablePhotoFields int
	VariablePhotoPrice  float64
	VariablePhotoStart  float64
}

func (t *TarjetaPVC) SetAmount() {
	// get amount
	fmt.Print("Ingrese la cantidad: ")
	fmt.Scanln(&t.Amount)
}

func (t *TarjetaPVC) SetTypeImpression() {
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

func (t *TarjetaPVC) SetVariableData() {
	// get variable data
	fmt.Print("Contiene campos variables? (y/n): ")
	var variableData string
	fmt.Scanln(&variableData)

	if strings.ToLower(variableData) == "y" {
		t.VariableData = true

		// get quantity of variable fields
		fmt.Print("Cuantos campos variables?: ")
		fmt.Scanln((&t.VariableDataFields))
	}
}

func (t *TarjetaPVC) SetVariablePhoto() {
	// get variable photo
	fmt.Print("Contiene foto variable? (y/n): ")
	var variablePhoto string
	fmt.Scanln(&variablePhoto)

	if strings.ToLower(variablePhoto) == "y" {
		t.VariablePhoto = true

		// get quantity of variable fields
		fmt.Print("Cuantas fotos variables?: ")
		fmt.Scanln((&t.VariablePhotoFields))
	}
}

func (t *TarjetaPVC) CalculateTotal() {
	// temp vars
	var variableData float64
	var variablePhoto float64

	// calculate total price
	if t.VariableData {
		variableData = (float64(t.VariableDataFields) * t.VariableDataPrice * float64(t.Amount)) + t.VariableDataStart
	}

	if t.VariablePhoto {
		variablePhoto = (float64(t.VariablePhotoFields) * t.VariablePhotoPrice * float64(t.Amount)) + t.VariablePhotoStart
	}

	t.PriceTotal = t.Price * float64(t.Amount)
	t.PriceTotal = t.PriceTotal + variableData + variablePhoto
}

func (t *TarjetaPVC) LoadValues() {
	// set the basic values
	t.SetAmount()
	t.SetTypeImpression()
	t.SetVariableData()
	t.SetVariablePhoto()

	// load config.ini file
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("Error loading config file")
		return
	}

	keyMap := make(map[int]float64)

	// ---------------- PRICE PER UNIT ----------------

	// load prices from config.ini on keyMap
	if t.TypeImpression == "SIMPLE FAZ" {
		for _, k := range cfg.Section("SIMPLEFAZ").Keys() {
			key, _ := strconv.Atoi(k.Name())
			value, _ := strconv.ParseFloat(k.Value(), 64)
			keyMap[key] = value
		}
	} else {
		for _, k := range cfg.Section("DOBLEFAZ").Keys() {
			key, _ := strconv.Atoi(k.Name())
			value, _ := strconv.ParseFloat(k.Value(), 64)
			keyMap[key] = value
		}
	}

	// use a for loop to get the closest key
	minDiff := math.MaxInt

	for key := range keyMap {
		diff := int(math.Abs(float64(key - t.Amount)))
		if diff < minDiff {
			minDiff = diff
			t.Price = keyMap[key]
		}
	}

	// ---------------- VARIABLE DATA ----------------

	if t.VariableData {
		// load prices from config.ini on keyMap
		for _, k := range cfg.Section("VARIABLEDATA").Keys() {
			key, _ := strconv.Atoi(k.Name())
			value, _ := strconv.ParseFloat(k.Value(), 64)
			keyMap[key] = value
		}

		// use a for loop to get the closest key
		minDiff := math.MaxInt

		for key := range keyMap {
			diff := int(math.Abs(float64(key - t.Amount)))
			if diff < minDiff {
				minDiff = diff
				t.VariableDataPrice = keyMap[key]
			}
		}

		// get variable data start cost
		t.VariableDataStart, _ = strconv.ParseFloat(cfg.Section("VARIABLEDATA").Key("costoInicio").Value(), 64)
	}

	// ---------------- VARIABLE PHOTO ----------------

	if t.VariablePhoto {
		// load prices from config.ini on keyMap
		for _, k := range cfg.Section("VARIABLEPHOTO").Keys() {
			key, _ := strconv.Atoi(k.Name())
			value, _ := strconv.ParseFloat(k.Value(), 64)
			keyMap[key] = value
		}

		// use a for loop to get the closest key
		minDiff := math.MaxInt

		for key := range keyMap {
			diff := int(math.Abs(float64(key - t.Amount)))
			if diff < minDiff {
				minDiff = diff
				t.VariablePhotoPrice = keyMap[key]
			}
		}

		// get variable data start cost
		t.VariablePhotoStart, _ = strconv.ParseFloat(cfg.Section("VARIABLEPHOTO").Key("costoInicio").Value(), 64)
	}

	t.CalculateTotal()
}
