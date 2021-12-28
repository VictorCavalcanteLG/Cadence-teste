package workflows

import (
	"fmt"

	"go.uber.org/cadence/workflow"
)

func Hello(ctx workflow.Context) error {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Hello world")
	fmt.Println("------------------------------------------------------------------")

	return nil
}
