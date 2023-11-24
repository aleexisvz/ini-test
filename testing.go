package main

import (
	"fmt"
	"ini-test/models"
)

func main() {
	var t models.TarjetaPVC

	t.LoadValues()

	fmt.Println("Cantidad: ", t.Amount)
	fmt.Println("Monto: ", t.Price)
	fmt.Println("Monto Total: ", t.PriceTotal)
	fmt.Println("Tipo de Impresion: ", t.TypeImpression)
	fmt.Println("Variable Data: ", t.VariableData)
	fmt.Println("Variable Data Fields: ", t.VariableDataFields)
	fmt.Println("Variable Data Price: ", t.VariableDataPrice)
	fmt.Println("Variable Data Start: ", t.VariableDataStart)
	fmt.Println("Variable Photo: ", t.VariablePhoto)
	fmt.Println("Variable Photo Fields: ", t.VariablePhotoFields)
	fmt.Println("Variable Photo Price: ", t.VariablePhotoPrice)
	fmt.Println("Variable Photo Start: ", t.VariablePhotoStart)
	fmt.Println("Relief: ", t.Relief)
	fmt.Println("Relief Price: ", t.ReliefPrice)
	fmt.Println("Relief Start: ", t.ReliefStart)
}
