package main

import (
	"fmt"
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

func calculatePrice(t *TarjetaPVC) {
	// vars
	var price int
	var amount int

	// load config file
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

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

		// set doblefaz price
		t.Price = cfg.Section("DOBLEFAZ").Key(t.Amount + "UNI").String()
		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)
		price = price * amount
		t.Price = strconv.Itoa(price)
	} else {
		t.TypeImpression = "SIMPLE FAZ"

		// set simplefaz price
		t.Price = cfg.Section("SIMPLEFAZ").Key(t.Amount + "UNI").String()
		price, _ = strconv.Atoi(t.Price)
		amount, _ = strconv.Atoi(t.Amount)
		price = price * amount
		t.Price = strconv.Itoa(price)
	}
}

func main() {
	var t TarjetaPVC
	calculatePrice(&t)

	fmt.Printf("%su. Tarjetas PVC %s te salen $%s", t.Amount, t.TypeImpression, t.Price)
}
