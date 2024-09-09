package helpers

import "math/rand"

func GenerateRandomNumeric(panjang int) string {
	layout := "1234567890"
	panjangLayout := len(layout)
	hasil := make([]byte, panjang)
	for i := range hasil {
		randomInt := rand.Intn(panjangLayout)
		hasil[i] = layout[randomInt]
	}
	return string(hasil)
}
