package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sparrc/go-ping"
)

const filename = "internetlog.txt"

func main() {
	eventTime := time.Now().String()
	message := fmt.Sprintf("%s: Monitoring Internet every 30 seconds.\n", eventTime)
	appendMessageToFile(message)
	fmt.Printf(message)
	watchInternet()
}

func watchInternet() {
	previousInternetState := true
	currentInternetState := true
	for {
		currentInternetState = isThereInternet()
		if previousInternetState != currentInternetState {
			eventTime := time.Now().String()
			var message string
			if currentInternetState == true {
				message = fmt.Sprintf("Regained Internet at: %s\n", eventTime)
			} else {
				message = fmt.Sprintf("Lost Internet at: %s\n", eventTime)
			}
			appendMessageToFile(message)
			fmt.Printf(message)
		}
		previousInternetState = currentInternetState
		time.Sleep(30 * time.Second)
	}
}

func appendMessageToFile(message string) {
	// Create file if it does not exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Unable to create :%s\n", filename)
		}
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Unable to open file: %s\n", err)
	}
	_, err = f.WriteString(message)
	if err != nil {
		fmt.Printf("Unable to write to file: %s\n", err)
	}
}

func isThereInternet() bool {
	pinger, err := ping.NewPinger("8.8.8.8")
	if err != nil {
		log.Printf("Unable to ping for some reason: %s\n", err)
		return false
	}
	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketsRecv == 1 {
		return true
	}
	return false
}
