package handlers

import (
	"creator-hub/models"
	"fmt"
	"strings"
)

// SelectionSortByTanggal mengurutkan berdasarkan tanggal dengan pilihan urutan.
func SelectionSortByTanggal(daftar *[KAPASITAS_MAKS]models.Konten, n int, urutan string) ([KAPASITAS_MAKS]models.Konten, int) {
	var sortedArr [KAPASITAS_MAKS]models.Konten
	copy(sortedArr[:], daftar[:n])

	for i := 0; i < n-1; i++ {
		indeksTukar := i
		for j := i + 1; j < n; j++ {
			// Logika perbandingan berdasarkan pilihan urutan
			var harusTukar bool
			if urutan == "asc" {
				// Ascending: cari tanggal paling awal (Before)
				harusTukar = sortedArr[j].TanggalPosting.Before(sortedArr[indeksTukar].TanggalPosting)
			} else { // desc
				// Descending: cari tanggal paling akhir (After)
				harusTukar = sortedArr[j].TanggalPosting.After(sortedArr[indeksTukar].TanggalPosting)
			}

			if harusTukar {
				indeksTukar = j
			}
		}
		sortedArr[i], sortedArr[indeksTukar] = sortedArr[indeksTukar], sortedArr[i]
	}
	return sortedArr, n
}

// InsertionSortByEngagement mengurutkan berdasarkan engagement dengan pilihan urutan.
func InsertionSortByEngagement(daftar *[KAPASITAS_MAKS]models.Konten, n int, urutan string) ([KAPASITAS_MAKS]models.Konten, int) {
	var sortedArr [KAPASITAS_MAKS]models.Konten
	copy(sortedArr[:], daftar[:n])

	for i := 1; i < n; i++ {
		kunci := sortedArr[i]
		j := i - 1

		// Logika perbandingan berdasarkan pilihan urutan
		var kondisiLoop bool
		if urutan == "asc" {
			// Ascending: engagement terkecil ke terbesar
			kondisiLoop = j >= 0 && sortedArr[j].TotalEngagement() > kunci.TotalEngagement()
		} else { // desc
			// Descending: engagement terbesar ke terkecil
			kondisiLoop = j >= 0 && sortedArr[j].TotalEngagement() < kunci.TotalEngagement()
		}

		for kondisiLoop {
			sortedArr[j+1] = sortedArr[j]
			j--
			// Update kondisi loop di dalam loop
			if urutan == "asc" {
				kondisiLoop = j >= 0 && sortedArr[j].TotalEngagement() > kunci.TotalEngagement()
			} else {
				kondisiLoop = j >= 0 && sortedArr[j].TotalEngagement() < kunci.TotalEngagement()
			}
		}
		sortedArr[j+1] = kunci
	}
	return sortedArr, n
}

// AnalisisKontenTerbaik menggunakan sorting descending.
func AnalisisKontenTerbaik(daftar *[KAPASITAS_MAKS]models.Konten, n int) {
	var kontenTerposting [KAPASITAS_MAKS]models.Konten
	var jumlahTerposting int

	for i := 0; i < n; i++ {
		if daftar[i].Status == "Sudah Posting" {
			kontenTerposting[jumlahTerposting] = daftar[i]
			jumlahTerposting++
		}
	}

	if jumlahTerposting == 0 {
		fmt.Println("Belum ada konten yang 'Sudah Posting' untuk dianalisis.")
		return
	}

	fmt.Println("🏆 Konten dengan Engagement Tertinggi:")
	diurutkan, jumlahUrut := InsertionSortByEngagement(&kontenTerposting, jumlahTerposting, "desc")
	LihatSemuaKonten(&diurutkan, jumlahUrut)
}

// SelectionSortByKategori adalah helper untuk Binary Search.
func SelectionSortByKategori(daftar *[KAPASITAS_MAKS]models.Konten, n int) {
	for i := 0; i < n-1; i++ {
		indeksMin := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(daftar[j].Kategori) < strings.ToLower(daftar[indeksMin].Kategori) {
				indeksMin = j
			}
		}
		daftar[i], daftar[indeksMin] = daftar[indeksMin], daftar[i]
	}
}
