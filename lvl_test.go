// Copyright © 2020 Hedzr Yeh.

package log

import (
	"fmt"
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

	Trace("11", 11)
	Debug("11", 11)
	Info("11", 11)
	Warn("11", 11)
	Error("11", 11)
	tztf1() // Fatal("11", 11)
	tztp1() // Panic("11", 11)
	Print("11", 11)
	Println("11", 11)

	Tracef("t9.%v", 11)
	Debugf("t9.%v", 11)
	Infof("t9.%v", 11)
	Warnf("t9.%v", 11)
	Errorf("t9.%v", 11)
	tttf1() // Fatalf("t9.%v", 11)
	tttp1() // Panicf("t9.%v", 11)
	Printf("t9.%v", 11)

	SetLogger(NewDummyLogger())
	AsLogger(AsL(GetLogger()))
	SetLevel(DebugLevel)
	GetLevel()

	if InDebugging() {
		println(MinimalEnv())
	}

	println(MinimalEnv())

	SetDebugMode(true)
	SetTraceMode(true)
}

func tztf1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	Fatal("panic")
}

func tztp1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	Panic("panic")
}

func tttf1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	Fatalf("panic")
}

func tttp1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("caught: %v\n", err)
		}
	}()

	Panicf("panic")
}
