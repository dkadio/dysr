package classes

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type CodesController struct {
	kv kvAdapter
}

type Code struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewCodesController() CodesController {
	bolt := NewBoltAdapterByString("codes.db", "codes")
	return CodesController{bolt}
}

func (c CodesController) CeateCode(gc *gin.Context) {

	url, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		log.Println("No Body Present")
		gc.JSON(http.StatusNoContent, nil)
	}
	log.Println(url)
	token := gc.GetHeader(headers.Authorization)

	log.Println("Call", token)
}

func (c CodesController) UpdateCode(gc *gin.Context) {
	log.Println("Call")
	id := gc.Param("id")
	fmt.Println("Some id to update", id)
}

//Query all codes for user request
func (c CodesController) GetCodes(gc *gin.Context) {
	log.Println("Call")
	gc.JSON(http.StatusOK, Code{"test", "test"})
}

func (c CodesController) GetCode(gc *gin.Context) {
	log.Println("Call")

	fmt.Println("Some id to update")
}

func createRandomKey(username string) string {
	h := sha256.New()
	h.Write([]byte(username + randSeq(10)))
	hs := base64.StdEncoding.EncodeToString(h.Sum(nil))

	rs := hs[:3] + hs[len(hs)-3:]
	return rs
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
