package application

import "github.com/Anti-Sh/go-rpn-calculator/internal/config"

type Application struct {
	config *config.Config
}

func NewApplication() *Application {
	return &Application{
		config: config.NewConfigFromEnv(),
	}
}

func (app *Application) Run() error {

	return nil
}

func (app *Application) RunServer() error {

	return nil
}
