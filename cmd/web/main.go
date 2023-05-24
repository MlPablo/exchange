package main

import (
	"context"
	"exchange/pkg"
	_http "exchange/pkg/http"
	"exchange/pkg/infrastructure/currency/currencyapi"
	"exchange/pkg/infrastructure/mail"
	"exchange/pkg/repository/filesysytem"
	"exchange/pkg/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logrus.Info("starting application...")
	// mailRepo := mem.NewMemoryRepository()
	mailRepo, err := filesysytem.NewFileSystemRepository(os.Getenv("FILE_STORE_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	mCfg, err := mail.NewConfig(
		os.Getenv("EMAIL_LOGIN"),
		os.Getenv("EMAIL_APP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mailPusher := mail.NewMailService(mCfg)

	currencyGetter := currencyapi.NewCurrencyApi(
		currencyapi.NewConfig(os.Getenv("CURR_API_KEY")),
		os.Getenv("CURR_URL"),
	)

	userMailService := services.NewEmailUserService(ctx, mailRepo)
	notifierService := services.NewNotificationService(ctx, mailRepo, currencyGetter, mailPusher)

	srvs := pkg.NewServices(currencyGetter, userMailService, notifierService)

	e := echo.New()
	e.Use(getServerLogger())
	_http.NewCurrencyHandler(e, srvs)

	go func() {
		if err := e.Start(":" + os.Getenv("SERVER_ADDR")); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	logrus.Info("application started =)")

	go syscallWait(cancel)
	<-ctx.Done()

	logrus.Info("application stopped.")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}

func getServerLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}
