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