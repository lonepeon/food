package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3" // sqlite3 adapter

	"github.com/lonepeon/food/internal/infrastructure/www"
	"github.com/lonepeon/golib/env"
	"github.com/lonepeon/golib/logger"
	"github.com/lonepeon/golib/sqlutil"
	"github.com/lonepeon/golib/web"
	"github.com/lonepeon/golib/web/authenticationstore"
	"github.com/lonepeon/golib/web/sessionstore"
)

type Config struct {
	SQLitePath           string `env:"FOOD_SQLITE_PATH,default=./food.sqlite"`
	WebAddress           string `env:"FOOD_WEB_ADDR,required=true"`
	SessionKey           string `env:"FOOD_SESSION_KEY,required=true"`
	AuthenticationPepper string `env:"FOOD_AUTH_PEPPER,required=true"`
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

	log.Infof("initialize database from %v", cfg.SQLitePath)
	db, err := initDatabase(log, cfg.SQLitePath)
	if err != nil {
		return fmt.Errorf("can't initialize database: %v", err)
	}

	sessionstore := sessionstore.NewSQLite(db, sessions.Options{
		HttpOnly: true,
		MaxAge:   1 * 60 * 60 * 24 * 2,
	}, []byte(cfg.SessionKey))

	authenticationBrowserStore := web.NewCurrentAuthenticatedUserSessionStore(sessionstore)
	authenticationBackendstore := authenticationstore.NewSQLite(db, cfg.AuthenticationPepper)

	authentication := web.NewAuthentication(authenticationBrowserStore, authenticationBackendstore, "templates/authentication/login.html.tmpl")

	if _, err := sqlutil.ExecuteMigrations(context.Background(), db, authenticationstore.Migrations()); err != nil {
		return fmt.Errorf("can't run authentication schema migrations: %v", err)
	}

	if _, err := authenticationBackendstore.Register("admin", "admin"); err != nil {
		if !errors.Is(err, web.ErrUserAlreadyExist) {
			return fmt.Errorf("can't create default user: %v", err)
		}
	}

	webServer := initWebServer(log, sessionstore)
	webServer.HandleFunc("GET", "/", www.RecipeIndex())
	webServer.HandleFunc("GET", "/admin/login/new", authentication.ShowLoginPage("/admin"))
	webServer.HandleFunc("POST", "/admin/login", authentication.Login("/admin"))
	webServer.HandleFunc("GET", "/admin/logout", authentication.Logout("/"))
	webServer.HandleFunc("GET", "/admin/ingredients", authentication.EnsureAuthentication("/admin/login/new", www.RecipeIndex()))

	return waitForServersShutdown(log, webServer, cfg.WebAddress)
}

func initWebServer(log *logger.Logger, sessionstore sessions.Store) *web.Server {
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

func initDatabase(log *logger.Logger, sqlitePath string) (*sql.DB, error) {
	options := make(url.Values)
	// busy_timeout is used to wait for a little bit when SQLite is backing up the data
	// instead of rejecting a transaction as soon as we access a locked row.
	options.Add("_busy_timeout", "3000")
	options.Add("_foreign_keys", "true")
	// journal_mode wal is enabled to let SQLite back up run properly
	options.Add("_journal_mode", "wal")
	// synchronous normal means fsync() call are only called when WAL becomes full
	options.Add("_synchronous", "NORMAL")

	db, err := sql.Open("sqlite3", sqlitePath+"?"+options.Encode())
	if err != nil {
		return nil, fmt.Errorf("can't open  sqlite file: %v", err)
	}

	sessionStoreMigrationsVersions, err := sqlutil.ExecuteMigrations(context.Background(), db, sessionstore.Migrations())
	if err != nil {
		return nil, fmt.Errorf("can't run session store migrations: %v", err)
	}
	log.Infof("database executed new sql session store migrations %s", strings.Join(sessionStoreMigrationsVersions, ", "))

	return db, nil
}
