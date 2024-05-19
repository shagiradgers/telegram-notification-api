package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Config interface {
	GetServerHost() (string, error)
	MustGetServerHost() string

	GetServerPort() (int, error)
	MustGetServerPort() int

	GetDatabaseConnectionString() (string, error)
	MustGetDatabaseConnectionString() string

	GetTelegramBotToken() (string, error)
	MustGetTelegramBotToken() string
}

type config struct {
	env        envValue
	projectDir string
	v          *viper.Viper
}

type configValue string

const (
	ServerHostValue          configValue = "server_host"
	ServerPortValue          configValue = "server_port"
	DatabaseConnectionString configValue = "database_connection_string"
	TelegramBotTokenString   configValue = "telegram_bot_token"
)

type envValue int

const (
	LocalEnv envValue = iota
	ProdEnv
)

func (c *config) envValueToConfigPath(env envValue) string {
	return map[envValue]string{
		LocalEnv: c.projectDir + "/.env/local_values.yml",
		ProdEnv:  c.projectDir + "/.env/local_values.yml",
	}[env]
}

func (c *config) GetServerHost() (string, error) {
	const op = "config.GetServerHost"
	v, err := c.getValueFromConfig(ServerHostValue)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return v.(string), nil
}

func (c *config) MustGetServerHost() string {
	v, err := c.getValueFromConfig(ServerHostValue)
	if err != nil {
		panic(err)
	}
	return v.(string)
}

func (c *config) GetServerPort() (int, error) {
	const op = "config.GetServerPort"
	v, err := c.getValueFromConfig(ServerPortValue)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return v.(int), err
}

func (c *config) MustGetServerPort() int {
	v, err := c.getValueFromConfig(ServerPortValue)
	if err != nil {
		panic(err)
	}
	return v.(int)
}

func (c *config) GetDatabaseConnectionString() (string, error) {
	const op = "config.GetDatabaseConnectionString"
	v, err := c.getValueFromConfig(DatabaseConnectionString)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return v.(string), err
}

func (c *config) MustGetDatabaseConnectionString() string {
	v, err := c.getValueFromConfig(DatabaseConnectionString)
	if err != nil {
		panic(err)
	}
	return v.(string)
}

func (c *config) GetTelegramBotToken() (string, error) {
	const op = "config.GetTelegramBotToken"
	v, err := c.getValueFromConfig(TelegramBotTokenString)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return v.(string), err
}

func (c *config) MustGetTelegramBotToken() string {
	v, err := c.getValueFromConfig(TelegramBotTokenString)
	if err != nil {
		panic(err)
	}
	return v.(string)
}

func (c *config) getValueFromConfig(val configValue) (any, error) {
	if c == nil {
		return "", errors.New("struct is nil")
	}
	if val == "" {
		return nil, errors.New("val is empty")
	}
	return c.v.Get(string(val)), nil
}

func NewConfig(env envValue) (Config, error) {
	const op = "config.NewConfig"

	c := &config{
		env:        env,
		v:          viper.New(),
		projectDir: "/Users/smingaraev/GolandProjects/telegram-notification-api",
	}
	c.v.SetConfigFile(c.envValueToConfigPath(env))
	if err := c.v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return c, nil
}
