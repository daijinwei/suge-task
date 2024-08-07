package main

import (
    "log"

    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"

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

    // This worker hosts both Workflow and Activity functions
    w := worker.New(c, app.MessageTaskQueue, worker.Options{})
    w.RegisterWorkflow(app.ParentWorkflow)
    w.RegisterWorkflow(app.ChildWorkflow)

    // Register consumer workflow
    w.RegisterWorkflow(app.ConsummerWorkFlow)

    // Start listening to the Task Queue
    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("unable to start Worker", err)
    }
}