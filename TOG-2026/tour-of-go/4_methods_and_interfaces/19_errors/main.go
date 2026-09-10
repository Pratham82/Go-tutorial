package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type CustomError struct {
	ErrorName     string
	ErrorCode     int
	ErrorTitle    string
	ErrorSubtitle string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("ErrorCid: %v, Error Name: %s, Subtitle: %s, Title: %s", e.ErrorCode, e.ErrorName, e.ErrorSubtitle, e.ErrorTitle)
}

func runErrorLog() error {

	return &CustomError{
		"Error 1",
		404,
		"Not found",
		"Value not found",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	if err := runErrorLog(); err != nil {
		fmt.Println(err)
	}
}
