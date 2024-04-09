package models

import "net/url"

type URLKey string

func (u *URLKey) Transform() (url.URL, error) {
	return url.URL{}, nil
}
