// Package drcodewrapper provides a wrapper around Sentry's SDK to simplify its configuration and initialization.
package drcodewrapper

import (
	"fmt"
	"github.com/getsentry/sentry-go"
)

// Config represents the configuration required to initialize Sentry.
// It includes parameters such as protocol, public key, host, port, project ID,
// and optional sampling rates for traces and profiles.
type Config struct {
	Protocol           string  // Protocol for the Sentry DSN (e.g., "https")
	PublicKey          string  // Public key for authenticating with Sentry
	Host               string  // Host where Sentry is hosted (e.g., "sentry.io")
	Port               int     // Port on which Sentry is accessible
	ProjectID          string  // Unique identifier for the Sentry project
	TracesSampleRate   float64 // Optional: Sample rate for tracing (default: 1.0)
	ProfilesSampleRate float64 // Optional: Sample rate for profiling (default: 1.0)
}

// constructDSN constructs the Data Source Name (DSN) required by Sentry to initialize the client.
// It combines the protocol, public key, host, port, and project ID into a single DSN string.
func constructDSN(config Config) string {
	return fmt.Sprintf("%s://%s@%s:%d/%s", config.Protocol, config.PublicKey, config.Host, config.Port, config.ProjectID)
}

// InitDrcode initializes the Sentry client with the provided configuration.
// It sets up Sentry for error tracking, tracing, and profiling based on the given Config.
// Returns an error if the initialization fails.
func InitDrcode(config Config) error {
	dsn := constructDSN(config)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:                dsn,
		TracesSampleRate:   config.TracesSampleRate,
		ProfilesSampleRate: config.ProfilesSampleRate,
	})

	return err
}
