package keygenerator

import (
	"math/rand"
	"time"

	"github.com/smakkking/url-shortener/internal/models"
)

// GenRandomString генерирует случайную строку из символов A-Za-z0-0-9_
func GenRandomString(size int) models.URLKey {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "_")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return models.URLKey(b)
}
