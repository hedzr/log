package closers_test

import (
	"github.com/hedzr/log/closers"
	"testing"
)

type redisHub struct {
}

func (s *redisHub) Close() {
	// close the connections to redis servers
	println("redis connections closed")
}

func TestCosers(t *testing.T) {

	defer closers.Close()

	closers.RegisterPeripheral(&redisHub{})

	closers.RegisterCloseFns(func() {
		// do some shutdown operations here
		println("close functor")
	})

	for _, ii := range closers.ClosersClosers() {
		println(ii)
	}

	closers.Closers().Close()
}
