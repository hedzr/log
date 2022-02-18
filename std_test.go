package log

import (
	"testing"
)

func TestTopLevel(t *testing.T) {
	Printf("hello")
	Infof("hello info")
	Warnf("hello warn")
	Errorf("hello error")
	Debugf("hello debug")
	Tracef("hello trace")
}

// TestNormal used logger directly so the caller is wrong.
// To get the right filename and line number of caller, see also TestTopLevel
func TestNormal(t *testing.T) {
	// config := log.NewLoggerConfigWith(true, "logrus", "trace")
	// logger := logrus.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}

func TestStdLogger(t *testing.T) {
	log := newStdLogger()
	log.SetLevel(TraceLevel)
	log.Tracef("trace")
	log.Debugf("debug")
	log.Infof("info")
	log.Warnf("warn")
	log.Errorf("error")
	// log.Fatalf("")
	l := AsL(log)
	l.Trace("trace")
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")

	log = newStdLoggerWith(TraceLevel)
	log.Printf("")
	log.SetLevel(InfoLevel)
	_ = log.GetLevel()
	log.Setup()
	tf(log)
	tp(log)
}
