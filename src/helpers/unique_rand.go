package helpers

import (
	"math/rand"
	"time"
)

type uniqueRand struct {
	generated map[int]bool
}

func NewUniqueRand() *uniqueRand {
	return &uniqueRand{make(map[int]bool)}
}

func (u *uniqueRand) Int(bounds int) int {
	source := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(source)
	for {
		i := newRand.Intn(bounds)
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
