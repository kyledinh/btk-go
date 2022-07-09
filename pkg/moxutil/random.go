package moxutil

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandomHyphenedKeyByLimit(limit int) string {
	terms := []string{"ANT", "BOO", "CAT", "DIP", "ENT", "FRO", "GIN", "HOP", "INK",
		"JAM", "KIP", "LAP", "MAP", "NET", "OAK"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(terms), func(i, j int) { terms[i], terms[j] = terms[j], terms[i] })

	if limit < 0 || limit > 10 {
		limit = 4
	}
	// sample return: "GIN-NET-BOO-OAK"
	return strings.Join(terms[:limit], "-")
}
