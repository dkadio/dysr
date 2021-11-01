package models

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/dkadio/dysr/util"
	"github.com/google/uuid"
	"strings"
	"time"
)

type CodeValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Code struct {
	UUID    string  `json:"-" path:"id" bson:"-"`
	Options Options `json:"options"`
	Key     string  `json:"key"`
	Value   string  `json:"value"`
}

type Options struct {
	Text       string `json:"text"`       //'https://github.com/ushelp/EasyQRCodeJS',
	Width      int    `json:"width"`      //256,
	Height     int    `json:"height"`     //256,
	ColorDark  string `json:"colorDark"`  //'#000000',
	ColorLight string `json:"colorLight"` //'#ffffff',
}

type CreateCode struct {
	Key     string  `json:"key"`
	Value   string  `json:"value"`
	Options Options `json:"options"`
}

type UserCode struct {
	UUID    string `json:"id" path:"id" bson:"_id"`
	User    string `json:"user"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
	Code    Code   `json:"code" bson:",inline"`
}

func NewUserCode(user, value string, options Options) UserCode {
	config := util.LoadConfig()
	userCode := UserCode{}
	userCode.UUID = uuid.New().String()
	userCode.User = user
	userCode.Created = time.Now().Unix()
	userCode.Updated = time.Now().Unix()
	userCode.Code.Value = value
	userCode.Code.Key = strings.ToUpper(createRandomKey(user))
	userCode.Code.Options = options
	userCode.Code.Options.Text = config.ServiceURL + "/" + userCode.Code.Key
	return userCode
}

func createRandomKey(username string) string {
	h := sha256.New()
	h.Write([]byte(username + uuid.New().String()))
	hs := base64.StdEncoding.EncodeToString(h.Sum(nil))

	rs := hs[:3] + hs[len(hs)-3:]
	return rs
}
