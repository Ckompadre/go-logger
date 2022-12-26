package loggger

import (
	"fmt"
	"time"
)

const (
	reset  string = "\033[0m"
	black  string = "\033[30m"
	red    string = "\033[31m"
	yellow string = "\033[33m"
	white  string = "\033[97m"
	bgred  string = "\033[41m\033[1m"
)

func timestamp() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.0000")
}

// Displays an information message
func Info(str string) {
	fmt.Print(white)
	fmt.Printf("I: [%s]: %s\n", timestamp(), str)
	fmt.Print(reset)
}

// Displays a warning message
func Warning(str string) {
	fmt.Print(yellow)
	fmt.Printf("W: [%s]: %s\n", timestamp(), str)
	fmt.Print(reset)
}

// Displays an error message
func Error(str string) {
	fmt.Print(red)
	fmt.Printf("E: [%s]: %s\n", timestamp(), str)
	fmt.Print(reset)
}

// Displays a fatal error message
func Fatal(str string) {
	fmt.Print(bgred)
	fmt.Printf("F: [%s]: %s\n", timestamp(), str)
	fmt.Print(reset)
}
