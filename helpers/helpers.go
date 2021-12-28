package helpers

import (
	"errors"

	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/tchannel"
)

const (
	domainName = "victor"
)

func NewCadenceClient(workflowClient workflowserviceclient.Interface) client.Client {
	return client.NewClient(workflowClient, domainName, &client.Options{})
}

func NewWorkflowClient(serviceNameCadenceClient string, serviceNameCadenceFrontend string) (workflowserviceclient.Interface, error) {
	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(serviceNameCadenceClient))
	if err != nil {
		return nil, err
	}
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: serviceNameCadenceClient,
		Outbounds: yarpc.Outbounds{
			serviceNameCadenceFrontend: {Unary: ch.NewSingleOutbound("127.0.0.1:7933")},
		},
	})

	if dispatcher == nil {
		return nil, errors.New("failed to create dispatcher")
	}

	if err := dispatcher.Start(); err != nil {
		panic(err)
	}

	return workflowserviceclient.New(dispatcher.ClientConfig(serviceNameCadenceFrontend)), nil
}
