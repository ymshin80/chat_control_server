package kafka

import (
	"chat_golang_control/config"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	cfg *config.Config
	consumer *kafka.Consumer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}

	var err error
	log.Println("URL=",cfg.Kafka.URL )
	log.Println("ClientID=",cfg.Kafka.GroupID )

	//TODO: consumer 로 변경
	if k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"group.id": cfg.Kafka.GroupID,
		"auto.offset.reset": "latest",
		"acks": "all",
	}); err != nil  {
		return nil , err
	} else {
		return k , nil
	}
}

func (k *Kafka) Pool(timeoutMs int) kafka.Event {
	return k.consumer.Poll(timeoutMs)
}

func (k *Kafka) RegisterSubTopic(topic string) error {
	if err := k.consumer.Subscribe(topic, nil); err != nil  {
		return err
	} else {
		return nil
	}
}

// func (k *Kafka) PublishEvent(topic string, value []byte ,ch chan kafka.Event) (kafka.Event, error) {
// 	if err := k.consumer.(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{
// 			Topic: &topic,
// 			Partition: kafka.PartitionAny,
// 		},
// 		Value: value,
// 	},ch); err != nil {
// 		return nil, err
// 	} else {
// 		return <- ch, nil
// 	}
// }