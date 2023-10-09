package config

import (
    "github.com/spf13/viper"
    "log"
)

type AppConfig struct {
    AppName     string
    LogLevel    string
    HTTPPort    string
    Version		string
}

var AppConfigInstance *AppConfig

func init() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    AppConfigInstance = &AppConfig{
        AppName:     viper.GetString("APP_NAME"),
        LogLevel:    viper.GetString("LOG_LEVEL"),
        HTTPPort:    viper.GetString("HTTP_PORT"),
        Version: viper.GetString("VERSION"),
    }
}
