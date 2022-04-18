package logs

type F map[string]any

type Printer func(msg string, data F)

type Logging struct {
	Debug Printer
	Info  Printer
	Error Printer
}

func Noop(msg string, data F) {
}

var Nolog = Logging{
	Debug: Noop,
	Error: Noop,
	Info:  Noop,
}
