package main

import (
	"fmt"
	"os"
	"time"
)

var fileName = "deneme.txt"

func WriteFile(item chan []byte, doneChannel chan bool, errChannel chan error) {
	file, err := os.Create(fileName)
	if err != nil {
		errChannel <- err
	}
	for {
		select {
		case <-doneChannel:
			return
		case data := <-item:
			_, err := file.Write(data)
			if err != nil {
				errChannel <- err
			}
			errChannel <- nil
		}
	}
}

func main() {
	errChannel := make(chan error)
	doneChannel := make(chan bool)
	item := make(chan []byte)

	items := []string{"baran", " can", " atbaş"}

	go WriteFile(item, doneChannel, errChannel)
	for _, v := range items {
		time.Sleep(time.Second)
		item <- []byte(v)
		err := <-errChannel
		if err != nil {
			fmt.Println(err)
		}
	}

	doneChannel <- true
}
