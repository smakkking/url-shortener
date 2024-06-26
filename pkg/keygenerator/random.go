package keygenerator

import (
	"math/rand"
	"time"
)

type RandomKeyGenerator struct {
}

type FixedKeyGenerator struct {
	Key string
}

func (f FixedKeyGenerator) GenRandomString(size int) string {
	return f.Key
}

// GenRandomString генерирует случайную строку из символов A-Za-z0-0-9_
func (r *RandomKeyGenerator) GenRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "____")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
