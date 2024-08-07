package app

import (
	"fmt"
	"errors"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// initConsumer initializes and configures a Kafka consumer instance.
func initConsumer(){
	var err error
	config := cluster.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	Consumer, err = cluster.NewConsumer(Brokers,GroupId,[]string{Topic},config)
	if err != nil {
		fmt.Printf("init consumer failed -> %v \n", err)
		panic(err.Error())
	}
	if Consumer == nil {
		panic(fmt.Sprintf("consumer is null. kafka info -> {brokers:%v, topic: %v, group: %v}", Brokers, Topic, GroupId))
	}
	fmt.Printf("init consumer successfully, consumer -> %v, topic -> %v, ", Consumer, Topic)
}

func init(){
	initConsumer()
}

// ConsumeMessage retrieves and processes a message from a Kafka topic using the provided cluster.Consumer instance.
//
// Parameters:
// - consumer: An instance of `cluster.Consumer` used to connect to Kafka and retrieve messages from topics. This
//   consumer should be properly initialized and configured before being passed to this function.
//
// Returns:
// - An error if there was an issue retrieving or processing the message. If the message is successfully consumed and
//   processed, the function returns nil to indicate that there were no errors.
func ConsumeMessage(consumer *cluster.Consumer) error {
	msg, ok := <-Consumer.Messages()
	if ok {
		fmt.Printf("consume kafka msg: %s \n", msg.Value)
		return nil
	}else{
		return errors.New("Bad messages")
	}
}

