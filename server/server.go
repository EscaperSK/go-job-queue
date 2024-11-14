package server

import (
	"net/http"
	"time"
	"webtest/queue"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", handle)

	router.Run("127.0.0.1:8080")
}

func handle(context *gin.Context) {
	queueId := context.Query("queue")

	processInQueue(queueId)

	context.String(http.StatusOK, "Executed on queue \""+queueId+"\"")
}

func processInQueue(queueId string) {
	await := make(chan struct{})

	queue.AddJob(queueId, func() {
		fakeJob()
		await <- struct{}{}
	})

	<-await
}

func fakeJob() {
	time.Sleep(time.Second * 3)
}
