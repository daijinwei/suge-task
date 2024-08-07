package app

import (
        "github.com/Shopify/sarama"
        cluster "github.com/bsm/sarama-cluster"
)

const (
	// ChildProduceWorkFlowFMT is a format string for constructing workflow IDs
	// for child production workflows. It expects an integer to be substituted
	// into the placeholder (%d) in the string.
	ChildProduceWorkFlowFMT = "ChildProduceWorkFlowID-%d"

	// ProduceWorkFlow represents the unique identifier for the production workflow.
	ProduceWorkFlow = "PRODUCE_WORK_FLOW"

	// ConsumeWorkFlow represents the unique identifier for the consumption workflow.
	ConsumeWorkFlow = "CONSUME_WORK_FLOW"

	// MessageTaskQueue represents the name of the task queue used for messaging.
	MessageTaskQueue = "SUGE_TASK_QUEUE"
)

var (
	// TemporalHstPort is the address and port number for the Temporal service host.
	TemporalHstPort =  "127.0.0.1:7233"

	// Count specifies the number of items or tasks.
	Count 	= 	200

	// Brokers is a slice of strings representing the addresses of Kafka broker nodes.
	Brokers = []string{"127.0.0.1:9092", "127.0.0.1:9093", "127.0.0.1:9094"}

	// Topic is the name of the Kafka topic used for publishing and consuming messages.
	Topic = "suge-test"

	// GroupId is the consumer group ID used for the Kafka consumer group to track offsets and manage message consumption.
	GroupId = "suge-test"

	// Producer is an instance of the Sarama.SyncProducer used for sending messages to Kafka topics.
	Producer sarama.SyncProducer

	// Consumer is an instance of the cluster.Consumer used for consuming messages from Kafka topics.
	Consumer *cluster.Consumer
)
