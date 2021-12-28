package main

import (
	"context"
	"time"

	"github.com/VictorCavalcanteLG/Cadence-teste/helpers"
	"github.com/VictorCavalcanteLG/Cadence-teste/workflows"
	"go.uber.org/cadence/client"
)

const (
	serviceNameCadenceClient   = "cadence-client"
	serviceNameCadenceFrontend = "cadence-frontend"
	domainName                 = "victor"
)

func main() {

	workflowClient, err := helpers.NewWorkflowClient(serviceNameCadenceClient, serviceNameCadenceFrontend)
	if err != nil {
		panic(err)
	}

	triggerClient := helpers.NewCadenceClient(workflowClient)

	_, err = triggerClient.StartWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:                           "workflowID",
		TaskList:                     "pocTasklist",
		ExecutionStartToCloseTimeout: 1 * time.Second,
	}, workflows.Hello)

	if err != nil {
		panic(err)
	}
}
