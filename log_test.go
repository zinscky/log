package log

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	log := New(Info, "test-app")

	log.Debug("DEBUG LOG")
	log.Info("INFO LOG")
	log.Warn("WARN LOG")
	log.Error("ERROR LOG")

	fmt.Println(log.String())
}
