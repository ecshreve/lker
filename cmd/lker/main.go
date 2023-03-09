package main

import (
	"fmt"
	"strings"

	"github.com/ecshreve/lker/pkg/server"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type LogFormat struct {
	log.TextFormatter
}

func (f *LogFormat) Format(entry *log.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)
	return []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m - [%s] - <%s> - %s\n", levelColor, strings.ToUpper(entry.Level.String()), entry.Time.Format(f.TimestampFormat), strings.TrimPrefix(entry.Caller.Function, "github.com/ecshreve/lker"), entry.Message)), nil
}

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
	colorPurple = 35
)

func getColorByLevel(level log.Level) int {
	switch level {
	case log.TraceLevel:
		return colorBlue
	case log.WarnLevel:
		return colorYellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		return colorRed
	case log.DebugLevel:
		return colorPurple
	default:
		return colorBlue
	}
}

func main() {

	log.SetFormatter(&LogFormat{log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	},
	})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	log.Trace("---> - enter")
	defer log.Trace("<--- - exit")

	s := server.NewServer()
	if err := s.Serve(); err != nil {
		log.Error("error returned from server", oops.Wrapf(err, "wrapped error from server"))
	}
}
