package helpers

import (
	"time"
)

func KonversiTanggalOutput(tglAwal string) (hasil string) {
	if tglAwal != "" {
		layoutAwal := time.RFC3339
		tglKonversi, _ := time.Parse(layoutAwal, tglAwal)

		layoutHasil := "2006-01-02 15:04:05"
		lokasi, _ := time.LoadLocation("Asia/Jakarta")
		hasil = tglKonversi.
			In(lokasi).
			Format(layoutHasil)
	}
	return
}

func KonversiTanggalDb(tglAwal string) (hasil time.Time) {
	if tglAwal != "" {
		layout := "2006-01-02 15:04:05"
		lokasi, _ := time.LoadLocation("Asia/Jakarta")
		hasil, _ = time.ParseInLocation(layout, tglAwal, lokasi)
	}
	return
}

func KonversiTanggalIssueOutput(tglAwal string) (hasil string) {
	if tglAwal != "" {
		layoutAwal := "2006-01-02T15:04:05.000-0700"
		tglKonversi, _ := time.Parse(layoutAwal, tglAwal)

		layoutHasil := "2006-01-02 15:04:05"
		lokasi, _ := time.LoadLocation("Asia/Jakarta")
		hasil = tglKonversi.
			In(lokasi).
			Format(layoutHasil)
	}
	return
}
