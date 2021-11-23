package ovo

// EnvironmentType is type for environment.
type EnvironmentType int8

// Available options for EnvironmentType.
const (
	Sandbox EnvironmentType = iota
	Production
)

var envURL = map[EnvironmentType]string{
	Sandbox:    "https://api.byte-stack.net/pos",
	Production: "https://api.ovo.id/pos",
}

var envLog = map[EnvironmentType]LogLevel{
	Sandbox:    LogDebug,
	Production: LogError,
}
