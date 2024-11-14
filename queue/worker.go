package queue

func runWorker(queue *queue) {
	go func() {
		for {
			job := <-queue.jobs
			job.run()
		}
	}()
}
