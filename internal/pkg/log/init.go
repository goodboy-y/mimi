package log

const DebugLogLevel = 1

const InfoLogLevel = 2

const WarnLogLevel = 3

const ErrorLogLevel = 4

func Init(level Level) {
	std = &logger{Level: level}
}

func LogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		Init(DebugLogLevel)
	case "warn":
		Init(WarnLogLevel)
	case "error":
		Init(ErrorLogLevel)
	default:
		Init(InfoLogLevel)
	}
}
