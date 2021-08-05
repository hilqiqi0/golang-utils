package strs

import (
	"math/rand"
	"time"
)

// Shuffle
func ShuffleStrList(s []string) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
