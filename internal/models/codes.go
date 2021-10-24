package models

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Code struct {
	UUID  string `json:"-" path:"id" bson:"-"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserCode struct {
	UUID    string `json:"id" path:"id" bson:"_id"`
	User    string `json:"user"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
	Code    Code   `json:"code" bson:",inline"`
}

func NewUserCode(user, value string) UserCode {
	userCode := UserCode{}
	userCode.UUID = uuid.New().String()
	userCode.User = user
	userCode.Created = time.Now().Unix()
	userCode.Updated = time.Now().Unix()
	userCode.Code.Value = value
	userCode.Code.Key = strings.ToUpper(createRandomKey(user))
	return userCode
}

func createRandomKey(username string) string {
	h := sha256.New()
	h.Write([]byte(username + uuid.New().String()))
	hs := base64.StdEncoding.EncodeToString(h.Sum(nil))

	rs := hs[:3] + hs[len(hs)-3:]
	return rs
}
