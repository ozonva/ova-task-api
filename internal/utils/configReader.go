package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type Configuration struct {
	ConfigVersion string `json:"configVersion"`
	Grpc          struct {
		Port int `json:"port"`
	} `json:"grpc"`
	Http struct {
		Port int `json:"port"`
	} `json:"http"`
	Db struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DataBase string `json:"dataBase"`
	} `json:"db"`
}

func ConfigCyclicReading(configFilePath string, configUpdatePeriodSeconds int, configUpdateHandle func(key string)) {
	var workingVersion string
	for {
		config, err := ReadConfig(configFilePath)
		if err != nil {
			log.Error().Err(err).Msg("read config error")
		} else {
			actualVersion := config.ConfigVersion
			if actualVersion != workingVersion {
				workingVersion = actualVersion
				configUpdateHandle(actualVersion)
			}
		}
		time.Sleep(time.Duration(configUpdatePeriodSeconds) * time.Second)
	}
}

func ReadConfig(configFilePath string) (*Configuration, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Error().Err(err).Msg("open config error")
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msg("config close error")
		}
	}(file)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Error().Err(err).Msg("config decode error")
		return nil, err
	}
	return &configuration, nil
}
