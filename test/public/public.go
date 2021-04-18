package main

import (
	"log"
	"rbmq/common"
	"rbmq/rabbitmq"
)

func main() {
	client := rabbitmq.NewRabbitMQ(common.RabbitUri, common.ExChangeName, common.ExChangeType, common.QueueName, common.RoutingKey)
	if err := client.Init(); err != nil {
		log.Println(err)
		return
	}
	err := client.Public("aabb", "mytest")
	common.FailOnError(err, "Failed to public")
}
