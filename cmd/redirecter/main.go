package main

import (
	"log"
	"net/http"

	"github.com/dkadio/dysr/internal/controllers"
	"github.com/dkadio/dysr/internal/models"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Lshortfile)
	rh := controllers.NewRedirectController()
	rh.Init()

	re2 := ".{6,6}"

	r := mux.NewRouter().SkipClean(true).UseEncodedPath()
	r.HandleFunc("/{code:"+re2+"}", redirect)

	http.ListenAndServe(":8001", r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	rh := controllers.NewRedirectController()
	code := r.URL.Path[1:]
	rd, err := rh.GetValueFor(code)
	if err == nil {
		req := models.NewRequestFrom(r, code)
		rh.InformRedirect(req)
	}
	http.Redirect(w, r, rd, http.StatusTemporaryRedirect)
}
