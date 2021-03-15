package solution

import (
	"math/rand"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Codec struct {
	base string
	urls map[string]string
}

func Constructor() Codec {
	return Codec{
		"http://tinyurl.com/",
		map[string]string{},
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	hash := make([]byte, 8)
	for i := range hash {
		hash[i] = charset[rand.Intn(len(charset))]
	}
	c.urls[string(hash)] = longUrl
	return c.base + string(hash)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.urls[shortUrl[len(c.base):]]
}
