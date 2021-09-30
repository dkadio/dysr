package classes

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestHashAndShort(t *testing.T) {
	s := hashAndShort("testingkajhsjdlkasbdkljas")
	fmt.Printf("%x", s)
}

func TestCreateShortVersion(t *testing.T) {
	s := newShortener("sha256", "url_test.db", "urls")
	u, _ := url.Parse("https://dysr.io/this/is/a/path")
	log.Println("url:", u)
	u1, err := s.createShortVersion(*u)
	log.Printf("%s is short for %s", u1, u.String())
	if err != nil {
		t.Fail()
	}
}
