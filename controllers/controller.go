package controllers

import (
	"app/models"
	"app/views"
	"fmt"
	"time"
)

func Pesan(customers *[]models.Customer, conjunctions *[]models.Conjunction, id int) {
	var input string
	views.ClearScreen()
	fmt.Print("Nama Customer : ")
	fmt.Scanln(&input)
	*customers = append(*customers, models.Customer{Id: id, Nama: input, Pembayaran: false})
	var menuId, kuantitas int
	for{
		views.ClearScreen()
		views.PrintListMenu()
		fmt.Print("Pilih Pesanan : ")
		fmt.Scanln(&menuId)
		fmt.Print("Kuantitas : ")
		fmt.Scanln(&kuantitas)
		con := models.Conjunction{CustomerId: id, MenuId: menuId, Kuantitas: kuantitas, Diproses: "", Diantar: "", Selesai: ""}
		go updateStatus(&con)
		*conjunctions = append(*conjunctions, con)
		fmt.Print("Akhiri Pesanan (y/n) : ")
		fmt.Scanln(&input)
		if input == "y" {
			break
		}
	}
	views.ClearScreen()
}
func TambahPesan(customers *[]models.Customer, conjunctions *[]models.Conjunction) {
	var input string
	views.ClearScreen()
	fmt.Print("Nama Customer : ")
	fmt.Scanln(&input)
	_, id, _ := findCustomerId(*customers, input)
	if id == 0 {
		views.ClearScreen()
		fmt.Println("Fitur Tambah Pesanan Tidak Tersedia, Pesan Terlebih Dahulu")
		return
	}
	var menuId, kuantitas int
	for{
		views.ClearScreen()
		views.PrintListMenu()
		fmt.Print("Pilih Pesanan : ")
		fmt.Scanln(&menuId)
		fmt.Print("Kuantitas : ")
		fmt.Scanln(&kuantitas)
		con := models.Conjunction{CustomerId: id, MenuId: menuId, Kuantitas: kuantitas, Diproses: "", Diantar: "", Selesai: ""}
		go updateStatus(&con)
		*conjunctions = append(*conjunctions, con)
		fmt.Print("Akhiri Pesanan? (y/n) : ")
		fmt.Scanln(&input)
		if input == "y" {
			break
		}
	}
	views.ClearScreen()
}
func CekPesanan(customers []models.Customer, conjunctions []models.Conjunction)  {
	var input string
	views.ClearScreen()
	fmt.Print("Nama Customer : ")
	fmt.Scanln(&input)
	_, id, c := findCustomerId(customers, input)
	if id == 0 {
		views.ClearScreen()
		fmt.Println("Fitur Cek Pesanan Tidak Tersedia, Pesan Terlebih Dahulu")
		return
	}
	views.ClearScreen()
	findAllOrderByCustomerId(conjunctions, c)
	time.Sleep(5*time.Second)
}
func Bayar(c *[]models.Customer)  {
	var input string
	views.ClearScreen()
	fmt.Print("Nama Customer : ")
	fmt.Scanln(&input)
	idx, id, _ := findCustomerId(*c, input)
	if id == 0 {
		views.ClearScreen()
		fmt.Println("Anda Belum Memesan Apapun")
		return
	}else {
		fmt.Print("Lakukan Pembayaran? (y/n) : ")
		fmt.Scanln(&input)
		if input == "y" {
			(*c)[idx].Pembayaran = true
			views.ClearScreen()
			fmt.Printf("Terima Kasih\033[32m %s \033[0mSudah Berkunjung", (*c)[idx].Nama)
		}
	}
	time.Sleep(5*time.Second)
}
func updateStatus(c *models.Conjunction)  {
	layout := "03:04:05 PM"
	var tProses, tAntar, tSelesai *time.Ticker
	tProses = time.NewTicker(5*time.Second)
	defer tProses.Stop()
    defer tAntar.Stop()
    defer tSelesai.Stop()
	select {
	case t:= <-tProses.C:
		println("tProses")
		c.Diproses = t.Format(layout)
		// tAntar = time.NewTicker(5*time.Second)
	}
	// select {
	// case t:= <-tAntar.C:
	// 	println("tAntar")
	// 	c.Diantar = t.Format(layout)
	// 	tSelesai = time.NewTicker(5*time.Second)
	// }
	// select {
	// case t:= <-tSelesai.C:
	// 	println("tSelesai")
	// 	c.Selesai = t.Format(layout)
	// }
}




func UpdateDatabase(c *[]models.Customer, con *[]models.Conjunction)  {
	encodeCustomer(c)
	encodeConjunction(con)
	encodeMenu()
	views.ClearScreen()
}
func TarikDatabase(c *[]models.Customer, con *[]models.Conjunction)  {
	decodeCustomer(c)
	decodeConjunction(con)
	views.ClearScreen()
}






func findAllOrderByCustomerId(con []models.Conjunction, c models.Customer)  {
	var history models.ShowCase
	history.Nama = c.Nama
	history.Pembayaran = c.Pembayaran
	for _, v := range con {
		if v.CustomerId == c.Id {
			history.Pesanan = append(history.Pesanan, models.Object{Nama: findNamaPesanan(v.MenuId), Kuantitas: v.Kuantitas, Diproses: v.Diproses, Diantar: v.Diantar, Selesai: v.Selesai})
		}
	}
	fmt.Println(string(encodeHistory(history)))
}
func findNamaPesanan(id int) string {
	for _, v := range models.ListMenu {
		if v.Id == id {
			return v.Nama
		}
	}
	return ""
}
func findCustomerId(c []models.Customer, input string) (int, int, models.Customer) {
	for i, v := range c {
		if v.Nama ==  input{
			return i, v.Id, v
		}
	}
	return -1, 0, models.Customer{}
}