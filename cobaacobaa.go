package main 

import "fmt"

type pengguna struct {
	nama string
	id int
}
const NMAX int = 10
type tabPengguna [NMAX] pengguna
type tabkategori [4] string
type tabPengeluaran [4] int

func main() {
	var pax, budget int
	var arrayUser tabPengguna
	var arrayKateogri tabkategori
	var arrayPengeluaran tabPengeluaran

	arrayKateogri = tabkategori {"akomodasi", "transportasi", "makanan", "hiburan"}
	fmt.Printf("Masukkan Jumlah Pengguna : ")
	fmt.Scan(&pax)
	fmt.Printf("Masukkan Budget Perjalanan : ")
	fmt.Scan(&budget)
	inputData(&arrayUser, &pax)
	tambahPengeluaran(arrayUser, &arrayKateogri, &arrayPengeluaran, pax)

}

func inputData (A *tabPengguna, n *int) {
	var i int

	for i = 0; i <= *n-1; i++{
		fmt.Printf("Masukkan Nama Pengguna : ")
		fmt.Scan(&A[i].nama)
		fmt.Printf("Masukkan ID Pengguna : ")
		fmt.Scan(&A[i].id)
	}
}

func tambahPengeluaran(A tabPengguna, B *tabkategori, C *tabPengeluaran, n int){
	var i, j, IDUser int
	var pilihan int

	fmt.Println()
	fmt.Println("Tambah Pengeluaran : ")
	fmt.Printf("Masukkan ID : ")
	fmt.Scan(&IDUser)
	fmt.Printf("Masukkan Kategori : (1. Akomodasi, 2. Transportasi, 3. Makanan, 4. Hiburan) : ")
	fmt.Scan(&pilihan)

	i = 0

	for i < n && IDUser != A[i].id {
		i++
	}

	j = pilihan - 1

	if j >= 0 && j < 4 && i < n {
		fmt.Printf("Berhasil mencatat untuk %s di kategori %s\n", A[i].nama, (*B)[j])
	} else {
		fmt.Println("Input tidak valid!")
	}

	fmt.Printf("Masukkan pengeluaran pada kategori %s : ", (*B)[j])
	fmt.Scan(&(*C)[j])
	fmt.Println(*C)
	
}
