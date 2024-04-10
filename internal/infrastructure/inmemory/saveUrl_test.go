package inmemory

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_SaveURL(t *testing.T) {
	// arrange
	storage := NewStorage()

	urlTest, _ := url.Parse("https://www.youtube.com/dsafsdf")
	urlTest2, _ := url.Parse("https://www.youtube.com/345345354")

	storage.aliasToURL["odfj2234kf"] = *urlTest
	storage.urlToAlias[*urlTest] = "odfj2234kf"

	testTable := []struct {
		name      string
		keyInput  string
		keyOutput string
		urlToSave url.URL
		wantErr   error
	}{
		{
			name:      "попытка вставки 2 раза одинаковой url, ключ должен остаться тем же",
			keyInput:  "sd23423423234",
			keyOutput: "odfj2234kf",
			urlToSave: *urlTest,
		},
		{
			name:      "успешная вставка",
			keyInput:  "23432",
			keyOutput: "23432",
			urlToSave: *urlTest2,
		},
	}

	// run
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			alias, err := storage.SaveURL(ctx, tt.keyInput, tt.urlToSave)
			assert.Equal(t, alias, tt.keyOutput)
			if !errors.Is(err, tt.wantErr) {
				t.Error()
			}
		})
	}
	// assert
}
