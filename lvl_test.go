// Copyright © 2020 Hedzr Yeh.

package log

import (
	"testing"
)

func TestLevels(t *testing.T) {
	for _, l := range AllLevels {
		t.Logf("level: %v", l)
	}

	for _, x := range []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC", "OFF", "XX"} {
		l, err := ParseLevel(x)
		if err == nil {
			t.Logf("level: %s => %v", x, l)
		} else if x != "XX" {
			t.Fatal(err)
		}

		var lex Level
		var ll = &lex
		if err = ll.UnmarshalText([]byte(x)); err != nil && x != "XX" {
			t.Fatal(err)
		}
	}

	l := Level(1000)
	t.Logf("level: %v => %v", 1000, l)
	if l.String() != "unknown" {
		t.Fatalf("expect 'unknown' level but got %q", l.String())
	}
}

func TestLog(t *testing.T) {
	_ = NewLoggerConfig()

	//var rootCmdX = &RootCommand{
	//	Command: Command{
	//		BaseOpt: BaseOpt{
	//			Name: "consul-tags",
	//		},
	//	},
	//}
	//
	//for _, x := range []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC", "OFF", "XX"} {
	//	Set("logger.level", x)
	//	_ = internalGetWorker().getWithLogexInitializor(DebugLevel)(&rootCmdX.Command, []string{})
	//}
	//
	//Set("logger.target", "journal")
	//Set("logger.format", "json")
	//_ = internalGetWorker().getWithLogexInitializor(DebugLevel)(&rootCmdX.Command, []string{})
}
