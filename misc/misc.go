package misc

import (
	//"github.com/ssorathia/go-ms-skel/app/logging"
	"github.com/Sirupsen/logrus"
	"github.com/amitm1/go-microsvc-skel/config"
	"gopkg.in/alexcesaro/statsd.v2"
)

type RequestHelpers struct {
	Logging   *logrus.Logger
	Statsd    *statsd.Client
	Config    *config.Config
	RequestId string
}
