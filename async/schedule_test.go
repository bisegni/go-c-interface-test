package async

import "testing"

func TestSchedulerInitDeinit(t *testing.T) {
	Init(1)
	Deinit()
}
