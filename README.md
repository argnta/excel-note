# CLI Password Manager (Excel-Based)

## Deskripsi
Program ini adalah aplikasi sederhana berbasis Command Line Interface (CLI) menggunakan bahasa Go untuk menyimpan dan mencari password aplikasi.

Penyimpanan data menggunakan file Excel (.xlsx) sebagai media latihan file handling dan struktur program dasar.  
Program ini dibuat untuk tujuan pembelajaran dan latihan awal.

---

## Fitur Utama
1. Menu interaktif berbasis CLI
2. Menyimpan nama aplikasi dan password
3. Mencari password berdasarkan nama aplikasi
4. Penyimpanan data menggunakan file Excel
5. Struktur program modular menggunakan package helper

---

## Struktur Program
.
├── helper
│ ├── add.go
│ ├── select.go
│ └── helper.go
├── Pass.xlsx
├── main.go
├── go.mod
├── go.sum
└── README.md

---

## Alur Program
1. Program dijalankan melalui terminal
2. User memilih menu yang tersedia
3. User menambahkan atau mencari password
4. Data disimpan dan dibaca dari file Excel

---

## Menu Aplikasi
CLI Pass Manager
Pilih menu
1 : Buat app password baru
2 : Cari Password
3 : Keluar

---

## Penyimpanan Data
- File: Pass.xlsx
- Kolom A: Nama aplikasi
- Kolom B: Password
- Data baru ditambahkan ke baris terakhir secara otomatis

---

## Cara Menjalankan
1. Pastikan Go sudah terinstall
2. Pastikan file Pass.xlsx tersedia
3. Jalankan perintah:

---

## Catatan Teknis
- Password disimpan dalam bentuk plain text
- Tidak menggunakan hashing atau encryption
- Program ini tidak ditujukan untuk penggunaan production

---

## Batasan Program
- Tidak ada autentikasi user
- Tidak ada enkripsi password
- Tidak cocok untuk data sensitif

---

## Pengembangan Lanjutan
- Menambahkan enkripsi password
- Menambahkan fitur update dan delete password
- Mengganti Excel dengan database

---

Project ini dibuat sebagai latihan awal untuk memahami dasar CLI application, input/output, dan file handling menggunakan Go.
