package datastore

import (
	"sync"

	"github.com/americanas-go/faas/datastore/aws/kinesis"
	"github.com/americanas-go/faas/datastore/aws/sns"
	"github.com/americanas-go/faas/datastore/aws/sqs"
	"github.com/americanas-go/faas/datastore/nats"
	"go.uber.org/fx"
)

var eventOnce sync.Once

func EventModule() fx.Option {

	options := fx.Options()

	eventOnce.Do(func() {

		value := EventProviderValue()

		switch value {
		case "kinesis":
			options = kinesis.Module()
		case "sns":
			options = sns.Module()
		case "sqs":
			options = sqs.Module()
		default:
			options = nats.Module()
		}

	})

	return options
}