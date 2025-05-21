package rabbitmq

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Client represents a RabbitMQ client
type Client struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queue    string
}

// NewClient creates a new RabbitMQ client
func NewClient(cfg *rabbitMQConf) (*Client, error) {
	// Connect to RabbitMQ
	connString := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.user, cfg.password, cfg.host, cfg.port)
	fmt.Println("con rabbit", connString)
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	if ch == nil {
		conn.Close()
		log.Printf("Channel is nil")
		return nil, err
	}

	// Declare an exchange
	err = ch.ExchangeDeclare(
		cfg.exchange, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare an exchange: %v", err)
	}

	// Declare a queue
	q, err := ch.QueueDeclare(
		cfg.queue, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		q.Name,
		"geofence.*",
		cfg.exchange,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue to exchange: %v", err)
	}

	log.Printf("Connected to RabbitMQ at %s:%d, exchange: %s, queue: %s",
		cfg.host, cfg.port, cfg.exchange, cfg.queue)

	return &Client{
		conn:     conn,
		channel:  ch,
		exchange: cfg.exchange,
		queue:    cfg.queue,
	}, nil
}

// Publish publishes a message to the exchange
func (c *Client) Publish(routingKey string, message []byte) error {
	log.Printf("\n####Published message to exchange %s with routing key %s", c.exchange, routingKey)
	err := c.channel.PublishWithContext(
		context.Background(),
		c.exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}

// Consume returns a channel of messages from the configured queue.
func (c *Client) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := c.channel.Consume(
		c.queue, // queue (should be "geofence_alerts")
		"",      // consumer
		true,    // auto-ack (set to false for manual ack in worker)
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register consumer: %v", err)
	}
	return msgs, nil
}

func (c *Client) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	log.Println("RabbitMQ connection closed")
}
