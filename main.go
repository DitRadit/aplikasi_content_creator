package main

import (
	"fmt"
)

const NMAX int = 100

type konten struct {
	Ide        string
	Platform   string
	Caption    string
	Hashtag    string
	Tanggal    string
	Engagement int
}

type kontenArray [NMAX]konten

var data kontenArray

var sizeKonten int = 0

func main() {
	menu()
}

// Menampilkan menu user
func menu() {
	var keyword string

	for {
		fmt.Println("\n=== Aplikasi AI Pembuat Konten Sosial Media ===")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Ubah Konten")
		fmt.Println("3. Hapus Konten")
		fmt.Println("4. Cari Konten dengan keyword (Sequential Search)")
		fmt.Println("5. Urutkan Konten berdasarkan Engagement (Selection Sort)")
		fmt.Println("6. Urutkan Konten berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("7. Rekomendasi Caption dan Hashtag")
		fmt.Println("8. Tampilkan Konten dengan Engagement Tertinggi pada Periode")
		fmt.Println("9. Tampilkan semua konten")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)
		// Membersihkan newline dari buffer
		var discard string
		fmt.Scanln(&discard)

		switch pilih {
		case 1:
			tambahKonten()
		case 2:
			// ubahKonten()
		case 3:
			// hapusKonten()
		case 4:
			fmt.Print("Masukkan kata kunci pencarian: ")
			fmt.Scanln(&keyword)
			cariKontenDenganKeywordSeqSearch(keyword, &data, sizeKonten) //sequential search
		case 5:
			selectionSortEngagement(&data, sizeKonten) // selection sort berdasarkan engagement
		case 6:
			insertionSortTanggal(&data, sizeKonten) // insertion sott berdasarkan tanggal
		case 7:
			//generateRekomendasiCaptiondanHashtag(data, sizeKonten) // mengenerate caption dan hashtag berdasarkan ide dan platform konten
		case 8:
			tampilkanSemuaKonten() // menampilkan semua konten
		case 0:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menambahkan konten dari inputan user
func tambahKonten() {
	var k konten
	var discard string
	fmt.Print("Masukkan Ide Konten (akhiri dengan -1): ")
	k.Ide = inputKalimatSampaiMinusSatu()

	fmt.Print("Masukkan Platform: ")
	fmt.Scan(&k.Platform)
	// Membersihkan newline dari buffer(karena ga bisa pake bufio)
	fmt.Scanln(&discard)

	fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&k.Tanggal)
	fmt.Scanln(&discard)

	fmt.Print("Masukkan Engagement: ")
	fmt.Scan(&k.Engagement)
	fmt.Scanln(&discard)

	if sizeKonten < NMAX {
		data[sizeKonten] = k
		sizeKonten++
		fmt.Println("Konten berhasil ditambahkan.")
	} else {
		fmt.Println("Daftar konten penuh!")
	}
}

// untuk memberikan nilai sebuah kalimat ke suatu variable, menggunakan for loop, kondisi kalimat selesai/berhenti adalah ketika user menginputkan -1
func inputKalimatSampaiMinusSatu() string {
	var kata, kalimat string
	fmt.Scan(&kata)
	for kata != "-1" {
		if kalimat == "" {
			kalimat = kata
		} else {
			kalimat += " " + kata
		}
		fmt.Scan(&kata)
	}
	return kalimat
}

// Menampilkan semua konten
func tampilkanSemuaKonten() {
	var i int
	if sizeKonten == 0 {
		fmt.Println("Belum ada konten yang ditambahkan.")
		return
	}

	fmt.Println("\n--- Daftar Semua Konten ---")
	for i = 0; i < sizeKonten; i++ {
		fmt.Printf("Konten #%d:\n", i+1)
		fmt.Printf("  Ide       : %s\n", data[i].Ide)
		fmt.Printf("  Platform  : %s\n", data[i].Platform)
		fmt.Printf("  Tanggal   : %s\n", data[i].Tanggal)
		fmt.Printf("  Engagement: %d\n", data[i].Engagement)
		fmt.Println()
	}
}

// Sequential search kata per kata
func seqSearchKataPerKata(kata string, text string) bool {
	var i int
	for i = 0; i <= len(text)-len(kata); i++ {
		var j int
		j = 0
		for j < len(kata) && text[i+j] == kata[j] {
			j++
		}
		if j == len(kata) {
			return true
		}
	}
	return false
}

// mencari konten dengan keyword dengan bantuan seqsearch kata per kata
func cariKontenDenganKeywordSeqSearch(keyword string, arr *kontenArray, n int) {
	var hasil [NMAX]konten
	var i, jumlahHasil int
	jumlahHasil = 0

	for i = 0; i < n; i++ {
		if seqSearchKataPerKata(keyword, arr[i].Ide) {
			if jumlahHasil < NMAX {
				hasil[jumlahHasil] = arr[i]
				jumlahHasil++
			}
		}
	}

	if jumlahHasil > 0 {
		fmt.Println("Konten yang mengandung keyword:")
		for i = 0; i < jumlahHasil; i++ {
			fmt.Printf("  Ide       : %s\n", hasil[i].Ide)
			fmt.Printf("  Platform  : %s\n", hasil[i].Platform)
			fmt.Printf("  Tanggal   : %s\n", hasil[i].Tanggal)
			fmt.Printf("  Engagement: %d\n", hasil[i].Engagement)
			fmt.Println()
		}
	} else {
		fmt.Println("Tidak ada konten yang sesuai dengan keyword.")
	}
}

// mengurutkan konten berdasarkan engagement tertinggi
func selectionSortEngagement(arr *kontenArray, n int) {
	var pass, idx, i int
	var temp konten

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if arr[idx].Engagement < arr[i].Engagement {
				idx = i
			}
			i = i + 1
		}
		temp = arr[pass-1]
		arr[pass-1] = arr[idx]
		arr[idx] = temp
		pass = pass + 1
	}

	fmt.Println("\nKonten berhasil diurutkan berdasarkan engagement dari tinggi ke rendah:")
	tampilkanSemuaKonten()
}

// Mengurutkan konten berdasarkan tanggal terbaru konten
func insertionSortTanggal(arr *kontenArray, n int) {
	var pass, i int
	var temp konten

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = arr[pass]
		for i > 0 && temp.Tanggal > arr[i-1].Tanggal {
			arr[i] = arr[i-1]
			i = i - 1
		}
		arr[i] = temp
		pass = pass + 1
	}

	fmt.Println("\nKonten berhasil diurutkan berdasarkan tanggal dari tinggi ke rendah:")
	tampilkanSemuaKonten()

}
