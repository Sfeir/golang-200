package utils

import (
	"github.com/Sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"time"
)

const (
	// LogStashFormatter is constant used to format logs as logstash format
	LogStashFormatter = "logstash"
	// TextFormatter is constant used to format logs as simple text format
	TextFormatter = "text"
)

// InitLog initializes the logrus logger
func InitLog(logLevel, formatter string) error {

	switch formatter {
	case LogStashFormatter:
		logrus.SetFormatter(&logrus_logstash.LogstashFormatter{
			TimestampFormat: time.RFC3339,
		})
	default:
		// TODO write the default case
		// TODO Set the formatter using the standard logrus TextFormatter with Forced colors and Full timestamp
	}

	// TODO set the standard output to os.Stdout

	// TODO parse the logLevel param

	// TODO check the parsing error
	// TODO if error occurs set the logger level as DebugLevel and return the error

	// TODO if no error occurred, set the parsed level as the logger level
	return nil
}
