package main

import (
	"context"
	"flag"
	"fmt"
	"locations/api"
	"locations/gen/locations"

	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		envF      = flag.String("env", "development", "Server host (valid values: development, production)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("port", "", "HTTP port (overrides host HTTP port specified in service design)")
		versionF  = flag.String("version", "v1", "API version")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[user-manager] ", log.Ltime)
	}

	// todo: Initialize the services.
	locationAPI := api.NewLocation()
	locationsEndpoints := locations.NewEndpoints(locationAPI)

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *envF {
	case "development":
		{
			addr := "http://0.0.0.0:8000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			api.HandleHTTPServer(ctx, u, locationsEndpoints, &wg, errc, logger, *dbgF)
		}

	case "production":
		{
			addr := "https://{version}.goa.design"
			addr = strings.Replace(addr, "{version}", *versionF, -1)
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":443"
			}
			api.HandleHTTPServer(ctx, u, locationsEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development|production)\n", *envF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
