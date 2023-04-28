package main

import (
	"runtime"
	"server-clients/client"
	"server-clients/server"
	"sync"
)

func main() {

	/* ******************************************************************************
	 * Configuration
	 * ******************************************************************************/
	runtime.GOMAXPROCS(runtime.NumCPU())
	numClients := 2
	var channelFinishSignals []chan bool
	for i := 0; i < numClients; i++ {
		channelFinishSignals = append(channelFinishSignals, make(chan bool))
	}

	/* ******************************************************************************
	 * Run http client(s)
	 * ******************************************************************************/
	waitClientsGroup := sync.WaitGroup{}

	for i := 0; i < numClients; i++ {
		waitClientsGroup.Add(1)
		go client.RunClient(&waitClientsGroup, i, channelFinishSignals[i])
	}

	/* ******************************************************************************
	 * Run a http server
	 * ******************************************************************************/
	waitServer := sync.WaitGroup{}
	waitServer.Add(1)
	go server.RunServer(&waitServer)

	/* ******************************************************************************
	 * Kill all sub goroutines
	 * ******************************************************************************/
	waitServer.Wait()
	for i := 0; i < numClients; i++ {
		channelFinishSignals[i] <- true
		close(channelFinishSignals[i])
	}

	/* ******************************************************************************
	 * Synchronization & Verifying all the go routines done
	 * ******************************************************************************/
	waitClientsGroup.Wait()

}
