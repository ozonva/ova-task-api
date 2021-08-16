package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type configuration struct {
	ConfigVersion string `json:"configVersion"`
}

func ConfigCyclicReading(configFilePath string, configUpdatePeriodSeconds int, configUpdateHandle func(key string)) {
	var workingVersion string
	for {
		config, err := readConfig(configFilePath)
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

func readConfig(configFilePath string) (*configuration, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println("open config error: ", err)
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	decoder := json.NewDecoder(file)
	configuration := configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("config decode error: ", err)
		return nil, err
	}
	return &configuration, nil
}
