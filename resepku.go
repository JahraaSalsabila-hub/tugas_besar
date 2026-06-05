package main

import "fmt"

const NMAX = 100

type Resep struct {
	nama      string
	kategori  string
	bahan     string
	durasi    int
	langkah   string
	jumlahCari int
}

var data [NMAX]Resep
var n int

func tambahResep() {
	if n >= NMAX {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\n=== Tambah Resep ===")

	fmt.Print("Nama Resep        : ")
	fmt.Scan(&data[n].nama)

	fmt.Print("Kategori          : ")
	fmt.Scan(&data[n].kategori)

	fmt.Print("Bahan Utama       : ")
	fmt.Scan(&data[n].bahan)

	fmt.Print("Durasi Memasak    : ")
	fmt.Scan(&data[n].durasi)

	fmt.Print("Langkah Memasak   : ")
	fmt.Scan(&data[n].langkah)

	data[n].jumlahCari = 0

	n++

	fmt.Println("Resep berhasil ditambahkan")
}

func tampilResep() {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data resep")
		return
	}

	fmt.Println("\n=== Daftar Resep ===")

	for i = 0; i < n; i++ {
		fmt.Println("---------------------------")
		fmt.Println("Nama Resep     :", data[i].nama)
		fmt.Println("Kategori       :", data[i].kategori)
		fmt.Println("Bahan Utama    :", data[i].bahan)
		fmt.Println("Durasi         :", data[i].durasi, "menit")
		fmt.Println("Langkah        :", data[i].langkah)
	}
}

func ubahResep() {
	var cari string
	var i int
	var ketemu bool

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Print("Masukkan nama resep yang ingin diubah : ")
	fmt.Scan(&cari)

	ketemu = false

	for i = 0; i < n; i++ {
		if data[i].nama == cari {
			ketemu = true

			fmt.Println("Masukkan data baru")

			fmt.Print("Nama Resep      : ")
			fmt.Scan(&data[i].nama)

			fmt.Print("Kategori        : ")
			fmt.Scan(&data[i].kategori)

			fmt.Print("Bahan Utama     : ")
			fmt.Scan(&data[i].bahan)

			fmt.Print("Durasi          : ")
			fmt.Scan(&data[i].durasi)

			fmt.Print("Langkah         : ")
			fmt.Scan(&data[i].langkah)

			fmt.Println("Data berhasil diubah")
		}
	}

	if !ketemu {
		fmt.Println("Resep tidak ditemukan")
	}
}

func hapusResep() {
	var cari string
	var i, j int
	var ketemu bool

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Print("Masukkan nama resep yang ingin dihapus : ")
	fmt.Scan(&cari)

	ketemu = false

	for i = 0; i < n; i++ {
		if data[i].nama == cari {
			ketemu = true

			for j = i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--

			fmt.Println("Data berhasil dihapus")
		}
	}

	if !ketemu {
		fmt.Println("Resep tidak ditemukan")
	}
}

func sequentialSearch() {
	var cari string
	var i int
	var ketemu bool

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Print("Masukkan bahan utama : ")
	fmt.Scan(&cari)

	ketemu = false

	fmt.Println("\nHasil pencarian:")

	for i = 0; i < n; i++ {
		if data[i].bahan == cari {
			ketemu = true
			data[i].jumlahCari++

			fmt.Println("---------------------------")
			fmt.Println("Nama Resep  :", data[i].nama)
			fmt.Println("Kategori    :", data[i].kategori)
			fmt.Println("Durasi      :", data[i].durasi, "menit")
		}
	}

	if !ketemu {
		fmt.Println("Resep tidak ditemukan")
	}
}

func insertionSortNama() {
	var i, j int
	var temp Resep

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].nama > temp.nama {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan nama")
}

func selectionSortDurasi() {
	var i, j, min int
	var temp Resep

	for i = 0; i < n-1; i++ {
		min = i

		for j = i + 1; j < n; j++ {
			if data[j].durasi < data[min].durasi {
				min = j
			}
		}

		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan durasi")
}

func binarySearch() {
	var cari string
	var left, right, mid int
	var ketemu bool

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	insertionSortNama()

	fmt.Print("Masukkan nama resep : ")
	fmt.Scan(&cari)

	left = 0
	right = n - 1
	ketemu = false

	for left <= right {
		mid = (left + right) / 2

		if data[mid].nama == cari {
			ketemu = true

			data[mid].jumlahCari++

			fmt.Println("\nData ditemukan")
			fmt.Println("Nama Resep   :", data[mid].nama)
			fmt.Println("Kategori     :", data[mid].kategori)
			fmt.Println("Bahan Utama  :", data[mid].bahan)
			fmt.Println("Durasi       :", data[mid].durasi, "menit")
			fmt.Println("Langkah      :", data[mid].langkah)

			break
		} else if data[mid].nama < cari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Resep tidak ditemukan")
	}
}

func statistik() {
	var i int
	var max int

	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Println("\n=== Statistik Resep ===")
	fmt.Println("Jumlah resep :", n)

	max = 0

	for i = 1; i < n; i++ {
		if data[i].jumlahCari > data[max].jumlahCari {
			max = i
		}
	}

	fmt.Println("Resep paling sering dicari :", data[max].nama)
	fmt.Println("Jumlah pencarian           :", data[max].jumlahCari)
}

func menu() {
	fmt.Println("\n===== APLIKASI RESEPKU =====")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Tampilkan Resep")
	fmt.Println("3. Ubah Resep")
	fmt.Println("4. Hapus Resep")
	fmt.Println("5. Sequential Search")
	fmt.Println("6. Binary Search")
	fmt.Println("7. Insertion Sort")
	fmt.Println("8. Selection Sort")
	fmt.Println("9. Statistik")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu : ")
}

func main() {
	var pilihan int

	for {
		menu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahResep()
		} else if pilihan == 2 {
			tampilResep()
		} else if pilihan == 3 {
			ubahResep()
		} else if pilihan == 4 {
			hapusResep()
		} else if pilihan == 5 {
			sequentialSearch()
		} else if pilihan == 6 {
			binarySearch()
		} else if pilihan == 7 {
			insertionSortNama()
		} else if pilihan == 8 {
			selectionSortDurasi()
		} else if pilihan == 9 {
			statistik()
		} else if pilihan == 0 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}
	}
}