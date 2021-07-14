package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/olivere/elastic"
	"gopkg.in/sohlich/elogrus.v3"

	esconf "github.com/mwy001/goland/pkg/conf/elasticsearch"
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
		lg.Formatter = &logrus.JSONFormatte{}

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

		l = lg.WithFields(logrus.Fields{"App": app, "Host": hostName(), "Env":envconf.CurrentEnvironment()})
	})

	return l
}

// L returns singleton instance of logger
func L() *logrus.Entry {
	return l
}
