package app 

import (
    "go.temporal.io/sdk/workflow"
)

// ConsumerWorkFlow processes messages from a Kafka topic within a Temporal workflow context.
//
// Parameters:
// - ctx: The Temporal workflow context used to manage and control the execution of the workflow. It provides access
//   to Temporal's workflow features and ensures that the workflow's state is maintained throughout the processing.
//
// Returns:
// - An error if there was an issue during message consumption or processing. If the function completes successfully,
//   it returns nil to indicate that there were no errors.
func ConsummerWorkFlow(ctx workflow.Context) error {
    logger := workflow.GetLogger(ctx)
    logger.Info("Starting ConsummerWorkFlow")

    future := workflow.ExecuteActivity(ctx, ConsumeMessagesActivity)
    if err := future.Get(ctx, nil); err != nil {
        logger.Error("Failed to execute consume messages activity:", err)
        return err
    }
    return nil
}

// ConsumeMessagesActivity consumes messages from a Kafka topic as part of a Temporal activity.
//
// This function is designed to be used as a Temporal activity that handles the consumption of messages from a Kafka
// topic. It is responsible for connecting to Kafka, polling for messages, and processing them as needed.
//
// Returns:
// - An error if there was an issue consuming messages or processing them. If the function completes successfully,
//   it returns nil to indicate that there were no errors.
func ConsumeMessagesActivity() error {
    return ConsumeMessage(Consumer)
}
