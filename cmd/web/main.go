package main

import (
	"context"
	"exchange/pkg"
	_http "exchange/pkg/http"
	"exchange/pkg/infrastructure/currency/currencyapi"
	"exchange/pkg/infrastructure/mail"
	"exchange/pkg/repository/mem"
	"exchange/pkg/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	ctx := context.Background()

	e := echo.New()
	mailRepo := mem.NewMemoryRepository()

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

	_http.NewCurrencyHandler(e, srvs)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
