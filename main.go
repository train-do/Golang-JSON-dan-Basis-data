package main

import (
	"app/controllers"
	"app/models"
	"app/views"
	"fmt"
)

func main() {
	id := 1
	var customers []models.Customer
	var conjunctions []models.Conjunction
	controllers.TarikDatabase(&customers, &conjunctions)
	appRuning:
	for {
		views.PrintMainMenu()
		fmt.Printf("%+v", conjunctions)
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			controllers.Pesan(&customers, &conjunctions, id)
			id++
		case "2":
			controllers.TambahPesan(&customers, &conjunctions)
		case "3":
			controllers.CekPesanan(customers, conjunctions)
		case "4":
			controllers.Bayar(&customers)
		case "0":
			controllers.UpdateDatabase(&customers, &conjunctions)
		case "exit":
			// views.PrintAll(customers, conjunctions)
			break appRuning
		}
		// ClearScreen()
	}
}
func init()  {
	views.ClearScreen()
}