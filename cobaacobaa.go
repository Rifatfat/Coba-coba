package main

import "fmt"

type pengguna struct {
	nama string
	id   int
}
type Pengeluaran struct {
	akomodasi    int
	transportasi int
	makanan      int
	hiburan      int
}

const PAX int = 10

type tabPengeluaran [PAX]Pengeluaran
type tabPengguna [PAX]pengguna

func main() {
	var pax, budget, opsi int
	var arrayUser tabPengguna
	var arrayPengeluaran tabPengeluaran

	fmt.Printf("Masukkan Jumlah Pengguna : ")
	fmt.Scan(&pax)
	fmt.Printf("Masukkan Budget Perjalanan : ")
	fmt.Scan(&budget)
	inputData(&arrayUser, &pax)
	for {
		tampilanMenu()
		fmt.Scan(&opsi)

		switch opsi {
		case 1:
			tambahPengeluaran(arrayUser, &arrayPengeluaran, pax)
		case 2:
			editPengeluaran(arrayUser, arrayPengeluaran, pax)

		case 3:
			hapusPengeluaran(arrayUser, &arrayPengeluaran, pax)
		case 4:
			lihatSemua(arrayUser, arrayPengeluaran, pax)
		case 5:
			cariPengeluaranSeq(arrayUser, arrayPengeluaran, pax)
		case 7:
			laporanAkhir(arrayUser, arrayPengeluaran, pax, budget)
		}
		if opsi == 8 {
			fmt.Println("terimaksih ya!")
			break
		}

	}

}
