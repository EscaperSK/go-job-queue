package queue

type job struct {
	action func()
}

func (j *job) run() {
	j.action()
}
