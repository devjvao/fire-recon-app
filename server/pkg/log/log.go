package log

import (
	"github.com/mattn/go-colorable"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"path/filepath"
)

// InitLog sets the log output to the desired folder.
func InitLog(logPath string, level logrus.Level) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, "misc.log"),
		MaxSize:    2,
		MaxAge:     10,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   false,
	}
	mw := io.MultiWriter(colorable.NewColorableStdout(), lumberjackLogger)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		PadLevelText:           true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	logrus.SetLevel(level)
}
