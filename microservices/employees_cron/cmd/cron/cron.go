package cron

import (
	"context"
	cronScheduler "github.com/emptyhopes/employees_cron/cmd/cron_scheduler"
	"github.com/emptyhopes/employees_cron/internal/provider"
	employeeProvider "github.com/emptyhopes/employees_cron/internal/provider/employees"
	"github.com/joho/godotenv"
)

type InterfaceCron interface {
	InitializeDependency(context.Context) error
	InitializeEnvironment(context.Context) error
	InitializeProvider(context.Context) error
	Run()
}

type cron struct {
	employeeProvider provider.InterfaceEmployeeProvider
}

var _ InterfaceCron = (*cron)(nil)

func NewCron(ctx context.Context) (*cron, error) {
	c := &cron{}

	err := c.InitializeDependency(ctx)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *cron) InitializeDependency(ctx context.Context) error {
	inits := []func(context.Context) error{
		c.InitializeEnvironment,
		c.InitializeProvider,
	}

	for _, function := range inits {
		err := function(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cron) InitializeEnvironment(_ context.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	return nil
}

func (c *cron) InitializeProvider(_ context.Context) error {
	c.employeeProvider = employeeProvider.NewEmployeeProvider()

	return nil
}

func (c *cron) Run() {
	service := c.employeeProvider.GetEmployeeService()

	cronScheduler.Run(service)
}
