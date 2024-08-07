package app

import (
	"go.temporal.io/sdk/workflow"
)

// ChildWorkflow executes a child workflow within a parent workflow in the Temporal system.
//
// Parameters:
// - ctx: The Temporal workflow context used to manage and control the execution of the child workflow.
//
// Returns:
// - An error if there was an issue executing the child workflow; otherwise, it returns nil to indicate successful execution.
func ChildWorkflow(ctx workflow.Context) (error) {
	return produceMessage(Producer)
}
