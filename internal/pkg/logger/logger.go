package logger

import (
	"github.com/Aoi-hosizora/ahlib-more/xlogrus"
	"github.com/Aoi-hosizora/ahlib-more/xrotation"
	"github.com/Aoi-hosizora/ahlib/xtime"
	"github.com/Aoi-hosizora/scut-academic-notifier/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"time"
)

var _logger *logrus.Logger

func Logger() *logrus.Logger {
	return _logger
}

func Setup() error {
	logger := logrus.New()
	level := logrus.WarnLevel
	if config.IsDebugMode() {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)
	logger.SetReportCaller(false)
	logger.SetFormatter(xlogrus.NewSimpleFormatter(
		xlogrus.WithTimestampFormat(time.RFC3339),
		xlogrus.WithTimeLocation(time.Local),
	))

	logName := config.Configs().Meta.LogName
	rotation, err := xrotation.New(
		logName+".%Y%m%d.log",
		xrotation.WithSymlinkFilename(logName+".current.log"),
		xrotation.WithRotationTime(24*time.Hour),
		xrotation.WithRotationMaxAge(15*24*time.Hour),
		xrotation.WithClock(xtime.Local),
	)
	if err != nil {
		return err
	}
	logger.AddHook(xlogrus.NewRotationHook(
		rotation,
		xlogrus.WithRotateLevel(level),
		xlogrus.WithRotateFormatter(xlogrus.RFC3339JsonFormatter()),
	))

	_logger = logger
	return nil
}
