package models

import (
	"net/http"
	"net/url"
)

type Request struct {
	CODE       string
	URL        *url.URL
	Proto      string
	Header     http.Header
	Host       string
	Trailer    http.Header
	RemoteAddr string
	RequestURI string
}

type CodeStats struct {
	Key    string `json:"key"`
	Date   string `json:"date"`
	CodeId string `json:"codeid"`
	Stats  Stats  `json:"stats"`
}

type Stats struct {
	Clicks ClickStats `json:"clicks"`
	Ips    []string   `json:"ips"`
}

type ClickStats struct {
	Total   int `json:"total"`
	Desktop int `json:"desktop"`
	Tablet  int `json:"tablet"`
	Mobile  int `json:"mobile"`
}

func NewRequestFrom(req *http.Request, code string) Request {
	return Request{CODE: code, URL: req.URL, Proto: req.Proto, Header: req.Header, Host: req.Host, Trailer: req.Trailer, RemoteAddr: req.RemoteAddr, RequestURI: req.RequestURI}
}
