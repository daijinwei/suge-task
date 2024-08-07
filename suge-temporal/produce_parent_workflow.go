package app 

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

// ParentWorkflow initiates a workflow in the Temporal system.
//
// Parameters:
// - ctx: The Temporal workflow context, which is used to manage and control the execution of the workflow.
//
// Returns:
// - A string that represents the result or ID of the workflow execution.
// - An error if there was an issue starting the workflow.
func ParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)
	for i := 0; i < Count; i++ {
		childProduceWorkFlowID := fmt.Sprintf(ChildProduceWorkFlowFMT,  i)
		cwo := workflow.ChildWorkflowOptions{
			WorkflowID: childProduceWorkFlowID,
		}
		ctx = workflow.WithChildOptions(ctx, cwo)
		var result string
		err := workflow.ExecuteChildWorkflow(ctx, ChildWorkflow).Get(ctx, &result)
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return "", err
		}

		logger.Info("Parent execution completed.", "Result", result)
	}
	return "call ok", nil
}
