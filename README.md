# Aplikasi AI Pembuat Konten Sosial Media

Aplikasi berbasis bahasa pemrograman Go (Golang) untuk pengelolaan konten media sosial sederhana. Aplikasi ini memungkinkan pengguna untuk menambahkan, mengubah, menghapus, mencari, mengurutkan, dan merekomendasikan konten berdasarkan tingkat engagement.

## Fitur

- Menambah konten baru dengan ide, platform, tanggal, dan engagement
- Mencari konten berdasarkan keyword dengan metode Sequential Search
- Mengurutkan konten berdasarkan engagement menggunakan Selection Sort
- Mengurutkan konten berdasarkan tanggal menggunakan Insertion Sort
- Menampilkan semua konten yang sudah tersimpan

## Struktur Data

- Menggunakan tipe data struct `konten` dengan atribut:
  - Ide (string)
  - Platform (string)
  - Caption (string)
  - Hashtag (string)
  - Tanggal (string, format YYYY-MM-DD)
  - Engagement (int)
- Data disimpan dalam array statis dengan kapasitas maksimum 100 elemen.

## Cara Menjalankan

1. Pastikan sudah menginstal Go di komputer Anda.
2. Simpan kode program dalam file `main.go`.
3. Jalankan perintah berikut di terminal:

   go run main.go

4. Ikuti menu yang muncul untuk mengoperasikan aplikasi.

## Contoh Input

- Menambahkan konten:
  - Ide: "Promo diskon lebaran -1"
  - Platform: Instagram
  - Tanggal: 2025-05-20
  - Engagement: 1500

## Lisensi

Program ini disediakan tanpa lisensi tertentu. Pengguna dapat menggunakannya untuk tujuan belajar dan pengembangan.

---

Jika ingin saya buatkan file `.md` ini siap untuk kamu unduh, beri tahu saja!
