// Package rfid provides structured tracing logs for Go routines.
package rfid

import "fmt"

func RFIDLog(queue, funcID, workerID, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("[%s][%s][%s] %s\n", queue, funcID, workerID, msg)
}

