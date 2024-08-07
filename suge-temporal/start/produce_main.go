package main

import (
    "context"
    "log"

    "go.temporal.io/sdk/client"

    "suge-message-temporal/app"
)

func main() {
    // Create the client object
    c, err := client.Dial(client.Options{
        HostPort: app.TemporalHstPort,
	})
    if err != nil {
        log.Fatalln("unable to create Temporal client", err)
    }
    defer c.Close()

    options := client.StartWorkflowOptions{
        ID:        app.ProduceWorkFlow,
        TaskQueue: app.MessageTaskQueue,
        CronSchedule: "* * * * *",	// Execute the workflow on average once per minute.
    }

    // Start the Workflow
    if _, err := c.ExecuteWorkflow(context.Background(), options, app.ParentWorkflow);  err != nil {
        log.Fatalln("unable to complete Workflow", err)
    }
}

