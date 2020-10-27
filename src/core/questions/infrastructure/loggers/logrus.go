package loggers

import (
	"../../domain"
	"github.com/sirupsen/logrus"
)

type questionLogger struct {
	logger *logrus.Logger
}

func (q questionLogger) QuestionError(err error) {
	q.logger.WithField("Type", "Question").
		Error(err)
}

func (q questionLogger) QuestionInfo(info string) {
	q.logger.WithField("Type", "Question").
		Info(info)
}

func NewLogrusQuestionLogger(logger *logrus.Logger) domain.QuestionLogger {
	return &questionLogger{
		logger,
	}
}
