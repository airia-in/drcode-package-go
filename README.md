### drcode-wrapper

**drcode-wrapper** is a Go package that provides a wrapper around Sentry, making it easy to configure and use Sentry's features, including profiling and error handling.

## Installation

To install the package, run:

bash

Copy code

`go get github.com/codewithnadeem14502/drcode-wrapper-go`

## Getting Started

**Basic Usage**

go

Copy code

<pre><code class="language-const">
package main

import (
    "github.com/codewithnadeem14502/drcode-wrapper-go"
    "log"
)

func main() {
    config := drcodewrapper.Config{
        Protocol:           "https",
        PublicKey:          "yourPublicKey",
        Host:               "sentry.io",
        Port:               443,
        ProjectID:          "yourProjectId",
        TracesSampleRate:   1.0,  // optional
        ProfilesSampleRate: 1.0,  // optional
    }

    err := drcodewrapper.InitDrcode(config)
    if err != nil {
        log.Fatalf("sentry.Init: %s", err)
    }

    // Your code here
}



</code></pre>

## Configuration

**Configuration Object**

- `Protocol` (string): The protocol to use (e.g., 'https').
- `PublicKey` (string): The public key for Sentry.
- `Host` (string): The Sentry host (e.g., 'sentry.io').
- `Port` (int): The port for the Sentry server.
- `ProjectID` (string): The Sentry project ID.
- `TracesSampleRate` (float64, optional): The sample rate for tracing (default: 1.0).
- `ProfilesSampleRate` (float64, optional): The sample rate for profiling (default: 1.0).

## API Reference

**Functions**

- **InitDrcode(config Config) error**

  Initializes Sentry with the provided configuration.

  - **config** (Config): The configuration object.

**Configuration Object**

The configuration object should contain the following fields:

- `Protocol` (string): The protocol to use (e.g., 'https').
- `PublicKey` (string): The public key for Sentry.
- `Host` (string): The Sentry host (e.g., 'sentry.io').
- `Port` (int): The port for the Sentry server.
- `ProjectID` (string): The Sentry project ID.
- `TracesSampleRate` (float64, optional): The sample rate for tracing (default: 1.0).
- `ProfilesSampleRate` (float64, optional): The sample rate for profiling (default: 1.0).

## License

This project is licensed under the MIT License.

## Version

0.1.2
