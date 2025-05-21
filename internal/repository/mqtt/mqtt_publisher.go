package mqtt

import (
	"fmt"
	"log"
)

func (c *Client) PublishLocation(vehicleID string, payload []byte) error {
	topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID)
	log.Println("####MQTT Subscribe", topic)
	token := c.Client.Publish(topic, 1, false, payload)
	token.Wait()
	return token.Error()
}
