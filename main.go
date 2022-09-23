package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goredis "github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/jakob-moeller-cloud/octi-sync-server/api"
	"github.com/jakob-moeller-cloud/octi-sync-server/config"
	"github.com/jakob-moeller-cloud/octi-sync-server/service/redis"
	"github.com/rs/zerolog"
	baseLogger "github.com/rs/zerolog/log"
)

// Func main should be as small as possible and do as little as possible by convention.
func main() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var logger zerolog.Logger

	switch cfg.LogSettings.Format {
	case config.LogSettingsFormatPretty:
		logger = zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
	case config.LogSettingsFormatJSON:
		fallthrough
	case config.LogSettingsFormatNone:
		fallthrough
	default:
		logger = baseLogger.With().Logger()
	}

	cfg.Logger = &logger

	uuid.EnableRandPool()

	// Run the server
	Run(cfg)
}

// Run will run the HTTP Server.
func Run(config *config.Config) {
	startUpContext, cancelStartUpContext := context.WithCancel(context.Background())
	defer cancelStartUpContext()

	clients, err := redis.NewClientsWithRegularPing(startUpContext, config, redis.ClientMutators{
		"default": nil,
		"shareClient": func(client *goredis.Client) *goredis.Client {
			return client.WithTimeout(1 * time.Hour)
		},
	})
	if err != nil {
		log.Print(err)

		return
	}

	config.Services.Accounts = &redis.Accounts{Client: clients["default"], ShareClient: clients["shareClient"]}
	config.Services.Modules = &redis.Modules{Client: clients["default"]}
	config.Services.Devices = &redis.Devices{Client: clients["default"]}

	// Define server options
	srv := &http.Server{
		Addr:              config.Server.Host + ":" + config.Server.Port,
		Handler:           api.New(startUpContext, config),
		ReadTimeout:       config.Server.Timeout.Read,
		ReadHeaderTimeout: config.Server.Timeout.Read,
		WriteTimeout:      config.Server.Timeout.Write,
		IdleTimeout:       config.Server.Timeout.Idle,
	}

	idleConsClosed := make(chan struct{})
	closeServer := func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint // We received an interrupt signal, shut down.
		// Set up a context to allow for graceful server shutdowns in the event
		// of an OS interrupt (defers the cancel just in case)
		ctx, cancel := context.WithTimeout(
			startUpContext,
			config.Server.Timeout.Server,
		)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			config.Logger.Warn().Msg("server shutdown error: " + err.Error())
		}

		close(idleConsClosed)
	}

	go closeServer()

	// service connections
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		config.Logger.Fatal().Msg("listen: %s" + err.Error())
	}

	<-idleConsClosed
	config.Logger.Info().Msg("server shut down finished")
}
