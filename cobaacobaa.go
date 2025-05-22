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
	tambahPengeluaran(arrayUser, &arrayKateogri, &arrayPengeluaran, pax)

}

func inputData(A *tabPengguna, n *int) {
	var i int

	for i = 0; i <= *n-1; i++ {
		fmt.Printf("Masukkan Nama Pengguna : ")
		fmt.Scan(&A[i].nama)
		fmt.Printf("Masukkan ID Pengguna : ")
		fmt.Scan(&A[i].id)
	}
}
func tampilanMenu() {

	fmt.Println("█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█")
	fmt.Println("█                      G O   B U D G E T                      █")
	fmt.Println("█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█")
	fmt.Println("             SELAMAT DATANG DI APLIKASI GO BUDGET             ")
	fmt.Println("               Aplikasi Pencatatan Keuangan Anda              ")
	fmt.Println("▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀")
	fmt.Println("█                                                             █")
	fmt.Println("█    ➤  1. Tambah Pengeluaran                                █")
	fmt.Println("█    ➤  2. Ubah Pengeluaran                                  █")
	fmt.Println("█    ➤  3. Hapus Pengeluaran                                 █")
	fmt.Println("█    ➤  4. Lihat Semua Pengeluaran                           █")
	fmt.Println("█    ➤  5. Cari Pengeluaran                                  █")
	fmt.Println("█    ➤  6. Urutkan Pengeluaran                               █")
	fmt.Println("█    ➤  7. Laporan Akhir                                     █")
	fmt.Println("█    ➤  8. Exit                                              █")
	fmt.Println("█                                                             █")
	fmt.Println("█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█")
	fmt.Println(" Tekan angka (1-8) lalu ENTER untuk memulai ")
}
func tambahPengeluaran(A tabPengguna, C *tabPengeluaran, n int) {
	var i, IDUser, jumlah int
	var pilihan int
	var kategori string

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

func editPengeluaran(pax int) {
	var id, pilihan, jumlah int
	fmt.Print("Masukkan ID pengguna: ")
	fmt.Scan(&IDUser) // unik id
	i = 0
	for IDUser == A[i].id {
		fmt.Println("Kategori yang ingin diedit: 1. Akomodasi 2. Transportasi 3. Makanan 4. Hiburan")
		fmt.Print("Masukkan kategori: ")
		fmt.Scan(&pilihan)
		fmt.Print("Masukkan nilai baru: ")
		fmt.Scan(&jumlah)

		switch pilihan {
		case 1:
			B[pilihan-1].akomodasi = jumlah
		case 2:
			B[pilihan-1].transportasi = jumlah
		case 3:
			B[pilihan-1].makanan = jumlah
		case 4:
			B[pilihan-1].hiburan = jumlah
		default:
			fmt.Println("Kategori tidak valid.")
		}
		i++
	}
}

// KURANG INI PENAMAAN VARIABEL KURANG MIX N MATCH WKWKWKKWWK
