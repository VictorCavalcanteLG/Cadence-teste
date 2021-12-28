package main

import (
	"github.com/VictorCavalcanteLG/Cadence-teste/helpers"
	"github.com/VictorCavalcanteLG/Cadence-teste/workflows"

	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
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

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	w := worker.New(workflowClient, domainName, "pocTasklist",
		worker.Options{
			Logger: logger,
		})

	workflow.Register(workflows.Hello)

	err = w.Run()
	if err != nil {
		panic(err)
	}
}
