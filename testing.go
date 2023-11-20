package main

import (
	"fmt"
	"ini-test/models"
)

func main() {
	var t models.TarjetaPVC

	t.LoadValues()

	fmt.Println(t)
}
