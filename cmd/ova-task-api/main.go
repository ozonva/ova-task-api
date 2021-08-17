package main

import (
	"fmt"
	"ozonva/ova-task-api/internal/utils"
)

const configFilePath = "configs/ova-task-api.config"
const configUpdatePeriodSeconds = 1

func main() {
	fmt.Println("Hey! I am the \"ova-task-api\" service")
	utils.ConfigCyclicReading(configFilePath, configUpdatePeriodSeconds, configUpdateHandle)
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
