package internal

import (
	"fmt"
	"time"
)

func Logger(level string, log string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), level, log)
}
