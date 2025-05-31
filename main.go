package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NMAX int = 100

type konten struct {
	Ide        string
	Platform   string
	Tanggal    string
	Engagement int
}

type tabKonten [NMAX]konten

func main() {
	menu()
}

// Menampilkan menu user
func menu() {
	var keyword string
	var data tabKonten
	var sizeKonten int
	sizeKonten = 0
	for {
		fmt.Println("\n=== Aplikasi AI Pembuat Konten Sosial Media ===")
		fmt.Println("+----+-------------------------------------------------------------+")
		fmt.Println("| No | Menu                                                        |")
		fmt.Println("+----+-------------------------------------------------------------+")
		fmt.Println("| 1  | Tambah Konten                                               |")
		fmt.Println("| 2  | Ubah Konten                                                 |")
		fmt.Println("| 3  | Hapus Konten                                                |")
		fmt.Println("| 4  | Cari Konten yang mirip dengan keyword (Sequential Search)   |")
		fmt.Println("| 5  | Cari Konten menurut tanggal           (Binary Search)       |")
		fmt.Println("| 6  | Urutkan Konten berdasarkan Engagement (Selection Sort)      |")
		fmt.Println("| 7  | Urutkan Konten berdasarkan Tanggal    (Insertion Sort)      |")
		fmt.Println("| 8  | Rekomendasi Caption dan Hashtag                             |")
		fmt.Println("| 9  | Tampilkan semua konten                                      |")
		fmt.Println("| 0  | Keluar                                                      |")
		fmt.Println("+----+-------------------------------------------------------------+")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)
		// Membersihkan newline dari buffer
		var discard string
		fmt.Scanln(&discard)

		switch pilih {
		case 1:
			tambahKonten(&data, &sizeKonten)
		case 2:
			ubahKonten(&data, sizeKonten)
		case 3:
			deleteKonten(&data, &sizeKonten)
		case 4:
			fmt.Print("Masukkan kata kunci pencarian: ")
			fmt.Scanln(&keyword)
			cariKontenDenganKeywordSeqSearch(keyword, data, sizeKonten) //sequential search
		case 5:
			fmt.Print("Masukkan kata kunci tanggal: ")
			fmt.Scanln(&keyword)
			printKontenByIndex(keyword, data, sizeKonten) //Binary search
		case 6:
			selectionSortEngagement(&data, sizeKonten) // selection sort berdasarkan engagement
		case 7:
			insertionSortTanggal(&data, sizeKonten) // insertion sott berdasarkan tanggal
			tampilkanSemuaKonten(data, sizeKonten)
		case 8:
			generateRekomendasiCaptiondanHashtag(data, sizeKonten) // mengenerate caption dan hashtag berdasarkan ide dan platform konten
		case 9:
			tampilkanSemuaKonten(data, sizeKonten) // menampilkan semua konten
		case 0:
			fmt.Println("Terima kasih, program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menambahkan konten dari inputan user dengan tampilan lebih rapi
func tambahKonten(data *tabKonten, sizeKonten *int) {
	if *sizeKonten >= NMAX {
		fmt.Println("Gagal menambahkan: daftar konten sudah penuh!")
		return
	}

	var k konten
	var discard string

	fmt.Println("\n+--------------------------------------+")
	fmt.Println("|         Tambah Konten Baru           |")
	fmt.Println("+--------------------------------------+")

	fmt.Print("Masukkan Ide Konten (akhiri dengan -1):\n> ")
	k.Ide = inputKalimatSampaiMinusSatu()

	fmt.Print("Masukkan Platform:\n> ")
	fmt.Scan(&k.Platform)
	fmt.Scanln(&discard)

	fmt.Print("Masukkan Tanggal (YYYY-MM-DD):\n> ")
	fmt.Scan(&k.Tanggal)
	fmt.Scanln(&discard)

	fmt.Print("Masukkan Engagement:\n> ")
	fmt.Scan(&k.Engagement)
	fmt.Scanln(&discard)

	data[*sizeKonten] = k
	*sizeKonten++

	fmt.Println("\nKonten berhasil ditambahkan ke dalam daftar!")
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
func tampilkanSemuaKonten(data tabKonten, sizeKonten int) {
	if sizeKonten == 0 {
		fmt.Println("\nBelum ada konten yang ditambahkan.")
		return
	}

	fmt.Println("\nDaftar Semua Konten")
	fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
	fmt.Println("| No   | Ide                                                                                                  | Platform    | Tanggal    | Engagement  |")
	fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")

	for i := 0; i < sizeKonten; i++ {
		fmt.Printf("| %-4d | %-100s | %-11s | %-10s | %-11d |\n",
			i+1, data[i].Ide, data[i].Platform, data[i].Tanggal, data[i].Engagement)
	}

	fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
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
func cariKontenDenganKeywordSeqSearch(keyword string, data tabKonten, n int) {
	var hasil [NMAX]konten
	var i, jumlahHasil int
	jumlahHasil = 0

	for i = 0; i < n; i++ {
		if seqSearchKataPerKata(keyword, data[i].Ide) {
			if jumlahHasil < NMAX {
				hasil[jumlahHasil] = data[i]
				jumlahHasil++
			}
		}
	}

	// Menampilkan hasil pencarian
	if jumlahHasil > 0 {
		fmt.Println("\nKonten yang mengandung keyword:", keyword)
		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
		fmt.Println("| No   | Ide                                                                                                  | Platform    | Tanggal    | Engagement  |")
		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")

		for i = 0; i < jumlahHasil; i++ {
			fmt.Printf("| %-4d | %-100s | %-11s | %-10s | %-11d |\n",
				i+1, hasil[i].Ide, hasil[i].Platform, hasil[i].Tanggal, hasil[i].Engagement)
		}

		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
	} else {
		fmt.Printf("\nTidak ada konten yang mengandung keyword: \"%s\"\n", keyword)
	}
}
func cariKontenDenganTanggalBinarySearch(keyword string, data tabKonten, n int) int {
	insertionSortTanggal(&data, n)
	var left, right, mid int
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if data[mid].Tanggal == keyword {
			return mid
		} else if data[mid].Tanggal > keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func printKontenByIndex(keyword string, data tabKonten, n int) {
	var idx int
	idx = cariKontenDenganTanggalBinarySearch(keyword, data, n)
	if idx != -1 {
		fmt.Println("\nKonten yang mengandung tanggal:", keyword)
		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
		fmt.Println("| No   | Ide                                                                                                  | Platform    | Tanggal    | Engagement  |")
		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")

		fmt.Printf("| %-4d | %-100s | %-11s | %-10s | %-11d |\n",
			idx+1, data[idx].Ide, data[idx].Platform, data[idx].Tanggal, data[idx].Engagement)
		fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
	} else {
		fmt.Printf("\nTidak ada konten yang mengandung tanggal: \"%s\"\n", keyword)
	}
}

// mengurutkan konten berdasarkan engagement tertinggi(Desc)
func selectionSortEngagement(data *tabKonten, n int) {
	var pass, idx, i int
	var temp konten

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if data[idx].Engagement < data[i].Engagement {
				idx = i
			}
			i = i + 1
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass = pass + 1
	}

	fmt.Println("\nKonten berhasil diurutkan berdasarkan engagement dari tinggi ke rendah:")
	tampilkanSemuaKonten(*data, n)
}

// Mengurutkan konten berdasarkan tanggal terbaru konten(Desc)
func insertionSortTanggal(data *tabKonten, n int) {
	var pass, i int
	var temp konten

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.Tanggal > data[i-1].Tanggal {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}

	fmt.Println("\nKonten berhasil diurutkan berdasarkan tanggal dari tinggi ke rendah:")

}

// Mengenerate caption dan hashtag berdasarkan ide dan hashtag dari konten yang memiliki engagement tertinggi
func generateRekomendasiCaptiondanHashtag(data tabKonten, n int) {
	var idx, i int
	var ide, platform, caption, kata string
	idx = cariEngagementTertinggi(data, n)

	if idx == -1 {
		fmt.Println("Konten tidak ditemukan")
		return
	}

	ide = data[idx].Ide
	platform = data[idx].Platform

	rand.Seed(time.Now().UnixNano())

	switch platform {
	case "Instagram":
		switch rand.Intn(10) {
		case 0:
			caption = fmt.Sprintf("🔥 %s sekarang trending di Instagram! Yuk intip!", ide)
		case 1:
			caption = fmt.Sprintf("📸 Jangan lewatkan! %s hanya di Instagram!", ide)
		case 2:
			caption = fmt.Sprintf("Bagikan %s ke story kamu sekarang juga!", ide)
		case 3:
			caption = fmt.Sprintf("Temukan %s di Instagram hari ini!", ide)
		case 4:
			caption = fmt.Sprintf("Yuk cek %s dan jadikan hari kamu lebih berwarna!", ide)
		case 5:
			caption = fmt.Sprintf("Jangan sampai ketinggalan %s di Instagram!", ide)
		case 6:
			caption = fmt.Sprintf("Inilah alasan kenapa %s jadi favorit!", ide)
		case 7:
			caption = fmt.Sprintf("Bergabung bersama kami dengan %s sekarang!", ide)
		case 8:
			caption = fmt.Sprintf("Cerita seru tentang %s ada di Instagram!", ide)
		case 9:
			caption = fmt.Sprintf("Instagram kamu akan lebih menarik dengan %s!", ide)
		}
	case "Twitter":
		switch rand.Intn(10) {
		case 0:
			caption = fmt.Sprintf("%s sedang ramai diperbincangkan. Ikutan yuk! 🐦", ide)
		case 1:
			caption = fmt.Sprintf("Trending topik: %s. Apa pendapatmu?", ide)
		case 2:
			caption = fmt.Sprintf("Tweet ini tentang %s wajib kamu baca!", ide)
		case 3:
			caption = fmt.Sprintf("Berita terbaru: %s, langsung dari Twitter.", ide)
		case 4:
			caption = fmt.Sprintf("Jangan lewatkan pembahasan %s di Twitter!", ide)
		case 5:
			caption = fmt.Sprintf("Ikuti diskusi seru tentang %s!", ide)
		case 6:
			caption = fmt.Sprintf("%s bikin heboh Twitter hari ini!", ide)
		case 7:
			caption = fmt.Sprintf("Suka %s? Yuk, ngobrol di Twitter!", ide)
		case 8:
			caption = fmt.Sprintf("Baca tweet trending %s sekarang juga!", ide)
		case 9:
			caption = fmt.Sprintf("Twitter hangat dengan topik %s!", ide)
		}
	case "Facebook":
		switch rand.Intn(10) {
		case 0:
			caption = fmt.Sprintf("👍 Simak %s dan bagikan ke teman Facebook kamu!", ide)
		case 1:
			caption = fmt.Sprintf("%s bisa jadi topik menarik buat diskusi!", ide)
		case 2:
			caption = fmt.Sprintf("Sudah lihat %s di Facebook hari ini?", ide)
		case 3:
			caption = fmt.Sprintf("Bagikan %s supaya teman kamu juga tahu!", ide)
		case 4:
			caption = fmt.Sprintf("Yuk, komentar tentang %s di Facebook!", ide)
		case 5:
			caption = fmt.Sprintf("Postingan terbaru: %s, jangan sampai kelewatan!", ide)
		case 6:
			caption = fmt.Sprintf("Facebook makin seru dengan %s!", ide)
		case 7:
			caption = fmt.Sprintf("Mari diskusi %s bersama komunitas Facebook!", ide)
		case 8:
			caption = fmt.Sprintf("Temukan hal menarik tentang %s di Facebook.", ide)
		case 9:
			caption = fmt.Sprintf("Facebook update: %s, langsung dari kami!", ide)
		}
	default:
		switch rand.Intn(10) {
		case 0:
			caption = fmt.Sprintf("Check this out: %s!", ide)
		case 1:
			caption = fmt.Sprintf("%s adalah topik hangat hari ini!", ide)
		case 2:
			caption = fmt.Sprintf("Kamu wajib tahu tentang %s!", ide)
		case 3:
			caption = fmt.Sprintf("Ini dia %s yang sedang ramai dibicarakan.", ide)
		case 4:
			caption = fmt.Sprintf("Temukan fakta menarik tentang %s!", ide)
		case 5:
			caption = fmt.Sprintf("Baca selengkapnya tentang %s di sini.", ide)
		case 6:
			caption = fmt.Sprintf("Jangan lewatkan info terbaru tentang %s!", ide)
		case 7:
			caption = fmt.Sprintf("Yuk, kenali %s lebih dekat.", ide)
		case 8:
			caption = fmt.Sprintf("Berita hangat: %s!", ide)
		case 9:
			caption = fmt.Sprintf("Update terbaru mengenai %s, baca sekarang!", ide)
		}
	}

	fmt.Println("\nRekomendasi Caption:")
	fmt.Printf("   \"%s\"\n", caption)

	// Generate Hashtag dari kata-kata dalam ide
	fmt.Println("\nRekomendasi Hashtag:")
	for i = 0; i < len(ide); i++ {
		if ide[i] == ' ' {
			if kata != "" {
				fmt.Printf("#%s ", kata)
				kata = ""
			}
		} else {
			kata += string(ide[i])
		}
	}
	if kata != "" {
		fmt.Printf("#%s", kata)
	}
	fmt.Println()
}

// fungsi mencari konten dengan engagement tertinggi dengan logika nilai ekstrim
func cariEngagementTertinggi(data tabKonten, n int) int {
	var maxValue int
	var maxIndex, i int

	if n == 0 {
		return -1
	}

	maxValue = data[0].Engagement
	maxIndex = 0

	for i = 1; i < n; i++ {
		if data[i].Engagement > maxValue {
			maxValue = data[i].Engagement
			maxIndex = i
		}
	}
	return maxIndex
}

func deleteKonten(data *tabKonten, n *int) {
	fmt.Println("\n+--------------------------------------+")
	fmt.Println("|         Hapus Konten                 |")
	fmt.Println("+--------------------------------------+")
	if *n == 0 {
		fmt.Println("Belum ada konten yang dapat dihapus.")
		return
	}

	tampilkanSemuaKonten(*data, *n)

	var pilihan int
	fmt.Print("Masukkan nomor konten yang ingin dihapus (0 untuk batal): ")
	fmt.Scanln(&pilihan)

	if pilihan == 0 {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	if pilihan < 1 || pilihan > *n {
		fmt.Println("Nomor tidak valid.")
		return
	}
	var idx int
	idx = pilihan - 1

	// Geser elemen setelah idx ke kiri
	var i int
	for i = idx; i < *n-1; i++ {
		(*data)[i] = (*data)[i+1]
	}

	*n--
	fmt.Println("Konten berhasil dihapus.")
}

func ubahKonten(data *tabKonten, sizeKonten int) {
	fmt.Println("\n+--------------------------------------+")
	fmt.Println("|           Ubah Konten                |")
	fmt.Println("+--------------------------------------+")
	if sizeKonten == 0 {
		fmt.Println("Belum ada konten yang ditambahkan.")
		return
	} else {
		tampilkanSemuaKonten(*data, sizeKonten)
		fmt.Println("Pilihlah nomor konten yang ingin di ubah :")
		var index int
		fmt.Scan(&index)
		if index == 0 {
			fmt.Println("Batal mengubah konten")
			return
		} else if index < 1 || index > sizeKonten {
			fmt.Println("Nomor konten tidak valid")
			return
		} else {
			index = index - 1
			for {
				fmt.Println("\n--- Sub-Menu Ubah Konten ---")
				fmt.Println("1. Ubah Ide")
				fmt.Println("2. Ubah Platform")
				fmt.Println("3. Ubah Tanggal")
				fmt.Println("4. Ubah Engagement")
				fmt.Println("0. Kembali ke Menu Utama")
				fmt.Print("Pilihan: ")

				var pilihan int
				var discard string
				fmt.Scan(&pilihan)
				fmt.Scanln(&discard)

				switch pilihan {
				case 1:
					fmt.Print("Masukkan ide baru (akhiri dengan -1):\n> ")
					data[index].Ide = inputKalimatSampaiMinusSatu()
					fmt.Println("Ide berhasil diubah.")
				case 2:
					fmt.Print("Masukkan Platform baru :\n> ")
					fmt.Scan(&data[index].Platform)
					fmt.Scanln(&discard)
					fmt.Println("Platform berhasil diubah.")
				case 3:
					fmt.Print("Masukkan Tanggal baru (YYYY-MM-DD) :\n> ")
					fmt.Scan(&data[index].Tanggal)
					fmt.Scanln(&discard)
					fmt.Println("Tanggal berhasil diubah.")
				case 4:
					fmt.Print("Masukkan Engagement baru :\n> ")
					fmt.Scan(&data[index].Engagement)
					fmt.Scanln(&discard)
					fmt.Println("Engagement berhasil diubah.")
				case 0:
					fmt.Println("Kembali ke menu utama.")
					fmt.Println("\nDaftar Konten Terbaru")
					fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
					fmt.Println("| No   | Ide                                                                                                  | Platform    | Tanggal    | Engagement  |")
					fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")

					fmt.Printf("| %-4d | %-100s | %-11s | %-10s | %-11d |\n", index+1, data[index].Ide, data[index].Platform, data[index].Tanggal, data[index].Engagement)
					fmt.Println("+------+------------------------------------------------------------------------------------------------------+-------------+------------+-------------+")
					return
				default:
					fmt.Println("Pilihan tidak valid")
				}
			}
		}
	}
}
