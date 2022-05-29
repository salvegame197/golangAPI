package util

import "salvegame197/golangApi/src/infra/common/util/error"

// Interface Util interface
type Interface interface {
}

// Util provide utilities
type Util struct {
	Error *error.Error
}

// Initialize initialize utilities
func Initialize() *Util {
	return &Util{Error: error.New()}
}
