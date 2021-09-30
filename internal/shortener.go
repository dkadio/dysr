package classes

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/dkadio/dysr/util"
	"log"
	"net/url"
)

type Shortener struct {
	hash string
	kv   kvAdapter
}

type kvAdapter interface {
	getValueFor(string) string
	putValueFor(string, string) error
}

func newShortener(hash, storeName, bucketName string) Shortener {
	config, _ := util.LoadConfig("./util")
	bolt := NewBoltAdapter(config)
	return Shortener{hash, bolt}
}

// gets a long version of a URL and returns the short version of a URL
func (s Shortener) GetShortUrl(u url.URL) url.URL {
	//check if its allready created and save in store
	str := s.kv.getValueFor(u.String()) // getValueFor(u.String(), s.storeName, s.bucketName)
	var u1 *url.URL
	var err error

	if str != "" {
		u1, err = url.Parse(str)
	} else {
		u2, _ := s.createShortVersion(u)
		u1, err = url.Parse(u2)
	}

	if err != nil {
		log.Println("This is not a valid domain:", u.String())
	}
	return *u1
}

//hashs long version an saves it to store
func (s Shortener) createShortVersion(longUrl url.URL) (string, error) {
	config, _ := util.LoadConfig("./util")
	su, _ := url.Parse(longUrl.String())
	su.Path = hashAndShort(longUrl.Path)
	su.Host = config.ServiceURL
	su.Scheme = config.ServiceProtocol
	err := s.writeToStore(su.Path, longUrl.String())
	return su.String(), err
}

func hashAndShort(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	hs := base64.StdEncoding.EncodeToString(h.Sum(nil))

	rs := hs[:3] + hs[len(hs)-3:]
	return rs
}

// gets
func (s Shortener) GetLongUrl(shortUrl url.URL) url.URL {
	//we expect that the long url is written to store
	log.Println("Try to get Long Url for", shortUrl.String())
	u, _ := url.Parse(s.readFromStore(shortUrl.String()))
	return *u
}

func (s Shortener) writeToStore(short, long string) error {
	log.Println("write to store", short, "=", long)
	if err := s.kv.putValueFor(short, long); err != nil {
		log.Fatal("Something went wrong at writing to store")
		return err
	}
	return nil
}

func (s Shortener) readFromStore(short string) string {
	return s.kv.getValueFor(short)
}
