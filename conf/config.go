package conf

import (
    "os"
    "time"

    "gopkg.in/yaml.v3"
)

type Config struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`

    JWT struct {
        Secret            string `yaml:"secret"`
        ExpirationMinutes int    `yaml:"expiration_minutes"`
    } `yaml:"jwt"`

    Database struct {
        DSN string `yaml:"dsn"`
    } `yaml:"database"`

    Route struct {
        PublicRoutes []string `yaml:"public_routes"`
    } `yaml:"-"`
}

var AppConfig *Config

func LoadConfig(appFile, routeFile string) error {
    data, err := os.ReadFile(appFile)
    if err != nil {
        return err
    }
    conf := &Config{}
    if err := yaml.Unmarshal(data, conf); err != nil {
        return err
    }

    routeData, err := os.ReadFile(routeFile)
    if err != nil {
        return err
    }
    var routes struct {
        PublicRoutes []string `yaml:"public_routes"`
    }
    if err := yaml.Unmarshal(routeData, &routes); err != nil {
        return err
    }
    conf.Route = routes

    AppConfig = conf
    return nil
}

func GetJWTDuration() time.Duration {
    return time.Duration(AppConfig.JWT.ExpirationMinutes) * time.Minute
}
