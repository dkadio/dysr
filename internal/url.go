package classes

import (
	"github.com/dkadio/dysr/util"
	"log"
	"net/url"
)

type Url struct {
	Long  url.URL
	Short url.URL
	util  Shortener
}

func NewShortUrl(u string) Url {
	config := util.LoadConfig()
	pu, _ := url.Parse(u)
	//read this from config
	util := newShortener(config.HASH, config.StoreName, config.BucketName)
	ru := Url{Short: *pu, util: util}
	ru.Long = ru.util.GetLongUrl(ru.Short)
	return ru
}

func NewLongUrl(u string) Url {
	config := util.LoadConfig()
	pu, _ := url.Parse(u)
	util := newShortener(config.HASH, config.StoreName, config.BucketName)
	ru := Url{Long: *pu, util: util}
	log.Println("Try to build short version for", ru.Long.String())
	ru.Short = ru.util.GetShortUrl(ru.Long)
	log.Println("Short Version is:", ru.Short.String())
	return ru
}
