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
		fmt.Println("8. Tampilkan semua konten")
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
			ubahKonten(&data, &sizeKonten)
		case 3:
			deleteKonten(&data, &sizeKonten)
		case 4:
			fmt.Print("Masukkan kata kunci pencarian: ")
			fmt.Scanln(&keyword)
			cariKontenDenganKeywordSeqSearch(keyword, &data, sizeKonten) //sequential search
		case 5:
			selectionSortEngagement(&data, sizeKonten) // selection sort berdasarkan engagement
		case 6:
			insertionSortTanggal(&data, sizeKonten) // insertion sott berdasarkan tanggal
		case 7:
			generateRekomendasiCaptiondanHashtag(data, sizeKonten) // mengenerate caption dan hashtag berdasarkan ide dan platform konten
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

// mengurutkan konten berdasarkan engagement tertinggi(Desc)
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

// Mengurutkan konten berdasarkan tanggal terbaru konten(Desc)
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

// Mengenerate caption dan hashtag berdasarkan ide dan hashtag dari konten yang memiliki engagement tertinggi
func generateRekomendasiCaptiondanHashtag(arr kontenArray, n int) {
	var idx, i int
	var ide, platform, caption, kata string
	idx = cariEngagementTertinggi(arr, n)

	if idx == -1 {
		fmt.Println("Konten tidak ditemukan")
		return
	}

	ide = arr[idx].Ide
	platform = arr[idx].Platform

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

	fmt.Println("Rekomendasi Caption :", caption)

	// Generate Hashtag dari kata-kata dalam ide
	fmt.Println("Rekomendasi Hashtag:")
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
func cariEngagementTertinggi(arr kontenArray, n int) int {
	var maxValue int
	var maxIndex, i int

	if n == 0 {
		return -1
	}

	maxValue = arr[0].Engagement
	maxIndex = 0

	for i = 1; i < n; i++ {
		if arr[i].Engagement > maxValue {
			maxValue = arr[i].Engagement
			maxIndex = i
		}
	}
	return maxIndex
}

func deleteKonten(data *kontenArray, size *int) {
	if *size == 0 {
		fmt.Println("Belum ada konten yang dapat dihapus.")
		return
	}

	tampilkanSemuaKonten()

	var pilihan int
	fmt.Print("Masukkan nomor konten yang ingin dihapus (0 untuk batal): ")
	fmt.Scanln(&pilihan)

	if pilihan == 0 {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	if pilihan < 1 || pilihan > *size {
		fmt.Println("Nomor tidak valid.")
		return
	}
	var idx int
	idx = pilihan - 1

	// Geser elemen setelah idx ke kiri
	for i := idx; i < *size-1; i++ {
		(*data)[i] = (*data)[i+1]
	}

	*size--
	fmt.Println("Konten berhasil dihapus.")
}

func ubahKonten(data *kontenArray, sizeKonten *int) {
	tampilkanSemuaKonten()

	if *sizeKonten == 0 {
		fmt.Println("Tidak ada konten yang dapat diubah")
		return
	}

	var noPilihan int
	fmt.Print("\nPilih nomor konten yang ingin diubah (0 untuk batal): ")
	fmt.Scanln(&noPilihan)

	if noPilihan < 0 || noPilihan > *sizeKonten {
		fmt.Println("Nomor tidak valid")
		return
	}

	if noPilihan == 0 {
		fmt.Println("Dibatalkan")
		return
	}
	var idx int
	idx = noPilihan - 1

	fmt.Println("\nMasukkan -1 untuk tidak mengubah nilai field (khusus integer)")
	fmt.Println("Biarkan kosong dan tekan Enter untuk tidak mengubah (khusus teks)")

	var input string

	// Ide Konten
	fmt.Printf("Ide Konten [%s]: ", (*data)[idx].Ide)
	input = inputKalimatSampaiMinusSatu()
	if input == "-1" {
		fmt.Println("Pengubahan dibatalkan")
		return
	}
	if input != "" {
		(*data)[idx].Ide = input
	}

	// Platform
	fmt.Printf("Platform [%s]: ", (*data)[idx].Platform)
	fmt.Scanln(&input)
	if input == "-1" {
		fmt.Println("Pengubahan dibatalkan")
		return
	}
	if input != "" {
		(*data)[idx].Platform = input
	}

	// Tanggal
	fmt.Printf("Tanggal [%s]: ", (*data)[idx].Tanggal)
	fmt.Scanln(&input)
	if input == "-1" {
		fmt.Println("Pengubahan dibatalkan")
		return
	}
	if input != "" {
		(*data)[idx].Tanggal = input
	}

	// Engagement
	fmt.Printf("Engagement [%d] (isi -1 untuk tidak mengubah): ", (*data)[idx].Engagement)
	var engagement int
	fmt.Scanln(&engagement)
	if engagement != -1 {
		(*data)[idx].Engagement = engagement
	}

	fmt.Println("Konten berhasil diperbarui!")
}
