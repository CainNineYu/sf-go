package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var Logger *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}
	defer logger.Sync()
	Logger = logger
}
func Setlogs(logLevel zapcore.Level) {
	defer func() {
		if err := recover(); err != nil {
			Logger.Fatal("Setlogs Error", zap.Any("error", err))
		}
	}()
	//Set basic log formats
	//CapitalLevelEncoder:  Converts the level to uppercase
	encoder := createEncoder("level", zapcore.CapitalLevelEncoder)
	encoderDB := createEncoder("", nil)
	console := createEncoder("level", zapcore.CapitalLevelEncoder)

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.WarnLevel && lvl >= logLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return (lvl > zapcore.WarnLevel && lvl != zapcore.DPanicLevel) && lvl >= logLevel
	})
	dpanicLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DPanicLevel
	})

	infoWriter := getWriter("./logs/info.logs")
	errorWriter := getWriter("./logs/error.logs")
	dbWriter := getWriter("./logs/database.logs")

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		zapcore.NewCore(encoderDB, zapcore.AddSync(dbWriter), dpanicLevel),
		zapcore.NewCore(console,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel),
	)
	if logLevel > zap.WarnLevel {
		Logger = zap.New(core, zap.AddCaller())
	} else {
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(logLevel))
	}

}

func getWriter(filename string) io.Writer {

	// Save the logs generated within seven days and split the logs every day
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		Logger.Fatal("getWriter Error", zap.Error(err))
	}
	return hook
}

func createEncoder(level string, encodeLevel zapcore.LevelEncoder) zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    level,
		EncodeLevel: encodeLevel,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:      "file",
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { enc.AppendInt64(int64(d) / 1000000) },
	})
}
