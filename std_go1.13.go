//go:build go1.13
// +build go1.13

package log

import (
	"io"
	"log"
)

func (s *stdLogger) GetOutput() (out io.Writer) { return log.Writer() }
