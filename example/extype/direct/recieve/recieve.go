package main

import (
	"github.com/streadway/amqp"
	"log"
	"rbmq/common"
	"strconv"
	"time"
)



func do(id int)  {
	// 创建连接
	conn, err := amqp.Dial(common.RabbitUri)
	common.FailOnError(err, "Failed to connect to Rabbitmq")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to create channel")

	// 声明交换机
	err = ch.ExchangeDeclare(
		"exdirecttest", // name
		"direct",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	common.FailOnError(err, "Failed to declare an exchange")

	// 声明队列
	q, err := ch.QueueDeclare(
		"quedirecttest",    // name
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
		"exdirecttest", // exchange
		false,
		nil,
	)
	ch.Qos(1, 0, false)
	workerName := "worker " + strconv.Itoa(id)
	// 接收消息
	msgs, err := ch.Consume(q.Name, workerName, true, false, false, false, nil)
	common.FailOnError(err, "Failed to register a consumer")

	log.Printf(" %d [*] Waiting for messages. To exit press CTRL+C", id)

	for msg := range msgs {
		log.Printf("%d Received a message: %s", id, msg.Body)
		time.Sleep(time.Second)
	}
}

func main()  {
	for i:=0; i<10; i++ {
		go do(i)
	}
	for {}
}