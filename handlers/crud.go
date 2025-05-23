package handlers

import (
	"creator-hub/models"
	"fmt"
)

// KAPASITAS_MAKS adalah ukuran tetap untuk array statis.
const KAPASITAS_MAKS = 100

// TambahKonten menambahkan konten baru ke array statis.
// Menggunakan pointer ke jumlah agar bisa mengubah nilainya secara langsung.
func TambahKonten(daftar *[KAPASITAS_MAKS]models.Konten, jumlah *int, nextID *int, judul, kategori, platform string) {
	if *jumlah >= KAPASITAS_MAKS {
		fmt.Println("❌ Gagal: Kapasitas penyimpanan konten sudah penuh.")
		return
	}

	kontenBaru := models.Konten{
		ID:       *nextID,
		JudulIde: judul,
		Kategori: kategori,
		Platform: platform,
		Status:   "Ide",
	}

	// Menambahkan data di indeks terakhir yang tersedia.
	daftar[*jumlah] = kontenBaru
	*jumlah++ // Menambah jumlah data yang tersimpan.
	*nextID++

	fmt.Println("\n✅ Ide konten berhasil ditambahkan!")
}

// LihatSemuaKonten menampilkan data dari array statis sebanyak n elemen.
func LihatSemuaKonten(daftar *[KAPASITAS_MAKS]models.Konten, n int) {
	if n == 0 {
		fmt.Println("Tidak ada konten untuk ditampilkan.")
		return
	}

	for i := 0; i < n; i++ {
		k := daftar[i] // Mengakses elemen array
		fmt.Printf("----------------------------------------\n")
		fmt.Printf("ID             : %d\n", k.ID)
		fmt.Printf("Judul          : %s\n", k.JudulIde)
		// ... (sisa output sama seperti sebelumnya) ...
		fmt.Printf("Kategori       : %s\n", k.Kategori)
		fmt.Printf("Platform       : %s\n", k.Platform)
		fmt.Printf("Status         : %s\n", k.Status)
		if k.Status != "Ide" && !k.TanggalPosting.IsZero() {
			fmt.Printf("Tanggal Posting: %s\n", k.TanggalPosting.Format("02 Jan 2006 15:04"))
		}
		if k.Status == "Sudah Posting" {
			fmt.Printf("Engagement     : Likes(%d), Komentar(%d), Share(%d) -> Total: %d\n",
				k.Interaksi.JumlahLike, k.Interaksi.JumlahKomentar, k.Interaksi.JumlahShare, k.TotalEngagement())
		}
	}
	fmt.Printf("----------------------------------------\n")
}

// HapusKonten menghapus konten dari array statis dengan menggeser elemen.
func HapusKonten(daftar *[KAPASITAS_MAKS]models.Konten, jumlah *int, indeks int) {
	if indeks < 0 || indeks >= *jumlah {
		fmt.Println("❌ Gagal: Indeks tidak valid.")
		return
	}

	// Menggeser semua elemen setelah 'indeks' ke kiri.
	for i := indeks; i < *jumlah-1; i++ {
		daftar[i] = daftar[i+1]
	}

	*jumlah-- // Mengurangi jumlah data yang tersimpan.

	// Optional: Membersihkan elemen terakhir yang sekarang menjadi duplikat.
	daftar[*jumlah] = models.Konten{}

	fmt.Println("\n✅ Konten berhasil dihapus.")
}
