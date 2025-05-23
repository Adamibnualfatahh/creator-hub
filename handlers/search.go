package handlers

import (
	"creator-hub/models"
	"strings"
)

// SequentialSearchByJudul mencari pada array statis.
func SequentialSearchByJudul(daftar *[KAPASITAS_MAKS]models.Konten, n int, keyword string) ([KAPASITAS_MAKS]models.Konten, int) {
	var hasil [KAPASITAS_MAKS]models.Konten
	var jumlahHasil int
	keywordLower := strings.ToLower(keyword)

	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(daftar[i].JudulIde), keywordLower) {
			hasil[jumlahHasil] = daftar[i]
			jumlahHasil++
		}
	}
	return hasil, jumlahHasil
}

// BinarySearchByKategori mencari pada array statis dan tanpa 'break'.
func BinarySearchByKategori(daftar *[KAPASITAS_MAKS]models.Konten, n int, kategori string) ([KAPASITAS_MAKS]models.Konten, int) {
	// Membuat salinan data untuk diurutkan
	var sortedArr [KAPASITAS_MAKS]models.Konten
	copy(sortedArr[:], daftar[:n])
	SelectionSortByKategori(&sortedArr, n) // Mengurutkan salinan

	var hasil [KAPASITAS_MAKS]models.Konten
	var jumlahHasil int
	kategoriLower := strings.ToLower(kategori)

	kiri, kanan := 0, n-1
	indeksDitemukan := -1

	// Loop berhenti jika ditemukan (indeksDitemukan != -1) atau jika area pencarian habis.
	for kiri <= kanan && indeksDitemukan == -1 {
		tengah := kiri + (kanan-kiri)/2
		kategoriTengah := strings.ToLower(sortedArr[tengah].Kategori)

		if kategoriTengah == kategoriLower {
			indeksDitemukan = tengah // Ditemukan
		} else if kategoriTengah < kategoriLower {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if indeksDitemukan != -1 {
		hasil[jumlahHasil] = sortedArr[indeksDitemukan]
		jumlahHasil++
		// Cek ke kiri dan kanan dari titik yang ditemukan
		i := indeksDitemukan - 1
		for i >= 0 && strings.ToLower(sortedArr[i].Kategori) == kategoriLower {
			hasil[jumlahHasil] = sortedArr[i]
			jumlahHasil++
			i--
		}
		j := indeksDitemukan + 1
		for j < n && strings.ToLower(sortedArr[j].Kategori) == kategoriLower {
			hasil[jumlahHasil] = sortedArr[j]
			jumlahHasil++
			j++
		}
	}
	return hasil, jumlahHasil
}
