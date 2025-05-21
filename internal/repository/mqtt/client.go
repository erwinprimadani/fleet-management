package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client represents an MQTT client
type Client struct {
	Client mqtt.Client
	Topic  string
}

// NewClient creates a new MQTT client
func NewClient(cfg *mqttConf) (*Client, error) {
	opts := mqtt.NewClientOptions()
	brokerURL := fmt.Sprintf("tcp://%s:%d", cfg.broker, cfg.port)
	opts.AddBroker(brokerURL)
	opts.SetClientID(cfg.clientID)

	if cfg.username != "" && cfg.password != "" {
		opts.SetUsername(cfg.username)
		opts.SetPassword(cfg.password)
	}
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(5 * time.Minute)

	opts.SetDefaultPublishHandler(defaultMessageHandler)
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectionLostHandler)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	log.Printf("\nConnected to MQTT broker at %s", brokerURL)

	return &Client{
		Client: client,
		Topic:  cfg.topic,
	}, nil
}

// Subscribe from the MQTT broker
func (c *Client) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	log.Println("####MQTT Subscribe", topic)
	return c.Client.Subscribe(topic, qos, callback)
}

// Disconnect disconnects from the MQTT broker
func (c *Client) Disconnect() {
	if c.Client.IsConnected() {
		c.Client.Disconnect(250)
		log.Println("Disconnected from MQTT broker")
	}
}

// Default message handler
func defaultMessageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message on topic: %s", msg.Topic())
}

// Connect handler
func connectHandler(client mqtt.Client) {
	log.Println("MQTT client connected")
}

// Connection lost handler
func connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("MQTT connection lost: %v", err)
}
