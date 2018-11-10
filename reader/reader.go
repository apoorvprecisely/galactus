package reader

import (
	"github.com/apoorvprecisely/galactus"
	"github.com/apoorvprecisely/galactus/surfer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)



func Start(config galactus.Config) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Reader.Uri,
		"group.id":          "galactusOne",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		return err
	}

	c.SubscribeTopics([]string{config.Reader.Topic}, nil)
	ss := surfer.New()
	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			return err
		}
		ss.OnEvent(msg.Value)
	}
	c.Close()
	return nil
}
