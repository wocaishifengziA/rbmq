package main

import (
	"github.com/streadway/amqp"
	"rbmq/common"
	"strconv"
)



func main()  {
	// 创建连接
	conn, err := amqp.Dial(common.RabbitUri)
	common.FailOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to create channel")

	// 声明交换机
	err = ch.ExchangeDeclare(
		"ex123", // name
		"fanout",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	common.FailOnError(err, "Failed to declare an exchange")

	// 声明队列
	q, err := ch.QueueDeclare(
		"que123",    // name
		false, // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	// bind
	err = ch.QueueBind(
		q.Name, // queue name
		"key123",     // routing key
		"ex123", // exchange
		false,
		nil,
	)

	// 发布消息
	bodyq := "Hello "
	for i := 0; i<10; i++ {
		j := i
		body := bodyq + strconv.Itoa(j)
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
		common.FailOnError(err, "Failed to Publish a message")
	}
}