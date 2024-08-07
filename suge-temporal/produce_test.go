package app

import (
	"testing"
)

func TestConsumeMessage(t *testing.T) {
	ConsumeMessage()
}

[root@cloud suge-temporal]# cat produce_test.go
package app

import (
"testing"
//"github.com/Shopify/sarama"
)

func TestProduceMessage(t *testing.T) {
	// Define the expected values
	//expectedPartition := int32(0)
	//expectedOffset := int64(1)
	// Call the produceMessage function
	err := produceMessage(Producer)

	// Check if any error occurred
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
