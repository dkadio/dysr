package main

import (
	"fmt"
	bib "github.com/dkadio/dysr/internal"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	re := `http.*`
	re2 := ".{6,}"

	r := mux.NewRouter().SkipClean(true).UseEncodedPath()
	s := r.PathPrefix("/api/v1/url").Subrouter()

	s.HandleFunc("/short/{url:"+re+"}", getShortUrL)
	r.HandleFunc("/{code:"+re2+"}", redirect)

	http.ListenAndServe(":8080", r)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	surl := bib.NewShortUrl(r.URL.String()[1:])
	fmt.Println("I will redirect this url to: ", surl.Long.String())

	http.Redirect(w, r, surl.Long.String(), http.StatusTemporaryRedirect)

	//now its time for statistics
}

func getShortUrL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := params["url"]
	fmt.Println("I will shorten this url:", url)
	//short this url
	fu := bib.NewLongUrl(url)
	fmt.Println("This is long: ", fu.Long.String())
	fmt.Println("This is Short: ", fu.Short.String())
	w.Write([]byte(fu.Short.String()))
}
