package server

var pages Page
var RoutingTable = map[string]interface{}{
	"/":      pages.pageRoot,
	"/hello": pages.pageHello,
}

func GetRoutingTable() map[string]interface{} {
	return RoutingTable
}
