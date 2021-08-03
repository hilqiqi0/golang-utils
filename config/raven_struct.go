package config

import (
	"github.com/getsentry/raven-go"
	"github.com/hilqiqi0/golang-utils/tools/errs"
)

func (c *ConfigEngine) SentryRavenInit(name string) error {
	sentryDSN := c.GetString(name)
	errs.CheckEmptyValue(sentryDSN)
	err := raven.SetDSN(sentryDSN)
	errs.CheckFatalErr(err)
	return err
}
