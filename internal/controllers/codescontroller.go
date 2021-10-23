package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	classes "github.com/dkadio/dysr/internal"
	dm "github.com/dkadio/dysr/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
)

type kvAdapter interface {
	GetValueFor(string) string
	PutValueFor(string, string) error
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type CodesController struct {
	adapter kvAdapter
}

func NewCodesController() CodesController {
	bolt := classes.NewBoltAdapterByString("codes.db", "codes")
	return CodesController{bolt}
}

func (c CodesController) CeateCode(gc *gin.Context) error {

	url, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		log.Println("No Body Present")
		gc.JSON(http.StatusNoContent, nil)
	}
	log.Println(url)
	token := gc.GetHeader(headers.Authorization)

	log.Println("Call", token)
	return nil
}

func (c CodesController) UpdateCode(gc *gin.Context) error {
	log.Println("Call")
	id := gc.Param("id")
	fmt.Println("Some id to update", id)
	return nil
}

//Query all codes for user request
func (c CodesController) GetCodes(gc *gin.Context, params *dm.Code) (dm.Code, error) {

	log.Println("Call Get Codes")

	return dm.Code{Key: "Test", Value: "Test"}, nil
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
