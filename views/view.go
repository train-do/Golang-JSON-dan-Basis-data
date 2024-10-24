package views

import (
	"app/models"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PrintMainMenu() {
	fmt.Println("1. Pesan")
	fmt.Println("2. Tambah Pesanan")
	fmt.Println("3. Cek Pesanan")
	fmt.Println("4. Bayar")
}
func PrintListMenu() {
	for _, v := range models.ListMenu {
		fmt.Printf("%d. %s\n", v.Id, v.Nama)
	}
}
func PrintAll(customer []models.Customer, conjunction []models.Conjunction) {
	fmt.Printf("%+v\n", customer)
	fmt.Printf("%+v\n", conjunction)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}