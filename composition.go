package main

import (
	"fmt"
	"time"
)

type BaseLogger struct {
	ServiceName string
}

func (b BaseLogger) Log(message string) {
	fmt.Printf("[%s] [%s] %s\n", time.Now().Format(time.RFC3339), b.ServiceName, message)
}

type FileLogger struct {
	BaseLogger //composition
	FilePath   string
}

func (f FileLogger) LogToFile(message string) {

	// In real life, write to a file instead of fmt.Println
	fmt.Printf("Writing log to %s: %s\n", f.FilePath, message)
}

type CloudLogger struct {
	BaseLogger
	CloudService string
}

func (c CloudLogger) LogToCloud(message string) {
	// In real life, push to AWS/Splunk
	fmt.Printf("Pushing log to %s: %s\n", c.CloudService, message)
}
