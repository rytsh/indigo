package common

import (
	"fmt"
	"log"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("[2006-01-02T15:04:05-0700]") + " " + string(bytes))
}

// SetCustomLog set usage
func SetCustomLog() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}
