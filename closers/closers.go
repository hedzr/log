package closers

import (
	"github.com/hedzr/log/basics"
	"sync/atomic"
)

// RegisterPeripheral adds a peripheral into our global closers set.
// a basics.Peripheral object is a closable instance.
func RegisterPeripheral(servers ...basics.Peripheral) { closers.RegisterPeripheral(servers...) }

// RegisterCloseFns adds a simple closure into our global closers set
func RegisterCloseFns(fns ...func()) { closers.RegisterCloseFns(fns...) }

// Close will cleanup all registered closers.
// You must make a call to Close before your app shutting down. For example:
//
//     func main() {
//         defer closers.Close()
//         // ...
//     }
func Close() {
	closers.Close()
}

// Closers returns the closers set as a basics.Peripheral
func Closers() basics.Peripheral { return closers }

// ClosersClosers returns the closers set as a basics.Peripheral array
func ClosersClosers() []basics.Peripheral { return closers.closers }

var closers = new(c)

type c struct {
	closers []basics.Peripheral
	closed  int32
}

func (s *c) RegisterPeripheral(servers ...basics.Peripheral) {
	s.closers = append(s.closers, servers...)
}

func (s *c) RegisterCloseFns(fns ...func()) {
	s.closers = append(s.closers, &w{fns})
}

type w struct {
	fns []func()
}

func (s *w) Close() {
	for _, c := range s.fns {
		if c != nil {
			c()
		}
	}
}

func (s *c) Close() {
	if atomic.CompareAndSwapInt32(&s.closed, 0, 1) {
		for _, c := range s.closers {
			// log.Debugf("  c.Close(), %v", c)
			if c != nil {
				c.Close()
			}
		}
	}
}
