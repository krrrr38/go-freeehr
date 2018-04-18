package freeehr

import (
	"os"
)

var debugEnable = os.Getenv("FREEE_DEBUG") == "1"

// DebugEnable return debug is enable or not
func DebugEnable() bool {
	return debugEnable
}
