/** Using Package Logrus from Sirupsen Golang */
package loggers

import (
	"os"

	packagesConfig "boilerplate-go/src/packages/config"

	"github.com/sirupsen/logrus"
)

type LogrusLog struct {
	*logrus.Logger
}

var logger = &LogrusLog{}

/** Setup Logrus */
func SetupLogrusLogger() {
	logger = &LogrusLog{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetOutput(os.Stdout)

	if packagesConfig.GetApplicationConfig().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func GetLogrusLogger() *LogrusLog {
	return logger
}
