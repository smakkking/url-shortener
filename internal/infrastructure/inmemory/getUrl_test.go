package inmemory

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_GetURL(t *testing.T) {
	// arrange
	storage := NewStorage()

	urlTest, _ := url.Parse("https://www.youtube.com/dsafsdf")

	storage.aliasToURL["odfj2234kf"] = *urlTest
	storage.urlToAlias[*urlTest] = "odfj2234kf"

	testTable := []struct {
		name    string
		key     string
		wantURL string
		wantErr error
	}{
		{
			name:    "OK",
			key:     "odfj2234kf",
			wantURL: "https://www.youtube.com/dsafsdf",
		},
		{
			name:    "No such url",
			key:     "234324",
			wantErr: ErrNotFound,
		},
	}

	// run
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			originalURL, err := storage.GetURL(ctx, tt.key)
			assert.Equal(t, originalURL.String(), tt.wantURL)
			if !errors.Is(err, tt.wantErr) {
				t.Error()
			}
		})
	}
	// assert

}
