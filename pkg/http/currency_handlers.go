package http

import (
	"context"
	"exchange/pkg"
	"exchange/pkg/domain"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type ExchangeHandler struct {
	services pkg.Services
}

type Controllers interface {
	RegisterHandlers()
}

func NewCurrencyHandler(e *echo.Echo, services *pkg.Services) {
	handler := &ExchangeHandler{
		services: *services,
	}
	e.GET("/rate", handler.GetBtcToUahCurrency)
	e.POST("/subscribe", handler.CreateMailSubscriber)
	e.POST("/sendEmails", handler.SendEmails)
}

func (e *ExchangeHandler) GetBtcToUahCurrency(c echo.Context) error {
	ctx := c.Request().Context()
	log.Print("got request")
	cur := domain.GetBitcoinToUAH()
	resp, err := e.services.CurrencyService.GetCurrency(ctx, cur)
	if err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	return c.JSON(http.StatusOK, resp)
}

func (e *ExchangeHandler) SendEmails(c echo.Context) error {
	go func() {
		if err := e.services.NotificatioinService.Notify(
			context.Background(),
			domain.DefaultNotification(),
		); err != nil {

			// TODO: ADD LOGGER WITH ERROR
		}
	}()

	return c.JSON(http.StatusOK, nil)
}

func (e *ExchangeHandler) CreateMailSubscriber(c echo.Context) error {
	ctx := c.Request().Context()

	email := domain.NewEmailUser(c.FormValue("email"))

	if err := email.Validate(); err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	if err := e.services.EmailUserService.NewEmailUser(ctx, email); err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServer:
		return http.StatusInternalServerError
	case domain.ErrAlreadyExist:
		return http.StatusConflict
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrInvalidStatus, domain.ErrBadRequst:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
