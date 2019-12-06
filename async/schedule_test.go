package async

import "testing"

import "time"

import "gotest.tools/assert"

type JobOne struct {
	count int
}

func (j *JobOne) execute() {
	for idx := 0; idx < 100; idx++ {
		j.count++
		time.Sleep(1 * time.Millisecond)
	}
}

func (j *JobOne) name() string {
	return "JobOne"
}

func TestSchedulerInitDeinit(t *testing.T) {
	job := JobOne{count: 0}
	Init(1, 2)
	Enqueue(&job)
	//give time to scheduler to execute the job
	time.Sleep(100 * time.Millisecond)
	Deinit()
	assert.Assert(t, job.count == 100)
}
