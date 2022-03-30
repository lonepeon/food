package main

import (
	"context"
	"embed"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/sessions"

	"github.com/lonepeon/food/internal/infrastructure/www"
	"github.com/lonepeon/golib/env"
	"github.com/lonepeon/golib/logger"
	"github.com/lonepeon/golib/web"
)

type Config struct {
	WebAddress       string `env:"FOOD_WEB_ADDR,required=true"`
	SessionStorePath string `env:"FOOD_SESSION_STORE_PATH,default=./tmp/sessions"`
	SessionKey       string `env:"FOOD_SESSION_KEY,required=true"`
}

//go:embed templates/*
var htmlTemplateFS embed.FS

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	log, closer := logger.NewLogger(os.Stdout)
	defer func() {
		if err := closer(); err != nil {
			fmt.Fprintf(os.Stderr, "can't flush logs: %v", err)
		}
	}()

	log = log.WithFields(logger.String("app-version", "HEAD"), logger.String("app-name", "food"))

	var cfg Config
	if err := env.Load(&cfg); err != nil {
		return fmt.Errorf("can't load config: %v", err)
	}

	var sessionstore = sessions.NewFilesystemStore(cfg.SessionStorePath, []byte(cfg.SessionKey))
	sessionstore.MaxLength(math.MaxInt64)

	webServer := initWebServer(log, sessionstore)
	webServer.HandleFunc("GET", "/", www.RecipeIndex())

	return waitForServersShutdown(log, webServer, cfg.WebAddress)
}

func initWebServer(log *logger.Logger, sessionstore *sessions.FilesystemStore) *web.Server {
	tmpl := web.TmplConfiguration{
		FS:                          htmlTemplateFS,
		Layout:                      "templates/layout.html.tmpl",
		ErrorLayout:                 "templates/layout.html.tmpl",
		RedirectionTemplate:         "templates/30x.html.tmpl",
		NotFoundTemplate:            "templates/404.html.tmpl",
		InternalServerErrorTemplate: "templates/500.html.tmpl",
		UnauthorizedTemplate:        "templates/401.html.tmpl",
	}

	return web.NewServer(log, tmpl, sessionstore)
}

func waitForServersShutdown(log *logger.Logger, webServer *web.Server, webAddress string) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	webErr := make(chan error)

	webServerStarter, webServerStopper := spawnWebServer(log, webServer, webAddress, webErr)
	go webServerStarter()

	var err error
	shutdownDuration := 20 * time.Second

	select {
	case <-sigs:
		if webErr := webServerStopper(shutdownDuration); err != nil {
			return fmt.Errorf("%v; %v", err, webErr)
		}
	case e := <-webErr:
		return e
	}

	return nil
}

func spawnWebServer(log *logger.Logger, server *web.Server, webAddress string, reporter chan<- error) (func(), func(time.Duration) error) {
	start := func() {
		log.Infof("starting web server on %s", webAddress)
		if err := server.ListenAndServe(webAddress); err != nil {
			reporter <- fmt.Errorf("web server failed: %v", err)
		}
	}

	stop := func(d time.Duration) error {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		return server.Shutdown(ctx)
	}

	return start, stop
}
