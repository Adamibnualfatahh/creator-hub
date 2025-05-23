package models

import "time"

// Engagement menampung data interaksi pada sebuah konten.
type Engagement struct {
	JumlahLike     int
	JumlahKomentar int
	JumlahShare    int
}

// Konten adalah struktur data utama untuk setiap ide konten.
type Konten struct {
	ID             int
	JudulIde       string
	Kategori       string
	Platform       string
	Status         string // "Ide", "Terjadwal", "Sudah Posting"
	TanggalPosting time.Time
	Interaksi      Engagement
}

// TotalEngagement adalah method untuk menghitung total interaksi.
func (k Konten) TotalEngagement() int {
	return k.Interaksi.JumlahLike + k.Interaksi.JumlahKomentar + k.Interaksi.JumlahShare
}
