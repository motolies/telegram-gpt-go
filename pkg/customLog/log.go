package customLog

import (
	"github.com/fatih/color"
	"log"
)

type Level string

const (
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

func ColorLog(message string, logLevel Level) {
	switch logLevel {
	case WARN:
		log.Println(color.YellowString("Warn: " + message))
	case ERROR:
		log.Println(color.RedString("Error: " + message))
	default:
		println(message)
	}
}
