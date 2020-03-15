package server

import (
	"sre-sms-server/middleware/auth"

	"github.com/robfig/cron/v3"
)

func init() {
	//tasks.Init()
	auth.LoadAuthData()
	c := cron.New()
	c.AddFunc("@every 10s", auth.LoadAuthData)
	c.Start()
}
