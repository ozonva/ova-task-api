package utils

import (
	"encoding/json"
	"fmt"
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
			fmt.Println("read config error: ", err)
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
		fmt.Println("open config error: ", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("config close error: ", err)
		}
	}(file)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("config decode error: ", err)
		return nil, err
	}
	return &configuration, nil
}
