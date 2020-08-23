// +build debug

package debug

import (
	"fmt"
	"log"
)

func Printf(format string, args ...interface{}) {
	log.Printf(fmt.Sprintf("DEBUG: %s", format), args...)
}
