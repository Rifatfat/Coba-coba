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
		case 6:
			urutkanPengeluaran(&arrayUser, &arrayPengeluaran, pax)
		case 7:
			laporanAkhir(arrayUser, arrayPengeluaran, pax, budget)
		}
		if opsi == 8 {
			fmt.Println("terimaksih ya!")
			break
		}

	}

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
	fmt.Printf("Masukkan Kategori : (1. Akomodasi, 2. Transportasi, 3. Makanan, 4. Hiburan) : ") // buat rapih
	fmt.Scan(&pilihan)
	//fmt.Printf("Masukkan Jumlah Pengeluaran yang mau di tambah : ")
	//fmt.Scan(&jumlah)
	for i = 0; i < n; i++ {
		if A[i].id == IDUser {
			switch pilihan {
			case 1:
				fmt.Println("Pilih Hotel:") // buat jadi rapih kotak kotak
				fmt.Println("1. Hotel Bintang 3 (Rp300.000/malam)")
				fmt.Println("2. Hotel Bintang 4 (Rp500.000/malam)")
				fmt.Println("3. Hotel Bintang 5 (Rp800.000/malam)")
				var hotelPilihan int
				var harga int
				fmt.Println("Masukan Pilihan : ")
				fmt.Scan(&hotelPilihan)

				switch hotelPilihan {
				case 1:
					harga = 300000
				case 2:
					harga = 500000
				case 3:
					harga = 800000
				default:
					fmt.Println("Pilihan hotel tidak valid.")
					return
				}

				fmt.Printf("Masukkan jumlah malam menginap: ")
				var malam int
				fmt.Scan(&malam)

				total := harga * malam
				C[i].akomodasi += total
				kategori = "akomodasi (Hotel Bintang " + fmt.Sprint(hotelPilihan) + ")"

				//C[i].akomodasi += jumlah
				//kategori = "akomodasi"
			case 2:
				var hargaTiket int
				var negara string
				// disini print negara tujuan dan harga nya (buat jadi rapih kotak kotak)
				fmt.Scan(&negara)
				switch negara {
				case "singapore", "Singapore":
					hargaTiket = 1000000
				case "jepang", "Jepang":
					hargaTiket = 7000000
				case "korea", "Korea":
					hargaTiket = 6500000
				case "thailand", "Thailand":
					hargaTiket = 3500000
				default:
					fmt.Println("Negara belum terdaftar, masukkan biaya manual.")
					fmt.Printf("Masukkan biaya tiket ke %s: ", negara)
					fmt.Scan(&hargaTiket)
				}

				C[i].transportasi += hargaTiket
				kategori = "transportasi ke " + negara
				//C[i].transportasi += jumlah
				//kategori = "transportasi"
			case 3:
				fmt.Println("Masukan jumlah : ")
				fmt.Scan(&jumlah)
				C[i].makanan += jumlah
				kategori = "makanan"
			case 4:
				fmt.Println("Masukan jumlah : ")
				fmt.Scan(&jumlah)
				C[i].hiburan += jumlah
				kategori = "hiburan"
			default:
				fmt.Println("Kategori tidak valid!")
				return
			}
			fmt.Printf("Berhasil mencatat untuk %s di kategori %s\n", A[i].nama, kategori) // kalo bisa dia tau saldo yang di input di budget
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}
func editPengeluaran(A tabPengguna, B tabPengeluaran, n int) {

	var pilihan, jumlah, IDUser, i int
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

//		fmt.Println("Kategori yang ingin diedit: 1. Akomodasi 2. Transportasi 3. Makanan 4. Hiburan")
//		fmt.Print("Masukkan kategori: ")
//		fmt.Scan(&pilihan)
//		fmt.Print("Masukkan nilai baru: ")
//		fmt.Scan(&jumlah)
//
//		switch pilihan {
//		case 1:
//			B[pilihan-1].akomodasi = jumlah
//		case 2:
//			B[pilihan-1].transportasi = jumlah
//		case 3:
//			B[pilihan-1].makanan = jumlah
//		case 4:
//			B[pilihan-1].hiburan = jumlah
//		default:
//			fmt.Println("Kategori tidak valid.")
//		}
//	}
func hapusPengeluaran(A tabPengguna, C *tabPengeluaran, n int) {
	var i, IDUser int // j int
	var pilihan int
	// var found bool = false
	var kategori string

	fmt.Println()
	fmt.Println("Hapus Pengeluaran ") // ubah tampilan jadi kotak kotak dan rapih
	fmt.Printf("Masukkan ID : ")
	fmt.Scan(&IDUser)
	fmt.Printf("Masukkan Kategori : (1. Akomodasi, 2. Transportasi, 3. Makanan, 4. Hiburan) : ") // ubah jadi kotak kotak
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
	fmt.Println("╔════════╦════════════════╦════════════╦══════════════╦══════════╦══════════╗") // rekomendasi itu mau di taro di mana ?
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
func cariPengeluaranSeq(A tabPengguna, B tabPengeluaran, n int) { // kalo pengeuarannya sama bagaimana ?
	var i int
	var maxakomodasi, maxtransportasi, maxmakanan, maxhiburan int
	//namaA := ""
	//maxPengeluaran := 0

	for i = 0; i < n-1; i++ {
		if B[i].akomodasi > B[i+1].akomodasi {
			maxakomodasi = B[i].akomodasi
			//namaA = A[i].nama // Simpan nama yang sesuai dengan nilai maks
		} else {
			maxakomodasi = B[i+1].akomodasi
			//namaA = A[i].nama
		}
	}

	i = 0

	for maxakomodasi != B[i].akomodasi {
		i++
	}

	fmt.Printf("Pengeluaran terbanyak pada kategori akomodasi sebesar %d oleh %s\n", maxakomodasi, A[i].nama)
	//----------------------------
	for i = 0; i < n-1; i++ {
		if B[i].transportasi > B[i+1].transportasi {
			maxtransportasi = B[i].transportasi
			//namaA = A[i].nama
		} else {
			maxtransportasi = B[i+1].transportasi
			//namaA = A[i].nama
		}
	}

	i = 0

	for maxtransportasi != B[i].transportasi {
		i++
	}

	fmt.Printf("Pengeluaran terbanyak pada kategori transportasi sebesar %d oleh %s\n", maxtransportasi, A[i].nama)
	//------------------------------------
	for i = 0; i < n-1; i++ {
		if B[i].makanan > B[i+1].makanan {
			maxmakanan = B[i].makanan
			//namaA = A[i].nama
		} else {
			maxmakanan = B[i+1].makanan
			//namaA = A[i].nama
		}
	}

	i = 0

	for maxmakanan != B[i].makanan {
		i++
	}

	fmt.Printf("Pengeluaran terbanyak pada kategori makanan sebesar %d oleh %s\n", maxmakanan, A[i].nama)
	//---------------------------------------
	for i = 0; i < n-1; i++ {
		if B[i].hiburan > B[i+1].hiburan {
			maxhiburan = B[i].hiburan
			//namaA = A[i].nama
		} else {
			maxhiburan = B[i+1].hiburan
			//namaA = A[i].nama
		}
	}

	i = 0

	for maxhiburan != B[i].hiburan {
		i++
	}

	fmt.Printf("Pengeluaran terbanyak pada hiburan akomodasi sebesar %d oleh %s\n", maxhiburan, A[i].nama)
}
func urutkanPengeluaran(daftarPengguna *tabPengguna, daftarPengeluaran *tabPengeluaran, jumlahPengguna int) {
	// binarry search
	var i, j int
	var penggunaSementara pengguna
	var pengeluaranSementara Pengeluaran

	for i = 1; i < jumlahPengguna; i++ {
		j = i
		for j > 0 {
			totalSekarang := daftarPengeluaran[j].akomodasi + daftarPengeluaran[j].transportasi + daftarPengeluaran[j].makanan + daftarPengeluaran[j].hiburan
			totalSebelumnya := daftarPengeluaran[j-1].akomodasi + daftarPengeluaran[j-1].transportasi + daftarPengeluaran[j-1].makanan + daftarPengeluaran[j-1].hiburan

			if totalSekarang < totalSebelumnya {
				penggunaSementara = daftarPengguna[j]
				daftarPengguna[j] = daftarPengguna[j-1]
				daftarPengguna[j-1] = penggunaSementara

				pengeluaranSementara = daftarPengeluaran[j]
				daftarPengeluaran[j] = daftarPengeluaran[j-1]
				daftarPengeluaran[j-1] = pengeluaranSementara
			}
			j--
		}
	}
	fmt.Println("\n╔════╦════════════════╦════════════════════╗")
	fmt.Println("║ ID ║ Nama           ║ Total Pengeluaran  ║")
	fmt.Println("╠════╬════════════════╬════════════════════╣")
	for i = 0; i < jumlahPengguna; i++ {
		total := daftarPengeluaran[i].akomodasi + daftarPengeluaran[i].transportasi + daftarPengeluaran[i].makanan + daftarPengeluaran[i].hiburan
		fmt.Printf("║ %-2d ║ %-14s ║ %-18d ║\n", daftarPengguna[i].id, daftarPengguna[i].nama, total)
	}
	fmt.Println("╚════╩════════════════╩════════════════════╝")
}
