package server

import "net/http"

type Page struct{}

func (p Page) pageRoot(w http.ResponseWriter, req *http.Request) {
	str := "Access to " + req.URL.Path
	w.Write([]byte(str))
}

func (p Page) pageHello(w http.ResponseWriter, req *http.Request) {
	str := "Hello, Client!"
	w.Write([]byte(str))
}
