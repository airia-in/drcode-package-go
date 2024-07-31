package drcodewrapper

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

type Config struct {
	Protocol           string
	PublicKey          string
	Host               string
	Port               int
	ProjectID          string
	TracesSampleRate   float64
	ProfilesSampleRate float64
}

func constructDSN(config Config) string {
	return fmt.Sprintf("%s://%s@%s:%d/%s", config.Protocol, config.PublicKey, config.Host, config.Port, config.ProjectID)
}

func InitDrcode(config Config) error {
	dsn := constructDSN(config)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:                dsn,
		TracesSampleRate:   config.TracesSampleRate,
		ProfilesSampleRate: config.ProfilesSampleRate,
	})

	return err
}