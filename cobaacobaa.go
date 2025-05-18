package main

func main () {

}
import "fmt"

type pengguna struct {
	nama string
	id   int
}

const NMAX int = 10

type tabPengguna [NMAX]pengguna
type tabkategori [4]string
type tabPengeluaran [4]int

func main() {
	var pax, budget int
	var arrayUser tabPengguna
	var arrayKateogri tabkategori
	var arrayPengeluaran tabPengeluaran

	arrayKateogri = tabkategori{"akomodasi", "transportasi", "makanan", "hiburan"}
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

func tambahPengeluaran(A tabPengguna, B *tabkategori, C *tabPengeluaran, n int) {
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

func editPengeluaran(pax int) {
	var id, pilihan, jumlah int
	fmt.Print("Masukkan ID pengguna: ")
	fmt.Scan(&id)
	if id < 0 || id >= pax {
		fmt.Println("ID tidak valid.")
		return
	}
	fmt.Println("Kategori yang ingin diedit: 1. Akomodasi 2. Transportasi 3. Makanan 4. Hiburan")
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&pilihan)
	fmt.Print("Masukkan nilai baru: ")
	fmt.Scan(&j)

	switch pilihan {
	case 1:
		akomodasi[id] = j // akomodasi[id] + j ( BISA DI BUAT KAYA GINI JUGA )
	case 2:
		transportasi[id] = j // transportasi[id] + j
	case 3:
		makanan[id] = j // makanan[id] = makanan[id] + j
	case 4:
		hiburan[id] = j // hiburan[id] = hiburan[id] + j
	default:
		fmt.Println("Kategori tidak valid.")
	}
}

func hapusPengeluaran(A tabPengguna, B *tabkategori, C *tabPengeluaran, n int) {
	var i, j, IDUser int
	var pilihan int

	fmt.Println()
	fmt.Println("Hapus Pengeluaran : ")
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

	(*C)[j] = 0
	fmt.Printf("Berhasil menghapus pengeluaran pada %s ", (*B)[j])
	fmt.Println(*C)

}
func lihatSemua(n int) {
	fmt.Println("╔════════╦════════════════╦════════════╦══════════════╦══════════╦══════════╗")
	fmt.Println("║   ID   ║      Nama      ║ Akomodasi  ║ Transportasi ║ Makanan  ║ Hiburan  ║")
	fmt.Println("╠════════╬════════════════╬════════════╬══════════════╬══════════╬══════════╣")

	for i := 0; i < n; i++ {
		fmt.Printf("║  %-4d ║ %-14s ║ %-10d ║ %-12d ║ %-8d ║ %-8d ║\n",
			i, pengguna[i], akomodasi[i], transportasi[i], makanan[i], hiburan[i])
	}

	fmt.Println("╚════════╩════════════════╩════════════╩══════════════╩══════════╩══════════╝")
}
