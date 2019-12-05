package async

import "log"

import "context"

var runnableChan chan Runnable
var ctx context.Context
var cancel context.CancelFunc

func execute(ctx context.Context, runnableChan <-chan Runnable) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Exit runnable scheduler")
			return

		case r := <-runnableChan:
			log.Printf("Executing new job %s\n", r.name())
			r.execute()
		}
	}
}

// Init ...
func Init(queueSize int) {
	runnableChan = make(chan Runnable, queueSize)
	ctx, cancel = context.WithCancel(context.Background())
	go execute(ctx, runnableChan)
}

// Deinit ...
func Deinit() {
	close(runnableChan)
}

// TryEnqueue impl.
func TryEnqueue(r Runnable) bool {
	select {
	case runnableChan <- r:
		return true
	default:
		return false
	}
}
