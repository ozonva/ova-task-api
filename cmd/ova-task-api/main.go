package main

import (
	"fmt"
	"sync"
	"time"
)

const configFilePath = "configs/ova-task-api.config"
const configUpdatePeriodSeconds = 1

func main() {
	timer := time.NewTicker(5 * time.Nanosecond)
	var buffer = make(chan int, 10)

	var wait = sync.WaitGroup{}
	wait.Add(1)

	go func() {
		defer wait.Done()
		time.Sleep(3 * time.Second)
		for {
			select {
			case entity, ok := <-buffer:
				fmt.Println(entity, ok)
				if ok == false {
					//return
				}
			case <-timer.C:
				fmt.Println("Timeout")
				break
			}
		}
	}()

	for i := 0; i >= 0; i++ {
		buffer <- i
		if i == 10 {
			close(buffer)
			break
		}
	}
	wait.Wait()
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
