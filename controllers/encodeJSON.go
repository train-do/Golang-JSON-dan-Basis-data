package controllers

import (
	"app/models"
	"encoding/json"
	"fmt"
	"os"
)

func encodeHistory(data models.ShowCase) []byte {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}
	return dataJSON
}

func encodeMenu() {
	file, err := os.Create("normalisasi/Menu.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	Menus := models.ListMenu

	if err := encoder.Encode(&Menus); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
func encodeCustomer(customers *[]models.Customer) {
	file, err := os.Create("normalisasi/Customer.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(customers); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
func encodeConjunction(conjunctions *[]models.Conjunction) {
	file, err := os.Create("normalisasi/Conjunction.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(conjunctions); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func decodeCustomer(customers *[]models.Customer) {
	file, err := os.Open("normalisasi/Customer.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&customers); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}
func decodeConjunction(conjunctions *[]models.Conjunction) {
	file, err := os.Open("normalisasi/Conjunction.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conjunctions); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}