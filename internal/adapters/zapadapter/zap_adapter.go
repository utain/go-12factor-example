package zapadapter

import (
	"github.com/utain/go/example/internal/logs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapAdapter() logs.Logging {
	logger, _ := zap.NewProduction(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
	defer logger.Sync()

	return logs.Logging{
		Debug: PrinterFactory(logger.Debug),
		Info:  PrinterFactory(logger.Info),
		Error: PrinterFactory(logger.Error),
	}
}

type zapPrinter func(msg string, fields ...zapcore.Field)

func PrinterFactory(printer zapPrinter) logs.Printer {
	return func(msg string, data logs.F) {
		fields := make([]zapcore.Field, len(data))
		i := 0
		for k, v := range data {
			fields[i] = zap.Any(k, v)
			i++
		}
		printer(msg, fields...)
	}
}
