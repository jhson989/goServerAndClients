package client

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

/* ******************************************************************************
 * Entry point of a client logic
 * ******************************************************************************/
func RunClient(waitGroup *sync.WaitGroup, clientID int, channelFinishSignal chan bool) {

	defer endOfProgram(waitGroup, clientID)

	time.Sleep(time.Duration(1) * time.Second)
	log.Printf("[Client-%d] client has started...\n", clientID)

	for {
		select {
		case <-channelFinishSignal:
			return
		default:
			delta := rand.Intn(5000) + 1000
			log.Printf("[Client-%d] client will sleep for %.3f s\n", clientID, float64(delta)*0.001)
			runClientRoutine(clientID)
			time.Sleep(time.Duration(delta) * time.Millisecond)
		}
	}

}

/* ******************************************************************************
 * Client logic
 * ******************************************************************************/
func runClientRoutine(clientID int) {
	response, err := http.Get("http://0.0.0.0:8000/")
	if err == nil {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("[Client-%d] response:\n  > %s\n", clientID, data)
	}
}

/* ******************************************************************************
 * Finalized client logic
 * ******************************************************************************/
func endOfProgram(waitGroup *sync.WaitGroup, clientID int) {

	defer waitGroup.Done()
	log.Printf("[Client-%d] client has finished...\n", clientID)

}
