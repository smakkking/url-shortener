package models

import "net/url"

type URLKey string

func (u *URLKey) Transform() url.URL {
	return url.URL{}
}
