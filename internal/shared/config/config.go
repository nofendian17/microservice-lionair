package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App         app         `json:"app" validate:"required"`
	Integration integration `json:"integration" validate:"required"`
	Storage     storage     `json:"storage"`
	Logger      logger      `json:"logger" validate:"required"`
}

type app struct {
	Environment        string `json:"environment" validate:"required"`
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description" validate:"required"`
	Port               int    `json:"port" validate:"required"`
	Version            string `json:"version" validate:"required"`
	HttpRequestTimeout int    `json:"httpRequestTimeout" validate:"required"`
	GrpcRequestTimeout int    `json:"grpcRequestTimeout" validate:"required"`
}

type integration struct {
	Credentials credentials `json:"credentials" validate:"required"`
	Url         string      `json:"url" validate:"required"`
	Target      string      `json:"target" validate:"required"`
	Service     service     `json:"service" validate:"required"`
}

type service struct {
	SessionCreate struct {
		Version string `json:"version" validate:"required"`
		Path    string `json:"path" validate:"required"`
	} `json:"sessionCreate"`
	SessionClose struct {
		Version string `json:"version" validate:"required"`
		Path    string `json:"path" validate:"required"`
	} `json:"sessionClose"`
	FlightMatrix struct {
		Version string `json:"version" validate:"required"`
		Path    string `json:"path" validate:"required"`
	} `json:"flightMatrix"`
	OTAAirSell struct {
		Version string `json:"version" validate:"required"`
		Path    string `json:"path" validate:"required"`
	} `json:"otaAirSell"`
}

type credentials struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Organization string `json:"organization" validate:"required"`
}

type storage struct {
	Redis redis `json:"redis"`
}

type redis struct {
	Address string `json:"address"`
	DB      int    `json:"db"`
}

type logger struct {
	FileLocation    string        `json:"fileLocation" validate:"required"`
	FileTdrLocation string        `json:"fileTdrLocation" validate:"required"`
	FileMaxAge      time.Duration `json:"fileMaxAge" validate:"required"`
	Stdout          bool          `json:"stdout"`
}

func New() *Config {
	path := "config.json"
	fmt.Println("Try NewConfig ... ", path)

	viper.SetConfigFile(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")

	// For test
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../../../configs")
	viper.AddConfigPath("../../../../configs")
	viper.AddConfigPath("../../../../../configs")
	viper.AddConfigPath("../../../../../../configs")
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("config loaded!")

	return &conf
}
