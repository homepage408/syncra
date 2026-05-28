package logger

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
}

type defaultLogger struct {
	level string
}

func New(level string) Logger {
	return &defaultLogger{level: level}
}

func (l *defaultLogger) Info(message string) {
	println("[INFO] " + message)
}

func (l *defaultLogger) Error(message string) {
	println("[ERROR] " + message)
}

func (l *defaultLogger) Debug(message string) {
	println("[DEBUG] " + message)
}
