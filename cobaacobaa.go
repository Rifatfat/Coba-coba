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
	fmt.Printf("Masukkan Jumlah Pengeluaran yang mau di tambah : ")
	fmt.Scan(&jumlah)
	for i = 0; i < n; i++ {
		if A[i].id == IDUser {
			switch pilihan {
			case 1:
				C[i].akomodasi += jumlah
				kategori = "akomodasi"
			case 2:
				C[i].transportasi += jumlah
				kategori = "transportasi"
			case 3:
				C[i].makanan += jumlah
				kategori = "makanan"
			case 4:
				C[i].hiburan += jumlah
				kategori = "hiburan"
			default:
				fmt.Println("Kategori tidak valid!")
				return
			}
			fmt.Printf("Berhasil mencatat untuk %s di kategori %s\n", A[i].nama, kategori)
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}
func editPengeluaran(A tabPengguna, B tabPengeluaran, n int) {

	var id, pilihan, jumlah, IDUser, i int
	fmt.Print("Masukkan ID pengguna: ")
	fmt.Scan(&id)

	for i < n && IDUser != A[i].id {
		i++
	}

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
}
func hapusPengeluaran(A tabPengguna, C *tabPengeluaran, n int) {
	var i, IDUser int // j int
	var pilihan int
	// var found bool = false
	var kategori string

	fmt.Println()
	fmt.Println("Hapus Pengeluaran : ")
	fmt.Printf("Masukkan ID : ")
	fmt.Scan(&IDUser)
	fmt.Printf("Masukkan Kategori : (1. Akomodasi, 2. Transportasi, 3. Makanan, 4. Hiburan) : ")
	fmt.Scan(&pilihan)
	for i = 0; i < n; i++ {
		if A[i].id == IDUser {
			switch pilihan {
			case 1:
				C[i].akomodasi = 0
				kategori = "akomodasi"
			case 2:
				C[i].transportasi = 0
				kategori = "transportasi"
			case 3:
				C[i].makanan = 0
				kategori = "makanan"
			case 4:
				C[i].hiburan = 0
				kategori = "hiburan"
			default:
				fmt.Println("Kategori tidak valid!")
				return
			}
			fmt.Printf("Berhasil menghapus untuk %s di kategori %s\n", A[i].nama, kategori)
			return
		}
	}
}
func lihatSemua(A tabPengguna, B tabPengeluaran, n int) {
	fmt.Println("╔════════╦════════════════╦════════════╦══════════════╦══════════╦══════════╗")
	fmt.Println("║   ID   ║      Nama      ║ Akomodasi  ║ Transportasi ║ Makanan  ║ Hiburan  ║")
	fmt.Println("╠════════╬════════════════╬════════════╬══════════════╬══════════╬══════════╣")

	for i := 0; i < n; i++ {
		fmt.Printf("║  %-4d ║ %-14s ║ %-10d ║ %-12d ║ %-8d ║ %-8d ║\n",
			A[i].id, A[i].nama, B[i].akomodasi, B[i].transportasi, B[i].makanan, B[i].hiburan)
	}

	fmt.Println("╚════════╩════════════════╩════════════╩══════════════╩══════════╩══════════╝")
}
func laporanAkhir(A tabPengguna, B tabPengeluaran, pax int, budget int) {
	fmt.Println("\n═══════════════════════════════════════════════════════════════════════")
	fmt.Println("                            LAPORAN AKHIR")
	fmt.Println("═══════════════════════════════════════════════════════════════════════")
	fmt.Printf("╔════╦════════════════╦════════════════════╦════════════╦════════════╗\n")
	fmt.Printf("║ ID ║ Nama           ║ Total Pengeluaran  ║ Budget     ║ Selisih    ║\n")
	fmt.Printf("╠════╬════════════════╬════════════════════╬════════════╬════════════╣\n")

	for i := 0; i < pax; i++ {
		total := B[i].akomodasi + B[i].transportasi + B[i].makanan + B[i].hiburan
		selisih := budget - total
		fmt.Printf("║ %-2d ║ %-14s ║ %-18d ║ %-10d ║ %-10d ║\n",
			A[i].id, A[i].nama, total, budget, selisih)
	}

	fmt.Printf("╚════╩════════════════╩════════════════════╩════════════╩════════════╝\n")
}
func cariPengeluaranSeq(A tabPengguna, B tabPengeluaran, n int) {
	var i int
	var maxakomodasi, maxtransportasi, maxmakanan, maxhiburan int
	namaA := ""
	maxPengeluaran := 0

	for i = 0; i < n; i++ {
		if B[i].akomodasi > maxPengeluaran {
			maxakomodasi = B[i].akomodasi
			namaA = A[i].nama // Simpan nama yang sesuai dengan nilai maks
		}
	}
	fmt.Printf("Pengeluaran terbanyak pada kategori akomodasi sebesar %d oleh %s\n", maxakomodasi, namaA)

	for i = 0; i < n; i++ {
		if B[i].transportasi > maxPengeluaran {
			maxtransportasi = B[i].transportasi
			namaA = A[i].nama
		}
	}
	fmt.Printf("Pengeluaran terbanyak pada kategori transportasi sebesar %d oleh %s\n", maxtransportasi, namaA)
	for i = 0; i < n; i++ {
		if B[i].makanan > maxPengeluaran {
			maxmakanan = B[i].makanan
			namaA = A[i].nama
		}
	}
	fmt.Printf("Pengeluaran terbanyak pada kategori makanan sebesar %d oleh %s\n", maxmakanan, namaA)

	for i = 0; i < n; i++ {
		if B[i].hiburan > maxPengeluaran {
			maxhiburan = B[i].hiburan
			namaA = A[i].nama
		}
	}
	fmt.Printf("Pengeluaran terbanyak pada kategori hiburan sebesar %d oleh %s\n", maxhiburan, namaA)
}
