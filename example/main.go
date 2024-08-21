package main

import (
	"exabytez/loki"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func LokiWorker(loki *loki.Loki, ch <-chan string, chExit chan struct{}) {
	fmt.Println("LokiWorker has been started")
	for {
		select {
		case msg := <-ch:
			loki.AddMessage(msg)
		case <-time.After(100 * time.Microsecond):
			loki.Fire()
		case <-chExit:
			loki.Fire()
			chExit <- struct{}{}
			fmt.Println("LokiWorker has been finished")
			return
		}

	}
}

func main() {
	lokiCH := make(chan string)
	lokiError := make(chan struct{})

	lokiClient := &loki.LokiClient{
		Host: "http://localhost:3100",
	}

	lokiOpt := &loki.LokiOpt{
		Client: lokiClient,
	}

	loki := loki.NewLoki(lokiOpt, loki.Label{"app": "goLokiTest7"})

	go LokiWorker(loki, lokiCH, lokiError)

	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(idx int) {
			time.Sleep(time.Duration(rand.Int32N(200) * int32(time.Millisecond)))
			lokiCH <- fmt.Sprintf("go[%d]: Is executed", idx)
			wg.Done()
		}(i)
	}
	wg.Wait()

	lokiError <- struct{}{}
	<-lokiError

	close(lokiCH)

}
