package application

import (
	"context"
	grpcServer "github.com/emptyhopes/employees/cmd/grpc_server"
	"github.com/emptyhopes/employees/internal/provider"
	employeeProvider "github.com/emptyhopes/employees/internal/provider/employees"
	"github.com/joho/godotenv"
)

type InterfaceApplication interface {
	InitializeDependency(context.Context) error
	InitializeEnvironment(context.Context) error
	InitializeProvider(context.Context) error
	Run()
}

type application struct {
	employeeProvider provider.InterfaceEmployeeProvider
}

var _ InterfaceApplication = (*application)(nil)

func NewApplication(ctx context.Context) (*application, error) {
	app := &application{}

	err := app.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *application) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.InitializeEnvironment,
		a.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *application) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (a *application) InitializeProvider(_ context.Context) error {
	a.employeeProvider = employeeProvider.NewEmployeeProvider()

	return nil
}

func (a *application) Run() {
	implementation := a.employeeProvider.GetEmployeeImplementation()

	grpcServer.Run(implementation)
}
