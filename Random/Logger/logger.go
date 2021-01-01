package main

import (
	"encoding/json"
	"fmt"
	"time"
	"os"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
)

func NewExample(){
	fmt.Printf("\n*** Using the Example logger\n\n")


	logger := zap.NewExample()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")  // would exit if uncommented
	logger.DPanic("This is a DPANIC message")
	//logger.Panic("This is a PANIC message")   // would exit if uncommented

	fmt.Println()

}

func NewDevelopment() {
	fmt.Printf("*** Using the Development logger\n\n")

	logger, _ := zap.NewDevelopment()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")   // would exit if uncommented
	// logger.DPanic("This is a DPANIC message") // would exit if uncommented
	//logger.Panic("This is a PANIC message")    // would exit if uncommented

	fmt.Println()
}

func NewProduction(){
	fmt.Printf("*** Using the Production logger\n\n")

	logger, _ := zap.NewProduction()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")   // would exit if uncommented
	logger.DPanic("This is a DPANIC message")
	// logger.Panic("This is a PANIC message")   // would exit if uncommented

	fmt.Println()


}

func SugarNewDevelopment() {
	fmt.Printf("*** Using the Sugar logger\n\n")

	logger, _ := zap.NewDevelopment()
	slogger := logger.Sugar()
	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s", "sprintf")
	slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)
	slogger.Debug("Debug() uses sprint")
	slogger.Warn("Warn() uses sprintf")
	slogger.Error("Error() allows tags")
	// slogger.Fatal("Fatal() allows tags")
	// slogger.DPanic("DPanic() allows tags")
	// slogger.Panic("Panic() allows tags")
}

func JsonLogger(){
	fmt.Println("*** Build a logger from a json ****")

	rawJSONConfig := []byte(`{
      "level": "info",
      "encoding": "console",
      "outputPaths": ["stdout", "/tmp/logs"],
      "errorOutputPaths": ["/tmp/errorlogs"],
      "initialFields": {"initFieldKey": "fieldValue"},
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "nameKey": "logger",
        "timeKey": "time",
        "callerKey": "logger",
        "stacktraceKey": "stacktrace",
        "callstackKey": "callstack",
        "errorKey": "error",
        "timeEncoder": "iso8601",
        "fileKey": "file",
        "levelEncoder": "capitalColor",
        "durationEncoder": "second",
        "callerEncoder": "full",
        "nameEncoder": "full",
        "sampling": {
            "initial": "3",
            "thereafter": "10"
        }
      }
    }`)

	config := zap.Config{}
	if err := json.Unmarshal(rawJSONConfig, &config); err != nil {
		panic(err)
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger.Debug("This is a DEBUG message")
	logger.Info("This should have an ISO8601 based time stamp")
	logger.Warn("This is a WARN message")
	logger.Error("This is an ERROR message")
	// logger.Fatal("This is a FATAL message")   // would exit if uncommented
	logger.DPanic("This is a DPANIC message") // would exit if uncommented

	const url = "http://example.com"
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))
	// {"level":"info","msg":"Failed to fetch URL.","url":"http://example.com","attempt":3,"backoff":"1"}
}

func GlobalLogger() {
	fmt.Printf("\n*** Using the global logger out of the box\n\n")

	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After replacing the global logger with a development logger\n\n")
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After replacing the global logger with a production logger\n\n")
	logger, _ = zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	zap.S().Infow("An info message", "iteration", 1)

	fmt.Printf("\n*** After undoing the last replacement of the global logger\n\n")
	undo()
	zap.S().Infow("An info message", "iteration", 1)

}

func CustomLogger(){
	/*
		// This would cause a panic: "panic: no encoder name specified"
		cfg := zap.Config{}
		logger, err := cfg.Build()
		if err != nil {
			panic(err)
		}
	*/

	/*
		// causes a panic: "panic: runtime error: invalid memory address or nil pointer dereference"
		logger, err := zap.Config{Encoding: "json"}.Build()
	*/

	/*
		// No output
		logger, err := zap.Config{Encoding: "json", Level: zap.NewAtomicLevelAt(zapcore.InfoLevel)}.Build()
	*/

	fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, no key specified\n\n")

	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
	}.Build()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

	fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, message key only specified\n\n")

	logger, _ = zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
		},
	}.Build()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

	fmt.Printf("\n*** Using a JSON encoder, at debug level, sending output to stdout, all possible keys specified\n\n")

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ = cfg.Build()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Info("This is an INFO message with fields", zap.String("region", "us-west"), zap.Int("id", 2))

	fmt.Printf("\n*** Same logger with console logging enabled instead\n\n")

	logger.WithOptions(
		zap.WrapCore(
			func(zapcore.Core) zapcore.Core {
				return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
			})).Info("This is an INFO message")
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func CustomMenCoder() {
	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}

	fmt.Printf("\n*** Using standard ISO8601 time encoder\n\n")

	// avoiding copying of atomic values
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, _ := cfg.Build()
	logger.Info("This should have an ISO8601 based time stamp")

	fmt.Printf("\n*** Using a custom time encoder\n\n")

	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a syslog style time stamp")

	fmt.Printf("\n*** Using a custom level encoder\n\n")

	cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder

	logger, _ = cfg.Build()
	logger.Info("This should have a interesting level name")
}

func main() {
	// NewExample()
	// NewDevelopment()
	// NewProduction()
	// SugarNewDevelopment()
	// JsonLogger()
	// GlobalLogger()
	// CustomLogger()
	CustomMenCoder()
}