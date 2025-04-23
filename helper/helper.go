package helper

import (
	"log"
	"time"
)

func LogInfo(message string, context ...any) {
	log.Printf("INFO [%s] %s | %v\n", time.Now().Format(time.RFC3339), message, context)
}

func LogWarning(message string, context ...any) {
	log.Printf("WARNING [%s] %s | %v\n", time.Now().Format(time.RFC3339), message, context)
}

func LogError(message string, context ...any) {
	log.Printf("ERROR [%s] %s | %v\n", time.Now().Format(time.RFC3339), message, context)
}

func LogCritical(message string, context ...any) {
	log.Fatalf("CRITICAL [%s] %s | %v\n", time.Now().Format(time.RFC3339), message, context)
}
