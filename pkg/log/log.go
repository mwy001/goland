package log

import (
	"net"
	"os"
	"sync"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"

	esconf "github.com/mwy001/goland/pkg/conf/elasticsearch"
	logstashconf "github.com/mwy001/goland/pkg/conf/logstash"

	envconf "github.com/mwy001/goland/pkg/conf/env"
)

var (
	l       *logrus.Entry
	onceLog sync.Once
)

func hostName() string {
	name, _ := os.Hostname()
	return name
}

// InitLogger initiate the logger for current application
func InitLogger(index string, app string) *logrus.Entry {
	onceLog.Do(func() {
		lg := logrus.New()
		lg.SetReportCaller(true)

		lg.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05-0700",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFile:  "class",
			},
		}

		if envconf.ElasticSearchLogEnabled() == "1" {
			client, err := elastic.NewClient(
				elastic.SetURL(esconf.Config().Es.Address),
				elastic.SetBasicAuth(esconf.Config().Es.User, esconf.Config().Es.Pass),
				elastic.SetSniff(false))

			if err != nil {
				l.Panic(err)
			}

			hook, err := elogrus.NewAsyncElasticHook(client, "", logrus.DebugLevel, index)
			if err != nil {
				l.Panic(err)
			}
			lg.AddHook(hook)
		}

		if envconf.LogstashLoggingEnabled() == "1" {
			conn, err := net.Dial("tcp", logstashconf.Config().Lc.LogstashDestinationURL)

			if err != nil {
				l.Panic(err)
			}

			hook := logrustash.New(conn, lg.Formatter)

			lg.AddHook(hook)
		}

		l = lg.WithFields(logrus.Fields{"app": app, "host": hostName(), "environment": envconf.CurrentEnvironment()})
	})

	return l
}

// L returns singleton instance of logger
func L() *logrus.Entry {
	return l
}
