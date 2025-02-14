package ports

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}
