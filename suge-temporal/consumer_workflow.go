package app

import (
    "fmt"
    "time"
    "context"
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
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: 24 * time.Hour, // 根据需求调整超时时间
        ScheduleToCloseTimeout: 24 * time.Hour, // 根据需求调整超时时间
    }

    // 启动多个消费者活动
    numConsumers := 3 // 设定消费者的数量
    for i := 0; i < numConsumers; i++ {
        ctx = workflow.WithActivityOptions(ctx, ao)
        // 执行长时间运行的活动
        if err := workflow.ExecuteActivity(ctx, ConsumeMessagesActivity).Get(ctx, nil); err != nil {
            return err
        }
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

func ConsumeMessagesActivity(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Done...")
            return ctx.Err() // 停止活动
        default:
            ConsumeMessage()
        }
    }
}