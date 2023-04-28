package server

import "net/http"

type Page struct{}

func (p Page) rootPage(w http.ResponseWriter, req *http.Request) {
	str := "Hello, Client! Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

var RoutingTable = map[string]interface{}{
	"/": Page.rootPage,
}

func GetRoutingTable() map[string]interface{} {
	return RoutingTable
}
