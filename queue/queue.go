package queue

type queue struct {
	id   string
	jobs chan job
}

var queues []*queue

func getQueue(id string) *queue {
	for _, queue := range queues {
		if queue.id == id {
			return queue
		}
	}

	queue := &queue{
		id:   id,
		jobs: make(chan job),
	}

	runWorker(queue)

	queues = append(queues, queue)

	return queue
}

func AddJob(queueId string, action func()) {
	queue := getQueue(queueId)

	queue.jobs <- job{action}
}
