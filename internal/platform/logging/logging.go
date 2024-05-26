// internal/platform/logging/logging.go
package logging

import (
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
