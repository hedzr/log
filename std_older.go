//go:build !go1.13
// +build !go1.13

package log

import (
	"io"
	"os"
)

func (s *stdLogger) GetOutput() (out io.Writer) { return os.Stderr }
