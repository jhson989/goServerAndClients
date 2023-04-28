package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

/* ******************************************************************************
 * Entry point of a server logic
 * ******************************************************************************/
func RunServer(waitGroup *sync.WaitGroup) {
	defer endOfServer(waitGroup)

	log.Printf("[Server] server has started...\n")
	log.Printf("[Server] server would be running for 10s\n")

	httpServer := runServerRoutine()
	time.Sleep(time.Duration(10) * time.Second)
	// Gracefully shutting down the server
	if err := httpServer.Shutdown(context.TODO()); err != nil {
		panic(err)
	}

}

/* ******************************************************************************
 * Server logic
 * ******************************************************************************/
func runServerRoutine() *http.Server {

	httpServer := &http.Server{Addr: ":8000"}
	//http.HandleFunc("/", pages.rootPage)
	http.HandleFunc("/", RoutingTable["/"].(func(http.ResponseWriter, *http.Request)))
	http.HandleFunc("/hello", RoutingTable["/hello"].(func(http.ResponseWriter, *http.Request)))

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return httpServer

}

/* ******************************************************************************
 * Finalized server logic
 * ******************************************************************************/
func endOfServer(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	log.Print("[Server] server has finished...\n")
	log.Print("[Server] try to kill all sub goroutines for clients ...\n")

}
