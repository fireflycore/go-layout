package logger

import "microservice-go/utils/logger/internal"

type Entrance struct {
}

// Info The normal log
func (Entrance) Info(log string) {
	internal.Logger("info", log)
}

// Success The result is a success log
func (Entrance) Success(log string) {
	internal.Logger("success", log)
}

// Warning The result is abnormal but does not affect running logs
func (Entrance) Warning(log string) {
	internal.Logger("warning", log)
}

// Error The result is abnormal and affects program operation
func (Entrance) Error(log string) {
	internal.Logger("error", log)
}
