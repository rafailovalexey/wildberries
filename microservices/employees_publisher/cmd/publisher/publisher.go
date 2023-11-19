package publisher

import (
	"context"
	natsPublisher "github.com/emptyhopes/employees_publisher/cmd/nats_publisher"
	"github.com/emptyhopes/employees_publisher/internal/provider"
	employeeProvider "github.com/emptyhopes/employees_publisher/internal/provider/employees"
	"github.com/joho/godotenv"
)

type InterfacePublisher interface {
	InitializeDependency(context.Context) error
	InitializeEnvironment(context.Context) error
	InitializeProvider(context.Context) error
	Run()
}

type publisher struct {
	employeeProvider provider.InterfaceEmployeeProvider
}

var _ InterfacePublisher = (*publisher)(nil)

func NewPublisher(ctx context.Context) (*publisher, error) {
	pub := &publisher{}

	err := pub.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return pub, nil
}

func (p *publisher) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		p.InitializeEnvironment,
		p.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *publisher) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (p *publisher) InitializeProvider(_ context.Context) error {
	p.employeeProvider = employeeProvider.NewEmployeeProvider()

	return nil
}

func (p *publisher) Run() {
	controller := p.employeeProvider.GetEmployeeController()

	natsPublisher.Run(controller)
}
