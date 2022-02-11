package log

// AsSystemdLogger _
func AsSystemdLogger(l L) SystemdLogger {
	return &asl{
		L: l,
	}
}

type asl struct {
	L
}

func (a *asl) Error(v ...interface{}) error {
	a.L.Error(v...)
	return nil
}

func (a *asl) Warning(v ...interface{}) error {
	a.L.Warn(v...)
	return nil
}

func (a *asl) Info(v ...interface{}) error {
	a.L.Info(v...)
	return nil
}

func (a *asl) Errorf(format string, args ...interface{}) error {
	if lf, ok := a.L.(LF); ok {
		lf.Errorf(format, args...)
	}
	return nil
}

func (a *asl) Warningf(format string, args ...interface{}) error {
	if lf, ok := a.L.(LF); ok {
		lf.Warnf(format, args...)
	}
	return nil
}

func (a *asl) Infof(format string, args ...interface{}) error {
	if lf, ok := a.L.(LF); ok {
		lf.Infof(format, args...)
	}
	return nil
}
