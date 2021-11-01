package main

import (
	"fmt"
	"net/http"

	"github.com/dkadio/dysr/internal/controllers"
	"github.com/gorilla/mux"
)

func main() {

	//TODO: This loads first all key values from mongo to bolt
	//TODO: Then redirect every received key to given value
	//TODO: reacts on CUD Code Options and updates bolt
	rh := controllers.NewRedirectController()
	rh.Init()

	re2 := ".{6,6}"

	r := mux.NewRouter().SkipClean(true).UseEncodedPath()
	r.HandleFunc("/{code:"+re2+"}", redirect)

	http.ListenAndServe(":8001", r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	rh := controllers.NewRedirectController()
	rd := rh.GetValueFor(r.URL.Path[1:])
	fmt.Println("Redirecting to: ", rd)

	http.Redirect(w, r, rd, http.StatusTemporaryRedirect)

	//now its time for statistics
}
