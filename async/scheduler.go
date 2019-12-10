package async

import (
	"context"
	"log"
	"sync"
)

var runnableChan chan Runnable
var ctx context.Context
var cancel context.CancelFunc
var wg sync.WaitGroup

func execute(ctx context.Context, runnableChan <-chan Runnable) {
	log.Println("Entering scheduler execute")
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("Leaving runnable scheduler")
			return

		case r := <-runnableChan:
			log.Printf("Executing new job '%s'\n", r.name())
			r.execute()
			if ctx.Err() != nil {
				log.Println("Scheduler error -> leaving runnable scheduler")
				return
			}
		}
	}

}

// Init ...
func Init(queueSize int, workerCount int) {
	runnableChan = make(chan Runnable, queueSize)
	ctx, cancel = context.WithCancel(context.Background())
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go execute(ctx, runnableChan)
	}

}

// Deinit ...
func Deinit() {
	//call cancel on context
	cancel()
	wg.Wait()
}

// Enqueue impl.
func Enqueue(r Runnable) bool {
	select {
	case runnableChan <- r:
		return true
	default:
		return false
	}
}
