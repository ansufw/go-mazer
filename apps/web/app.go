package web

import (
	"fmt"
	"log"

	"github.com/ansufw/go-mazer/apps/web/router"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode mode   `mapstructure:"mode"`
}

type RootConfig struct {
	Services map[string]ServiceConfig `mapstructure:"app"`
}

type App struct {
	app    *fiber.App
	config *ServiceConfig
}

type mode string

const (
	ModeDev  mode = "dev"
	ModeProd mode = "prod"
)

func loadConfig(serviceName string) (*ServiceConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var rootConfig RootConfig
	if err := viper.Unmarshal(&rootConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	serviceConfig, exists := rootConfig.Services[serviceName]
	if !exists {
		return nil, fmt.Errorf("service '%s' not found in config", serviceName)
	}

	return &serviceConfig, nil
}

func New(serviceName string) (*App, error) {
	config, err := loadConfig(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	app := fiber.New()

	router.Route(app)

	return &App{app: app, config: config}, nil
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.config.Port)

	// DELME debug
	if a.config.Mode == ModeDev {
		log.Println("DEBUG: Running in dev mode")
	}
	if a.config.Mode == ModeProd {
		log.Println("DEBUG: Running in prod mode")
	}

	return a.app.Listen(addr)
}
