package log

import "testing"

func TestNewLoggerConfig(t *testing.T) {
	log := NewDummyLogger()
	log.Tracef("")
	log.Debugf("")
	log.Infof("")
	log.Warnf("")
	log.Errorf("")
	log.Fatalf("")
	log.Printf("")
	log.SetLevel(InfoLevel)
	_ = log.GetLevel()
	log.Setup()
}
