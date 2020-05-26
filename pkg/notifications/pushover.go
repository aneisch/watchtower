package notifications

import (
	t "github.com/containrrr/watchtower/pkg/types"
	"github.com/johntdyer/slackrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
  "github.com/gregdel/pushover"
)

const (
	pushoverType = "pushover"
)

type pushoverTypeNotifier struct {
	pushover.Pushover
}

func newPushoverkNotifier(c *cobra.Command, acceptedLogLevels []log.Level) t.Notifier {
	flags := c.PersistentFlags()

	app, _ := flags.GetString("notification-pushover-app")

	n := &pushoverTypeNotifier{
		Pushover: pushover.Pushover{
			token:        app,
		},
	}

	log.AddHook(n)
	return n
}

func (s *pushoverTypeNotifier) StartNotification() {}

func (s *pushoverTypeNotifier) SendNotification() {}
