package classes

import (
	"log"
	"testing"
)

func TestNewLongUrl(t *testing.T) {
	url := NewLongUrl("https://notdysr.io/izibitzi")
	if url.Short.String() == "" {
		log.Println("no Short version is set for url", url.Long.String())
		t.Fail()
	}
}

func TestUrls(t *testing.T) {
	url2 := NewLongUrl("https://notdysr.io/izibitzi")
	url := NewShortUrl(url2.Short.String())

	if url2.Long.String() != url.Long.String() {
		t.Fail()
	}
	if url2.Short.String() != url.Short.String() {
		t.Fail()
	}
}
