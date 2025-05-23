package main

import (
	"bufio"
	"creator-hub/handlers"
	"creator-hub/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Menggunakan konstanta untuk ukuran array, sesuai spesifikasi.
const KAPASITAS_MAKS = 100

// Variabel global hanya untuk array utama.
var daftarKonten [KAPASITAS_MAKS]models.Konten

func main() {
	// Variabel state (jumlah data, ID berikutnya) dideklarasikan di main, bukan global.
	var jumlahKonten int = 0
	var nextID int = 1

	siapkanDataAwal(&nextID, &jumlahKonten)

	fmt.Println("========================================")
	fmt.Println("Selamat Datang di Creator Hub!")
	fmt.Println("========================================")

	var keluar bool = false
	for !keluar { // Loop utama tanpa 'break'.
		tampilkanMenuUtama()
		pilihan := bacaInputInt("Pilihan Anda: ")

		switch pilihan {
		case 1:
			menuTambahKonten(&jumlahKonten, &nextID)
		case 2:
			fmt.Println("\n--- Daftar Semua Konten ---")
			handlers.LihatSemuaKonten(&daftarKonten, jumlahKonten)
		case 3:
			menuUbahKonten(&jumlahKonten)
		case 4:
			menuHapusKonten(&jumlahKonten)
		case 5:
			menuCariKonten(jumlahKonten)
		case 6:
			menuUrutkanKonten(jumlahKonten)
		case 7:
			handlers.AnalisisKontenTerbaik(&daftarKonten, jumlahKonten)
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan Creator Hub!")
			keluar = true // Mengubah flag untuk keluar dari loop.
		default:
			fmt.Println("❌ Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

// --- Fungsi-fungsi untuk interaksi di main ---

func tampilkanMenuUtama() {
	fmt.Println("\n--- MENU UTAMA ---")
	fmt.Println("1. Tambah Ide Konten")
	fmt.Println("2. Lihat Semua Konten")
	fmt.Println("3. Ubah Detail Konten")
	fmt.Println("4. Hapus Konten")
	fmt.Println("5. Cari Konten...")
	fmt.Println("6. Urutkan Daftar Konten...")
	fmt.Println("7. Analisis Konten Terbaik")
	fmt.Println("0. Keluar")
}

func menuTambahKonten(jumlah *int, nextID *int) {
	fmt.Println("\n--- Menambah Ide Konten Baru ---")
	judul := bacaInputString("Masukkan Judul Ide: ")
	kategori := bacaInputString("Masukkan Kategori: ")
	platform := bacaInputString("Masukkan Platform: ")
	handlers.TambahKonten(&daftarKonten, jumlah, nextID, judul, kategori, platform)
}

func cariKontenByID(id int, n int) (int, bool) {
	var i int = 0
	var ditemukan bool = false
	for i < n && !ditemukan {
		if daftarKonten[i].ID == id {
			ditemukan = true
		} else {
			i++
		}
	}
	if ditemukan {
		return i, true
	}
	return -1, false
}

func menuUbahKonten(jumlah *int) {
	fmt.Println("\n--- Mengubah Detail Konten ---")
	id := bacaInputInt("Masukkan ID konten yang ingin diubah: ")
	indeks, ditemukan := cariKontenByID(id, *jumlah)

	if !ditemukan {
		fmt.Println("❌ Konten dengan ID tersebut tidak ditemukan.")
		return
	}

	kontenLama := daftarKonten[indeks]
	fmt.Printf("\nMengubah data untuk: '%s'\n", kontenLama.JudulIde)
	fmt.Println("Tip: Tekan Enter untuk melewati field yang tidak ingin diubah.")

	// 1. Meminta input baru. Prompt diubah untuk memberi tahu user bisa melewati.
	judulBaru := bacaInputString(fmt.Sprintf("Judul Baru (semula: %s): ", kontenLama.JudulIde))
	kategoriBaru := bacaInputString(fmt.Sprintf("Kategori Baru (semula: %s): ", kontenLama.Kategori))
	platformBaru := bacaInputString(fmt.Sprintf("Platform Baru (semula: %s): ", kontenLama.Platform))

	fmt.Println("Pilih Status Baru (1: Ide, 2: Terjadwal, 3: Sudah Posting): ")
	pilihanStatus := bacaInputInt(fmt.Sprintf("Pilihan Status (semula: %s): ", kontenLama.Status))

	var statusBaru string
	var tanggalBaru time.Time
	var likesBaru, komentarBaru, shareBaru int

	var statusValid bool = true
	switch pilihanStatus {
	case 1:
		statusBaru = "Ide"
		// Menggunakan nilai lama jika status tidak berubah
		tanggalBaru = kontenLama.TanggalPosting
		likesBaru = kontenLama.Interaksi.JumlahLike
		komentarBaru = kontenLama.Interaksi.JumlahKomentar
		shareBaru = kontenLama.Interaksi.JumlahShare
	case 2:
		statusBaru = "Terjadwal"
		fmt.Println("Masukkan Tanggal & Jam Posting (format: YYYY-MM-DD HH:MM): ")

		var inputTanggalValid bool = false
		for !inputTanggalValid {
			tglStr := bacaInputString(fmt.Sprintf("Tanggal (semula: %s): ", kontenLama.TanggalPosting.Format("2006-01-02 15:04")))
			if tglStr == "" { // Jika input tanggal kosong, pakai tanggal lama
				tanggalBaru = kontenLama.TanggalPosting
				inputTanggalValid = true
			} else {
				parsedTgl, err := time.Parse("2006-01-02 15:04", tglStr)
				if err == nil {
					tanggalBaru = parsedTgl
					inputTanggalValid = true
				} else {
					fmt.Println("❌ Format tanggal salah, coba lagi.")
				}
			}
		}
	case 3:
		statusBaru = "Sudah Posting"
		if kontenLama.TanggalPosting.IsZero() {
			tanggalBaru = time.Now()
		} else {
			tanggalBaru = kontenLama.TanggalPosting
		}

		fmt.Println("Masukkan Data Engagement:")
		likesBaru = bacaInputInt(fmt.Sprintf("Jumlah Like (semula: %d): ", kontenLama.Interaksi.JumlahLike))
		komentarBaru = bacaInputInt(fmt.Sprintf("Jumlah Komentar (semula: %d): ", kontenLama.Interaksi.JumlahKomentar))
		shareBaru = bacaInputInt(fmt.Sprintf("Jumlah Share (semula: %d): ", kontenLama.Interaksi.JumlahShare))
	default:
		fmt.Println("Pilihan status tidak valid, perubahan dibatalkan.")
		statusValid = false
	}

	// 3. Update data di array jika status yang dipilih valid
	if statusValid {
		// --- PERUBAHAN DI SINI ---
		// Update field hanya jika input dari pengguna tidak kosong.
		if judulBaru != "" {
			daftarKonten[indeks].JudulIde = judulBaru
		}
		if kategoriBaru != "" {
			daftarKonten[indeks].Kategori = kategoriBaru
		}
		if platformBaru != "" {
			daftarKonten[indeks].Platform = platformBaru
		}
		// --- AKHIR PERUBAHAN ---

		// Logika update untuk status dan field terkaitnya tetap berjalan
		daftarKonten[indeks].Status = statusBaru
		daftarKonten[indeks].TanggalPosting = tanggalBaru
		daftarKonten[indeks].Interaksi.JumlahLike = likesBaru
		daftarKonten[indeks].Interaksi.JumlahKomentar = komentarBaru
		daftarKonten[indeks].Interaksi.JumlahShare = shareBaru

		fmt.Println("\n✅ Detail konten berhasil diubah!")
	}
}

func menuHapusKonten(jumlah *int) {
	fmt.Println("\n--- Menghapus Konten ---")
	id := bacaInputInt("Masukkan ID konten yang akan dihapus: ")
	indeks, ditemukan := cariKontenByID(id, *jumlah)
	if !ditemukan {
		fmt.Println("❌ Konten dengan ID tersebut tidak ditemukan.")
		return
	}
	konfirmasi := bacaInputString(fmt.Sprintf("Yakin ingin menghapus '%s'? (y/n): ", daftarKonten[indeks].JudulIde))
	if strings.ToLower(konfirmasi) == "y" {
		handlers.HapusKonten(&daftarKonten, jumlah, indeks)
	} else {
		fmt.Println("\nPenghapusan dibatalkan.")
	}
}

func menuCariKonten(n int) {
	fmt.Println("\n--- Menu Pencarian Konten ---")
	fmt.Println("1. Cari berdasarkan Judul (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Kategori (Binary Search)")
	pilihan := bacaInputInt("Pilihan Anda: ")
	if pilihan == 1 {
		keyword := bacaInputString("Masukkan kata kunci judul: ")
		hasil, jHasil := handlers.SequentialSearchByJudul(&daftarKonten, n, keyword)
		fmt.Printf("\n🔎 Hasil Pencarian untuk '%s':\n", keyword)
		handlers.LihatSemuaKonten(&hasil, jHasil)
	} else if pilihan == 2 {
		kategori := bacaInputString("Masukkan nama kategori yang dicari: ")
		hasil, jHasil := handlers.BinarySearchByKategori(&daftarKonten, n, kategori)
		fmt.Printf("\n🔎 Hasil Pencarian untuk Kategori '%s':\n", kategori)
		handlers.LihatSemuaKonten(&hasil, jHasil)
	}
}

func menuUrutkanKonten(n int) {
	fmt.Println("\n--- Menu Pengurutan Konten ---")
	fmt.Println("1. Urutkan berdasarkan Tanggal Posting (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Engagement Tertinggi (Insertion Sort)")
	pilihanSort := bacaInputInt("Pilihan Anda: ")

	fmt.Println("Pilih Urutan (1: Ascending, 2: Descending):")
	pilihanUrutan := bacaInputInt("Pilihan Urutan: ")
	var urutan string
	if pilihanUrutan == 1 {
		urutan = "asc"
	} else {
		urutan = "desc"
	}

	var hasil [handlers.KAPASITAS_MAKS]models.Konten
	var jHasil int

	if pilihanSort == 1 {
		fmt.Printf("\n📊 Hasil Pengurutan berdasarkan Tanggal Posting (%s):\n", urutan)
		hasil, jHasil = handlers.SelectionSortByTanggal(&daftarKonten, n, urutan)
	} else if pilihanSort == 2 {
		fmt.Printf("\n🏆 Hasil Pengurutan berdasarkan Engagement (%s):\n", urutan)
		hasil, jHasil = handlers.InsertionSortByEngagement(&daftarKonten, n, urutan)
	}
	handlers.LihatSemuaKonten(&hasil, jHasil)
}

// --- Helper untuk input dan data awal ---

func bacaInputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func bacaInputInt(prompt string) int {
	var inputValid bool = false
	var inputInt int
	for !inputValid {
		inputStr := bacaInputString(prompt)
		parsedInt, err := strconv.Atoi(inputStr)
		if err == nil {
			inputInt = parsedInt
			inputValid = true
		} else {
			fmt.Println("❌ Input tidak valid, harap masukkan angka.")
		}
	}
	return inputInt
}

func siapkanDataAwal(nextID, jumlah *int) {
	// Menambahkan data awal ke array statis
	handlers.TambahKonten(&daftarKonten, jumlah, nextID, "Review Laptop Gaming Murah", "Teknologi", "YouTube")
	handlers.TambahKonten(&daftarKonten, jumlah, nextID, "Cara Membuat Kopi Susu", "Kuliner", "TikTok")
	handlers.TambahKonten(&daftarKonten, jumlah, nextID, "Belajar Dasar Go", "Edukasi", "YouTube")
	// Menambahkan data engagement dan tanggal secara manual untuk data awal
	daftarKonten[0].Status = "Sudah Posting"
	daftarKonten[0].TanggalPosting, _ = time.Parse("2006-01-02", "2024-05-20")
	daftarKonten[0].Interaksi = models.Engagement{JumlahLike: 150, JumlahKomentar: 30, JumlahShare: 15}

	daftarKonten[1].Status = "Sudah Posting"
	daftarKonten[1].TanggalPosting, _ = time.Parse("2006-01-02", "2024-04-15")
	daftarKonten[1].Interaksi = models.Engagement{JumlahLike: 500, JumlahKomentar: 80, JumlahShare: 120}
}
