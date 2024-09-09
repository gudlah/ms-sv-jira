package helpers

import "math"

func Masking(teks string) string {
	maskingPercentage := float64(50)
	panjangTeks := float64(len(teks))
	panjangMasking := float64((panjangTeks * maskingPercentage) / 100)
	indeksAwalFloat := (panjangTeks - panjangMasking) / 2
	var indeksAwalInt int
	if indeksAwalFloat-float64(int(indeksAwalFloat)) > 0.5 {
		indeksAwalInt = int(math.Ceil(indeksAwalFloat))
	} else {
		indeksAwalInt = int(math.Floor(indeksAwalFloat))
	}
	indeksAkhir := indeksAwalInt + int(panjangMasking)

	rs := []rune(teks)
	for i := 0; i < len(rs); i++ {
		if i > (indeksAwalInt-1) && i < indeksAkhir {
			rs[i] = '*'
		}
	}
	return string(rs)
}
