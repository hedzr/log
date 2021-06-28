package closers

import (
	"fmt"
	"github.com/hedzr/log/basics"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func SetupCloseHandlerAndWait(wg *sync.WaitGroup, closers ...basics.Peripheral) {
	setupCloseHandler(closers...)
	wg.Wait()
}

func SetupCloseHandlerAndEnterLoop(closers ...basics.Peripheral) {
	enterLoop(setupCloseHandler(closers...))
}

// setupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupCloseHandler(onFinish ...basics.Peripheral) chan struct{} {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		for _, f := range onFinish {
			f.Close()
		}
		closers.Close()
		//os.Exit(0)
		close(done)
	}()
	return done
}

func enterLoop(done chan struct{}) {
	for {
		select {
		case <-done:
			return
		}
	}
}
