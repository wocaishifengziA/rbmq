package main

import (
	"github.com/streadway/amqp"
	"log"
	"rbmq/common"
)

const rbUri = "amqp://root:root@192.168.0.111:5672/"

func main()  {
	// 创建连接

	conn, err := amqp.Dial(rbUri)
	common.FailOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to create channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	common.FailOnError(err, "Failed to declare a queue")
	// 接收消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


