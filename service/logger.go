package service

import (
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func initLog() {
	log = logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  time.RFC3339Nano,
	})

	log.SetOutput(os.Stdout)

	switch strings.ToLower(strings.ToLower(os.Getenv("CONFIG_LOG_LEVEL"))) {
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}

func Log(level string, label string, message string) {
	switch strings.ToLower(level) {
	case "panic":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Panic(message)
	case "fatal":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Fatal(message)
	case "error":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Error(message)
	case "warn":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Warn(message)
	case "debug":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Debug(message)
	case "tarce":
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Trace(message)
	default:
		log.WithFields(logrus.Fields{
			"service": strings.ToLower(os.Getenv("CONFIG_SERVICE_NAME")),
			"label":   label,
		}).Info(message)
	}
}
