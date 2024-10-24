package models

type Customer struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Pembayaran bool   `json:"sudahBayar"`
}
type Conjunction struct {
	CustomerId int    `json:"customerId"`
	MenuId     int    `json:"menuId"`
	Kuantitas  int    `json:"qty"`
	Diproses   string `json:"waktuProses"`
	Diantar    string `json:"waktuAntar"`
	Selesai    string `json:"waktuSelesai"`
}
type Menu struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}
type ShowCase struct {
	Nama       string   `json:"nama"`
	Pembayaran bool     `json:"sudahBayar"`
	Pesanan    []Object `json:"listPesanan"`
}
type Object struct {
	Nama      string `json:"nama"`
	Kuantitas int    `json:"qty"`
	Diproses  string `json:"waktuProses,omitempty"`
	Diantar   string `json:"waktuAntar,omitempty"`
	Selesai   string `json:"waktuSelesai,omitempty"`
}

var ListMenu = []Menu{{1, "Minuman"}, {2, "Makanan"}}