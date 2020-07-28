// Copyright © 2020 Hedzr Yeh.

package log

type dummyLogger struct{}

// NewDummyLogger return a dummy logger
func NewDummyLogger() Logger {
	return &dummyLogger{}
}

func (d *dummyLogger) Tracef(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Debugf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Infof(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Warnf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Errorf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Fatalf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Printf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) SetLevel(lvl Level) {
	//panic("implement me")
}

func (d *dummyLogger) GetLevel() Level {
	//panic("implement me")
	return InfoLevel // logex.GetLevel()
}

func (d *dummyLogger) Setup() {
	//panic("implement me")
}
