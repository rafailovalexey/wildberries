package subscriber

import (
	"context"
	"github.com/emptyhopes/employees/internal/provider"
	employeeProvider "github.com/emptyhopes/employees_subscriber/internal/provider/employees"
	"github.com/joho/godotenv"
)

type InterfaceSubscriber interface {
	InitializeDependency(context.Context) error
	InitializeEnvironment(context.Context) error
	InitializeProvider(context.Context) error
	Run()
}

type subscriber struct {
	employeeProvider provider.InterfaceEmployeeProvider
}

var _ InterfaceSubscriber = (*subscriber)(nil)

func NewSubscriber(ctx context.Context) (*subscriber, error) {
	sub := &subscriber{}

	err := sub.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *subscriber) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		s.InitializeEnvironment,
		s.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *subscriber) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (s *subscriber) InitializeProvider(_ context.Context) error {
	s.employeeProvider = employeeProvider.NewEmployeeProvider()

	return nil
}

func (s *subscriber) Run() {
	controller := s.employeeProvider.GetEmployeeController()

	natsSubscriber.Run(controller)
}
