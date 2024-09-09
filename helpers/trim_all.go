package helpers

import "strings"

func TrimAll(teks string) string {
	return strings.
		Join(strings.Fields(teks), " ")
}
