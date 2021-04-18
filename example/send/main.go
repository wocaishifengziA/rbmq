package main

import (
	"github.com/streadway/amqp"
	"log"
	"rbmq/common"
)



func main()  {
	// 创建连接
	conn, err := amqp.Dial(common.RabbitUri)
	common.FailOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to create channel")

	// 声明队列
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	common.FailOnError(err, "Failed to declare a queue")

	// 发布消息
	body := "Hello World!"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(body),
	})

	if ch.Confirm(false) != nil {
		log.Println("bad")
	}else {
		log.Println("ok")
	}
	common.FailOnError(err, "Failed to Publish a message")
}
