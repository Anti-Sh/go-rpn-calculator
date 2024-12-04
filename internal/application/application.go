package application

import "os"

type Config struct {
	Port string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func NewApplication() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (app *Application) Run() error {

	return nil
}

func (app *Application) RunServer() error {

	return nil
}
