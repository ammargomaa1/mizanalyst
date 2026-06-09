package mizanlyst_logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func init() {
	log.SetFlags(0) // We handle our own formatting
}

// Log writes a formatted log message prefixed with a timestamp and [MIZANLYST] tag.
// It supports all standard fmt formatting verbs.
//
// Usage:
//
//	mizanlyst_logger.Log("there's an error with id %d and message is %s and object is %+v", x, y, z)
func Log(format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	caller := getCallerInfo()
	message := fmt.Sprintf(format, args...)
	log.Printf("[MIZANLYST] %s | %s | %s", timestamp, caller, message)
}

// getCallerInfo returns the file and line number of the caller.
func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}

	// Extract just the last two path segments for readability
	parts := strings.Split(file, "/")
	if len(parts) > 2 {
		file = strings.Join(parts[len(parts)-2:], "/")
	}

	return fmt.Sprintf("%s:%d", file, line)
}
