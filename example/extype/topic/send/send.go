package main

import (
	"github.com/streadway/amqp"
	"log"
	"rbmq/common"
	"strconv"
)

const (
	EXCHANGE = "ext"
	QUEUE = "quet"
)

func main()  {
	// 创建连接
	conn, err := amqp.Dial(common.RabbitUri)
	common.FailOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to create channel")

	//// 声明交换机
	//err = ch.ExchangeDeclare(
	//	"extopictest", // name
	//	"topic",      // type
	//	true,         // durable
	//	false,        // auto-deleted
	//	false,        // internal
	//	false,        // no-wait
	//	nil,          // arguments
	//)
	//common.FailOnError(err, "Failed to declare an exchange")
	//
	//// 声明队列
	//q, err := ch.QueueDeclare(
	//	"quetopictest",    // name
	//	false, // durable
	//	false, // delete when usused
	//	false,  // exclusive
	//	false, // no-wait
	//	nil,   // arguments
	//)
	//common.FailOnError(err, "Failed to declare a queue")
	//
	//// bind
	//err = ch.QueueBind(
	//	q.Name, // queue name
	//	"*.b.*",     // routing key
	//	"extopictest", // exchange
	//	false,
	//	nil,
	//)

	// 发布消息
	bodyq := "Hello"
	for i := 0; i<10; i++ {
		j := i
		body := "a.b." + bodyq + strconv.Itoa(j)
		log.Println(body)
		err = ch.Publish(EXCHANGE, body, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
		common.FailOnError(err, "Failed to Publish a message")
	}
}
