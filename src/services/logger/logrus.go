package logger

import (
	"github.com/natefinch/lumberjack"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
	logger *logrus.Logger
}

func (l logger) ReturnLogrus() *logrus.Logger {
	return l.logger
}

type Logger interface {
	GeneralInfo(info string)
	GeneralError(info string, err error)
	ReturnLogrus() *logrus.Logger
}

type Event struct {
	id      int
	message string
}

var (
	generalInfoMessage  = Event{1, "General --- "}
	generalErrorMessage = Event{2, "General error ---- "}
)

func NewLogrusLogger() Logger {

	var baseLogger = logrus.New()
	baseLogger.Formatter = &logrus.JSONFormatter{}

	//Options: Debug, Info
	if os.Getenv("LOGLEVEL") == "debug" {
		baseLogger.Level = logrus.DebugLevel
	}
	if os.Getenv("LOGLEVEL") == "info" {
		baseLogger.Level = logrus.InfoLevel
	}
	path := os.Getenv("DIR_PATH") + os.Getenv("LOG_PATH")

	year, month, day := time.Now().Date()
	f, err := os.OpenFile(path+"/log_"+strconv.Itoa(day)+"_"+month.String()+"_"+strconv.Itoa(year), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		path, err = os.Getwd()
		f, err = os.OpenFile(path+"/log_file", os.O_WRONLY|os.O_CREATE, 0755)
	}
	mSize, _ := strconv.Atoi(os.Getenv("MAX_SIZE"))
	mBackups, _ := strconv.Atoi(os.Getenv("MAX_BACKUPS"))
	mAge, _ := strconv.Atoi(os.Getenv("MAX_AGE"))

	if err == nil {
		baseLogger.SetOutput(&lumberjack.Logger{
			Filename:   f.Name(),
			MaxSize:    mSize,    // megabytes after which new file is created
			MaxAge:     mAge,     //days
			MaxBackups: mBackups, // number of backups
		})
	}

	return &logger{
		baseLogger,
	}
}

func (l logger) GeneralInfo(info string) {
	l.logger.WithField("ID", generalInfoMessage.id).
		Debug(generalInfoMessage.message, info)
}
func (l logger) GeneralError(info string, err error) {
	l.logger.WithField("ID", generalErrorMessage.id).
		WithField("Info", info).Errorf(generalErrorMessage.message, err)
}
