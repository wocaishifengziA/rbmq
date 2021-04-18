package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	URI string
	conn *amqp.Connection
	pubChannel *amqp.Channel
	subChannel *amqp.Channel

	QueueName    string
	ExChangeName string
	ExChangeType string
	RoutingKey   string
}

type MqExchange struct {
	ExName string
	ExType string
	VHost string
}


// mq 连接
func (c *RabbitMQ) Connect() error {
	var err error
	if conn, err := amqp.Dial(c.URI); err != nil {
		fmt.Println("MQ连接失败，err: ", err)
		return err
	}else {
		c.conn = conn
	}

	if c.pubChannel, err = c.conn.Channel(); err != nil {
		fmt.Println("MQ创建pub channel失败，err: ", err)
		return err
	}
	if c.subChannel, err = c.conn.Channel(); err != nil {
		fmt.Println("MQ创建sub channel失败，err: ", err)
		return err
	}
	return nil
}

// init exchange
func (c *RabbitMQ) InitExchange() error {
	if err := c.pubChannel.ExchangeDeclare(c.ExChangeName, c.ExChangeType, true, false, false, false, nil); err != nil {
		fmt.Println("exchange init failed!")
		return err
	}
	return nil
}

// init queue
func (c *RabbitMQ) InitQueue() error {
	_, err := c.pubChannel.QueueDeclare(c.QueueName, true, false, false, false, nil)
	if err != nil {
		fmt.Println("queue init failed!")
		return err
	}
	// queue bind exchange
	if err := c.pubChannel.QueueBind(c.QueueName, c.RoutingKey, c.ExChangeName, false, nil); err != nil {
		fmt.Println("绑定队列失败")
	}
	return nil
}

func (c *RabbitMQ) Init() error {
	var err error
	if err = c.Connect(); err != nil {
		return err
	}
	if err = c.InitExchange(); err != nil {
		return err
	}
	if err = c.InitQueue(); err != nil {
		return err
	}
	return nil
}

func (c *RabbitMQ) Public(topic string, message string) error {
	err := c.pubChannel.Publish(c.ExChangeName, topic, false, false, amqp.Publishing{
		Body:         []byte(message),
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *RabbitMQ) Sub () {

	data, err := c.subChannel.Consume(c.QueueName, "jake", false, false, false, false, nil)
	if err != nil {
		fmt.Println("接收失败")
	}
	for d := range data {
		println(d.Body)
	}
}

func NewRabbitMQ(uri string, exChangeName string, exChangeType string, queueName string, routingKey string)  *RabbitMQ {
	return &RabbitMQ{
		URI:          uri,
		ExChangeName: exChangeName,
		ExChangeType: exChangeType,
		QueueName: queueName,
		RoutingKey: routingKey,
	}
}
