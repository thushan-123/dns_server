package encodeDecode

import (
	"strings"
)

func EncodeDomain(domain string) []byte {
	parts := strings.Split(domain, ".")
	var encoded []byte

	for _, p := range parts {
		encoded = append(encoded, byte(len(p)))
		encoded = append(encoded, []byte(p)...)
	}

	return append(encoded, 0)
}
