package app

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// initProducer initializes and configures a Sarama SyncProducer instance.
//
// Returns:
// - The initialized Sarama SyncProducer instance, which can be used to send messages to Kafka.
// - An error if there was an issue initializing the producer, such as connection failures or invalid configuration.
func initProducer() {
	var err error
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	brokers := Brokers
	if Producer, err = sarama.NewSyncProducer(brokers,config); err != nil {
		fmt.Printf("init Producer failed -> %v \n", err)
		panic(err)
	}
	fmt.Println("init Producer successfully")
}


// produceMessage sends a message to a Kafka topic using the provided Sarama SyncProducer.
//
// Parameters:
// - producer: An instance of sarama.SyncProducer used to send messages to Kafka. This producer should be properly
//   initialized and connected to the Kafka cluster before calling this function.
//
// Returns:
// - An error if there was an issue sending the message; otherwise, it returns nil to indicate that the message was
//   successfully sent.
func produceMessage(producer sarama.SyncProducer) error {
	part, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic:     Topic,
		Value:     sarama.ByteEncoder("this is test message."),
	})

	if err != nil {
		fmt.Println("Send message failed, err:", err)
		return err
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", part, offset)
	return nil
}

func init(){
	initProducer()
}
